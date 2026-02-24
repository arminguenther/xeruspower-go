// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

// Package jsonrpc provides an [Agent] that speaks JSON-RPC 2.0 over HTTP and implements [idl.Caller].
// It also initializes registries for decoding [idl.Object] and [idl.ValueObject] types.
package jsonrpc

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"sync"
	"sync/atomic"

	bulk "github.com/arminguenther/xeruspower-go/bulkrpc/bulkrequest"
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
)

// A TypeCoder can state its type as an [idl.TypeCode].
type TypeCoder interface {
	TypeCode() idl.TypeCode
}

// CodeFor looks up the [idl.TypeCode] of a given type.
// Works with any [idl.Object] or [idl.ValueObject].
func CodeFor[T TypeCoder]() (idl.TypeCode, bool) {
	key := reflect.TypeFor[T]()
	if newFn, ok := valobj.TypeRegistry[key]; ok {
		return newFn().TypeCode(), true
	}
	if newFn, ok := object.TypeRegistry[key]; ok {
		return newFn("", nil).TypeCode(), true
	}
	return idl.TypeCode{}, false
}

// NewAgent creates a new Agent.
// Valid protocol schemes are "http" and "https".
// A host is specified by hostname or IP address and can include a port.
// Returned Agent uses username and password for HTTP Basic Authentication.
func NewAgent(scheme, host, username, password string) *Agent {
	return &Agent{Scheme: scheme, Host: host, Username: username, Password: password}
}

var _ idl.Caller = (*Agent)(nil)

// An Agent speaks JSON-RPC 2.0 over HTTP.
// It stores connection details and user credentials.
//
// Use [NewAgent] to create an Agent with HTTP Basic Authentication.
// Set a session Token to use session authentication instead.
// Requests are sent with [http.DefaultClient] unless Client is set.
//
// Agents are safe for concurrent use by multiple goroutines.
// An Agent may bundle concurrent calls in a single bulk request to decrease transport overhead.
// An individual request can be forced with a context created by [WithoutBulk].
type Agent struct {
	// Scheme specifies the protocol, either "http" or "https".
	Scheme string
	// Host specifies the hostname or IP address of the RPC server.
	// It may include a port number.
	Host string
	// Username is used for HTTP Basic Authentication.
	Username string
	// Password is used for HTTP Basic Authentication.
	Password string
	// Token can be used to authenticate through a session.
	// If Token is set, Username and Password are ignored.
	Token string
	// Client sends out JSON-RPC requests as HTTP POSTs.
	// If nil, http.DefaultClient is used.
	Client *http.Client
	// once initializes startSem and more channels
	// on the first Call that allows bulk processing.
	once sync.Once
	// startSem is a semaphore on callMore goroutines.
	startSem chan struct{}
	// more passes additional calls to callMore,
	// to merge them into a single bulk request.
	more chan methodCall
}

type bulkOff struct{}

var bulkOffKey = bulkOff{}

// WithoutBulk returns a derived context that points to the parent context.
// The derived context forces [Agent.Call] to make an individual request.
func WithoutBulk(parent context.Context) context.Context {
	return context.WithValue(parent, bulkOffKey, struct{}{})
}

type methodCall struct {
	ctx     context.Context
	rid     string
	request bulk.JsonObject
	retCh   chan<- callReturnValues
}

type callReturnValues struct {
	result any
	err    error
}

// requestID is increased on each JSON-RPC request.
var requestID atomic.Int64

