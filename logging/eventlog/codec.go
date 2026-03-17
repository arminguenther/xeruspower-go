// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package eventlog

import (
	"github.com/arminguenther/xeruspower-go/v40300/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding/valobj"
)

func (e *_ClearedEvent) Decode(value map[string]any, caller idl.Caller) error {
	e.UserEvent = valobj.For[userevent.UserEvent]()
	err := e.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}
