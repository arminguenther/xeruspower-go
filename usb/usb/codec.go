// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package usb

import (
	"github.com/arminguenther/xeruspower-go/v40010/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40010/idl"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding/valobj"
)

func (u *Device) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["bus"] = u.Bus
	j0["device"] = u.Device
	j0["vendorId"] = u.VendorId
	j0["productId"] = u.ProductId
	return j0
}

func (u *Device) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("bus", j0)
	if err != nil {
		return err
	}
	u.Bus, err = encoding.AsInt32(j0["bus"])
	if err != nil {
		return err
	}
	err = encoding.In("device", j0)
	if err != nil {
		return err
	}
	u.Device, err = encoding.AsInt32(j0["device"])
	if err != nil {
		return err
	}
	err = encoding.In("vendorId", j0)
	if err != nil {
		return err
	}
	u.VendorId, err = encoding.AsInt32(j0["vendorId"])
	if err != nil {
		return err
	}
	err = encoding.In("productId", j0)
	if err != nil {
		return err
	}
	u.ProductId, err = encoding.AsInt32(j0["productId"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 1)
	j0["hostPortsEnabled"] = s.HostPortsEnabled
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("hostPortsEnabled", j0)
	if err != nil {
		return err
	}
	s.HostPortsEnabled, err = encoding.Is[bool](j0["hostPortsEnabled"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_SettingsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.UserEvent = valobj.For[userevent.UserEvent]()
	err := s.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldSettings", value)
	if err != nil {
		return err
	}
	err = s.oldSettings.Decode(value["oldSettings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newSettings", value)
	if err != nil {
		return err
	}
	err = s.newSettings.Decode(value["newSettings"], caller)
	if err != nil {
		return err
	}
	return nil
}
