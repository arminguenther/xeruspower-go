// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package sensor

import (
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/idl/event"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
)

func (t *TypeSpec) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["readingtype"] = t.Readingtype
	j0["type"] = t.Type
	j0["unit"] = t.Unit
	return j0
}

func (t *TypeSpec) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("readingtype", j0)
	if err != nil {
		return err
	}
	t.Readingtype, err = encoding.AsInt32(j0["readingtype"])
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	t.Type, err = encoding.AsInt32(j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("unit", j0)
	if err != nil {
		return err
	}
	t.Unit, err = encoding.AsInt32(j0["unit"])
	if err != nil {
		return err
	}
	return nil
}

func (t *_TypeSpecChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	t.Event = valobj.For[event.Event]()
	err := t.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldTypeSpec", value)
	if err != nil {
		return err
	}
	err = t.oldTypeSpec.Decode(value["oldTypeSpec"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newTypeSpec", value)
	if err != nil {
		return err
	}
	err = t.newTypeSpec.Decode(value["newTypeSpec"], caller)
	if err != nil {
		return err
	}
	return nil
}
