// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package testdisplay

import (
	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding"
)

func (i *DisplayInfo) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["type"] = i.Type
	j0["address"] = i.Address
	s1 := make([]any, 0, len(i.Options))
	for k1, v1 := range i.Options {
		s1 = append(s1, map[string]any{"key": k1, "value": v1})
	}
	j0["options"] = s1
	j0["orientation"] = i.Orientation
	return j0
}

func (i *DisplayInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	i.Type, err = encoding.Is[string](j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("address", j0)
	if err != nil {
		return err
	}
	i.Address, err = encoding.Is[string](j0["address"])
	if err != nil {
		return err
	}
	err = encoding.In("options", j0)
	if err != nil {
		return err
	}
	i1, e1, l1 := encoding.AsMapItems(j0["options"])
	i.Options = make(map[string]string, l1)
	for a1, b1 := range i1 {
		var k1 string
		k1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		var v1 string
		v1, err = encoding.Is[string](b1)
		if err != nil {
			return err
		}
		i.Options[k1] = v1
	}
	err = e1()
	if err != nil {
		return err
	}
	err = encoding.In("orientation", j0)
	if err != nil {
		return err
	}
	i.Orientation, err = encoding.AsEnum[DisplayOrientation](j0["orientation"])
	if err != nil {
		return err
	}
	return nil
}
