// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package datapushservice

// IsKnown reports whether the EntryType is a known value.
func (e EntryType) IsKnown() bool {
	return e >= SENSORLIST && e <= AUDITLOG
}

//go:generate stringer -output=enum_strings.go -type=EntryType
