// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package tacplusmanager

import (
	"github.com/arminguenther/xeruspower-go/v40100/idl"
	"github.com/arminguenther/xeruspower-go/v40100/internal/encoding"
)

func (s *ServerSettings) Encode() map[string]any {
	j0 := make(map[string]any, 7)
	j0["server"] = s.Server
	j0["port"] = s.Port
	j0["timeoutSeconds"] = s.TimeoutSeconds
	j0["retries"] = s.Retries
	j0["sharedSecret"] = s.SharedSecret
	j0["authenType"] = s.AuthenType
	j0["disableAccounting"] = s.DisableAccounting
	return j0
}

func (s *ServerSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("server", j0)
	if err != nil {
		return err
	}
	s.Server, err = encoding.Is[string](j0["server"])
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
	err = encoding.In("timeoutSeconds", j0)
	if err != nil {
		return err
	}
	s.TimeoutSeconds, err = encoding.AsInt32(j0["timeoutSeconds"])
	if err != nil {
		return err
	}
	err = encoding.In("retries", j0)
	if err != nil {
		return err
	}
	s.Retries, err = encoding.AsInt32(j0["retries"])
	if err != nil {
		return err
	}
	err = encoding.In("sharedSecret", j0)
	if err != nil {
		return err
	}
	s.SharedSecret, err = encoding.Is[string](j0["sharedSecret"])
	if err != nil {
		return err
	}
	err = encoding.In("authenType", j0)
	if err != nil {
		return err
	}
	s.AuthenType, err = encoding.AsEnum[AuthenType](j0["authenType"])
	if err != nil {
		return err
	}
	err = encoding.In("disableAccounting", j0)
	if err != nil {
		return err
	}
	s.DisableAccounting, err = encoding.Is[bool](j0["disableAccounting"])
	if err != nil {
		return err
	}
	return nil
}
