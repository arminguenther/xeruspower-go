// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package controller

import (
	"github.com/arminguenther/xeruspower-go/v40040/idl"
	"github.com/arminguenther/xeruspower-go/v40040/idl/event"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/valobj"
)

func (c *CtrlStatistic) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["mainCSumErrCnt"] = c.MainCSumErrCnt
	j0["subCSumErrCnt"] = c.SubCSumErrCnt
	j0["timeoutCnt"] = c.TimeoutCnt
	j0["resetCnt"] = c.ResetCnt
	j0["emResetCnt"] = c.EmResetCnt
	return j0
}

func (c *CtrlStatistic) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("mainCSumErrCnt", j0)
	if err != nil {
		return err
	}
	c.MainCSumErrCnt, err = encoding.AsInt32(j0["mainCSumErrCnt"])
	if err != nil {
		return err
	}
	err = encoding.In("subCSumErrCnt", j0)
	if err != nil {
		return err
	}
	c.SubCSumErrCnt, err = encoding.AsInt32(j0["subCSumErrCnt"])
	if err != nil {
		return err
	}
	err = encoding.In("timeoutCnt", j0)
	if err != nil {
		return err
	}
	c.TimeoutCnt, err = encoding.AsInt32(j0["timeoutCnt"])
	if err != nil {
		return err
	}
	err = encoding.In("resetCnt", j0)
	if err != nil {
		return err
	}
	c.ResetCnt, err = encoding.AsInt32(j0["resetCnt"])
	if err != nil {
		return err
	}
	err = encoding.In("emResetCnt", j0)
	if err != nil {
		return err
	}
	c.EmResetCnt, err = encoding.AsInt32(j0["emResetCnt"])
	if err != nil {
		return err
	}
	return nil
}

func (m *MetaData) Encode() map[string]any {
	j0 := make(map[string]any, 10)
	j0["type"] = m.Type
	j0["address"] = m.Address
	j0["magic"] = m.Magic
	j0["versionAvailable"] = m.VersionAvailable
	j0["fwAppVersion"] = m.FwAppVersion
	j0["fwBootVersion"] = m.FwBootVersion
	j0["hwVersion"] = m.HwVersion
	j0["serial"] = m.Serial
	j0["haveResetCnt"] = m.HaveResetCnt
	j0["haveEmResetCnt"] = m.HaveEmResetCnt
	return j0
}

func (m *MetaData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	m.Type, err = encoding.AsEnum[Type](j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("address", j0)
	if err != nil {
		return err
	}
	m.Address, err = encoding.Is[string](j0["address"])
	if err != nil {
		return err
	}
	err = encoding.In("magic", j0)
	if err != nil {
		return err
	}
	m.Magic, err = encoding.AsInt32(j0["magic"])
	if err != nil {
		return err
	}
	err = encoding.In("versionAvailable", j0)
	if err != nil {
		return err
	}
	m.VersionAvailable, err = encoding.Is[bool](j0["versionAvailable"])
	if err != nil {
		return err
	}
	err = encoding.In("fwAppVersion", j0)
	if err != nil {
		return err
	}
	m.FwAppVersion, err = encoding.AsInt32(j0["fwAppVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("fwBootVersion", j0)
	if err != nil {
		return err
	}
	m.FwBootVersion, err = encoding.AsInt32(j0["fwBootVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("hwVersion", j0)
	if err != nil {
		return err
	}
	m.HwVersion, err = encoding.AsInt32(j0["hwVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("serial", j0)
	if err != nil {
		return err
	}
	m.Serial, err = encoding.Is[string](j0["serial"])
	if err != nil {
		return err
	}
	err = encoding.In("haveResetCnt", j0)
	if err != nil {
		return err
	}
	m.HaveResetCnt, err = encoding.Is[bool](j0["haveResetCnt"])
	if err != nil {
		return err
	}
	err = encoding.In("haveEmResetCnt", j0)
	if err != nil {
		return err
	}
	m.HaveEmResetCnt, err = encoding.Is[bool](j0["haveEmResetCnt"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_StatusChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.Event = valobj.For[event.Event]()
	err := s.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldStatus", value)
	if err != nil {
		return err
	}
	s.oldStatus, err = encoding.AsEnum[Status](value["oldStatus"])
	if err != nil {
		return err
	}
	err = encoding.In("newStatus", value)
	if err != nil {
		return err
	}
	s.newStatus, err = encoding.AsEnum[Status](value["newStatus"])
	if err != nil {
		return err
	}
	return nil
}

func (m *_MetaDataChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	m.Event = valobj.For[event.Event]()
	err := m.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldMetaData", value)
	if err != nil {
		return err
	}
	err = m.oldMetaData.Decode(value["oldMetaData"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newMetaData", value)
	if err != nil {
		return err
	}
	err = m.newMetaData.Decode(value["newMetaData"], caller)
	if err != nil {
		return err
	}
	return nil
}
