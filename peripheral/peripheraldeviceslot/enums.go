// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package peripheraldeviceslot

// IsKnown reports whether the PortType is a known value.
func (p PortType) IsKnown() bool {
	return p >= ONBOARD && p <= GATEWAY_SENSOR
}

// Fallback returns a known PortType used as a fallback if an
// unknown value is encountered while decoding.
// The fallback is [UNSPECIFIED].
func (p PortType) Fallback() PortType {
	return UNSPECIFIED
}

//go:generate stringer -output=enum_strings.go -type=PortType
