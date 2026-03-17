// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package security

import (
	"github.com/arminguenther/xeruspower-go/v40010/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40010/idl"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40010/internal/encoding/valobj"
)

func (i *IpfwRule) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["ipMask"] = i.IpMask
	j0["policy"] = i.Policy
	return j0
}

func (i *IpfwRule) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("ipMask", j0)
	if err != nil {
		return err
	}
	i.IpMask, err = encoding.Is[string](j0["ipMask"])
	if err != nil {
		return err
	}
	err = encoding.In("policy", j0)
	if err != nil {
		return err
	}
	i.Policy, err = encoding.AsEnum[IpfwPolicy](j0["policy"])
	if err != nil {
		return err
	}
	return nil
}

func (i *IpFw) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["enabled"] = i.Enabled
	j0["defaultPolicyIn"] = i.DefaultPolicyIn
	j0["defaultPolicyOut"] = i.DefaultPolicyOut
	s1 := make([]any, 0, len(i.RuleSetIn))
	for _, e1 := range i.RuleSetIn {
		s1 = append(s1, e1.Encode())
	}
	j0["ruleSetIn"] = s1
	s2 := make([]any, 0, len(i.RuleSetOut))
	for _, e2 := range i.RuleSetOut {
		s2 = append(s2, e2.Encode())
	}
	j0["ruleSetOut"] = s2
	return j0
}

func (i *IpFw) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("enabled", j0)
	if err != nil {
		return err
	}
	i.Enabled, err = encoding.Is[bool](j0["enabled"])
	if err != nil {
		return err
	}
	err = encoding.In("defaultPolicyIn", j0)
	if err != nil {
		return err
	}
	i.DefaultPolicyIn, err = encoding.AsEnum[IpfwPolicy](j0["defaultPolicyIn"])
	if err != nil {
		return err
	}
	err = encoding.In("defaultPolicyOut", j0)
	if err != nil {
		return err
	}
	i.DefaultPolicyOut, err = encoding.AsEnum[IpfwPolicy](j0["defaultPolicyOut"])
	if err != nil {
		return err
	}
	err = encoding.In("ruleSetIn", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["ruleSetIn"])
	if err != nil {
		return err
	}
	i.RuleSetIn = make([]IpfwRule, 0, len(s1))
	for _, a1 := range s1 {
		var e1 IpfwRule
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		i.RuleSetIn = append(i.RuleSetIn, e1)
	}
	err = encoding.In("ruleSetOut", j0)
	if err != nil {
		return err
	}
	var s2 []any
	s2, err = encoding.Is[[]any](j0["ruleSetOut"])
	if err != nil {
		return err
	}
	i.RuleSetOut = make([]IpfwRule, 0, len(s2))
	for _, a2 := range s2 {
		var e2 IpfwRule
		err = e2.Decode(a2, caller)
		if err != nil {
			return err
		}
		i.RuleSetOut = append(i.RuleSetOut, e2)
	}
	return nil
}

func (r *RoleAccessRule) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["startIp"] = r.StartIp
	j0["endIp"] = r.EndIp
	j0["roleId"] = r.RoleId
	j0["policy"] = r.Policy
	return j0
}

func (r *RoleAccessRule) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("startIp", j0)
	if err != nil {
		return err
	}
	r.StartIp, err = encoding.Is[string](j0["startIp"])
	if err != nil {
		return err
	}
	err = encoding.In("endIp", j0)
	if err != nil {
		return err
	}
	r.EndIp, err = encoding.Is[string](j0["endIp"])
	if err != nil {
		return err
	}
	err = encoding.In("roleId", j0)
	if err != nil {
		return err
	}
	r.RoleId, err = encoding.AsInt32(j0["roleId"])
	if err != nil {
		return err
	}
	err = encoding.In("policy", j0)
	if err != nil {
		return err
	}
	r.Policy, err = encoding.AsEnum[RoleAccessPolicy](j0["policy"])
	if err != nil {
		return err
	}
	return nil
}

func (r *RoleAccessControl) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["enabled"] = r.Enabled
	j0["defaultPolicy"] = r.DefaultPolicy
	s1 := make([]any, 0, len(r.Rules))
	for _, e1 := range r.Rules {
		s1 = append(s1, e1.Encode())
	}
	j0["rules"] = s1
	return j0
}

