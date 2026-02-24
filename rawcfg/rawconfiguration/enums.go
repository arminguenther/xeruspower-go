// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package rawconfiguration

// IsKnown reports whether the Status is a known value.
func (s Status) IsKnown() bool {
	return s >= UNKNOWN && s <= UPDATE_FAILED
}

//go:generate stringer -output=enum_strings.go -type=Status
