// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package resmon

// IsKnown reports whether the EntryType is a known value.
func (t EntryType) IsKnown() bool {
	return t >= GLOBAL_CPU_USAGE && t <= PROC_COUNT
}

//go:generate stringer -output=enum_strings.go -type=EntryType
