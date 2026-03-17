// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package idl

import "context"

// A Caller calls object methods.
type Caller interface {
	// Call calls a method on the object with resource ID rid.
	// Method parameters are passed by name.
	Call(ctx context.Context, rid string, method string, params map[string]any) (result any, err error)
}

// Object is an object that may have methods.
type Object interface {
	Caller() Caller     // Caller calls the object methods.
	RID() string        // RID is the resource identifier.
	TypeCode() TypeCode // TypeCode describes the object type.
}

// NewObject creates a new plain Object.
func NewObject(rid string, caller Caller) Object {
	return &_Object{rid: rid, caller: caller}
}

type _Object struct {
	rid    string
	caller Caller
}

func (o *_Object) TypeCode() TypeCode {
	return TypeCode{
		Name:  "idl.Object",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (o *_Object) RID() string {
	return o.rid
}

func (o *_Object) Caller() Caller {
	return o.caller
}
