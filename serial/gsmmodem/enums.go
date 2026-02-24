// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package gsmmodem

// IsKnown reports whether the SimSecurityStatus is a known value.
func (s SimSecurityStatus) IsKnown() bool {
	return s >= UNLOCKED && s <= UNKNOWN
}

//go:generate stringer -output=enum_strings.go -type=SimSecurityStatus
