// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package corectrl

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40510/tfw/scannerctrl"
)

func init() {
	object.Register(NewCoreCtrl)
}

type _CoreCtrl struct {
	idl.Object
}

// NewCoreCtrl returns a new CoreCtrl interface for the object at given RID.
func NewCoreCtrl(rid string, caller idl.Caller) CoreCtrl {
	return &_CoreCtrl{idl.NewObject(rid, caller)}
}

func (c *_CoreCtrl) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "tfw.CoreCtrl",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_CoreCtrl) GetScanner(ctx context.Context, in0 string) (scannerctrl.ScannerCtrl, error) {
	var ret scannerctrl.ScannerCtrl
	val, err := c.Caller().Call(ctx, c.RID(), "getScanner", map[string]any{
		"name": in0,
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
	ret, err = object.As[scannerctrl.ScannerCtrl](res["_ret_"], c.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (c *_CoreCtrl) StartAllScanners(ctx context.Context) error {
	_, err := c.Caller().Call(ctx, c.RID(), "startAllScanners", nil)
	return err
}

func (c *_CoreCtrl) StopAllScanners(ctx context.Context) error {
	_, err := c.Caller().Call(ctx, c.RID(), "stopAllScanners", nil)
	return err
}

func (c *_CoreCtrl) StartAllBackWorkers(ctx context.Context) error {
	_, err := c.Caller().Call(ctx, c.RID(), "startAllBackWorkers", nil)
	return err
}

func (c *_CoreCtrl) StopAllBackWorkers(ctx context.Context) error {
	_, err := c.Caller().Call(ctx, c.RID(), "stopAllBackWorkers", nil)
	return err
}
