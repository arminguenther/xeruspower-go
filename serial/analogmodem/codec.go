// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package analogmodem

import (
	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/idl/event"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/valobj"
)

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["dialInEnabled"] = s.DialInEnabled
	j0["ringsUntilAnswer"] = s.RingsUntilAnswer
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("dialInEnabled", j0)
	if err != nil {
		return err
	}
	s.DialInEnabled, err = encoding.Is[bool](j0["dialInEnabled"])
	if err != nil {
		return err
	}
	err = encoding.In("ringsUntilAnswer", j0)
	if err != nil {
		return err
	}
	s.RingsUntilAnswer, err = encoding.AsInt32(j0["ringsUntilAnswer"])
	if err != nil {
		return err
	}
	return nil
}

func (d *_DialInEvent) Decode(value map[string]any, caller idl.Caller) error {
	d.Event = valobj.For[event.Event]()
	err := d.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("number", value)
	if err != nil {
		return err
	}
	d.number, err = encoding.Is[string](value["number"])
	if err != nil {
		return err
	}
	return nil
}

func (c *_CallReceivedEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.DialInEvent = valobj.For[DialInEvent]()
	err := c.DialInEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (c *_CallEndedEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.DialInEvent = valobj.For[DialInEvent]()
	err := c.DialInEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("disconnectedRemotely", value)
	if err != nil {
		return err
	}
	c.disconnectedRemotely, err = encoding.Is[bool](value["disconnectedRemotely"])
	if err != nil {
		return err
	}
	return nil
}
