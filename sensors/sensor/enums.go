// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package sensor

// IsKnown reports whether the OnOffState is a known value.
func (o OnOffState) IsKnown() bool {
	return o >= OFF && o <= ON
}

// IsKnown reports whether the OpenClosedState is a known value.
func (o OpenClosedState) IsKnown() bool {
	return o >= OPEN && o <= CLOSED
}

// IsKnown reports whether the NormalAlarmedState is a known value.
func (n NormalAlarmedState) IsKnown() bool {
	return n >= NORMAL && n <= ALARMED
}

// IsKnown reports whether the OkFaultState is a known value.
func (o OkFaultState) IsKnown() bool {
	return o >= OK && o <= FAULT
}

//go:generate stringer -output=enum_strings.go -type=OnOffState,OpenClosedState,NormalAlarmedState,OkFaultState
