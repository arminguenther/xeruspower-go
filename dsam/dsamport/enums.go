// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package dsamport

// IsKnown reports whether the DeviceInterfaceType is a known value.
func (d DeviceInterfaceType) IsKnown() bool {
	return d >= DEV_IFTYPE_AUTO && d <= DEV_IFTYPE_DCE
}

// IsKnown reports whether the Parity is a known value.
func (p Parity) IsKnown() bool {
	return p >= PARITY_NONE && p <= PARITY_EVEN
}

// IsKnown reports whether the FlowControl is a known value.
func (f FlowControl) IsKnown() bool {
	return f >= FLOW_CTRL_NONE && f <= FLOW_CTRL_SOFTWARE
}

// IsKnown reports whether the State is a known value.
func (s State) IsKnown() bool {
	return s >= STATE_AVAILABLE && s <= STATE_BUSY
}

//go:generate stringer -output=enum_strings.go -type=DeviceInterfaceType,FlowControl,State,Parity
