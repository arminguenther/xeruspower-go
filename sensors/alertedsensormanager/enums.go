// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package alertedsensormanager

// IsKnown reports whether the AlertState is a known value.
func (a AlertState) IsKnown() bool {
	return a >= UNAVAILABLE && a <= WARNED
}

//go:generate stringer -output=enum_strings.go -type=AlertState
