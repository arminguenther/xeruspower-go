// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package sx

import (
	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/object"
)

func init() {
	object.Register(NewSx)
}

type _Sx struct {
	idl.Object
}

// NewSx returns a new Sx interface for the object at given RID.
func NewSx(rid string, caller idl.Caller) Sx {
	return &_Sx{idl.NewObject(rid, caller)}
}

func (s *_Sx) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sx.Sx",
		Major: 1, Submajor: 0, Minor: 0,
	}
}
