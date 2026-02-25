// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package security

// IsKnown reports whether the IpfwPolicy is a known value.
func (i IpfwPolicy) IsKnown() bool {
	return i >= ACCEPT && i <= REJECT
}

// IsKnown reports whether the RoleAccessPolicy is a known value.
func (r RoleAccessPolicy) IsKnown() bool {
	return r >= ALLOW && r <= DENY
}

// IsKnown reports whether the SSHHostKeyType is a known value.
func (s SSHHostKeyType) IsKnown() bool {
	return s >= SSH_HOST_KEY_TYPE_RSA && s <= SSH_HOST_KEY_TYPE_ED25519
}

// IsKnown reports whether the SSHKeyFingerprintType is a known value.
func (s SSHKeyFingerprintType) IsKnown() bool {
	return s >= SSH_KEY_FPRINT_TYPE_MD5_HEX && s <= SSH_KEY_FPRINT_TYPE_UNKNOWN
}

//go:generate stringer -output=enum_strings.go -type=SSHHostKeyType,SSHKeyFingerprintType,IpfwPolicy,RoleAccessPolicy
