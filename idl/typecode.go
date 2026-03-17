// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package idl

import (
	"fmt"
	"strconv"
	"strings"
)

// TypeCode describes an IDL type.
type TypeCode struct {
	Name     string // Name of the type
	Major    int    // Major version
	Submajor int    // Submajor version
	Minor    int    // Minor version
}

// ParseTypeCode parses a string-encoded TypeCode.
// The expected pattern of the string is:
//
//	Name:Major.Submajor.Minor
func ParseTypeCode(s string) (code TypeCode, ok bool) {
	name, rem, _ := strings.Cut(s, ":")
	subs := strings.Split(rem, ".")
	if len(subs) != 3 {
		return code, false
	}
	var major, submajor, minor int
	var err error
	for i, ver := range []*int{&major, &submajor, &minor} {
		*ver, err = strconv.Atoi(subs[i])
		if err != nil {
			return code, false
		}
	}
	return TypeCode{name, major, submajor, minor}, true
}

// String returns the TypeCode as a string.
// It matches the pattern expected by [ParseTypeCode].
func (t TypeCode) String() string {
	return fmt.Sprintf("%s:%d.%d.%d", t.Name, t.Major, t.Submajor, t.Minor)
}

// StringNoVersions returns the TypeCode as a string like [TypeCode.String],
// but without version suffixes on the type itself and all scoping types.
// See [ParseTypeCode] and [TypeCode.SegmentCodes] for version suffix patterns.
func (t TypeCode) StringNoVersions() string {
	var subs []string
	for _, code := range t.SegmentCodes() {
		subs = append(subs, code.Name)
	}
	return strings.Join(subs, ".")
}

// IsCompatible reports whether a type version (server) is compatible with another (client).
// Identical, call compatible and backward compatible versions count as compatible.
//
// A version is compatible with another if their Majors are equal and
// its Minor is equal to or greater than the other's.
// Scoping types must be compatible respectively.
func (t TypeCode) IsCompatible(client TypeCode) bool {
	codes := t.SegmentCodes()
	others := client.SegmentCodes()
	if len(codes) != len(others) {
		return false
	}
	for i := range codes {
		if codes[i].Name == others[i].Name &&
			codes[i].Major == others[i].Major &&
			// any Submajor
			codes[i].Minor >= others[i].Minor {
			continue
		}
		return false
	}
	return true
}

// SegmentCodes splits Name into segments delimited by dots,
// parses them and returns a slice of one TypeCode per segment.
// Major, Submajor and Minor apply to the last segment.
// Each TypeCode.Name will be unqualified.
//
// Segments represent IDL modules or types.
// If a segment does not specify its version, it is 1.0.0.
// Scoping versioned types use following pattern:
//
//	Name_Major_Submajor_Minor
func (t TypeCode) SegmentCodes() []TypeCode {
	parse := func(segment string) TypeCode {
		code := TypeCode{segment, 1, 0, 0}
		subs := strings.Split(segment, "_")
		if len(subs) < 4 {
			return code
		}
		var major, submajor, minor int
		var err error
		for i, ver := range []*int{&minor, &submajor, &major} {
			*ver, err = strconv.Atoi(subs[len(subs)-1-i])
			if err != nil {
				return code
			}
		}
		name := strings.Join(subs[:len(subs)-3], "_")
		return TypeCode{name, major, submajor, minor}
	}
	subs := strings.Split(t.Name, ".")
	codes := make([]TypeCode, 0, len(subs))
	for _, seg := range subs[:len(subs)-1] {
		codes = append(codes, parse(seg))
	}
	self := TypeCode{subs[len(subs)-1], t.Major, t.Submajor, t.Minor}
	return append(codes, self)
}
