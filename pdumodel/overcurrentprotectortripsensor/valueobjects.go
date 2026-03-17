// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package overcurrentprotectortripsensor

import (
	"github.com/arminguenther/xeruspower-go/v40032/idl"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40032/pdumodel/outlet"
	"github.com/arminguenther/xeruspower-go/v40032/sensors/statesensor"
)

func init() {
	valobj.Register(func() TripStateChangedEvent { return &_TripStateChangedEvent{} })
}

type _TripStateChangedEvent struct {
	statesensor.StateChangedEvent
	tripCause       outlet.Outlet
	tripInformation TripEventInformation
}

func (t *_TripStateChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.OverCurrentProtectorTripSensor_1_1_11.TripStateChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (t *_TripStateChangedEvent) TripCause() outlet.Outlet {
	return t.tripCause
}

func (t *_TripStateChangedEvent) TripInformation() TripEventInformation {
	return t.tripInformation
}

func (t *_TripStateChangedEvent) isTripStateChangedEvent() {}
