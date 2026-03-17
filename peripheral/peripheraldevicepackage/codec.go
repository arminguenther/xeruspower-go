// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package peripheraldevicepackage

import (
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/idl/event"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/peripheral/peripheraldeviceslot"
)

func (p *PackageInfo) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["state"] = p.State
	s1 := make([]any, 0, len(p.Position))
	for _, e1 := range p.Position {
		s1 = append(s1, e1.Encode())
	}
	j0["position"] = s1
	j2 := make(map[string]any, 5)
	j2["serial"] = p.HwInfo.Serial
	j2["packageClass"] = p.HwInfo.PackageClass
	j2["model"] = p.HwInfo.Model
	j2["minDowngradeVersion"] = p.HwInfo.MinDowngradeVersion
	j2["revision"] = p.HwInfo.Revision
	j0["hwInfo"] = j2
	j3 := make(map[string]any, 3)
	j3["compileDate"] = p.FwInfo.CompileDate.Unix()
	j4 := make(map[string]any, 3)
	j4["majorNumber"] = p.FwInfo.Version.MajorNumber
	j4["minorNumber"] = p.FwInfo.Version.MinorNumber
	j4["bootloaderVersion"] = p.FwInfo.Version.BootloaderVersion
	j3["version"] = j4
	j3["updateDate"] = p.FwInfo.UpdateDate.Unix()
	j0["fwInfo"] = j3
	return j0
}

func (p *PackageInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("state", j0)
	if err != nil {
		return err
	}
	p.State, err = encoding.AsEnum[PackageInfoState](j0["state"])
	if err != nil {
		return err
	}
	err = encoding.In("position", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["position"])
	if err != nil {
		return err
	}
	p.Position = make([]peripheraldeviceslot.PosElement, 0, len(s1))
	for _, a1 := range s1 {
		var e1 peripheraldeviceslot.PosElement
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		p.Position = append(p.Position, e1)
	}
	err = encoding.In("hwInfo", j0)
	if err != nil {
		return err
	}
	j2, err := encoding.Is[map[string]any](j0["hwInfo"])
	if err != nil {
		return err
	}
	err = encoding.In("serial", j2)
	if err != nil {
		return err
	}
	p.HwInfo.Serial, err = encoding.Is[string](j2["serial"])
	if err != nil {
		return err
	}
	err = encoding.In("packageClass", j2)
	if err != nil {
		return err
	}
	p.HwInfo.PackageClass, err = encoding.Is[string](j2["packageClass"])
	if err != nil {
		return err
	}
	err = encoding.In("model", j2)
	if err != nil {
		return err
	}
	p.HwInfo.Model, err = encoding.Is[string](j2["model"])
	if err != nil {
		return err
	}
	err = encoding.In("minDowngradeVersion", j2)
	if err != nil {
		return err
	}
	p.HwInfo.MinDowngradeVersion, err = encoding.AsInt32(j2["minDowngradeVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("revision", j2)
	if err != nil {
		return err
	}
	p.HwInfo.Revision, err = encoding.Is[string](j2["revision"])
	if err != nil {
		return err
	}
	err = encoding.In("fwInfo", j0)
	if err != nil {
		return err
	}
	j3, err := encoding.Is[map[string]any](j0["fwInfo"])
	if err != nil {
		return err
	}
	err = encoding.In("compileDate", j3)
	if err != nil {
		return err
	}
	p.FwInfo.CompileDate, err = encoding.AsTime(j3["compileDate"])
	if err != nil {
		return err
	}
	err = encoding.In("version", j3)
	if err != nil {
		return err
	}
	j4, err := encoding.Is[map[string]any](j3["version"])
	if err != nil {
		return err
	}
	err = encoding.In("majorNumber", j4)
	if err != nil {
		return err
	}
	p.FwInfo.Version.MajorNumber, err = encoding.AsInt32(j4["majorNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("minorNumber", j4)
	if err != nil {
		return err
	}
	p.FwInfo.Version.MinorNumber, err = encoding.AsInt32(j4["minorNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("bootloaderVersion", j4)
	if err != nil {
		return err
	}
	p.FwInfo.Version.BootloaderVersion, err = encoding.AsInt32(j4["bootloaderVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("updateDate", j3)
	if err != nil {
		return err
	}
	p.FwInfo.UpdateDate, err = encoding.AsTime(j3["updateDate"])
	if err != nil {
		return err
	}
	return nil
}

func (v *_BatteryPoweredDevicePackageVoltageChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	v.Event = valobj.For[event.Event]()
	err := v.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldVoltage", value)
	if err != nil {
		return err
	}
	v.oldVoltage, err = encoding.AsFloat64(value["oldVoltage"])
	if err != nil {
		return err
	}
	err = encoding.In("newVoltage", value)
	if err != nil {
		return err
	}
	v.newVoltage, err = encoding.AsFloat64(value["newVoltage"])
	if err != nil {
		return err
	}
	return nil
}
