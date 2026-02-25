// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package panel

// IsKnown reports whether the LabelingScheme is a known value.
func (l LabelingScheme) IsKnown() bool {
	return l >= SEQUENTIAL && l <= ODD_EVEN
}

//go:generate stringer -output=enum_strings.go -type=LabelingScheme
