// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package encoding

import (
	"encoding/json"
	"testing"
)

func TestIs(t *testing.T) {
	b, err := Is[bool](true)
	if err != nil {
		t.Error(err)
	}
	if !b {
		t.Errorf("expected true")
	}
	str0 := "asdf"
	str1, err := Is[string](str0)
	if err != nil {
		t.Error(err)
	}
	if str1 != str0 {
		t.Errorf("expected %v, got %v", str0, str1)
	}
	num0 := json.Number("1")
	num1, err := Is[json.Number](num0)
	if err != nil {
		t.Error(err)
	}
	if num1 != num0 {
		t.Errorf("expected %v, got %v", num0, num1)
	}
	sl0 := []any{"foo", "bar", "ok"}
	sl1, err := Is[[]any](sl0)
	if err != nil {
		t.Error(err)
	}
	if len(sl1) != len(sl0) {
		t.Errorf("expected slice len %v, got %v", len(sl0), len(sl1))
	}
	for i := range sl0 {
		if sl0[i] != sl1[i] {
			t.Errorf("expected %v, got %v", sl0[i], sl1[i])
		}
	}
	_, err = Is[map[string]any](map[string]any{"label": "1", "on": true})
	if err != nil {
		t.Error(err)
	}
}

func TestAsByte(t *testing.T) {
	var v any
	v = json.Number("0")
	b, err := AsByte(v)
	if err != nil || b != 0 {
		t.Fatal(b, err)
	}
	v = json.Number("255")
	b, err = AsByte(v)
	if err != nil || b != 255 {
		t.Fatal(b, err)
	}
	v = json.Number("-1")
	b, err = AsByte(v)
	if err == nil {
		t.Fatalf("should fail: in=%v out=%v", v, b)
	}
	v = json.Number("256")
	b, err = AsByte(v)
	if err == nil {
		t.Fatalf("should fail: in=%v out=%v", v, b)
	}
	v = "number"
	b, err = AsByte(v)
	if err == nil {
		t.Fatalf("should fail: in=%v out=%v", v, b)
	}
}

func TestIn(t *testing.T) {
	err := In("rid", map[string]any{"rid": "/model/pdu/0"})
	if err != nil {
		t.Error(err)
	}
	err = In("type", map[string]any{"rid": "/model/pdu/0"})
	if err == nil {
		t.Error("should fail")
	}
}

func TestAsMapItems(t *testing.T) {
	keys := []json.Number{"3", "5", "1", "1"}
	values := []string{"a", "s", "d", "f"}
	v := make([]any, 4)
	for i := range 4 {
		v[i] = map[string]any{"key": keys[i], "value": values[i]}
	}
	iter, errCb, length := AsMapItems(v)
	if length != 4 {
		t.Error(length)
	}
	n := 0
	for a, b := range iter {
		if a != keys[n] {
			t.Errorf("key #%d: expected %v, got %v\n", n, keys[n], a)
		}
		if b != values[n] {
			t.Errorf("value #%d: expected %v, got %v\n", n, keys[n], b)
		}
		n++
	}
	err := errCb()
	if err != nil {
		t.Fatal(err)
	}
	v2 := make([]any, 4)
	for i := range 4 {
		var key string
		switch i {
		case 1:
			key = "kee"
		default:
			key = "key"
		}
		v2[i] = map[string]any{key: keys[i], "value": values[i]}
	}
	iter, errCb, length = AsMapItems(v2)
	if length != 4 {
		t.Error(length)
	}
	i := 0
	for a, b := range iter {
		if a != keys[i] {
			t.Errorf("key #%d: expected %v, got %v\n", i, keys[i], a)
		}
		if b != values[i] {
			t.Errorf("value #%d: expected %v, got %v\n", i, keys[i], b)
		}
		i++
	}
	if errCb() == nil {
		t.Error("expected iterator error")
	}
}

type Color int

const (
	Red Color = iota
	Green
	Blue
)

func (c Color) IsKnown() bool {
	return c >= Red && c <= Blue
}

type Shape int

const (
	Triangle Shape = iota
	Circle
	Other
	Square
)

func (s Shape) IsKnown() bool {
	return s >= Triangle && s <= Square
}

func (s Shape) Fallback() Shape {
	return Other
}

func TestAsEnum(t *testing.T) {
	for i, test := range []struct {
		number json.Number
		color  Color
	}{
		{json.Number("0"), Red},
		{json.Number("1"), Green},
		{json.Number("2"), Blue},
	} {
		color, err := AsEnum[Color](test.number)
		if err != nil {
			t.Fatal(err)
		}
		if color != test.color {
			t.Fatalf("#%d: expected %v was %v", i, test.color, color)
		}
	}
	for i, test := range []struct {
		number json.Number
		shape  Shape
	}{
		{json.Number("-1"), Other},
		{json.Number("0"), Triangle},
		{json.Number("1"), Circle},
		{json.Number("2"), Other},
		{json.Number("3"), Square},
		{json.Number("4"), Other},
	} {
		shape, err := AsEnum[Shape](test.number)
		if err != nil {
			t.Fatal(err)
		}
		if shape != test.shape {
			t.Fatalf("#%d: expected %v was %v", i, test.shape, shape)
		}
	}
}
