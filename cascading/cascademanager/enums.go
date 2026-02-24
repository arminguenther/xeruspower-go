// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package cascademanager

// IsKnown reports whether the Role is a known value.
func (r Role) IsKnown() bool {
	return r >= STANDALONE && r <= LINK_UNIT
}

// IsKnown reports whether the LinkUnitType is a known value.
func (l LinkUnitType) IsKnown() bool {
	return l >= NETWORK && l <= ZIGBEE
}

// IsKnown reports whether the LinkUnitStatus is a known value.
func (l LinkUnitStatus) IsKnown() bool {
	return l >= UNKNOWN && l <= PAIRING
}

// Fallback returns a known LinkUnitStatus used as a fallback if an
// unknown value is encountered while decoding.
// The fallback is [UNKNOWN].
func (l LinkUnitStatus) Fallback() LinkUnitStatus {
	return UNKNOWN
}

//go:generate stringer -output=enum_strings.go -type=Role,LinkUnitType,LinkUnitStatus
