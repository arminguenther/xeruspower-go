// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package services

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/object"
)

func init() {
	object.Register(NewServices)
}

type _Services struct {
	idl.Object
}

// NewServices returns a new Services interface for the object at given RID.
func NewServices(rid string, caller idl.Caller) Services {
	return &_Services{idl.NewObject(rid, caller)}
}

func (s *_Services) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "net.Services",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_Services) GetSettings(ctx context.Context) ([]ServiceSettings, error) {
	var out0 []ServiceSettings
	val, err := s.Caller().Call(ctx, s.RID(), "getSettings", nil)
	if err != nil {
		return out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return out0, err
	}
	err = encoding.In("servicesSettings", res)
	if err != nil {
		return out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["servicesSettings"])
	if err != nil {
		return out0, err
	}
	out0 = make([]ServiceSettings, 0, len(s0))
	for _, a0 := range s0 {
		var e0 ServiceSettings
		err = e0.Decode(a0, s.Caller())
		if err != nil {
			return out0, err
		}
		out0 = append(out0, e0)
	}
	return out0, nil
}

func (s *_Services) SetSettings(ctx context.Context, in0 []ServiceSettings) (int32, error) {
	var ret int32
	s0 := make([]any, 0, len(in0))
	for _, e0 := range in0 {
		s0 = append(s0, e0.Encode())
	}
	val, err := s.Caller().Call(ctx, s.RID(), "setSettings", map[string]any{
		"servicesSettings": s0,
	})
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}
