// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package zigbeemanager

// IsKnown reports whether the DongleState is a known value.
func (d DongleState) IsKnown() bool {
	return d >= NOT_INIT && d <= IN_MODEM_STATE
}

//go:generate stringer -output=enum_strings.go -type=DongleState
