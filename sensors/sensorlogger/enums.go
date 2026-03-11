// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package sensorlogger

// IsKnown reports whether the LoggerBackupState is a known value.
func (b LoggerBackupState) IsKnown() bool {
	return b >= STOPPED && b <= SUSPENDED
}

//go:generate stringer -output=enum_strings.go -type=LoggerBackupState
