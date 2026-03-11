// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package servermonitor

// IsKnown reports whether the ServerReachability is a known value.
func (s ServerReachability) IsKnown() bool {
	return s >= WAITING && s <= ERROR
}

// IsKnown reports whether the ServerPowerState is a known value.
func (s ServerPowerState) IsKnown() bool {
	return s >= UNKNOWN && s <= SHUTTING_DOWN
}

// IsKnown reports whether the ServerPowerControlResult is a known value.
func (s ServerPowerControlResult) IsKnown() bool {
	return s >= NO_ERROR && s <= POWER_CHECK_TIMEOUT
}

// IsKnown reports whether the ServerPowerCheckMethod is a known value.
func (s ServerPowerCheckMethod) IsKnown() bool {
	return s >= TIMER && s <= POWER_DROP
}

//go:generate stringer -output=enum_strings.go -type=ServerReachability,ServerPowerState,ServerPowerCheckMethod,ServerPowerControlResult
