// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package controller

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding/object"
)

func init() {
	object.Register(NewController)
}

type _Controller struct {
	idl.Object
}

// NewController returns a new Controller interface for the object at given RID.
func NewController(rid string, caller idl.Caller) Controller {
	return &_Controller{idl.NewObject(rid, caller)}
}

func (c *_Controller) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Controller",
		Major: 5, Submajor: 0, Minor: 1,
	}
}

func (c *_Controller) GetCommunicationStatus(ctx context.Context) (Status, error) {
	var ret Status
	val, err := c.Caller().Call(ctx, c.RID(), "getCommunicationStatus", nil)
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
	ret, err = encoding.AsEnum[Status](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (c *_Controller) GetStatistics(ctx context.Context) (CtrlStatistic, error) {
	var ret CtrlStatistic
	val, err := c.Caller().Call(ctx, c.RID(), "getStatistics", nil)
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

func (c *_Controller) GetMetaData(ctx context.Context) (MetaData, error) {
	var ret MetaData
	val, err := c.Caller().Call(ctx, c.RID(), "getMetaData", nil)
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

func (c *_Controller) Reset(ctx context.Context) (int32, error) {
	var ret int32
	val, err := c.Caller().Call(ctx, c.RID(), "reset", nil)
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
