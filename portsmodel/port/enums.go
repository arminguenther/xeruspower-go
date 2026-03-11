// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package port

// IsKnown reports whether the DetectionType is a known value.
func (d DetectionType) IsKnown() bool {
	return d >= AUTO && d <= DISABLED
}

// IsKnown reports whether the DeviceTypeId is a known value.
func (d DeviceTypeId) IsKnown() bool {
	return d >= UNSPECIFIED && d <= GATEWAY_SENSOR
}

// Fallback returns a known DeviceTypeId used as a fallback if an
// unknown value is encountered while decoding.
// The fallback is [OTHER].
func (d DeviceTypeId) Fallback() DeviceTypeId {
	return OTHER
}

//go:generate stringer -output=enum_strings.go -type=DetectionType,DeviceTypeId
