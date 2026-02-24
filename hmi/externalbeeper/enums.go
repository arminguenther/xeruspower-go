// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package externalbeeper

// IsKnown reports whether the State is a known value.
func (s State) IsKnown() bool {
	return s >= OFF && s <= ALARMING
}

//go:generate stringer -output=enum_strings.go -type=State
