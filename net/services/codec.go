// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package services

import (
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
)

func (s *ServiceSettings) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["service"] = s.Service
	j0["enable"] = s.Enable
	j0["port"] = s.Port
	return j0
}

func (s *ServiceSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("service", j0)
	if err != nil {
		return err
	}
	s.Service, err = encoding.Is[string](j0["service"])
	if err != nil {
		return err
	}
	err = encoding.In("enable", j0)
	if err != nil {
		return err
	}
	s.Enable, err = encoding.Is[bool](j0["enable"])
	if err != nil {
		return err
	}
	err = encoding.In("port", j0)
	if err != nil {
		return err
	}
	s.Port, err = encoding.AsInt32(j0["port"])
	if err != nil {
		return err
	}
	return nil
}
