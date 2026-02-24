// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package modbus

// IsKnown reports whether the Parity is a known value.
func (p Parity) IsKnown() bool {
	return p >= NONE && p <= ODD
}

//go:generate stringer -output=enum_strings.go -type=Parity