func (r *RoleAccessControl) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("enabled", j0)
	if err != nil {
		return err
	}
	r.Enabled, err = encoding.Is[bool](j0["enabled"])
	if err != nil {
		return err
	}
	err = encoding.In("defaultPolicy", j0)
	if err != nil {
		return err
	}
	r.DefaultPolicy, err = encoding.AsEnum[RoleAccessPolicy](j0["defaultPolicy"])
	if err != nil {
		return err
	}
	err = encoding.In("rules", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["rules"])
	if err != nil {
		return err
	}
	r.Rules = make([]RoleAccessRule, 0, len(s1))
	for _, a1 := range s1 {
		var e1 RoleAccessRule
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		r.Rules = append(r.Rules, e1)
	}
	return nil
}

func (b *BlockSettings) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["maxFailedLogins"] = b.MaxFailedLogins
	j0["blockTimeout"] = b.BlockTimeout
	j0["failedLoginTimeout"] = b.FailedLoginTimeout
	return j0
}

func (b *BlockSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("maxFailedLogins", j0)
	if err != nil {
		return err
	}
	b.MaxFailedLogins, err = encoding.AsInt32(j0["maxFailedLogins"])
	if err != nil {
		return err
	}
	err = encoding.In("blockTimeout", j0)
	if err != nil {
		return err
	}
	b.BlockTimeout, err = encoding.AsInt32(j0["blockTimeout"])
	if err != nil {
		return err
	}
	err = encoding.In("failedLoginTimeout", j0)
	if err != nil {
		return err
	}
	b.FailedLoginTimeout, err = encoding.AsInt32(j0["failedLoginTimeout"])
	if err != nil {
		return err
	}
	return nil
}

func (p *PasswordSettings) Encode() map[string]any {
	j0 := make(map[string]any, 10)
	j0["enableAging"] = p.EnableAging
	j0["agingInterval"] = p.AgingInterval
	j0["enableStrongReq"] = p.EnableStrongReq
	j0["minPwLength"] = p.MinPwLength
	j0["maxPwLength"] = p.MaxPwLength
	j0["enforceLower"] = p.EnforceLower
	j0["enforceUpper"] = p.EnforceUpper
	j0["enforceNumeric"] = p.EnforceNumeric
	j0["enforceSpecial"] = p.EnforceSpecial
	j0["pwHistoryDepth"] = p.PwHistoryDepth
	return j0
}

func (p *PasswordSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("enableAging", j0)
	if err != nil {
		return err
	}
	p.EnableAging, err = encoding.Is[bool](j0["enableAging"])
	if err != nil {
		return err
	}
	err = encoding.In("agingInterval", j0)
	if err != nil {
		return err
	}
	p.AgingInterval, err = encoding.AsInt32(j0["agingInterval"])
	if err != nil {
		return err
	}
	err = encoding.In("enableStrongReq", j0)
	if err != nil {
		return err
	}
	p.EnableStrongReq, err = encoding.Is[bool](j0["enableStrongReq"])
	if err != nil {
		return err
	}
	err = encoding.In("minPwLength", j0)
	if err != nil {
		return err
	}
	p.MinPwLength, err = encoding.AsInt32(j0["minPwLength"])
	if err != nil {
		return err
	}
	err = encoding.In("maxPwLength", j0)
	if err != nil {
		return err
	}
	p.MaxPwLength, err = encoding.AsInt32(j0["maxPwLength"])
	if err != nil {
		return err
	}
	err = encoding.In("enforceLower", j0)
	if err != nil {
		return err
	}
	p.EnforceLower, err = encoding.Is[bool](j0["enforceLower"])
	if err != nil {
		return err
	}
	err = encoding.In("enforceUpper", j0)
	if err != nil {
		return err
	}
	p.EnforceUpper, err = encoding.Is[bool](j0["enforceUpper"])
	if err != nil {
		return err
	}
	err = encoding.In("enforceNumeric", j0)
	if err != nil {
		return err
	}
	p.EnforceNumeric, err = encoding.Is[bool](j0["enforceNumeric"])
	if err != nil {
		return err
	}
	err = encoding.In("enforceSpecial", j0)
	if err != nil {
		return err
	}
	p.EnforceSpecial, err = encoding.Is[bool](j0["enforceSpecial"])
	if err != nil {
		return err
	}
	err = encoding.In("pwHistoryDepth", j0)
	if err != nil {
		return err
	}
	p.PwHistoryDepth, err = encoding.AsInt32(j0["pwHistoryDepth"])
	if err != nil {
		return err
	}
	return nil
}

