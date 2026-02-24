// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package alertedsensormanager

// IsKnown reports whether the AlertState is a known value.
func (a AlertState) IsKnown() bool {
	return a >= UNAVAILABLE && a <= WARNED
}

//go:generate stringer -output=enum_strings.go -type=AlertState
