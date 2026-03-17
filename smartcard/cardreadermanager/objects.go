// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package cardreadermanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40010/idl"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40010/smartcard/cardreader"
)

func init() {
	object.Register(NewCardReaderManager)
}

type _CardReaderManager struct {
	idl.Object
}

// NewCardReaderManager returns a new CardReaderManager interface for the object at given RID.
func NewCardReaderManager(rid string, caller idl.Caller) CardReaderManager {
	return &_CardReaderManager{idl.NewObject(rid, caller)}
}

func (c *_CardReaderManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "smartcard.CardReaderManager",
		Major: 1, Submajor: 0, Minor: 5,
	}
}

func (c *_CardReaderManager) GetCardReaders(ctx context.Context) ([]cardreader.CardReader, error) {
	var ret []cardreader.CardReader
	val, err := c.Caller().Call(ctx, c.RID(), "getCardReaders", nil)
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
	var s0 []any
	s0, err = encoding.Is[[]any](res["_ret_"])
	if err != nil {
		return ret, err
	}
	ret = make([]cardreader.CardReader, 0, len(s0))
	for _, a0 := range s0 {
		var e0 cardreader.CardReader
		e0, err = object.As[cardreader.CardReader](a0, c.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (c *_CardReaderManager) GetCardReaderById(ctx context.Context, in0 string) (cardreader.CardReader, error) {
	var ret cardreader.CardReader
	val, err := c.Caller().Call(ctx, c.RID(), "getCardReaderById", map[string]any{
		"readerId": in0,
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
	ret, err = object.As[cardreader.CardReader](res["_ret_"], c.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (c *_CardReaderManager) SetCardReaderSettings(ctx context.Context, in0 string, in1 CardReaderSettings) (int32, error) {
	var ret int32
	val, err := c.Caller().Call(ctx, c.RID(), "setCardReaderSettings", map[string]any{
		"position": in0,
		"setting":  in1.Encode(),
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

func (c *_CardReaderManager) GetAllCardReaderSettings(ctx context.Context) (map[string]CardReaderSettings, error) {
	var ret map[string]CardReaderSettings
	val, err := c.Caller().Call(ctx, c.RID(), "getAllCardReaderSettings", nil)
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
	i0, e0, l0 := encoding.AsMapItems(res["_ret_"])
	ret = make(map[string]CardReaderSettings, l0)
	for a0, b0 := range i0 {
		var k0 string
		k0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, err
		}
		var v0 CardReaderSettings
		err = v0.Decode(b0, c.Caller())
		if err != nil {
			return ret, err
		}
		ret[k0] = v0
	}
	err = e0()
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (c *_CardReaderManager) GetSupportedCardFormats(ctx context.Context) ([]string, error) {
	var ret []string
	val, err := c.Caller().Call(ctx, c.RID(), "getSupportedCardFormats", nil)
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
	var s0 []any
	s0, err = encoding.Is[[]any](res["_ret_"])
	if err != nil {
		return ret, err
	}
	ret = make([]string, 0, len(s0))
	for _, a0 := range s0 {
		var e0 string
		e0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}