func (s *SSHSettings) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["allowPasswordAuth"] = s.AllowPasswordAuth
	j0["allowPublicKeyAuth"] = s.AllowPublicKeyAuth
	return j0
}

func (s *SSHSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("allowPasswordAuth", j0)
	if err != nil {
		return err
	}
	s.AllowPasswordAuth, err = encoding.Is[bool](j0["allowPasswordAuth"])
	if err != nil {
		return err
	}
	err = encoding.In("allowPublicKeyAuth", j0)
	if err != nil {
		return err
	}
	s.AllowPublicKeyAuth, err = encoding.Is[bool](j0["allowPublicKeyAuth"])
	if err != nil {
		return err
	}
	return nil
}

func (s *SSHKeyFingerprint) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["fingerprint"] = s.Fingerprint
	j0["type"] = s.Type
	return j0
}

func (s *SSHKeyFingerprint) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("fingerprint", j0)
	if err != nil {
		return err
	}
	s.Fingerprint, err = encoding.Is[string](j0["fingerprint"])
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	s.Type, err = encoding.AsEnum[SSHKeyFingerprintType](j0["type"])
	if err != nil {
		return err
	}
	return nil
}

func (s *SSHHostKey) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["key"] = s.Key
	j0["type"] = s.Type
	s1 := make([]any, 0, len(s.Fingerprints))
	for _, e1 := range s.Fingerprints {
		s1 = append(s1, e1.Encode())
	}
	j0["fingerprints"] = s1
	return j0
}

func (s *SSHHostKey) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("key", j0)
	if err != nil {
		return err
	}
	s.Key, err = encoding.Is[string](j0["key"])
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	s.Type, err = encoding.AsEnum[SSHHostKeyType](j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("fingerprints", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["fingerprints"])
	if err != nil {
		return err
	}
	s.Fingerprints = make([]SSHKeyFingerprint, 0, len(s1))
	for _, a1 := range s1 {
		var e1 SSHKeyFingerprint
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		s.Fingerprints = append(s.Fingerprints, e1)
	}
	return nil
}

func (r *RestrictedServiceAgreement) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["enabled"] = r.Enabled
	j0["banner"] = r.Banner
	return j0
}

func (r *RestrictedServiceAgreement) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("enabled", j0)
	if err != nil {
		return err
	}
	r.Enabled, err = encoding.Is[bool](j0["enabled"])
	if err != nil {
		return err
	}
	err = encoding.In("banner", j0)
	if err != nil {
		return err
	}
	r.Banner, err = encoding.Is[string](j0["banner"])
	if err != nil {
		return err
	}
	return nil
}

func (t *TpmInfo) Encode() map[string]any {
	j0 := make(map[string]any, 1)
	j0["detected"] = t.Detected
	return j0
}

func (t *TpmInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("detected", j0)
	if err != nil {
		return err
	}
	t.Detected, err = encoding.Is[bool](j0["detected"])
	if err != nil {
		return err
	}
	return nil
}

func (p *_PasswordSettingsChanged) Decode(value map[string]any, caller idl.Caller) error {
	p.UserEvent = valobj.For[userevent.UserEvent]()
	err := p.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldSettings", value)
	if err != nil {
		return err
	}
	err = p.oldSettings.Decode(value["oldSettings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newSettings", value)
	if err != nil {
		return err
	}
	err = p.newSettings.Decode(value["newSettings"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (f *_FrontPanelPrivilegesChanged) Decode(value map[string]any, caller idl.Caller) error {
	f.UserEvent = valobj.For[userevent.UserEvent]()
	err := f.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldPrivileges", value)
	if err != nil {
		return err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](value["oldPrivileges"])
	if err != nil {
		return err
	}
	f.oldPrivileges = make([]string, 0, len(s0))
	for _, a0 := range s0 {
		var e0 string
		e0, err = encoding.Is[string](a0)
		if err != nil {
			return err
		}
		f.oldPrivileges = append(f.oldPrivileges, e0)
	}
	err = encoding.In("newPrivileges", value)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](value["newPrivileges"])
	if err != nil {
		return err
	}
	f.newPrivileges = make([]string, 0, len(s1))
	for _, a1 := range s1 {
		var e1 string
		e1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		f.newPrivileges = append(f.newPrivileges, e1)
	}
	return nil
}
