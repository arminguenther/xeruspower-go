// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package assetstrip

// IsKnown reports whether the State is a known value.
func (s State) IsKnown() bool {
	return s >= DISCONNECTED && s <= AVAILABLE
}

// IsKnown reports whether the TagType is a known value.
func (t TagType) IsKnown() bool {
	return t >= SINGLE && t <= NONE
}

// IsKnown reports whether the CascadeState is a known value.
func (c CascadeState) IsKnown() bool {
	return c >= CASCADE_ACTIVE && c <= CASCADE_FIRMWARE_UPDATE
}

// IsKnown reports whether the FirmwareUpdateState is a known value.
func (f FirmwareUpdateState) IsKnown() bool {
	return f >= UPDATE_STARTED && f <= UPDATE_FAILED
}

//go:generate stringer -output=enum_strings.go -type=TagType,CascadeState,State,FirmwareUpdateState
