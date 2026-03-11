// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package poselement

import (
	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding"
)

func (p *PosElement) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["portType"] = p.PortType
	j0["port"] = p.Port
	return j0
}

func (p *PosElement) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("portType", j0)
	if err != nil {
		return err
	}
	p.PortType, err = encoding.AsEnum[PortType](j0["portType"])
	if err != nil {
		return err
	}
	err = encoding.In("port", j0)
	if err != nil {
		return err
	}
	p.Port, err = encoding.Is[string](j0["port"])
	if err != nil {
		return err
	}
	return nil
}
