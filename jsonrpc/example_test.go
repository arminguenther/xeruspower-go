// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package jsonrpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/arminguenther/xeruspower-go/v40220/idl/event"
	"github.com/arminguenther/xeruspower-go/v40220/pdumodel/pdu"
	"github.com/arminguenther/xeruspower-go/v40220/session/sessionmanager"
)

func ExampleCodeFor() {
	code, ok := CodeFor[event.Event]()
	fmt.Println(code, ok)
	// Output: idl.Event:1.0.0 true
}

func ExampleAgent() {
	agent := NewAgent("https", "10.0.42.2", "admin", "raritan")
	_ = pdu.NewPdu("/model/pdu/0", agent)
}

func ExampleAgent_insecureSkipVerify() {
	// disable certificate verification if TLS is not set up yet
	agent := NewAgent("https", "10.0.42.2", "admin", "raritan")
	tp := http.DefaultTransport.(*http.Transport).Clone()
	tp.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	agent.Client = &http.Client{Transport: tp}
}

func ExampleAgent_sessionToken() {
	agent := NewAgent("https", "10.0.42.2", "admin", "raritan")
	mgr := sessionmanager.NewSessionManager("/session", agent)
	ret, _, token, err := mgr.NewSession(context.Background())
	if err != nil {
		panic(err)
	}
	switch ret {
	case 0:
		agent.Token = token
	case 1:
		panic("session creation denied due to single login limitation")
	default:
		panic(fmt.Sprintf("session creation failed; ret %d", ret))
	}
}
