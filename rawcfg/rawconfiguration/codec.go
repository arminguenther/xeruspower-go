// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package rawconfiguration

import (
	"github.com/arminguenther/xeruspower-go/v40510/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/valobj"
)

func (r *_RawConfigDownloadedEvent) Decode(value map[string]any, caller idl.Caller) error {
	r.UserEvent = valobj.For[userevent.UserEvent]()
	err := r.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (r *_RawConfigUpdatedEvent) Decode(value map[string]any, caller idl.Caller) error {
	r.UserEvent = valobj.For[userevent.UserEvent]()
	err := r.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}
