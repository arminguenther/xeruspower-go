// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package ade

import (
	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding"
)

func (m *MetaData) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["adeType"] = m.AdeType
	j0["channels"] = m.Channels
	j0["currentDivider"] = m.CurrentDivider
	j0["voltageDivider"] = m.VoltageDivider
	j0["energyDivider"] = m.EnergyDivider
	return j0
}

func (m *MetaData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("adeType", j0)
	if err != nil {
		return err
	}
	m.AdeType, err = encoding.Is[string](j0["adeType"])
	if err != nil {
		return err
	}
	err = encoding.In("channels", j0)
	if err != nil {
		return err
	}
	m.Channels, err = encoding.AsInt32(j0["channels"])
	if err != nil {
		return err
	}
	err = encoding.In("currentDivider", j0)
	if err != nil {
		return err
	}
	m.CurrentDivider, err = encoding.AsFloat64(j0["currentDivider"])
	if err != nil {
		return err
	}
	err = encoding.In("voltageDivider", j0)
	if err != nil {
		return err
	}
	m.VoltageDivider, err = encoding.AsFloat64(j0["voltageDivider"])
	if err != nil {
		return err
	}
	err = encoding.In("energyDivider", j0)
	if err != nil {
		return err
	}
	m.EnergyDivider, err = encoding.AsFloat64(j0["energyDivider"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Sample) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["vrms"] = s.Vrms
	j0["irms"] = s.Irms
	j0["watt"] = s.Watt
	j0["va"] = s.Va
	j0["wh"] = s.Wh
	j0["vah"] = s.Vah
	return j0
}

func (s *Sample) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("vrms", j0)
	if err != nil {
		return err
	}
	s.Vrms, err = encoding.AsInt64(j0["vrms"])
	if err != nil {
		return err
	}
	err = encoding.In("irms", j0)
	if err != nil {
		return err
	}
	s.Irms, err = encoding.AsInt64(j0["irms"])
	if err != nil {
		return err
	}
	err = encoding.In("watt", j0)
	if err != nil {
		return err
	}
	s.Watt, err = encoding.AsInt64(j0["watt"])
	if err != nil {
		return err
	}
	err = encoding.In("va", j0)
	if err != nil {
		return err
	}
	s.Va, err = encoding.AsInt64(j0["va"])
	if err != nil {
		return err
	}
	err = encoding.In("wh", j0)
	if err != nil {
		return err
	}
	s.Wh, err = encoding.AsInt64(j0["wh"])
	if err != nil {
		return err
	}
	err = encoding.In("vah", j0)
	if err != nil {
		return err
	}
	s.Vah, err = encoding.AsInt64(j0["vah"])
	if err != nil {
		return err
	}
	return nil
}
