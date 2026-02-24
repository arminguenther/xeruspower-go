// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package modbuscfg

import (
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
)

func init() {
	object.Register(NewModbusCfg)
}

type _ModbusCfg struct {
	idl.Object
}

// NewModbusCfg returns a new ModbusCfg interface for the object at given RID.
func NewModbusCfg(rid string, caller idl.Caller) ModbusCfg {
	return &_ModbusCfg{idl.NewObject(rid, caller)}
}

func (m *_ModbusCfg) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.ModbusCfg",
		Major: 1, Submajor: 0, Minor: 0,
	}
}
