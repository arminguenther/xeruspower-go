// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package radiusmanager

// IsKnown reports whether the AuthType is a known value.
func (a AuthType) IsKnown() bool {
	return a >= PAP && a <= MSCHAPv2
}

//go:generate stringer -output=enum_strings.go -type=AuthType
