// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package circuit

// IsKnown reports whether the Type is a known value.
func (t Type) IsKnown() bool {
	return t >= ONE_PHASE_LN && t <= THREE_PHASE
}

//go:generate stringer -output=enum_strings.go -type=Type
