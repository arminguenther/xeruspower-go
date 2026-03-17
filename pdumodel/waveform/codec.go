// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package waveform

import (
	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding"
)

func (w *Waveform) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["voltage"] = encoding.NonNilSlice(w.Voltage)
	j0["current"] = encoding.NonNilSlice(w.Current)
	j0["sampleRate"] = w.SampleRate
	return j0
}

func (w *Waveform) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("voltage", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["voltage"])
	if err != nil {
		return err
	}
	w.Voltage = make([]float64, 0, len(s1))
	for _, a1 := range s1 {
		var e1 float64
		e1, err = encoding.AsFloat64(a1)
		if err != nil {
			return err
		}
		w.Voltage = append(w.Voltage, e1)
	}
	err = encoding.In("current", j0)
	if err != nil {
		return err
	}
	var s2 []any
	s2, err = encoding.Is[[]any](j0["current"])
	if err != nil {
		return err
	}
	w.Current = make([]float64, 0, len(s2))
	for _, a2 := range s2 {
		var e2 float64
		e2, err = encoding.AsFloat64(a2)
		if err != nil {
			return err
		}
		w.Current = append(w.Current, e2)
	}
	err = encoding.In("sampleRate", j0)
	if err != nil {
		return err
	}
	w.SampleRate, err = encoding.AsInt32(j0["sampleRate"])
	if err != nil {
		return err
	}
	return nil
}