// Call calls a method with given parameters on the remote object identified by the resource ID.
func (a *Agent) Call(ctx context.Context, rid string, method string, params map[string]any) (result any, err error) {
	id := requestID.Add(1)
	req := bulk.JsonObject{"jsonrpc": "2.0", "id": id, "method": method}
	if params != nil {
		req["params"] = params
	}
	if ctx.Value(bulkOffKey) != nil {
		var resp bulk.JsonObject
		resp, err = a.do(ctx, rid, req)
		if err != nil {
			return nil, err
		}
		return callResults(req, resp)
	}
	a.once.Do(func() {
		a.startSem = make(chan struct{}, 1)
		a.more = make(chan methodCall)
	})
	ch := make(chan callReturnValues, 1)
	call := methodCall{ctx: ctx, rid: rid, request: req, retCh: ch}
	select {
	case a.startSem <- struct{}{}:
		go a.callMore(call)
	case a.more <- call:
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	select {
	case ret := <-ch:
		return ret.result, ret.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// collectCalls returns a slice of available calls starting with call0.
// The number of total calls is limited to 200 to reduce server load.
func (a *Agent) collectCalls(call0 methodCall) []methodCall {
	calls := []methodCall{call0}
	for len(calls) < 200 {
		select {
		case call := <-a.more:
			calls = append(calls, call)
		default:
			return calls
		}
	}
	return calls
}

// callMore does call0 and more calls if available.
// If more are available, calls are done in a single JSON-RPC bulk request.
// Result and error of each call are sent to the respective return channel.
func (a *Agent) callMore(call0 methodCall) {
	defer func() { <-a.startSem }()
	calls := a.collectCalls(call0)
	if len(calls) == 1 {
		var ret callReturnValues
		var resp bulk.JsonObject
		resp, ret.err = a.do(call0.ctx, call0.rid, call0.request)
		if ret.err == nil {
			ret.result, ret.err = callResults(call0.request, resp)
		}
		call0.retCh <- ret
		return
	}
	ctx, cancel := context.WithCancel(WithoutBulk(context.Background()))
	defer cancel()
	go func() {
		for _, call := range calls {
			select {
			case <-call.ctx.Done():
			case <-ctx.Done():
				return
			}
		}
		cancel()
	}()
	requests := make([]bulk.BulkRequestRequest, len(calls))
	for i, call := range calls {
		requests[i] = bulk.BulkRequestRequest{Rid: call.rid, Json: call.request}
	}
	responses, errBulk := bulk.NewBulkRequest("/bulk", a).PerformBulk(ctx, requests)
	makeRet := func(i int) (ret callReturnValues) {
		if errBulk != nil {
			ret.err = errBulk
			return ret
		}
		if i >= len(responses) {
			ret.err = errors.New("response missing in bulk response")
			return ret
		}
		if responses[i].Statcode != http.StatusOK {
			ret.err = HTTPError(responses[i].Statcode)
			return ret
		}
		ret.result, ret.err = callResults(requests[i].Json, responses[i].Json)
		return ret
	}
	for i := range calls {
		select {
		case <-calls[i].ctx.Done():
			continue
		default:
			calls[i].retCh <- makeRet(i)
		}
	}
}

// do sends a JSON-RPC request and returns a JSON-RPC response.
func (a *Agent) do(ctx context.Context, rid string, request bulk.JsonObject) (response bulk.JsonObject, err error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	u := &url.URL{Scheme: a.Scheme, Host: a.Host, Path: rid}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-Client-Type", "Go RPC")
	switch a.Token {
	case "":
		httpReq.SetBasicAuth(a.Username, a.Password)
	default:
		httpReq.Header.Set("X-Sessiontoken", a.Token)
	}
	client := http.DefaultClient
	if a.Client != nil {
		client = a.Client
	}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer func() { _ = httpResp.Body.Close() }()
	if httpResp.StatusCode != http.StatusOK {
		return nil, HTTPError(httpResp.StatusCode)
	}
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.UseNumber()
	err = decoder.Decode(&response)
	return response, err
}

// callResults returns the result of a JSON-RPC request and response pair.
// It returns an error if the:
//   - response is malformed
//   - response protocol version is not 2.0
//   - request and response IDs are not equal
//   - response holds a JSON-RPC error
func callResults(request, response bulk.JsonObject) (any, error) {
	version, err := encoding.InAndIs[string]("jsonrpc", response)
	if err != nil {
		return nil, err
	}
	if version != "2.0" {
		return nil, errors.New("version not 2.0")
	}
	if err = encoding.In("id", response); err != nil {
		return nil, err
	}
	id, err := encoding.AsInt64(response["id"])
	switch {
	case err != nil:
		if _, ok := response["error"]; ok && response["id"] == nil {
			break
		}
		return nil, err
	case id != request["id"].(int64):
		return nil, fmt.Errorf("wrong response(%d) for request(%d)", id, request["id"])
	}
	if result, ok := response["result"]; ok {
		return result, nil
	}
	errorObj, err := encoding.InAndIs[map[string]any]("error", response)
	if err != nil {
		return nil, err
	}
	if err = encoding.In("code", errorObj); err != nil {
		return nil, err
	}
	code, err := encoding.AsInt(errorObj["code"])
	if err != nil {
		return nil, err
	}
	msg, err := encoding.InAndIs[string]("message", errorObj)
	if err != nil {
		return nil, err
	}
	return nil, RPCError{code, msg}
}

// RPCError is a JSON-RPC protocol error.
type RPCError struct {
	Code    int
	Message string
}

func (e RPCError) Error() string {
	return fmt.Sprintf("JSON-RPC error %d: %s", e.Code, e.Message)
}

// An HTTPError is an error on the HTTP protocol level.
// Any status code other than [http.StatusOK] is seen as an error.
type HTTPError int

func (e HTTPError) Error() string {
	return fmt.Sprintf("HTTP status %d: %s", e, http.StatusText(int(e)))
}
