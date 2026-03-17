// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package gatewaysensormanager

// IsKnown reports whether the SnmpSecurityLevel is a known value.
func (s SnmpSecurityLevel) IsKnown() bool {
	return s >= NO_AUTH_NO_PRIV && s <= AUTH_PRIV
}

// IsKnown reports whether the SnmpAuthProtocol is a known value.
func (s SnmpAuthProtocol) IsKnown() bool {
	return s >= MD5 && s <= SHA512
}

// IsKnown reports whether the SnmpPrivProtocol is a known value.
func (s SnmpPrivProtocol) IsKnown() bool {
	return s >= DES && s <= AES256_3DES
}

// IsKnown reports whether the EncodingType is a known value.
func (e EncodingType) IsKnown() bool {
	return e >= BOOL && e <= IEEE754
}

// IsKnown reports whether the Interpretation is a known value.
func (i Interpretation) IsKnown() bool {
	return i >= DEFAULT && i <= STATE_OFF
}

// IsKnown reports whether the ModbusEndianness is a known value.
func (m ModbusEndianness) IsKnown() bool {
	return m >= MODBUS_BIG_ENDIAN && m <= MODBUS_LITTLE_ENDIAN
}

// IsKnown reports whether the FeedbackObjectFeedbackState is a known value.
func (f FeedbackObjectFeedbackState) IsKnown() bool {
	return f >= UNSPECIFIED && f <= GOOD
}

//go:generate stringer -output=enum_strings.go -type=SnmpAuthProtocol,FeedbackObjectFeedbackState,EncodingType,ModbusEndianness,SnmpPrivProtocol,Interpretation,SnmpSecurityLevel
