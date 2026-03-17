// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package serversslcert

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40200/idl"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding/object"
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
		Major: 3, Submajor: 0, Minor: 0,
	}
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
