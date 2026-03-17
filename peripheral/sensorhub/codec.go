// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package sensorhub

import (
	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40020/portsmodel/portfuse"
)

func (h *HubPortInfo) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["hubPort"] = h.HubPort
	j0["fuse"] = object.ToMap(h.Fuse)
	return j0
}

func (h *HubPortInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("hubPort", j0)
	if err != nil {
		return err
	}
	h.HubPort, err = encoding.Is[string](j0["hubPort"])
	if err != nil {
		return err
	}
	err = encoding.In("fuse", j0)
	if err != nil {
		return err
	}
	h.Fuse, err = object.As[portfuse.PortFuse](j0["fuse"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (d *DeviceInfo) Encode() map[string]any {
	j0 := make(map[string]any, 8)
	j0["serial"] = d.Serial
	j0["model"] = d.Model
	j0["upstreamType"] = d.UpstreamType
	j0["position"] = d.Position.Encode()
	j0["protocolVersion"] = d.ProtocolVersion
	j0["bootVersion"] = d.BootVersion
	j0["appVersion"] = d.AppVersion
	s1 := make([]any, 0, len(d.HubPortInfos))
	for _, e1 := range d.HubPortInfos {
		s1 = append(s1, e1.Encode())
	}
	j0["hubPortInfos"] = s1
	return j0
}

func (d *DeviceInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("serial", j0)
	if err != nil {
		return err
	}
	d.Serial, err = encoding.Is[string](j0["serial"])
	if err != nil {
		return err
	}
	err = encoding.In("model", j0)
	if err != nil {
		return err
	}
	d.Model, err = encoding.Is[string](j0["model"])
	if err != nil {
		return err
	}
	err = encoding.In("upstreamType", j0)
	if err != nil {
		return err
	}
	d.UpstreamType, err = encoding.AsEnum[UpstreamType](j0["upstreamType"])
	if err != nil {
		return err
	}
	err = encoding.In("position", j0)
	if err != nil {
		return err
	}
	err = d.Position.Decode(j0["position"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("protocolVersion", j0)
	if err != nil {
		return err
	}
	d.ProtocolVersion, err = encoding.AsInt32(j0["protocolVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("bootVersion", j0)
	if err != nil {
		return err
	}
	d.BootVersion, err = encoding.AsInt32(j0["bootVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("appVersion", j0)
	if err != nil {
		return err
	}
	d.AppVersion, err = encoding.AsInt32(j0["appVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("hubPortInfos", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["hubPortInfos"])
	if err != nil {
		return err
	}
	d.HubPortInfos = make([]HubPortInfo, 0, len(s1))
	for _, a1 := range s1 {
		var e1 HubPortInfo
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		d.HubPortInfos = append(d.HubPortInfos, e1)
	}
	return nil
}
