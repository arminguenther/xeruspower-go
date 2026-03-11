// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package scep

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40411/cert/serversslcert"
	"github.com/arminguenther/xeruspower-go/v40411/idl"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding/object"
)

func init() {
	object.Register(NewScep)
}

type _Scep struct {
	idl.Object
}

// NewScep returns a new Scep interface for the object at given RID.
func NewScep(rid string, caller idl.Caller) Scep {
	return &_Scep{idl.NewObject(rid, caller)}
}

func (s *_Scep) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "cert.Scep",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_Scep) EnableScep(ctx context.Context, in0 EnrollmentData, in1 Settings) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "enableScep", map[string]any{
		"initialEnrollmentData": in0.Encode(),
		"scepSettings":          in1.Encode(),
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

func (s *_Scep) DiscardScep(ctx context.Context) error {
	_, err := s.Caller().Call(ctx, s.RID(), "discardScep", nil)
	return err
}

func (s *_Scep) GetInitialEnrollmentData(ctx context.Context) (EnrollmentData, error) {
	var ret EnrollmentData
	val, err := s.Caller().Call(ctx, s.RID(), "getInitialEnrollmentData", nil)
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
	err = ret.Decode(res["_ret_"], s.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_Scep) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := s.Caller().Call(ctx, s.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], s.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_Scep) GetStatus(ctx context.Context) (Status, error) {
	var ret Status
	val, err := s.Caller().Call(ctx, s.RID(), "getStatus", nil)
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
	err = ret.Decode(res["_ret_"], s.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_Scep) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setSettings", map[string]any{
		"scepSettings": in0.Encode(),
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

func (s *_Scep) GetSupportedKeyInfos(ctx context.Context) ([]serversslcert.KeyInfo, error) {
	var ret []serversslcert.KeyInfo
	val, err := s.Caller().Call(ctx, s.RID(), "getSupportedKeyInfos", nil)
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
	var s0 []any
	s0, err = encoding.Is[[]any](res["_ret_"])
	if err != nil {
		return ret, err
	}
	ret = make([]serversslcert.KeyInfo, 0, len(s0))
	for _, a0 := range s0 {
		var e0 serversslcert.KeyInfo
		err = e0.Decode(a0, s.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}
