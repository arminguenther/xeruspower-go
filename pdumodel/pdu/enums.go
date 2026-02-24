// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package pdu

// IsKnown reports whether the Orientation is a known value.
func (p Orientation) IsKnown() bool {
	return p >= PO_NONE && p <= PO_TOPFEED
}

// IsKnown reports whether the StartupState is a known value.
func (s StartupState) IsKnown() bool {
	return s >= SS_ON && s <= SS_LASTKNOWN
}

//go:generate stringer -output=enum_strings.go -type=Orientation,StartupState
