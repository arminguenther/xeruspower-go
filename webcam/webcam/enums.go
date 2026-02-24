// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package webcam

// IsKnown reports whether the PixelFormat is a known value.
func (p PixelFormat) IsKnown() bool {
	return p >= MJPEG && p <= YUV
}

// IsKnown reports whether the PowerLineFrequency is a known value.
func (p PowerLineFrequency) IsKnown() bool {
	return p >= NOT_SUPPORTED && p <= DISABLED
}

//go:generate stringer -output=enum_strings.go -type=PowerLineFrequency,PixelFormat
