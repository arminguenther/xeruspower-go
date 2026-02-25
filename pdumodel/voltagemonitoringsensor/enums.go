// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package voltagemonitoringsensor

// IsKnown reports whether the EventType is a known value.
func (e EventType) IsKnown() bool {
	return e >= DIP && e <= UNKNOWN
}

// Fallback returns a known EventType used as a fallback if an
// unknown value is encountered while decoding.
// The fallback is [UNKNOWN].
func (e EventType) Fallback() EventType {
	return UNKNOWN
}

//go:generate stringer -output=enum_strings.go -type=EventType
