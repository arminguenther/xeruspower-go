// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package inlets

import (
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
)

func (i *Info) Encode() map[string]any {
	j0 := make(map[string]any, 1)
	j0["numberOfInlets"] = i.NumberOfInlets
	return j0
}

func (i *Info) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("numberOfInlets", j0)
	if err != nil {
		return err
	}
	i.NumberOfInlets, err = encoding.AsInt32(j0["numberOfInlets"])
	if err != nil {
		return err
	}
	return nil
}
