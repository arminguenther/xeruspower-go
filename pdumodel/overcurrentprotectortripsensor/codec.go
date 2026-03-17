// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package overcurrentprotectortripsensor

import (
	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40220/pdumodel/outlet"
	"github.com/arminguenther/xeruspower-go/v40220/sensors/statesensor"
)

func (t *TripEventInformation) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["cause"] = object.ToMap(t.Cause)
	j0["timestamp"] = t.Timestamp.Unix()
	j0["current"] = t.Current
	return j0
}

func (t *TripEventInformation) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("cause", j0)
	if err != nil {
		return err
	}
	t.Cause, err = object.As[outlet.Outlet](j0["cause"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("timestamp", j0)
	if err != nil {
		return err
	}
	t.Timestamp, err = encoding.AsTime(j0["timestamp"])
	if err != nil {
		return err
	}
	err = encoding.In("current", j0)
	if err != nil {
		return err
	}
	t.Current, err = encoding.AsFloat64(j0["current"])
	if err != nil {
		return err
	}
	return nil
}

func (t *_TripStateChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	t.StateChangedEvent = valobj.For[statesensor.StateChangedEvent]()
	err := t.StateChangedEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("tripCause", value)
	if err != nil {
		return err
	}
	t.tripCause, err = object.As[outlet.Outlet](value["tripCause"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("tripInformation", value)
	if err != nil {
		return err
	}
	err = t.tripInformation.Decode(value["tripInformation"], caller)
	if err != nil {
		return err
	}
	return nil
}
