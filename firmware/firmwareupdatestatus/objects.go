// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package firmwareupdatestatus

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40410/idl"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding/object"
)

func init() {
	object.Register(NewFirmwareUpdateStatus)
}

type _FirmwareUpdateStatus struct {
	idl.Object
}

// NewFirmwareUpdateStatus returns a new FirmwareUpdateStatus interface for the object at given RID.
func NewFirmwareUpdateStatus(rid string, caller idl.Caller) FirmwareUpdateStatus {
	return &_FirmwareUpdateStatus{idl.NewObject(rid, caller)}
}

func (f *_FirmwareUpdateStatus) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "firmware.FirmwareUpdateStatus",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (f *_FirmwareUpdateStatus) GetStatus(ctx context.Context) (UpdateStatus, error) {
	var ret UpdateStatus
	val, err := f.Caller().Call(ctx, f.RID(), "getStatus", nil)
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
	err = ret.Decode(res["_ret_"], f.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}
