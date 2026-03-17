// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package tacplusmanager

// IsKnown reports whether the AuthenType is a known value.
func (a AuthenType) IsKnown() bool {
	return a >= ASCII && a <= MSCHAP
}

//go:generate stringer -output=enum_strings.go -type=AuthenType
