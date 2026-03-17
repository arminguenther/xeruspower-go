// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package fitness

import (
	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding"
)

func (d *DataEntry) Encode() map[string]any {
	j0 := make(map[string]any, 7)
	j0["id"] = d.Id
	j0["value"] = d.Value
	j0["maxValue"] = d.MaxValue
	j0["worstValue"] = d.WorstValue
	j0["thresholdValue"] = d.ThresholdValue
	j0["rawValue"] = d.RawValue
	j0["flags"] = d.Flags
	return j0
}

func (d *DataEntry) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("id", j0)
	if err != nil {
		return err
	}
	d.Id, err = encoding.Is[string](j0["id"])
	if err != nil {
		return err
	}
	err = encoding.In("value", j0)
	if err != nil {
		return err
	}
	d.Value, err = encoding.AsInt32(j0["value"])
	if err != nil {
		return err
	}
	err = encoding.In("maxValue", j0)
	if err != nil {
		return err
	}
	d.MaxValue, err = encoding.AsInt32(j0["maxValue"])
	if err != nil {
		return err
	}
	err = encoding.In("worstValue", j0)
	if err != nil {
		return err
	}
	d.WorstValue, err = encoding.AsInt32(j0["worstValue"])
	if err != nil {
		return err
	}
	err = encoding.In("thresholdValue", j0)
	if err != nil {
		return err
	}
	d.ThresholdValue, err = encoding.AsInt32(j0["thresholdValue"])
	if err != nil {
		return err
	}
	err = encoding.In("rawValue", j0)
	if err != nil {
		return err
	}
	d.RawValue, err = encoding.AsInt64(j0["rawValue"])
	if err != nil {
		return err
	}
	err = encoding.In("flags", j0)
	if err != nil {
		return err
	}
	d.Flags, err = encoding.AsInt32(j0["flags"])
	if err != nil {
		return err
	}
	return nil
}

func (e *ErrorLogEntry) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["id"] = e.Id
	j0["value"] = e.Value
	j0["thresholdValue"] = e.ThresholdValue
	j0["rawValue"] = e.RawValue
	j0["powerOnHours"] = e.PowerOnHours
	j0["timeStampUTC"] = e.TimeStampUTC.Unix()
	return j0
}

func (e *ErrorLogEntry) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("id", j0)
	if err != nil {
		return err
	}
	e.Id, err = encoding.Is[string](j0["id"])
	if err != nil {
		return err
	}
	err = encoding.In("value", j0)
	if err != nil {
		return err
	}
	e.Value, err = encoding.AsInt32(j0["value"])
	if err != nil {
		return err
	}
	err = encoding.In("thresholdValue", j0)
	if err != nil {
		return err
	}
	e.ThresholdValue, err = encoding.AsInt32(j0["thresholdValue"])
	if err != nil {
		return err
	}
	err = encoding.In("rawValue", j0)
	if err != nil {
		return err
	}
	e.RawValue, err = encoding.AsInt64(j0["rawValue"])
	if err != nil {
		return err
	}
	err = encoding.In("powerOnHours", j0)
	if err != nil {
		return err
	}
	e.PowerOnHours, err = encoding.AsInt32(j0["powerOnHours"])
	if err != nil {
		return err
	}
	err = encoding.In("timeStampUTC", j0)
	if err != nil {
		return err
	}
	e.TimeStampUTC, err = encoding.AsTime(j0["timeStampUTC"])
	if err != nil {
		return err
	}
	return nil
}
