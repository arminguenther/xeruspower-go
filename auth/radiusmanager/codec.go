// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package radiusmanager

import (
	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding"
)

func (s *ServerSettings) Encode() map[string]any {
	j0 := make(map[string]any, 10)
	j0["id"] = s.Id
	j0["server"] = s.Server
	j0["sharedSecret"] = s.SharedSecret
	j0["udpAuthPort"] = s.UdpAuthPort
	j0["udpAccountPort"] = s.UdpAccountPort
	j0["timeout"] = s.Timeout
	j0["retries"] = s.Retries
	j0["authType"] = s.AuthType
	j0["disableAccounting"] = s.DisableAccounting
	j0["messageAuthenticatorOptional"] = s.MessageAuthenticatorOptional
	return j0
}

func (s *ServerSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("id", j0)
	if err != nil {
		return err
	}
	s.Id, err = encoding.Is[string](j0["id"])
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
	err = encoding.In("sharedSecret", j0)
	if err != nil {
		return err
	}
	s.SharedSecret, err = encoding.Is[string](j0["sharedSecret"])
	if err != nil {
		return err
	}
	err = encoding.In("udpAuthPort", j0)
	if err != nil {
		return err
	}
	s.UdpAuthPort, err = encoding.AsInt32(j0["udpAuthPort"])
	if err != nil {
		return err
	}
	err = encoding.In("udpAccountPort", j0)
	if err != nil {
		return err
	}
	s.UdpAccountPort, err = encoding.AsInt32(j0["udpAccountPort"])
	if err != nil {
		return err
	}
	err = encoding.In("timeout", j0)
	if err != nil {
		return err
	}
	s.Timeout, err = encoding.AsInt32(j0["timeout"])
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
	err = encoding.In("authType", j0)
	if err != nil {
		return err
	}
	s.AuthType, err = encoding.AsEnum[AuthType](j0["authType"])
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
	err = encoding.In("messageAuthenticatorOptional", j0)
	if err != nil {
		return err
	}
	s.MessageAuthenticatorOptional, err = encoding.Is[bool](j0["messageAuthenticatorOptional"])
	if err != nil {
		return err
	}
	return nil
}
