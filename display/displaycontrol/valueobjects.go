// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package displaycontrol

import (
	"github.com/arminguenther/xeruspower-go/v40040/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40040/idl"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() DefaultViewChangedEvent { return &_DefaultViewChangedEvent{} })
}

type _DefaultViewChangedEvent struct {
	userevent.UserEvent
	newView DefaultViewItem
}

func (d *_DefaultViewChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "display.DisplayControl_1_0_1.DefaultViewChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DefaultViewChangedEvent) NewView() DefaultViewItem {
	return d.newView
}

func (d *_DefaultViewChangedEvent) isDefaultViewChangedEvent() {}
