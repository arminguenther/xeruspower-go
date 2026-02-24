// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package dsamdevice

import (
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
)

func (f *FirmwareVersion) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["major"] = f.Major
	j0["minor"] = f.Minor
	return j0
}

func (f *FirmwareVersion) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("major", j0)
	if err != nil {
		return err
	}
	f.Major, err = encoding.AsInt32(j0["major"])
	if err != nil {
		return err
	}
	err = encoding.In("minor", j0)
	if err != nil {
		return err
	}
	f.Minor, err = encoding.AsInt32(j0["minor"])
	if err != nil {
		return err
	}
	return nil
}

func (i *Info) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["dsamNumber"] = i.DsamNumber
	j0["serialNumber"] = i.SerialNumber
	j0["portCount"] = i.PortCount
	j0["hardwareVersion"] = i.HardwareVersion
	j0["firmwareVersion"] = i.FirmwareVersion.Encode()
	return j0
}

func (i *Info) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("dsamNumber", j0)
	if err != nil {
		return err
	}
	i.DsamNumber, err = encoding.AsInt32(j0["dsamNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("serialNumber", j0)
	if err != nil {
		return err
	}
	i.SerialNumber, err = encoding.Is[string](j0["serialNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("portCount", j0)
	if err != nil {
		return err
	}
	i.PortCount, err = encoding.AsInt32(j0["portCount"])
	if err != nil {
		return err
	}
	err = encoding.In("hardwareVersion", j0)
	if err != nil {
		return err
	}
	i.HardwareVersion, err = encoding.AsInt32(j0["hardwareVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("firmwareVersion", j0)
	if err != nil {
		return err
	}
	err = i.FirmwareVersion.Decode(j0["firmwareVersion"], caller)
	if err != nil {
		return err
	}
	return nil
}
