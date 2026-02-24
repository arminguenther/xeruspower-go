// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package bluetooth

// IsKnown reports whether the Antenna is a known value.
func (a Antenna) IsKnown() bool {
	return a >= Internal && a <= External
}

//go:generate stringer -output=enum_strings.go -type=Antenna
