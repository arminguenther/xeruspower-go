// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package security

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
)

func init() {
	object.Register(NewSecurity)
}

type _Security struct {
	idl.Object
}

// NewSecurity returns a new Security interface for the object at given RID.
func NewSecurity(rid string, caller idl.Caller) Security {
	return &_Security{idl.NewObject(rid, caller)}
}

func (s *_Security) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "security.Security",
		Major: 5, Submajor: 0, Minor: 1,
	}
}

func (s *_Security) GetHttpRedirSettings(ctx context.Context) (bool, error) {
	var ret bool
	val, err := s.Caller().Call(ctx, s.RID(), "getHttpRedirSettings", nil)
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
	ret, err = encoding.Is[bool](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_Security) SetHttpRedirSettings(ctx context.Context, in0 bool) error {
	_, err := s.Caller().Call(ctx, s.RID(), "setHttpRedirSettings", map[string]any{
		"http2httpsRedir": in0,
	})
	return err
}

func (s *_Security) IsHstsEnabled(ctx context.Context) (bool, error) {
	var ret bool
	val, err := s.Caller().Call(ctx, s.RID(), "isHstsEnabled", nil)
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
	ret, err = encoding.Is[bool](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_Security) SetHstsEnabled(ctx context.Context, in0 bool) error {
	_, err := s.Caller().Call(ctx, s.RID(), "setHstsEnabled", map[string]any{
		"enable": in0,
	})
	return err
}

func (s *_Security) GetIpFwSettings(ctx context.Context) (IpFw, error) {
	var ret IpFw
	val, err := s.Caller().Call(ctx, s.RID(), "getIpFwSettings", nil)
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

func (s *_Security) SetIpFwSettings(ctx context.Context, in0 IpFw) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setIpFwSettings", map[string]any{
		"ipFw": in0.Encode(),
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

func (s *_Security) GetIpV6FwSettings(ctx context.Context) (IpFw, error) {
	var ret IpFw
	val, err := s.Caller().Call(ctx, s.RID(), "getIpV6FwSettings", nil)
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

func (s *_Security) SetIpV6FwSettings(ctx context.Context, in0 IpFw) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setIpV6FwSettings", map[string]any{
		"ipV6Fw": in0.Encode(),
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

func (s *_Security) GetRoleAccessControlSettings(ctx context.Context) (RoleAccessControl, error) {
	var ret RoleAccessControl
	val, err := s.Caller().Call(ctx, s.RID(), "getRoleAccessControlSettings", nil)
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

func (s *_Security) SetRoleAccessControlSettings(ctx context.Context, in0 RoleAccessControl) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setRoleAccessControlSettings", map[string]any{
		"settings": in0.Encode(),
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

func (s *_Security) GetRoleAccessControlSettingsV6(ctx context.Context) (RoleAccessControl, error) {
	var ret RoleAccessControl
	val, err := s.Caller().Call(ctx, s.RID(), "getRoleAccessControlSettingsV6", nil)
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

func (s *_Security) SetRoleAccessControlSettingsV6(ctx context.Context, in0 RoleAccessControl) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setRoleAccessControlSettingsV6", map[string]any{
		"settings": in0.Encode(),
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

func (s *_Security) GetBlockSettings(ctx context.Context) (BlockSettings, error) {
	var ret BlockSettings
	val, err := s.Caller().Call(ctx, s.RID(), "getBlockSettings", nil)
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

func (s *_Security) SetBlockSettings(ctx context.Context, in0 BlockSettings) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setBlockSettings", map[string]any{
		"settings": in0.Encode(),
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

func (s *_Security) GetPwSettings(ctx context.Context) (PasswordSettings, error) {
	var ret PasswordSettings
	val, err := s.Caller().Call(ctx, s.RID(), "getPwSettings", nil)
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

func (s *_Security) SetPwSettings(ctx context.Context, in0 PasswordSettings) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setPwSettings", map[string]any{
		"pwSettings": in0.Encode(),
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

func (s *_Security) GetIdleTimeoutSettings(ctx context.Context) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "getIdleTimeoutSettings", nil)
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

func (s *_Security) SetIdleTimeoutSettings(ctx context.Context, in0 int32) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setIdleTimeoutSettings", map[string]any{
		"idleTimeout": in0,
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

func (s *_Security) GetSingleLoginLimitation(ctx context.Context) (bool, error) {
	var ret bool
	val, err := s.Caller().Call(ctx, s.RID(), "getSingleLoginLimitation", nil)
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
	ret, err = encoding.Is[bool](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_Security) SetSingleLoginLimitation(ctx context.Context, in0 bool) error {
	_, err := s.Caller().Call(ctx, s.RID(), "setSingleLoginLimitation", map[string]any{
		"singleLogin": in0,
	})
	return err
}

func (s *_Security) GetSSHSettings(ctx context.Context) (SSHSettings, error) {
	var ret SSHSettings
	val, err := s.Caller().Call(ctx, s.RID(), "getSSHSettings", nil)
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

func (s *_Security) SetSSHSettings(ctx context.Context, in0 SSHSettings) error {
	_, err := s.Caller().Call(ctx, s.RID(), "setSSHSettings", map[string]any{
		"settings": in0.Encode(),
	})
	return err
}

func (s *_Security) GetSSHHostKeys(ctx context.Context) ([]SSHHostKey, error) {
	var ret []SSHHostKey
	val, err := s.Caller().Call(ctx, s.RID(), "getSSHHostKeys", nil)
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
	ret = make([]SSHHostKey, 0, len(s0))
	for _, a0 := range s0 {
		var e0 SSHHostKey
		err = e0.Decode(a0, s.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (s *_Security) GetRestrictedServiceAgreement(ctx context.Context) (RestrictedServiceAgreement, error) {
	var ret RestrictedServiceAgreement
	val, err := s.Caller().Call(ctx, s.RID(), "getRestrictedServiceAgreement", nil)
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

func (s *_Security) SetRestrictedServiceAgreement(ctx context.Context, in0 RestrictedServiceAgreement) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setRestrictedServiceAgreement", map[string]any{
		"settings": in0.Encode(),
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

func (s *_Security) GetSupportedFrontPanelPrivileges(ctx context.Context) ([]string, error) {
	var ret []string
	val, err := s.Caller().Call(ctx, s.RID(), "getSupportedFrontPanelPrivileges", nil)
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
	ret = make([]string, 0, len(s0))
	for _, a0 := range s0 {
		var e0 string
		e0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (s *_Security) GetFrontPanelPrivileges(ctx context.Context) ([]string, error) {
	var ret []string
	val, err := s.Caller().Call(ctx, s.RID(), "getFrontPanelPrivileges", nil)
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
	ret = make([]string, 0, len(s0))
	for _, a0 := range s0 {
		var e0 string
		e0, err = encoding.Is[string](a0)
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (s *_Security) SetFrontPanelPrivileges(ctx context.Context, in0 []string) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setFrontPanelPrivileges", map[string]any{
		"privileges": encoding.NonNilSlice(in0),
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

func (s *_Security) SetDefaultAdminAccountPassword(ctx context.Context, in0 string, in1 bool) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setDefaultAdminAccountPassword", map[string]any{
		"password":                 in0,
		"disableStrongPasswordReq": in1,
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

func (s *_Security) SetAdminAccountPasswordHash(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "setAdminAccountPasswordHash", map[string]any{
		"passwordHash": in0,
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

func (s *_Security) IsSecureBootActive(ctx context.Context) (bool, error) {
	var ret bool
	val, err := s.Caller().Call(ctx, s.RID(), "isSecureBootActive", nil)
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
	ret, err = encoding.Is[bool](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_Security) GetTpmInfo(ctx context.Context) (TpmInfo, error) {
	var ret TpmInfo
	val, err := s.Caller().Call(ctx, s.RID(), "getTpmInfo", nil)
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

func (s *_Security) GetActiveFipsSettings(ctx context.Context) (FipsSettings, error) {
	var ret FipsSettings
	val, err := s.Caller().Call(ctx, s.RID(), "getActiveFipsSettings", nil)
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

func (s *_Security) GetPersistentFipsSettings(ctx context.Context) (FipsSettings, error) {
	var ret FipsSettings
	val, err := s.Caller().Call(ctx, s.RID(), "getPersistentFipsSettings", nil)
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

func (s *_Security) SetPersistentFipsSettings(ctx context.Context, in0 FipsSettings) error {
	_, err := s.Caller().Call(ctx, s.RID(), "setPersistentFipsSettings", map[string]any{
		"settings": in0.Encode(),
	})
	return err
}
