// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package overcurrentprotector

// IsKnown reports whether the Type is a known value.
func (t Type) IsKnown() bool {
	return t >= BREAKER_1POLE && t <= RCBO_4POLE
}

//go:generate stringer -output=enum_strings.go -type=Type
