// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package scannerctrl

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/object"
)

func init() {
	object.Register(NewScannerCtrl)
}

type _ScannerCtrl struct {
	idl.Object
}

// NewScannerCtrl returns a new ScannerCtrl interface for the object at given RID.
func NewScannerCtrl(rid string, caller idl.Caller) ScannerCtrl {
	return &_ScannerCtrl{idl.NewObject(rid, caller)}
}

func (s *_ScannerCtrl) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "tfw.ScannerCtrl",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_ScannerCtrl) GetScanInterval(ctx context.Context) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "getScanInterval", nil)
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

func (s *_ScannerCtrl) SetScanInterval(ctx context.Context, in0 int32) error {
	_, err := s.Caller().Call(ctx, s.RID(), "setScanInterval", map[string]any{
		"interval": in0,
	})
	return err
}
