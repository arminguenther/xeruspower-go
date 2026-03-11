// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package idl

// ValueObject is an object that may have elements.
type ValueObject interface {
	TypeCode() TypeCode
	Encode() map[string]any
	Decode(value map[string]any, caller Caller) error
	isValueObject()
}

// NewValueObject creates a new plain ValueObject.
func NewValueObject() ValueObject {
	return &_ValueObject{}
}

type _ValueObject struct{}

func (v *_ValueObject) TypeCode() TypeCode {
	return TypeCode{
		Name:  "idl.ValueObject",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (v *_ValueObject) Encode() map[string]any {
	return make(map[string]any)
}

func (v *_ValueObject) Decode(_ map[string]any, _ Caller) error {
	return nil
}

func (v *_ValueObject) isValueObject() {}
