// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package portfuse

// IsKnown reports whether the Status is a known value.
func (s Status) IsKnown() bool {
	return s >= UNKNOWN && s <= GOOD
}

//go:generate stringer -output=enum_strings.go -type=Status
