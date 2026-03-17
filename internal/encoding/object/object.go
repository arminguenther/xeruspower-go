// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

// Package object handles encoding/decoding of [idl.Object] types.
package object

import (
	"reflect"

	"github.com/arminguenther/xeruspower-go/v40200/idl"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding"
)

func init() {
	Register(idl.NewObject)
}

// ToMap returns a map description of an [idl.Object] that correctly encodes
// as JSON object or JSON null when passed to [encoding/json.Marshal].
func ToMap(object idl.Object) map[string]any {
	if object == nil {
		return nil
	}
	return map[string]any{"type": object.TypeCode().String(), "rid": object.RID()}
}

// CodeRegistry stores functions created by [Register].
// Key is the string:
//
//	fn().TypeCode().StringNoVersions()
var CodeRegistry = make(map[string]func(string, idl.Caller) idl.Object)

// TypeRegistry stores functions created by [Register].
// Key is the type parameter passed to [Register].
var TypeRegistry = make(map[reflect.Type]func(string, idl.Caller) idl.Object)

// Register registers a function to create a new [idl.Object] of type O.
// The function gets wrapped and stored in [CodeRegistry] and [TypeRegistry].
func Register[O idl.Object](newO func(string, idl.Caller) O) {
	code := newO("", nil).TypeCode()
	fn := func(rid string, caller idl.Caller) idl.Object { return newO(rid, caller) }
	CodeRegistry[code.StringNoVersions()] = fn
	TypeRegistry[reflect.TypeFor[O]()] = fn
}

// As asserts that v is an encoded [idl.Object] of type O.
// O needs to be a type already registered with [Register].
func As[O idl.Object](v any, caller idl.Caller) (O, error) {
	var obj O
	if v == nil {
		return obj, nil
	}
	json, err := encoding.Is[map[string]any](v)
	if err != nil {
		return obj, err
	}
	rid, err := encoding.InAndIs[string]("rid", json)
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
	newFn, ok := CodeRegistry[code.StringNoVersions()]
	if ok {
		obj, ok = newFn(rid, caller).(O)
		if ok && code.IsCompatible(obj.TypeCode()) {
			return obj, nil
		}
	}
	newFn = TypeRegistry[reflect.TypeFor[O]()]
	return newFn(rid, caller).(O), nil
}
