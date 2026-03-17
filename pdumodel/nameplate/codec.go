// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package nameplate

import (
	"github.com/arminguenther/xeruspower-go/v40200/idl"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding"
)

func (r *Rating) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["current"] = r.Current
	j0["decimalCurrent"] = r.DecimalCurrent
	j0["minVoltage"] = r.MinVoltage
	j0["maxVoltage"] = r.MaxVoltage
	return j0
}

func (r *Rating) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("current", j0)
	if err != nil {
		return err
	}
	r.Current, err = encoding.AsInt32(j0["current"])
	if err != nil {
		return err
	}
	err = encoding.In("decimalCurrent", j0)
	if err != nil {
		return err
	}
	r.DecimalCurrent, err = encoding.AsFloat32(j0["decimalCurrent"])
	if err != nil {
		return err
	}
	err = encoding.In("minVoltage", j0)
	if err != nil {
		return err
	}
	r.MinVoltage, err = encoding.AsInt32(j0["minVoltage"])
	if err != nil {
		return err
	}
	err = encoding.In("maxVoltage", j0)
	if err != nil {
		return err
	}
	r.MaxVoltage, err = encoding.AsInt32(j0["maxVoltage"])
	if err != nil {
		return err
	}
	return nil
}

func (n *Nameplate) Encode() map[string]any {
	j0 := make(map[string]any, 7)
	j0["manufacturer"] = n.Manufacturer
	j0["brand"] = n.Brand
	j0["model"] = n.Model
	j0["partNumber"] = n.PartNumber
	j0["serialNumber"] = n.SerialNumber
	j1 := make(map[string]any, 4)
	j1["voltage"] = n.Rating.Voltage
	j1["current"] = n.Rating.Current
	j1["frequency"] = n.Rating.Frequency
	j1["power"] = n.Rating.Power
	j0["rating"] = j1
	j0["imageFileURL"] = n.ImageFileURL
	return j0
}

func (n *Nameplate) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("manufacturer", j0)
	if err != nil {
		return err
	}
	n.Manufacturer, err = encoding.Is[string](j0["manufacturer"])
	if err != nil {
		return err
	}
	err = encoding.In("brand", j0)
	if err != nil {
		return err
	}
	n.Brand, err = encoding.Is[string](j0["brand"])
	if err != nil {
		return err
	}
	err = encoding.In("model", j0)
	if err != nil {
		return err
	}
	n.Model, err = encoding.Is[string](j0["model"])
	if err != nil {
		return err
	}
	err = encoding.In("partNumber", j0)
	if err != nil {
		return err
	}
	n.PartNumber, err = encoding.Is[string](j0["partNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("serialNumber", j0)
	if err != nil {
		return err
	}
	n.SerialNumber, err = encoding.Is[string](j0["serialNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("rating", j0)
	if err != nil {
		return err
	}
	j1, err := encoding.Is[map[string]any](j0["rating"])
	if err != nil {
		return err
	}
	err = encoding.In("voltage", j1)
	if err != nil {
		return err
	}
	n.Rating.Voltage, err = encoding.Is[string](j1["voltage"])
	if err != nil {
		return err
	}
	err = encoding.In("current", j1)
	if err != nil {
		return err
	}
	n.Rating.Current, err = encoding.Is[string](j1["current"])
	if err != nil {
		return err
	}
	err = encoding.In("frequency", j1)
	if err != nil {
		return err
	}
	n.Rating.Frequency, err = encoding.Is[string](j1["frequency"])
	if err != nil {
		return err
	}
	err = encoding.In("power", j1)
	if err != nil {
		return err
	}
	n.Rating.Power, err = encoding.Is[string](j1["power"])
	if err != nil {
		return err
	}
	err = encoding.In("imageFileURL", j0)
	if err != nil {
		return err
	}
	n.ImageFileURL, err = encoding.Is[string](j0["imageFileURL"])
	if err != nil {
		return err
	}
	return nil
}
