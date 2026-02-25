// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package peripheralg2production

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
)

func init() {
	object.Register(NewG2Production)
}

type _G2Production struct {
	idl.Object
}

// NewG2Production returns a new G2Production interface for the object at given RID.
func NewG2Production(rid string, caller idl.Caller) G2Production {
	return &_G2Production{idl.NewObject(rid, caller)}
}

func (g *_G2Production) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.G2Production",
		Major: 4, Submajor: 0, Minor: 0,
	}
}

func (g *_G2Production) ProgramSerialNumber(ctx context.Context, in0 string, in1 string) (int32, error) {
	var ret int32
	val, err := g.Caller().Call(ctx, g.RID(), "programSerialNumber", map[string]any{
		"romcode": in0,
		"serial":  in1,
	})
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
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}
