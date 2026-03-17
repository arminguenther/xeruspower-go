// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package port

// IsKnown reports whether the DetectionType is a known value.
func (d DetectionType) IsKnown() bool {
	return d >= AUTO && d <= DISABLED
}

//go:generate stringer -output=enum_strings.go -type=DetectionType
