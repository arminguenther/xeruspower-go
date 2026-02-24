// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package pole

// IsKnown reports whether the PowerLine is a known value.
func (p PowerLine) IsKnown() bool {
	return p >= L1 && p <= MINUS
}

//go:generate stringer -output=enum_strings.go -type=PowerLine
