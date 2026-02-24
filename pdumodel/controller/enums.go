// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package controller

// IsKnown reports whether the Status is a known value.
func (s Status) IsKnown() bool {
	return s >= OK && s <= FIRMWARE_UPDATE
}

// IsKnown reports whether the Type is a known value.
func (t Type) IsKnown() bool {
	return t >= OUTLET_CTRL && t <= METER_CTRL
}

//go:generate stringer -output=enum_strings.go -type=Status,Type
