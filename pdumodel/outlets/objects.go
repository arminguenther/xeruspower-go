// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package outlets

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/object"
)

func init() {
	object.Register(NewOutlets)
}

type _Outlets struct {
	idl.Object
}

// NewOutlets returns a new Outlets interface for the object at given RID.
func NewOutlets(rid string, caller idl.Caller) Outlets {
	return &_Outlets{idl.NewObject(rid, caller)}
}

func (o *_Outlets) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Outlets",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (o *_Outlets) GetInfo(ctx context.Context) (Info, error) {
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
