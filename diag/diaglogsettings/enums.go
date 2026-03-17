// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package diaglogsettings

// IsKnown reports whether the LogLevel is a known value.
func (l LogLevel) IsKnown() bool {
	return l >= LOG_LEVEL_NONE && l <= LOG_LEVEL_TRACE
}

//go:generate stringer -output=enum_strings.go -type=LogLevel
