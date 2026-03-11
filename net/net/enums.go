// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package net

// IsKnown reports whether the PortForwardingRole is a known value.
func (p PortForwardingRole) IsKnown() bool {
	return p >= PRIMARY_UNIT && p <= EXPANSION_UNIT
}

// IsKnown reports whether the IpConfigMethod is a known value.
func (i IpConfigMethod) IsKnown() bool {
	return i >= STATIC && i <= AUTO
}

// IsKnown reports whether the InterfaceType is a known value.
func (i InterfaceType) IsKnown() bool {
	return i >= ETHERNET && i <= BRIDGE
}

// IsKnown reports whether the InterfaceOpState is a known value.
func (i InterfaceOpState) IsKnown() bool {
	return i >= NOT_PRESENT && i <= UP
}

// IsKnown reports whether the EapOuterAuthMethod is a known value.
func (e EapOuterAuthMethod) IsKnown() bool {
	return e >= EAP_PEAP && e <= EAP_TLS
}

// IsKnown reports whether the EapInnerAuthMethod is a known value.
func (e EapInnerAuthMethod) IsKnown() bool {
	return e >= INNER_EAP_MSCHAPv2 && e <= INNER_EAP_TLS
}

// IsKnown reports whether the EapStatus is a known value.
func (e EapStatus) IsKnown() bool {
	return e >= EAP_STATUS_DISABLED && e <= EAP_STATUS_SUCCESS
}

// IsKnown reports whether the EthSpeed is a known value.
func (e EthSpeed) IsKnown() bool {
	return e >= SPEED_AUTO && e <= SPEED_MBIT_1000
}

// IsKnown reports whether the EthDuplexMode is a known value.
func (e EthDuplexMode) IsKnown() bool {
	return e >= DUPLEX_MODE_AUTO && e <= DUPLEX_MODE_FULL
}

// IsKnown reports whether the EthAuthType is a known value.
func (e EthAuthType) IsKnown() bool {
	return e >= ETH_AUTH_NONE && e <= ETH_AUTH_EAP
}

// IsKnown reports whether the WlanSecProtocol is a known value.
func (w WlanSecProtocol) IsKnown() bool {
	return w >= WPA2 && w <= WPA2
}

// IsKnown reports whether the WlanAuthType is a known value.
func (w WlanAuthType) IsKnown() bool {
	return w >= WLAN_AUTH_NONE && w <= WLAN_AUTH_EAP
}

// IsKnown reports whether the WlanChannelWidth is a known value.
func (w WlanChannelWidth) IsKnown() bool {
	return w >= CHANNEL_WIDTH_UNKNOWN && w <= CHANNEL_WIDTH_160
}

//go:generate stringer -output=enum_strings.go -type=EthDuplexMode,PortForwardingRole,EthAuthType,EapInnerAuthMethod,WlanAuthType,EthSpeed,InterfaceType,EapStatus,InterfaceOpState,EapOuterAuthMethod,IpConfigMethod,WlanChannelWidth,WlanSecProtocol
