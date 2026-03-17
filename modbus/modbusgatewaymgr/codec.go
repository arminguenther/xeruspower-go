// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package modbusgatewaymgr

import (
	"github.com/arminguenther/xeruspower-go/v40032/idl"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding"
)

func (r *GatewayMgrRtuSettings) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["defaultAddr"] = r.DefaultAddr
	j0["speed"] = r.Speed
	j0["parity"] = r.Parity
	return j0
}

func (r *GatewayMgrRtuSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("defaultAddr", j0)
	if err != nil {
		return err
	}
	r.DefaultAddr, err = encoding.AsByte(j0["defaultAddr"])
	if err != nil {
		return err
	}
	err = encoding.In("speed", j0)
	if err != nil {
		return err
	}
	r.Speed, err = encoding.AsInt32(j0["speed"])
	if err != nil {
		return err
	}
	err = encoding.In("parity", j0)
	if err != nil {
		return err
	}
	r.Parity, err = encoding.AsByte(j0["parity"])
	if err != nil {
		return err
	}
	return nil
}

func (s *GatewayMgrSettings) Encode() map[string]any {
	j0 := make(map[string]any, 1)
	j0["rtu"] = s.Rtu.Encode()
	return j0
}

func (s *GatewayMgrSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("rtu", j0)
	if err != nil {
		return err
	}
	err = s.Rtu.Decode(j0["rtu"], caller)
	if err != nil {
		return err
	}
	return nil
}
