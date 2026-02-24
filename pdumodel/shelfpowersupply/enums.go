// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package shelfpowersupply

// IsKnown reports whether the PowerState is a known value.
func (p PowerState) IsKnown() bool {
	return p >= PS_OFF && p <= PS_ON
}

//go:generate stringer -output=enum_strings.go -type=PowerState
