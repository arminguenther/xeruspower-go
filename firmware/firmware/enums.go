// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package firmware

// IsKnown reports whether the UpdateHistoryStatus is a known value.
func (u UpdateHistoryStatus) IsKnown() bool {
	return u >= SUCCESSFUL && u <= INCOMPLETE
}

// IsKnown reports whether the ImageState is a known value.
func (i ImageState) IsKnown() bool {
	return i >= NONE && i <= COMPLETE
}

// IsKnown reports whether the UpdateFlags is a known value.
func (u UpdateFlags) IsKnown() bool {
	return u >= CROSS_OEM && u <= ALLOW_UNTRUSTED
}

//go:generate stringer -output=enum_strings.go -type=UpdateFlags,UpdateHistoryStatus,ImageState
