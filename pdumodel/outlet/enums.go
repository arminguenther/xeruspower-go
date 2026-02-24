// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package outlet

// IsKnown reports whether the PowerState is a known value.
func (p PowerState) IsKnown() bool {
	return p >= PS_OFF && p <= PS_ON
}

// IsKnown reports whether the StartupState is a known value.
func (s StartupState) IsKnown() bool {
	return s >= SS_ON && s <= SS_PDUDEF
}

//go:generate stringer -output=enum_strings.go -type=StartupState,PowerState
