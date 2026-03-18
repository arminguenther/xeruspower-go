// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package resmon

import (
	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding"
)

func (e *Entry) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["type"] = e.Type
	j0["name"] = e.Name
	j0["value"] = e.Value
	return j0
}

func (e *Entry) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	e.Type, err = encoding.AsEnum[EntryType](j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	e.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("value", j0)
	if err != nil {
		return err
	}
	e.Value, err = encoding.AsInt64(j0["value"])
	if err != nil {
		return err
	}
	return nil
}
