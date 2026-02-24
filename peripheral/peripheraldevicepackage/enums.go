// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package peripheraldevicepackage

// IsKnown reports whether the PackageInfoState is a known value.
func (s PackageInfoState) IsKnown() bool {
	return s >= NORMAL && s <= CONFIG_ERROR
}

//go:generate stringer -output=enum_strings.go -type=PackageInfoState
