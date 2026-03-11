// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package outlets

import (
	"github.com/arminguenther/xeruspower-go/v40412/idl"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding"
)

func (i *Info) Encode() map[string]any {
	j0 := make(map[string]any, 1)
	j0["numberOfOutlets"] = i.NumberOfOutlets
	return j0
}

func (i *Info) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("numberOfOutlets", j0)
	if err != nil {
		return err
	}
	i.NumberOfOutlets, err = encoding.AsInt32(j0["numberOfOutlets"])
	if err != nil {
		return err
	}
	return nil
}
