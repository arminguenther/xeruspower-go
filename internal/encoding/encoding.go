// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

// Package encoding provides utility for encoding/decoding
// IDL types to/from their JSON representation.
package encoding

import (
	"encoding/json"
	"fmt"
	"iter"
	"strconv"
	"time"

	"github.com/arminguenther/xeruspower-go/v40010/idl"
)

// JSON types are those which [json.Decoder.Decode]
// stores into an interface value when decoding JSON after
// [json.Decoder.UseNumber] was called.
//
// The [json.Decoder] stores nil for JSON null.
type JSON interface {
	// true/false | string | array | object | number
	bool | string | []any | map[string]any | json.Number
}

// Is asserts that v holds a value of [JSON] type T.
func Is[T JSON](v any) (t T, err error) {
	t, ok := v.(T)
	if !ok {
		return t, fmt.Errorf("type assertion: %v is %T not %T", v, v, t)
	}
	return t, nil
}

// In returns an error if the key is not in m.
func In(key string, m map[string]any) error {
	if _, ok := m[key]; !ok {
		return fmt.Errorf("key not found: %q", key)
	}
	return nil
}

// InAndIs checks that m contains a value for the key like [In],
// and if so asserts that the value is of type T like [Is].
func InAndIs[T JSON](key string, m map[string]any) (T, error) {
	if err := In(key, m); err != nil {
		var ret T
		return ret, err
	}
	return Is[T](m[key])
}

// AsByte asserts that v is a [json.Number] and parses it as a byte.
func AsByte(v any) (byte, error) {
	n, err := Is[json.Number](v)
	if err != nil {
		return 0, err
	}
	u, err := strconv.ParseUint(string(n), 10, 8)
	return byte(u), err
}

// AsInt32 asserts that v is a [json.Number] and parses it as an int32.
func AsInt32(v any) (int32, error) {
	n, err := Is[json.Number](v)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(string(n), 10, 32)
	return int32(i), err
}

// AsInt64 asserts that v is a [json.Number] and parses it as an int64.
func AsInt64(v any) (int64, error) {
	n, err := Is[json.Number](v)
	if err != nil {
		return 0, err
	}
	return n.Int64()
}

// AsInt asserts that v is a [json.Number] and parses it as an int.
func AsInt(v any) (int, error) {
	n, err := Is[json.Number](v)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(n))
}

// AsFloat32 asserts that v is a [json.Number] and parses it as a float32.
func AsFloat32(v any) (float32, error) {
	n, err := Is[json.Number](v)
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseFloat(string(n), 32)
	return float32(f), err
}

// AsFloat64 asserts that v is a [json.Number] and parses it as a float64.
func AsFloat64(v any) (float64, error) {
	n, err := Is[json.Number](v)
	if err != nil {
		return 0, err
	}
	return n.Float64()
}

// AsTime asserts that v is a [json.Number] and parses it as a [time.Time].
// The number is interpreted as seconds since the Unix epoch.
func AsTime(v any) (time.Time, error) {
	sec, err := AsInt64(v)
	return time.Unix(sec, 0), err
}

// Enum describes an enumeration type.
type Enum interface {
	~int
	IsKnown() bool
}

// A Fallbacker has a known Enum value as fallback.
type Fallbacker[E Enum] interface {
	Fallback() E
}

// AsEnum asserts that v is a [json.Number] and parses it as [Enum] type E.
// For unknown values, a fallback is returned if E is a [Fallbacker].
func AsEnum[E Enum](v any) (E, error) {
	var e E
	i, err := AsInt(v)
	if err != nil {
		return e, err
	}
	e = E(i)
	if e.IsKnown() {
		return e, nil
	}
	if f, ok := any(e).(Fallbacker[E]); ok {
		return f.Fallback(), nil
	}
	return e, fmt.Errorf("unknown enum value: %v", e)
}

// AsTypeCode asserts that v is a string and parses it as an [idl.TypeCode].
func AsTypeCode(v any) (idl.TypeCode, error) {
	s, err := Is[string](v)
	if err != nil {
		return idl.TypeCode{}, err
	}
	code, ok := idl.ParseTypeCode(s)
	if !ok {
		return code, fmt.Errorf("invalid typecode syntax: %q", s)
	}
	return code, nil
}

// AsMapItems asserts that v is a container slice of map items with "key" and "value".
// It returns an iterator over those key-value pairs that stops on the first error,
// a callback to retrieve this error, and the slice length.
func AsMapItems(v any) (items iter.Seq2[any, any], errFn func() error, length int) {
	slice, err := Is[[]any](v)
	items = func(yield func(any, any) bool) {
		if err != nil {
			return
		}
		var m map[string]any
		for i := range slice {
			m, err = Is[map[string]any](slice[i])
			if err != nil {
				return
			}
			err = In("key", m)
			if err != nil {
				return
			}
			err = In("value", m)
			if err != nil {
				return
			}
			if !yield(m["key"], m["value"]) {
				return
			}
		}
	}
	return items, func() error { return err }, len(slice)
}

// SimpleEncoding types are those of which non-nil slices can
// be passed to [json.Marshal] without further encoding work.
//
// Note that byte is not included, as [json.Marshal] has
// special handling for []byte.
type SimpleEncoding interface {
	// boolean | string | int | long | float | double | enumeration
	bool | string | int32 | int64 | float32 | float64 | ~int
}

// NonNilSlice returns the passed slice as is if non-nil.
// Else a new empty slice of same type is returned.
// This forces [json.Marshal] to encode it as a JSON array instead of as JSON null.
func NonNilSlice[T SimpleEncoding](slice []T) []T {
	if slice == nil {
		return make([]T, 0)
	}
	return slice
}
