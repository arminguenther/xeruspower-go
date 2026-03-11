// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package inlets

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40410/idl"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding/object"
)

func init() {
	object.Register(NewInlets)
}

type _Inlets struct {
	idl.Object
}

// NewInlets returns a new Inlets interface for the object at given RID.
func NewInlets(rid string, caller idl.Caller) Inlets {
	return &_Inlets{idl.NewObject(rid, caller)}
}

func (i *_Inlets) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Inlets",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_Inlets) GetInfo(ctx context.Context) (Info, error) {
	var ret Info
	val, err := i.Caller().Call(ctx, i.RID(), "getInfo", nil)
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	err = ret.Decode(res["_ret_"], i.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}
