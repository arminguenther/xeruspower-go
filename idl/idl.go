// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

// Package idl defines some of the inherent IDL types.
// Remaining ones are mapped to standard Go types:
//
//	IDL               Go
//	boolean           bool
//	string            string
//	byte              byte
//	int               int32
//	long              int64
//	float             float32
//	double            float64
//	vector<T>         []T
//	map<K, V>         map[K]V
//	time              time.Time
//	typecode          idl.TypeCode
//	Object            idl.Object
//	_                 idl.ValueObject
package idl
