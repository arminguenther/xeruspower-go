// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package scep

// IsKnown reports whether the EnrollmentRunStatus is a known value.
func (e EnrollmentRunStatus) IsKnown() bool {
	return e >= NOT_RUN && e <= FAILED
}

//go:generate stringer -output=enum_strings.go -type=EnrollmentRunStatus
