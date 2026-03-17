// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package assetstripconfig

// IsKnown reports whether the ScanMode is a known value.
func (s ScanMode) IsKnown() bool {
	return s >= SCANMODE_DISABLED && s <= SCANMODE_BOTH
}

// IsKnown reports whether the NumberingMode is a known value.
func (n NumberingMode) IsKnown() bool {
	return n >= TOP_DOWN && n <= BOTTOM_UP
}

// IsKnown reports whether the Orientation is a known value.
func (o Orientation) IsKnown() bool {
	return o >= TOP_CONNECTOR && o <= BOTTOM_CONNECTOR
}

// IsKnown reports whether the LEDOperationMode is a known value.
func (l LEDOperationMode) IsKnown() bool {
	return l >= LED_OPERATION_MANUAL && l <= LED_OPERATION_AUTO
}

// IsKnown reports whether the LEDMode is a known value.
func (l LEDMode) IsKnown() bool {
	return l >= LED_MODE_ON && l <= LED_MODE_BLINK_SLOW
}

//go:generate stringer -output=enum_strings.go -type=ScanMode,LEDMode,Orientation,LEDOperationMode,NumberingMode
