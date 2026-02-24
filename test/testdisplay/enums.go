// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package testdisplay

// IsKnown reports whether the DisplayOrientation is a known value.
func (o DisplayOrientation) IsKnown() bool {
	return o >= NORMAL && o <= RIGHT
}

// IsKnown reports whether the DisplayTestStatus is a known value.
func (t DisplayTestStatus) IsKnown() bool {
	return t >= TEST_IDLE && t <= TEST_FAILED
}

//go:generate stringer -output=enum_strings.go -type=DisplayTestStatus,DisplayOrientation
