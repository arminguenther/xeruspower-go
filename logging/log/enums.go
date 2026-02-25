// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package log

// IsKnown reports whether the RangeDirection is a known value.
func (r RangeDirection) IsKnown() bool {
	return r >= FORWARD && r <= BACKWARD
}

//go:generate stringer -output=enum_strings.go -type=RangeDirection
