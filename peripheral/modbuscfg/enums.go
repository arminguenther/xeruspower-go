// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package modbuscfg

// IsKnown reports whether the SerialSettingsParity is a known value.
func (p SerialSettingsParity) IsKnown() bool {
	return p >= NONE && p <= ODD
}

// IsKnown reports whether the ModbusFunction is a known value.
func (m ModbusFunction) IsKnown() bool {
	return m >= COIL && m <= INPUT_REGISTER
}

// IsKnown reports whether the SpecificModbusErrors is a known value.
func (s SpecificModbusErrors) IsKnown() bool {
	return s >= ERROR_BADCRC && s <= ERROR_OTHER
}

//go:generate stringer -output=enum_strings.go -type=SerialSettingsParity,ModbusFunction,SpecificModbusErrors
