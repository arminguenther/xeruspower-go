// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package crestron

import (
	"github.com/arminguenther/xeruspower-go/v40000/idl"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding"
)

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 1)
	j0["enableXioCloudConnection"] = s.EnableXioCloudConnection
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("enableXioCloudConnection", j0)
	if err != nil {
		return err
	}
	s.EnableXioCloudConnection, err = encoding.Is[bool](j0["enableXioCloudConnection"])
	if err != nil {
		return err
	}
	return nil
}
