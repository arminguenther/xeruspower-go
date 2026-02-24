// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package serversslcert

// IsKnown reports whether the KeyType is a known value.
func (k KeyType) IsKnown() bool {
	return k >= KEY_TYPE_UNKNOWN && k <= KEY_TYPE_ECDSA
}

// IsKnown reports whether the EllipticCurve is a known value.
func (e EllipticCurve) IsKnown() bool {
	return e >= EC_CURVE_UNKNOWN && e <= EC_CURVE_NIST_P521
}

//go:generate stringer -output=enum_strings.go -type=KeyType,EllipticCurve
