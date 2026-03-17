// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package modbus

import (
	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding"
)

func (c *Capabilities) Encode() map[string]any {
	j0 := make(map[string]any, 1)
	j0["hasModbusSerial"] = c.HasModbusSerial
	return j0
}

func (c *Capabilities) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("hasModbusSerial", j0)
	if err != nil {
		return err
	}
	c.HasModbusSerial, err = encoding.Is[bool](j0["hasModbusSerial"])
	if err != nil {
		return err
	}
	return nil
}

func (t *TcpSettings) Encode() map[string]any {
	j0 := make(map[string]any, 1)
	j0["readonly"] = t.Readonly
	return j0
}

func (t *TcpSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("readonly", j0)
	if err != nil {
		return err
	}
	t.Readonly, err = encoding.Is[bool](j0["readonly"])
	if err != nil {
		return err
	}
	return nil
}

func (s *SerialSettings) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["enabled"] = s.Enabled
	j0["baudrate"] = s.Baudrate
	j0["parity"] = s.Parity
	j0["stopbits"] = s.Stopbits
	j0["readonly"] = s.Readonly
	return j0
}

func (s *SerialSettings) Decode(v any, caller idl.Caller) error {
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
	err = encoding.In("baudrate", j0)
	if err != nil {
		return err
	}
	s.Baudrate, err = encoding.AsInt32(j0["baudrate"])
	if err != nil {
		return err
	}
	err = encoding.In("parity", j0)
	if err != nil {
		return err
	}
	s.Parity, err = encoding.AsEnum[Parity](j0["parity"])
	if err != nil {
		return err
	}
	err = encoding.In("stopbits", j0)
	if err != nil {
		return err
	}
	s.Stopbits, err = encoding.AsInt32(j0["stopbits"])
	if err != nil {
		return err
	}
	err = encoding.In("readonly", j0)
	if err != nil {
		return err
	}
	s.Readonly, err = encoding.Is[bool](j0["readonly"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["tcp"] = s.Tcp.Encode()
	j0["serial"] = s.Serial.Encode()
	j0["primaryUnitId"] = s.PrimaryUnitId
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("tcp", j0)
	if err != nil {
		return err
	}
	err = s.Tcp.Decode(j0["tcp"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("serial", j0)
	if err != nil {
		return err
	}
	err = s.Serial.Decode(j0["serial"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("primaryUnitId", j0)
	if err != nil {
		return err
	}
	s.PrimaryUnitId, err = encoding.AsInt32(j0["primaryUnitId"])
	if err != nil {
		return err
	}
	return nil
}
