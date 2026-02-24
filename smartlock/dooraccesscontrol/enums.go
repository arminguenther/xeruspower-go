// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package dooraccesscontrol

// IsKnown reports whether the DoorAccessDenialReason is a known value.
func (d DoorAccessDenialReason) IsKnown() bool {
	return d >= DENIED_NO_MATCHING_RULE && d <= DENIED_CONDITIONS_TIMEOUT
}

//go:generate stringer -output=enum_strings.go -type=DoorAccessDenialReason
