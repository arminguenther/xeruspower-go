// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package cascade

import (
	"github.com/arminguenther/xeruspower-go/v40412/idl"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding"
)

func (i *Info) Encode() map[string]any {
	j0 := make(map[string]any, 1)
	j0["pduIds"] = encoding.NonNilSlice(i.PduIds)
	return j0
}

func (i *Info) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("pduIds", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["pduIds"])
	if err != nil {
		return err
	}
	i.PduIds = make([]int32, 0, len(s1))
	for _, a1 := range s1 {
		var e1 int32
		e1, err = encoding.AsInt32(a1)
		if err != nil {
			return err
		}
		i.PduIds = append(i.PduIds, e1)
	}
	return nil
}
