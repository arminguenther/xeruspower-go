// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package serialport

// IsKnown reports whether the PortState is a known value.
func (p PortState) IsKnown() bool {
	return p >= CONSOLE && p <= DISCONNECTED
}

// IsKnown reports whether the DetectionType is a known value.
func (d DetectionType) IsKnown() bool {
	return d >= AUTOMATIC && d <= FORCE_GSMMODEM
}

// IsKnown reports whether the BaudRate is a known value.
func (b BaudRate) IsKnown() bool {
	return b >= BR1200 && b <= BR115200
}

//go:generate stringer -output=enum_strings.go -type=DetectionType,PortState,BaudRate
