// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package internalbeeper

// IsKnown reports whether the State is a known value.
func (s State) IsKnown() bool {
	return s >= OFF && s <= ON_ACTIVATION
}

//go:generate stringer -output=enum_strings.go -type=State
