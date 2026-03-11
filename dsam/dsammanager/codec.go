// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package dsammanager

import (
	"github.com/arminguenther/xeruspower-go/v40510/dsam/dsamdevice"
	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/idl/event"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/valobj"
)

func (d *_DsamAttachedEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.Event = valobj.For[event.Event]()
	err := d.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("info", value)
	if err != nil {
		return err
	}
	err = d.info.Decode(value["info"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("device", value)
	if err != nil {
		return err
	}
	d.device, err = object.As[dsamdevice.DsamDevice](value["device"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (d *_DsamDetachedEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.Event = valobj.For[event.Event]()
	err := d.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("info", value)
	if err != nil {
		return err
	}
	err = d.info.Decode(value["info"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (d *_DsamControllerChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.Event = valobj.For[event.Event]()
	err := d.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("dsamNumber", value)
	if err != nil {
		return err
	}
	d.dsamNumber, err = encoding.AsInt32(value["dsamNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("reset", value)
	if err != nil {
		return err
	}
	d.reset, err = encoding.Is[bool](value["reset"])
	if err != nil {
		return err
	}
	err = encoding.In("resetReason", value)
	if err != nil {
		return err
	}
	d.resetReason, err = encoding.Is[string](value["resetReason"])
	if err != nil {
		return err
	}
	return nil
}
