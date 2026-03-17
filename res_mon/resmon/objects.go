// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package resmon

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding/object"
)

func init() {
	object.Register(NewResMon)
}

type _ResMon struct {
	idl.Object
}

// NewResMon returns a new ResMon interface for the object at given RID.
func NewResMon(rid string, caller idl.Caller) ResMon {
	return &_ResMon{idl.NewObject(rid, caller)}
}

func (r *_ResMon) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "res_mon.ResMon",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_ResMon) GetDataEntries(ctx context.Context) ([]Entry, error) {
	var out0 []Entry
	val, err := r.Caller().Call(ctx, r.RID(), "getDataEntries", nil)
	if err != nil {
		return out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return out0, err
	}
	err = encoding.In("entries", res)
	if err != nil {
		return out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["entries"])
	if err != nil {
		return out0, err
	}
	out0 = make([]Entry, 0, len(s0))
	for _, a0 := range s0 {
		var e0 Entry
		err = e0.Decode(a0, r.Caller())
		if err != nil {
			return out0, err
		}
		out0 = append(out0, e0)
	}
	return out0, nil
}
