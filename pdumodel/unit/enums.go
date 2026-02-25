// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package unit

// IsKnown reports whether the Orientation is a known value.
func (o Orientation) IsKnown() bool {
	return o >= NORMAL && o <= FLIPPED
}

//go:generate stringer -output=enum_strings.go -type=Orientation
