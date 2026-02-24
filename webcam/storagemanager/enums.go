// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package storagemanager

// IsKnown reports whether the StorageType is a known value.
func (s StorageType) IsKnown() bool {
	return s >= LOCAL && s <= NFS
}

// IsKnown reports whether the Direction is a known value.
func (d Direction) IsKnown() bool {
	return d >= ASCENDING && d <= DESCENDING
}

// IsKnown reports whether the StorageStatus is a known value.
func (s StorageStatus) IsKnown() bool {
	return s >= INITIALIZING && s <= READY
}

//go:generate stringer -output=enum_strings.go -type=StorageType,StorageStatus,Direction
