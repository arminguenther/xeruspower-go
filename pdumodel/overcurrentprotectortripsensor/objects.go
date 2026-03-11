// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package overcurrentprotectortripsensor

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40413/pdumodel/outlet"
	"github.com/arminguenther/xeruspower-go/v40413/pdumodel/waveform"
	"github.com/arminguenther/xeruspower-go/v40413/sensors/statesensor"
)

func init() {
	object.Register(NewOverCurrentProtectorTripSensor)
}

type _OverCurrentProtectorTripSensor struct {
	statesensor.StateSensor
}

// NewOverCurrentProtectorTripSensor returns a new OverCurrentProtectorTripSensor interface for the object at given RID.
func NewOverCurrentProtectorTripSensor(rid string, caller idl.Caller) OverCurrentProtectorTripSensor {
	return &_OverCurrentProtectorTripSensor{statesensor.NewStateSensor(rid, caller)}
}

func (o *_OverCurrentProtectorTripSensor) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.OverCurrentProtectorTripSensor",
		Major: 1, Submajor: 1, Minor: 14,
	}
}

func (o *_OverCurrentProtectorTripSensor) GetTripCause(ctx context.Context) (outlet.Outlet, error) {
	var ret outlet.Outlet
	val, err := o.Caller().Call(ctx, o.RID(), "getTripCause", nil)
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
	ret, err = object.As[outlet.Outlet](res["_ret_"], o.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (o *_OverCurrentProtectorTripSensor) GetTripEventInformation(ctx context.Context) (TripEventInformation, error) {
	var ret TripEventInformation
	val, err := o.Caller().Call(ctx, o.RID(), "getTripEventInformation", nil)
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
	err = ret.Decode(res["_ret_"], o.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (o *_OverCurrentProtectorTripSensor) GetTripWaveform(ctx context.Context) (waveform.Waveform, error) {
	var ret waveform.Waveform
	val, err := o.Caller().Call(ctx, o.RID(), "getTripWaveform", nil)
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
	err = ret.Decode(res["_ret_"], o.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}
