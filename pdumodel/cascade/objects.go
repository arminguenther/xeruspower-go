// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package cascade

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
)

func init() {
	object.Register(NewCascade)
}

type _Cascade struct {
	idl.Object
}

// NewCascade returns a new Cascade interface for the object at given RID.
func NewCascade(rid string, caller idl.Caller) Cascade {
	return &_Cascade{idl.NewObject(rid, caller)}
}

func (c *_Cascade) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Cascade",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_Cascade) GetInfo(ctx context.Context) (Info, error) {
	var ret Info
	val, err := c.Caller().Call(ctx, c.RID(), "getInfo", nil)
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
	err = ret.Decode(res["_ret_"], c.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}
