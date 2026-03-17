// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package rawconfiguration

import (
	"context"
	"time"

	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/object"
)

func init() {
	object.Register(NewRawConfiguration)
}

type _RawConfiguration struct {
	idl.Object
}

// NewRawConfiguration returns a new RawConfiguration interface for the object at given RID.
func NewRawConfiguration(rid string, caller idl.Caller) RawConfiguration {
	return &_RawConfiguration{idl.NewObject(rid, caller)}
}

func (r *_RawConfiguration) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "rawcfg.RawConfiguration",
		Major: 1, Submajor: 0, Minor: 1,
	}
}

func (r *_RawConfiguration) GetStatus(ctx context.Context) (Status, time.Time, error) {
	var out0 Status
	var out1 time.Time
	val, err := r.Caller().Call(ctx, r.RID(), "getStatus", nil)
	if err != nil {
		return out0, out1, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return out0, out1, err
	}
	err = encoding.In("status", res)
	if err != nil {
		return out0, out1, err
	}
	out0, err = encoding.AsEnum[Status](res["status"])
	if err != nil {
		return out0, out1, err
	}
	err = encoding.In("timeStamp", res)
	if err != nil {
		return out0, out1, err
	}
	out1, err = encoding.AsTime(res["timeStamp"])
	if err != nil {
		return out0, out1, err
	}
	return out0, out1, nil
}
