// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package bulkconfiguration

// IsKnown reports whether the Status is a known value.
func (s Status) IsKnown() bool {
	return s >= UNKNOWN && s <= RESTORE_FAILED
}

// IsKnown reports whether the FilterType is a known value.
func (f FilterType) IsKnown() bool {
	return f >= WHITELIST && f <= BLACKLIST
}

//go:generate stringer -output=enum_strings.go -type=Status,FilterType
