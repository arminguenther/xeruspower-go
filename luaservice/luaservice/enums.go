// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package luaservice

// IsKnown reports whether the ScriptStateExecState is a known value.
func (e ScriptStateExecState) IsKnown() bool {
	return e >= STAT_NEW && e <= STAT_RESTARTING
}

// IsKnown reports whether the ScriptStateExitType is a known value.
func (e ScriptStateExitType) IsKnown() bool {
	return e >= EXIT_CODE && e <= SIGNAL
}

//go:generate stringer -output=enum_strings.go -type=ScriptStateExitType,ScriptStateExecState
