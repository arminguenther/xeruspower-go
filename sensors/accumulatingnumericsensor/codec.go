// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package accumulatingnumericsensor

import (
	"github.com/arminguenther/xeruspower-go/v40000/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40000/idl"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding/valobj"
)

func (r *_ResetEvent) Decode(value map[string]any, caller idl.Caller) error {
	r.UserEvent = valobj.For[userevent.UserEvent]()
	err := r.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldReading", value)
	if err != nil {
		return err
	}
	err = r.oldReading.Decode(value["oldReading"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newReading", value)
	if err != nil {
		return err
	}
	err = r.newReading.Decode(value["newReading"], caller)
	if err != nil {
		return err
	}
	return nil
}
