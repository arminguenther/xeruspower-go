// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

// Package valobj handles encoding/decoding of [idl.ValueObject] types.
package valobj

import (
	"reflect"

	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding"
)

func init() {
	Register(idl.NewValueObject)
}

// ToMap returns a map description of an [idl.ValueObject] that correctly encodes
// as JSON object or JSON null when passed to [encoding/json.Marshal].
func ToMap(object idl.ValueObject) map[string]any {
	if object == nil {
		return nil
	}
	return map[string]any{"type": object.TypeCode().String(), "value": object.Encode()}
}

// CodeRegistry stores functions created by [Register].
// Key is the string:
//
//	fn().TypeCode().StringNoVersions()
var CodeRegistry = make(map[string]func() idl.ValueObject)

// TypeRegistry stores functions created by [Register].
// Key is the type parameter passed to [Register].
var TypeRegistry = make(map[reflect.Type]func() idl.ValueObject)

// Register registers a function to create a new [idl.ValueObject] of type V.
// The function gets wrapped and stored in [CodeRegistry] and [TypeRegistry].
func Register[V idl.ValueObject](newV func() V) {
	code := newV().TypeCode()
	fn := func() idl.ValueObject { return newV() }
	CodeRegistry[code.StringNoVersions()] = fn
	TypeRegistry[reflect.TypeFor[V]()] = fn
}

// As asserts that v is an encoded [idl.ValueObject] of type V.
// V needs to be a type already registered with [Register].
func As[V idl.ValueObject](v any, caller idl.Caller) (V, error) {
	var obj V
	if v == nil {
		return obj, nil
	}
	json, err := encoding.Is[map[string]any](v)
	if err != nil {
		return obj, err
	}
	err = encoding.In("type", json)
	if err != nil {
		return obj, err
	}
	code, err := encoding.AsTypeCode(json["type"])
	if err != nil {
		return obj, err
	}
	value, err := encoding.InAndIs[map[string]any]("value", json)
	if err != nil {
		return obj, err
	}
	newFn, ok := CodeRegistry[code.StringNoVersions()]
	if ok {
		obj, ok = newFn().(V)
		if ok && code.IsCompatible(obj.TypeCode()) {
			err = obj.Decode(value, caller)
			return obj, err
		}
	}
	obj = For[V]()
	err = obj.Decode(value, caller)
	return obj, err
}

// For returns a new ValueObject of type V.
// V needs to be a type already registered with [Register].
func For[V idl.ValueObject]() V {
	return TypeRegistry[reflect.TypeFor[V]()]().(V)
}
