// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package transferswitch

// IsKnown reports whether the Type is a known value.
func (t Type) IsKnown() bool {
	return t >= STS && t <= HTS
}

// IsKnown reports whether the TransferReason is a known value.
func (t TransferReason) IsKnown() bool {
	return t >= REASON_UNKNOWN && t <= REASON_INTERNAL_FAILURE
}

//go:generate stringer -output=enum_strings.go -type=Type,TransferReason
