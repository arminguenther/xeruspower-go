// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package event

import (
	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() Event { return &_Event{} })
}

type _Event struct {
	idl.ValueObject
	source idl.Object
}

func (e *_Event) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "idl.Event",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (e *_Event) Source() idl.Object {
	return e.source
}

func (e *_Event) isEvent() {}
