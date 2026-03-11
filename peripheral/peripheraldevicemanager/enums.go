// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package peripheraldevicemanager

// IsKnown reports whether the DeviceManagerZCoordMode is a known value.
func (z DeviceManagerZCoordMode) IsKnown() bool {
	return z >= RACKUNITS && z <= FREEFORM
}

// IsKnown reports whether the DeviceManagerDeviceFirmwareUpdateState is a known value.
func (d DeviceManagerDeviceFirmwareUpdateState) IsKnown() bool {
	return d >= UPDATE_STARTED && d <= UPDATE_FAILED
}

//go:generate stringer -output=enum_strings.go -type=DeviceManagerDeviceFirmwareUpdateState,DeviceManagerZCoordMode
