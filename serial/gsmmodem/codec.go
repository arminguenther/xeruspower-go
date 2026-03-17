// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package gsmmodem

import (
	"github.com/arminguenther/xeruspower-go/v40032/idl"
	"github.com/arminguenther/xeruspower-go/v40032/idl/event"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding/valobj"
)

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["pin"] = s.Pin
	j0["smsc"] = s.Smsc
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("pin", j0)
	if err != nil {
		return err
	}
	s.Pin, err = encoding.Is[string](j0["pin"])
	if err != nil {
		return err
	}
	err = encoding.In("smsc", j0)
	if err != nil {
		return err
	}
	s.Smsc, err = encoding.Is[string](j0["smsc"])
	if err != nil {
		return err
	}
	return nil
}

func (i *Information) Encode() map[string]any {
	j0 := make(map[string]any, 10)
	j0["imei"] = i.Imei
	j0["imsi"] = i.Imsi
	j0["manufacturer"] = i.Manufacturer
	j0["model"] = i.Model
	j0["revision"] = i.Revision
	j0["ownNumber"] = i.OwnNumber
	j0["simSmsc"] = i.SimSmsc
	j0["networkName"] = i.NetworkName
	j0["serviceProviderName"] = i.ServiceProviderName
	j0["receptionLevel"] = i.ReceptionLevel
	return j0
}

func (i *Information) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("imei", j0)
	if err != nil {
		return err
	}
	i.Imei, err = encoding.Is[string](j0["imei"])
	if err != nil {
		return err
	}
	err = encoding.In("imsi", j0)
	if err != nil {
		return err
	}
	i.Imsi, err = encoding.Is[string](j0["imsi"])
	if err != nil {
		return err
	}
	err = encoding.In("manufacturer", j0)
	if err != nil {
		return err
	}
	i.Manufacturer, err = encoding.Is[string](j0["manufacturer"])
	if err != nil {
		return err
	}
	err = encoding.In("model", j0)
	if err != nil {
		return err
	}
	i.Model, err = encoding.Is[string](j0["model"])
	if err != nil {
		return err
	}
	err = encoding.In("revision", j0)
	if err != nil {
		return err
	}
	i.Revision, err = encoding.Is[string](j0["revision"])
	if err != nil {
		return err
	}
	err = encoding.In("ownNumber", j0)
	if err != nil {
		return err
	}
	i.OwnNumber, err = encoding.Is[string](j0["ownNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("simSmsc", j0)
	if err != nil {
		return err
	}
	i.SimSmsc, err = encoding.Is[string](j0["simSmsc"])
	if err != nil {
		return err
	}
	err = encoding.In("networkName", j0)
	if err != nil {
		return err
	}
	i.NetworkName, err = encoding.Is[string](j0["networkName"])
	if err != nil {
		return err
	}
	err = encoding.In("serviceProviderName", j0)
	if err != nil {
		return err
	}
	i.ServiceProviderName, err = encoding.Is[string](j0["serviceProviderName"])
	if err != nil {
		return err
	}
	err = encoding.In("receptionLevel", j0)
	if err != nil {
		return err
	}
	i.ReceptionLevel, err = encoding.AsInt32(j0["receptionLevel"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_SimSecurityStatusChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.Event = valobj.For[event.Event]()
	err := s.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("newSimStatus", value)
	if err != nil {
		return err
	}
	s.newSimStatus, err = encoding.AsEnum[SimSecurityStatus](value["newSimStatus"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_SimPinUpdatedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.Event = valobj.For[event.Event]()
	err := s.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("newPin", value)
	if err != nil {
		return err
	}
	s.newPin, err = encoding.Is[string](value["newPin"])
	if err != nil {
		return err
	}
	return nil
}
