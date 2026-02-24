// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package transferswitch

// IsKnown reports whether the Type is a known value.
func (t Type) IsKnown() bool {
	return t >= STS && t <= HTS
}

// IsKnown reports whether the TransferReason is a known value.
func (t TransferReason) IsKnown() bool {
	return t >= REASON_UNKNOWN && t <= REASON_BYPASS_ACTIVE
}

// Fallback returns a known TransferReason used as a fallback if an
// unknown value is encountered while decoding.
// The fallback is [REASON_UNKNOWN].
func (t TransferReason) Fallback() TransferReason {
	return REASON_UNKNOWN
}

//go:generate stringer -output=enum_strings.go -type=Type,TransferReason
