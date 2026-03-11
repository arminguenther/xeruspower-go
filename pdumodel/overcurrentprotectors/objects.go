// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package overcurrentprotectors

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40411/idl"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding/object"
)

func init() {
	object.Register(NewOverCurrentProtectors)
}

type _OverCurrentProtectors struct {
	idl.Object
}

// NewOverCurrentProtectors returns a new OverCurrentProtectors interface for the object at given RID.
func NewOverCurrentProtectors(rid string, caller idl.Caller) OverCurrentProtectors {
	return &_OverCurrentProtectors{idl.NewObject(rid, caller)}
}

func (o *_OverCurrentProtectors) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.OverCurrentProtectors",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (o *_OverCurrentProtectors) GetInfo(ctx context.Context) (Info, error) {
	var ret Info
	val, err := o.Caller().Call(ctx, o.RID(), "getInfo", nil)
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
	err = ret.Decode(res["_ret_"], o.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}
