// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package userevent

import (
	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/idl/event"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/valobj"
)

func (u *_UserEvent) Decode(value map[string]any, caller idl.Caller) error {
	u.Event = valobj.For[event.Event]()
	err := u.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("actUserName", value)
	if err != nil {
		return err
	}
	u.actUserName, err = encoding.Is[string](value["actUserName"])
	if err != nil {
		return err
	}
	err = encoding.In("actIpAddr", value)
	if err != nil {
		return err
	}
	u.actIpAddr, err = encoding.Is[string](value["actIpAddr"])
	if err != nil {
		return err
	}
	return nil
}
