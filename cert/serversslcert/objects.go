// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package serversslcert

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40411/idl"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding/object"
)

func init() {
	object.Register(NewServerSSLCert)
}

type _ServerSSLCert struct {
	idl.Object
}

// NewServerSSLCert returns a new ServerSSLCert interface for the object at given RID.
func NewServerSSLCert(rid string, caller idl.Caller) ServerSSLCert {
	return &_ServerSSLCert{idl.NewObject(rid, caller)}
}

func (s *_ServerSSLCert) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "cert.ServerSSLCert",
		Major: 4, Submajor: 0, Minor: 1,
	}
}

func (s *_ServerSSLCert) GetSupportedKeyInfos(ctx context.Context) ([]KeyInfo, error) {
	var ret []KeyInfo
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
	ret = make([]KeyInfo, 0, len(s0))
	for _, a0 := range s0 {
		var e0 KeyInfo
		err = e0.Decode(a0, s.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (s *_ServerSSLCert) GenerateUnsignedKeyPair(ctx context.Context, in0 ReqInfo, in1 string) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "generateUnsignedKeyPair", map[string]any{
		"reqInfo":   in0.Encode(),
		"challenge": in1,
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

func (s *_ServerSSLCert) GenerateSelfSignedKeyPair(ctx context.Context, in0 ReqInfo, in1 int32) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "generateSelfSignedKeyPair", map[string]any{
		"reqInfo": in0.Encode(),
		"days":    in1,
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

func (s *_ServerSSLCert) DeletePending(ctx context.Context) error {
	_, err := s.Caller().Call(ctx, s.RID(), "deletePending", nil)
	return err
}

func (s *_ServerSSLCert) GetInfo(ctx context.Context) (Info, error) {
	var out0 Info
	val, err := s.Caller().Call(ctx, s.RID(), "getInfo", nil)
	if err != nil {
		return out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return out0, err
	}
	err = encoding.In("info", res)
	if err != nil {
		return out0, err
	}
	err = out0.Decode(res["info"], s.Caller())
	if err != nil {
		return out0, err
	}
	return out0, nil
}

func (s *_ServerSSLCert) GetActiveCertChainPEM(ctx context.Context) (string, error) {
	var ret string
	val, err := s.Caller().Call(ctx, s.RID(), "getActiveCertChainPEM", nil)
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
	ret, err = encoding.Is[string](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_ServerSSLCert) GetActiveKeyPEM(ctx context.Context, in0 string) (string, error) {
	var ret string
	val, err := s.Caller().Call(ctx, s.RID(), "getActiveKeyPEM", map[string]any{
		"keyPassword": in0,
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
	ret, err = encoding.Is[string](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_ServerSSLCert) GetPendingRequestPEM(ctx context.Context) (string, error) {
	var ret string
	val, err := s.Caller().Call(ctx, s.RID(), "getPendingRequestPEM", nil)
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
	ret, err = encoding.Is[string](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_ServerSSLCert) GetPendingCertChainPEM(ctx context.Context) (string, error) {
	var ret string
	val, err := s.Caller().Call(ctx, s.RID(), "getPendingCertChainPEM", nil)
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
	ret, err = encoding.Is[string](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_ServerSSLCert) GetPendingKeyPEM(ctx context.Context, in0 string) (string, error) {
	var ret string
	val, err := s.Caller().Call(ctx, s.RID(), "getPendingKeyPEM", map[string]any{
		"keyPassword": in0,
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
	ret, err = encoding.Is[string](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_ServerSSLCert) SetPendingCertChainPEM(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setPendingCertChainPEM", map[string]any{
		"certChain": in0,
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

func (s *_ServerSSLCert) SetPendingKeyAndCertChainPEM(ctx context.Context, in0 string, in1 string, in2 string) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setPendingKeyAndCertChainPEM", map[string]any{
		"key":         in0,
		"certChain":   in1,
		"keyPassword": in2,
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

func (s *_ServerSSLCert) InstallPendingKeyPair(ctx context.Context) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "installPendingKeyPair", nil)
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

func (s *_ServerSSLCert) GenerateScepKey(ctx context.Context, in0 KeyInfo) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "generateScepKey", map[string]any{
		"keyInfo": in0.Encode(),
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

func (s *_ServerSSLCert) InstallScepCertChain(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "installScepCertChain", map[string]any{
		"certChain": in0,
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
