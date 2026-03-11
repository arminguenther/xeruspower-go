// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package sessionmanager

// IsKnown reports whether the CloseReason is a known value.
func (c CloseReason) IsKnown() bool {
	return c >= CLOSE_REASON_LOGOUT && c <= CLOSE_REASON_FORCED_DISCONNECT
}

//go:generate stringer -output=enum_strings.go -type=CloseReason
