// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package user

// IsKnown reports whether the SnmpV3SecLevel is a known value.
func (s SnmpV3SecLevel) IsKnown() bool {
	return s >= NO_AUTH_NO_PRIV && s <= AUTH_PRIV
}

// IsKnown reports whether the SnmpV3AuthProto is a known value.
func (s SnmpV3AuthProto) IsKnown() bool {
	return s >= MD5 && s <= SHA512
}

// IsKnown reports whether the SnmpV3PrivProto is a known value.
func (s SnmpV3PrivProto) IsKnown() bool {
	return s >= DES && s <= AES256_3DES
}

// IsKnown reports whether the TemperatureEnum is a known value.
func (t TemperatureEnum) IsKnown() bool {
	return t >= DEG_C && t <= DEG_F
}

// IsKnown reports whether the LengthEnum is a known value.
func (l LengthEnum) IsKnown() bool {
	return l >= METER && l <= FEET
}

// IsKnown reports whether the PressureEnum is a known value.
func (p PressureEnum) IsKnown() bool {
	return p >= PASCAL && p <= PSI
}

//go:generate stringer -output=enum_strings.go -type=SnmpV3PrivProto,TemperatureEnum,PressureEnum,SnmpV3SecLevel,LengthEnum,SnmpV3AuthProto
