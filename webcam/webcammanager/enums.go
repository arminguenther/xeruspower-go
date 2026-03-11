// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package webcammanager

// IsKnown reports whether the Priority is a known value.
func (p Priority) IsKnown() bool {
	return p >= VERY_LOW && p <= VERY_HIGH
}

//go:generate stringer -output=enum_strings.go -type=Priority
