// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package smtp

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding/object"
)

func init() {
	object.Register(NewSmtp)
}

type _Smtp struct {
	idl.Object
}

// NewSmtp returns a new Smtp interface for the object at given RID.
func NewSmtp(rid string, caller idl.Caller) Smtp {
	return &_Smtp{idl.NewObject(rid, caller)}
}

func (s *_Smtp) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "devsettings.Smtp",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (s *_Smtp) GetConfiguration(ctx context.Context) (Configuration, error) {
	var ret Configuration
	val, err := s.Caller().Call(ctx, s.RID(), "getConfiguration", nil)
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

func (s *_Smtp) SetConfiguration(ctx context.Context, in0 Configuration) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setConfiguration", map[string]any{
		"cfg": in0.Encode(),
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

func (s *_Smtp) TestConfiguration(ctx context.Context, in0 Configuration, in1 []string) (TestResult, error) {
	var ret TestResult
	val, err := s.Caller().Call(ctx, s.RID(), "testConfiguration", map[string]any{
		"cfg":        in0.Encode(),
		"recipients": encoding.NonNilSlice(in1),
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
	err = ret.Decode(res["_ret_"], s.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}
