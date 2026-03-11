// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package cfg

import (
	"github.com/arminguenther/xeruspower-go/v40410/idl"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding"
)

func (k *KeyValue) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["key"] = k.Key
	j0["value"] = k.Value
	return j0
}

func (k *KeyValue) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("key", j0)
	if err != nil {
		return err
	}
	k.Key, err = encoding.Is[string](j0["key"])
	if err != nil {
		return err
	}
	err = encoding.In("value", j0)
	if err != nil {
		return err
	}
	k.Value, err = encoding.Is[string](j0["value"])
	if err != nil {
		return err
	}
	return nil
}
