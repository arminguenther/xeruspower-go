// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package powerbus

import (
	"github.com/arminguenther/xeruspower-go/v40412/idl"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40412/sensors/numericsensor"
)

func (s *Sensors) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["current"] = object.ToMap(s.Current)
	j0["voltage"] = object.ToMap(s.Voltage)
	j0["activePower"] = object.ToMap(s.ActivePower)
	j0["activeEnergy"] = object.ToMap(s.ActiveEnergy)
	return j0
}

func (s *Sensors) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("current", j0)
	if err != nil {
		return err
	}
	s.Current, err = object.As[numericsensor.NumericSensor](j0["current"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("voltage", j0)
	if err != nil {
		return err
	}
	s.Voltage, err = object.As[numericsensor.NumericSensor](j0["voltage"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("activePower", j0)
	if err != nil {
		return err
	}
	s.ActivePower, err = object.As[numericsensor.NumericSensor](j0["activePower"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("activeEnergy", j0)
	if err != nil {
		return err
	}
	s.ActiveEnergy, err = object.As[numericsensor.NumericSensor](j0["activeEnergy"], caller)
	if err != nil {
		return err
	}
	return nil
}
