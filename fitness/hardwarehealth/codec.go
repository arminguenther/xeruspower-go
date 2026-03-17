// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package hardwarehealth

import (
	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/idl/event"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/valobj"
)

func (f *Failure) Encode() map[string]any {
	j0 := make(map[string]any, 8)
	j0["creationOrderId"] = f.CreationOrderId
	j0["componentId"] = f.ComponentId
	j0["type"] = f.Type
	j0["description"] = f.Description
	j0["isAsserted"] = f.IsAsserted
	j0["lastAssertTimeStamp"] = f.LastAssertTimeStamp.Unix()
	j0["lastDeassertTimeStamp"] = f.LastDeassertTimeStamp.Unix()
	j0["assertCount"] = f.AssertCount
	return j0
}

func (f *Failure) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("creationOrderId", j0)
	if err != nil {
		return err
	}
	f.CreationOrderId, err = encoding.AsInt32(j0["creationOrderId"])
	if err != nil {
		return err
	}
	err = encoding.In("componentId", j0)
	if err != nil {
		return err
	}
	f.ComponentId, err = encoding.Is[string](j0["componentId"])
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	f.Type, err = encoding.AsInt32(j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("description", j0)
	if err != nil {
		return err
	}
	f.Description, err = encoding.Is[string](j0["description"])
	if err != nil {
		return err
	}
	err = encoding.In("isAsserted", j0)
	if err != nil {
		return err
	}
	f.IsAsserted, err = encoding.Is[bool](j0["isAsserted"])
	if err != nil {
		return err
	}
	err = encoding.In("lastAssertTimeStamp", j0)
	if err != nil {
		return err
	}
	f.LastAssertTimeStamp, err = encoding.AsTime(j0["lastAssertTimeStamp"])
	if err != nil {
		return err
	}
	err = encoding.In("lastDeassertTimeStamp", j0)
	if err != nil {
		return err
	}
	f.LastDeassertTimeStamp, err = encoding.AsTime(j0["lastDeassertTimeStamp"])
	if err != nil {
		return err
	}
	err = encoding.In("assertCount", j0)
	if err != nil {
		return err
	}
	f.AssertCount, err = encoding.AsInt32(j0["assertCount"])
	if err != nil {
		return err
	}
	return nil
}

func (f *_FailureStatusChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	f.Event = valobj.For[event.Event]()
	err := f.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("componentId", value)
	if err != nil {
		return err
	}
	f.componentId, err = encoding.Is[string](value["componentId"])
	if err != nil {
		return err
	}
	err = encoding.In("failureType", value)
	if err != nil {
		return err
	}
	f.failureType, err = encoding.AsInt32(value["failureType"])
	if err != nil {
		return err
	}
	err = encoding.In("isAsserted", value)
	if err != nil {
		return err
	}
	f.isAsserted, err = encoding.Is[bool](value["isAsserted"])
	if err != nil {
		return err
	}
	return nil
}
