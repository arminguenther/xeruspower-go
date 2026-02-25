// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package datetime

// IsKnown reports whether the Protocol is a known value.
func (p Protocol) IsKnown() bool {
	return p >= STATIC && p <= NTP
}

//go:generate stringer -output=enum_strings.go -type=Protocol
