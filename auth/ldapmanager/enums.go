// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package ldapmanager

// IsKnown reports whether the ServerType is a known value.
func (s ServerType) IsKnown() bool {
	return s >= ACTIVE_DIRECTORY && s <= OPEN_LDAP
}

// IsKnown reports whether the SecurityProtocol is a known value.
func (s SecurityProtocol) IsKnown() bool {
	return s >= SEC_PROTO_NONE && s <= SEC_PROTO_STARTTLS
}

//go:generate stringer -output=enum_strings.go -type=ServerType,SecurityProtocol
