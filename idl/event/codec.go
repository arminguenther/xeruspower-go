// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package event

import (
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/internal/encoding/valobj"
)

func (e *_Event) Decode(value map[string]any, caller idl.Caller) error {
	e.ValueObject = valobj.For[idl.ValueObject]()
	var err error
	err = encoding.In("source", value)
	if err != nil {
		return err
	}
	e.source, err = object.As[idl.Object](value["source"], caller)
	if err != nil {
		return err
	}
	return nil
}
