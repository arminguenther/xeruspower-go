// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package cardreader

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40010/idl"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding/object"
)

func init() {
	object.Register(NewCardReader)
}

type _CardReader struct {
	idl.Object
}

// NewCardReader returns a new CardReader interface for the object at given RID.
func NewCardReader(rid string, caller idl.Caller) CardReader {
	return &_CardReader{idl.NewObject(rid, caller)}
}

func (c *_CardReader) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartcard.CardReader",
		Major: 1, Submajor: 0, Minor: 4,
	}
}

func (c *_CardReader) GetMetaData(ctx context.Context) (MetaData, error) {
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

func (c *_CardReader) GetCardInformation(ctx context.Context) (int32, CardInformation, error) {
	var ret int32
	var out0 CardInformation
	val, err := c.Caller().Call(ctx, c.RID(), "getCardInformation", nil)
	if err != nil {
		return ret, out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, out0, err
	}
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("cardInfo", res)
	if err != nil {
		return ret, out0, err
	}
	err = out0.Decode(res["cardInfo"], c.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}
