// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package eventengine

// IsKnown reports whether the EventType is a known value.
func (t EventType) IsKnown() bool {
	return t >= STATE && t <= TRIGGER
}

// IsKnown reports whether the EngineEventDescType is a known value.
func (t EngineEventDescType) IsKnown() bool {
	return t >= NODE && t <= LEAF
}

// IsKnown reports whether the EngineConditionOp is a known value.
func (o EngineConditionOp) IsKnown() bool {
	return o >= AND && o <= XOR
}

// IsKnown reports whether the EngineConditionMatchType is a known value.
func (m EngineConditionMatchType) IsKnown() bool {
	return m >= ASSERTED && m <= BOTH
}

//go:generate stringer -output=enum_strings.go -type=EventType,EngineConditionOp,EngineConditionMatchType,EngineEventDescType
