// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package portforwardingautosetup

// IsKnown reports whether the ExpansionUnitState is a known value.
func (p ExpansionUnitState) IsKnown() bool {
	return p >= PF_AUTO_SETUP_UNIT_ALREADY_EXISTS && p <= PF_AUTO_SETUP_UNIT_ERROR
}

// IsKnown reports whether the RunningState is a known value.
func (p RunningState) IsKnown() bool {
	return p >= PF_AUTO_SETUP_NONE && p <= PF_AUTO_SETUP_FINISHED_SUCCESSFULLY
}

//go:generate stringer -output=enum_strings.go -type=RunningState,ExpansionUnitState
