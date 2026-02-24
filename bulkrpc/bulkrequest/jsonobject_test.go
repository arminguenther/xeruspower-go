// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package bulkrequest

import (
	"encoding/json"
	"testing"
)

func TestJsonObject_Encode(t *testing.T) {
	var obj JsonObject
	bytes, _ := json.Marshal(obj.Encode())
	if string(bytes) != "{}" {
		t.Fatalf("expected {}, got %v", string(bytes))
	}
	obj = JsonObject{"foo": "bar"}
	bytes, _ = json.Marshal(obj.Encode())
	enc := `{"foo":"bar"}`
	if string(bytes) != enc {
		t.Fatalf("expected %v, got %v", enc, string(bytes))
	}
}

func TestJsonObject_Decode(t *testing.T) {
	var obj JsonObject
	v := map[string]any{"enabled": true}
	err := obj.Decode(v, nil)
	if err != nil {
		t.Fatal(err)
	}
	enabled, ok := obj["enabled"]
	if !ok {
		t.Fatal("missing key")
	}
	if enabled != true {
		t.Fatalf("wrong value")
	}
	v = nil
	err = obj.Decode(v, nil)
	if err != nil {
		t.Fatal(err)
	}
	if obj.Decode(nil, nil) == nil {
		t.Fatal("expected error")
	}
}
