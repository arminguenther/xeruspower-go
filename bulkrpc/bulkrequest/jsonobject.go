// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package bulkrequest

import (
	"github.com/arminguenther/xeruspower-go/v40412/idl"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding"
)

// JsonObject is a plain JSON object used in JSON-RPC.
//
// Requests:
// All members should be ready for [encoding/json.Marshal].
//
// Responses:
// All members are Go values that [encoding/json.Decoder] stores
// into interface values after [encoding/json.Decoder.UseNumber] was called:
//
//	JSON        Go
//	null        nil
//	true/false  bool
//	string      string
//	number      json.Number
//	array       []any
//	object      map[string]any
type JsonObject map[string]any

// Encode returns the JsonObject ready for [encoding/json.Marshal].
func (j *JsonObject) Encode() JsonObject {
	if *j == nil {
		return make(JsonObject)
	}
	return *j
}

// Decode sets j to a JSON object stored in an interface value by [encoding/json.Decoder].
func (j *JsonObject) Decode(v any, _ idl.Caller) error {
	obj, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	*j = obj
	return nil
}
