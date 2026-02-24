// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package authmanager

// IsKnown reports whether the Type is a known value.
func (t Type) IsKnown() bool {
	return t >= LOCAL && t <= LDAP
}

//go:generate stringer -output=enum_strings.go -type=Type
