// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package jsonrpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	bulk "github.com/arminguenther/xeruspower-go/v40413/bulkrpc/bulkrequest"
)

func TestCallResults(t *testing.T) {
	var id int64 = 42
	req := bulk.JsonObject{"jsonrpc": "2.0", "id": id, "method": "getMetaData"}
	resp := bulk.JsonObject{"jsonrpc": "2.0", "id": json.Number("42"), "result": true}
	res, err := callResults(req, resp)
	if err != nil {
		t.Error(err)
	}
	if res != true {
		t.Error("result should be true")
	}
	msg := "parsing failed"
	_, err = callResults(req, bulk.JsonObject{"jsonrpc": "2.0", "id": nil, "error": map[string]any{
		"code":    json.Number("-32700"),
		"message": msg,
	}})
	var e RPCError
	if ok := errors.As(err, &e); err == nil || !ok {
		t.Errorf("expected RPC error, got %v", err)
	}
	if e.Code != -32700 {
		t.Errorf("expected code %d, got %d", -32700, e.Code)
	}
	if e.Message != msg {
		t.Errorf("expected message %q, got %q", msg, e.Message)
	}
	_, err = callResults(req, bulk.JsonObject{"jsonrpc": "3.0"})
	if err == nil || err.Error() != "version not 2.0" {
		t.Errorf("expected version error, got %v", err)
	}
	_, err = callResults(req, bulk.JsonObject{"jsonrpc": "2.0", "id": json.Number("7"), "result": true})
	if err == nil || err.Error() != fmt.Sprintf("wrong response(%d) for request(%d)", 7, id) {
		t.Errorf("expected ID error, got %v", err)
	}
}
