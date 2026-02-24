// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package assetstriplogger

// IsKnown reports whether the RecordType is a known value.
func (r RecordType) IsKnown() bool {
	return r >= EMPTY && r <= ASSET_STRIP_STATE_CHANGED
}

//go:generate stringer -output=enum_strings.go -type=RecordType
