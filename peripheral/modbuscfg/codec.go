// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package modbuscfg

import (
	"github.com/arminguenther/xeruspower-go/v40040/idl"
	"github.com/arminguenther/xeruspower-go/v40040/internal/encoding"
)

func (s *SerialSettings) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["baud"] = s.Baud
	j0["parity"] = s.Parity
	j0["dataBits"] = s.DataBits
	j0["stopBits"] = s.StopBits
	return j0
}

func (s *SerialSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("baud", j0)
	if err != nil {
		return err
	}
	s.Baud, err = encoding.AsInt32(j0["baud"])
	if err != nil {
		return err
	}
	err = encoding.In("parity", j0)
	if err != nil {
		return err
	}
	s.Parity, err = encoding.AsEnum[SerialSettingsParity](j0["parity"])
	if err != nil {
		return err
	}
	err = encoding.In("dataBits", j0)
	if err != nil {
		return err
	}
	s.DataBits, err = encoding.AsInt32(j0["dataBits"])
	if err != nil {
		return err
	}
	err = encoding.In("stopBits", j0)
	if err != nil {
		return err
	}
	s.StopBits, err = encoding.AsInt32(j0["stopBits"])
	if err != nil {
		return err
	}
	return nil
}
