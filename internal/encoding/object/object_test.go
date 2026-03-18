// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package object_test

import (
	"strings"
	"testing"

	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40033/jsonrpc"
	"github.com/arminguenther/xeruspower-go/v40033/pdumodel/edevice"
	"github.com/arminguenther/xeruspower-go/v40033/pdumodel/outlet"
	"github.com/arminguenther/xeruspower-go/v40033/pdumodel/pdu"
)

func TestAs(t *testing.T) {
	encode := func(code string) map[string]any {
		return map[string]any{"rid": "/some/rid", "type": code}
	}
	// Pdu as expected
	c, _ := jsonrpc.CodeFor[pdu.Pdu]()
	code := c.String()
	_, err := object.As[pdu.Pdu](encode(code), nil)
	if err != nil {
		t.Fatal(err)
	}
	// Outlet extends EDevice
	c, _ = jsonrpc.CodeFor[outlet.Outlet]()
	code = c.String()
	dev, err := object.As[edevice.EDevice](encode(code), nil)
	if err != nil {
		t.Fatal(err)
	}
	_, ok := dev.(outlet.Outlet)
	if !ok {
		t.Fatalf("should be an Outlet proxy: %T", dev)
	}
	// invalid code syntax
	codes := []string{
		"pdumodel.Pdu_1.0.0",
		"pdumodel.Pdu:1.0.0.0",
		"pdumodel.Pdu:A.0.0",
		"pdumodel.Pdu:1.B.0",
		"pdumodel.Pdu:1.0.C",
	}
	for _, code := range codes {
		_, err = object.As[pdu.Pdu](encode(code), nil)
		if err == nil {
			t.Fatal("expected syntax error")
		}
		if !strings.HasPrefix(err.Error(), "invalid typecode syntax:") {
			t.Fatal("expected syntax error, got:", err)
		}
	}
	// Outlet proxy does not satisfy Pdu interface, fall back to Pdu proxy
	c, _ = jsonrpc.CodeFor[outlet.Outlet]()
	code = c.String()
	_, err = object.As[pdu.Pdu](encode(code), nil)
	if err != nil {
		t.Fatal(err)
	}
	// incompatible Outlet version, fall back to EDevice proxy
	code = "pdumodel.Outlet:1.0.0"
	_, err = object.As[edevice.EDevice](encode(code), nil)
	if err != nil {
		t.Fatal(err)
	}
	// unknown code, fall back to Pdu proxy
	code = "unknown.RocketShip:1.0.0"
	_, err = object.As[pdu.Pdu](encode(code), nil)
	if err != nil {
		t.Fatal(err)
	}
}
