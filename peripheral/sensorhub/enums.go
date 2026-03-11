// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package sensorhub

// IsKnown reports whether the UpstreamType is a known value.
func (u UpstreamType) IsKnown() bool {
	return u >= BUILTIN && u <= REMOTE_HUB
}

//go:generate stringer -output=enum_strings.go -type=UpstreamType
