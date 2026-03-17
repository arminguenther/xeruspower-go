// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package eventservice

import (
	"github.com/arminguenther/xeruspower-go/v40010/idl"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding/object"
)

func (e *ChannelEventSelect) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["type"] = e.Type.String()
	j0["src"] = object.ToMap(e.Src)
	return j0
}

func (e *ChannelEventSelect) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	e.Type, err = encoding.AsTypeCode(j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("src", j0)
	if err != nil {
		return err
	}
	e.Src, err = object.As[idl.Object](j0["src"], caller)
	if err != nil {
		return err
	}
	return nil
}
