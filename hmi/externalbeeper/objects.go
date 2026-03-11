// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package externalbeeper

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/object"
)

func init() {
	object.Register(NewExternalBeeper)
}

type _ExternalBeeper struct {
	idl.Object
}

// NewExternalBeeper returns a new ExternalBeeper interface for the object at given RID.
func NewExternalBeeper(rid string, caller idl.Caller) ExternalBeeper {
	return &_ExternalBeeper{idl.NewObject(rid, caller)}
}

func (e *_ExternalBeeper) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "hmi.ExternalBeeper",
		Major: 1, Submajor: 0, Minor: 1,
	}
}

func (e *_ExternalBeeper) GetState(ctx context.Context) (State, error) {
	var ret State
	val, err := e.Caller().Call(ctx, e.RID(), "getState", nil)
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
	ret, err = encoding.AsEnum[State](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (e *_ExternalBeeper) Alarm(ctx context.Context) error {
	_, err := e.Caller().Call(ctx, e.RID(), "alarm", nil)
	return err
}

func (e *_ExternalBeeper) On(ctx context.Context) error {
	_, err := e.Caller().Call(ctx, e.RID(), "on", nil)
	return err
}

func (e *_ExternalBeeper) Off(ctx context.Context) error {
	_, err := e.Caller().Call(ctx, e.RID(), "off", nil)
	return err
}
