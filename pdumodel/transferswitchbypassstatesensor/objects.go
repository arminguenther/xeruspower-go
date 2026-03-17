// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package transferswitchbypassstatesensor

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40300/sensors/statesensor"
)

func init() {
	object.Register(NewTransferSwitchBypassStateSensor)
}

type _TransferSwitchBypassStateSensor struct {
	statesensor.StateSensor
}

// NewTransferSwitchBypassStateSensor returns a new TransferSwitchBypassStateSensor interface for the object at given RID.
func NewTransferSwitchBypassStateSensor(rid string, caller idl.Caller) TransferSwitchBypassStateSensor {
	return &_TransferSwitchBypassStateSensor{statesensor.NewStateSensor(rid, caller)}
}

func (t *_TransferSwitchBypassStateSensor) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.TransferSwitchBypassStateSensor",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (t *_TransferSwitchBypassStateSensor) AcknowledgeFault(ctx context.Context) error {
	_, err := t.Caller().Call(ctx, t.RID(), "acknowledgeFault", nil)
	return err
}
