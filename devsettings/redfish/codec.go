// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package redfish

import (
	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
)

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 1)
	j0["enabled"] = s.Enabled
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("enabled", j0)
	if err != nil {
		return err
	}
	s.Enabled, err = encoding.Is[bool](j0["enabled"])
	if err != nil {
		return err
	}
	return nil
}
