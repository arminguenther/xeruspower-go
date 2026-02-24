// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package user

import (
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
)

func (s *SnmpV3Settings) Encode() map[string]any {
	j0 := make(map[string]any, 10)
	j0["enabled"] = s.Enabled
	j0["secLevel"] = s.SecLevel
	j0["authProtocol"] = s.AuthProtocol
	j0["usePasswordAsAuthPassphrase"] = s.UsePasswordAsAuthPassphrase
	j0["haveAuthPassphrase"] = s.HaveAuthPassphrase
	j0["authPassphrase"] = s.AuthPassphrase
	j0["privProtocol"] = s.PrivProtocol
	j0["useAuthPassphraseAsPrivPassphrase"] = s.UseAuthPassphraseAsPrivPassphrase
	j0["havePrivPassphrase"] = s.HavePrivPassphrase
	j0["privPassphrase"] = s.PrivPassphrase
	return j0
}

func (s *SnmpV3Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("enabled", j0)
	if err != nil {
		return err
	}
	s.Enabled, err = encoding.Is[bool](j0["enabled"])
	if err != nil {
		return err
	}
	err = encoding.In("secLevel", j0)
	if err != nil {
		return err
	}
	s.SecLevel, err = encoding.AsEnum[SnmpV3SecLevel](j0["secLevel"])
	if err != nil {
		return err
	}
	err = encoding.In("authProtocol", j0)
	if err != nil {
		return err
	}
	s.AuthProtocol, err = encoding.AsEnum[SnmpV3AuthProto](j0["authProtocol"])
	if err != nil {
		return err
	}
	err = encoding.In("usePasswordAsAuthPassphrase", j0)
	if err != nil {
		return err
	}
	s.UsePasswordAsAuthPassphrase, err = encoding.Is[bool](j0["usePasswordAsAuthPassphrase"])
	if err != nil {
		return err
	}
	err = encoding.In("haveAuthPassphrase", j0)
	if err != nil {
		return err
	}
	s.HaveAuthPassphrase, err = encoding.Is[bool](j0["haveAuthPassphrase"])
	if err != nil {
		return err
	}
	err = encoding.In("authPassphrase", j0)
	if err != nil {
		return err
	}
	s.AuthPassphrase, err = encoding.Is[string](j0["authPassphrase"])
	if err != nil {
		return err
	}
	err = encoding.In("privProtocol", j0)
	if err != nil {
		return err
	}
	s.PrivProtocol, err = encoding.AsEnum[SnmpV3PrivProto](j0["privProtocol"])
	if err != nil {
		return err
	}
	err = encoding.In("useAuthPassphraseAsPrivPassphrase", j0)
	if err != nil {
		return err
	}
	s.UseAuthPassphraseAsPrivPassphrase, err = encoding.Is[bool](j0["useAuthPassphraseAsPrivPassphrase"])
	if err != nil {
		return err
	}
	err = encoding.In("havePrivPassphrase", j0)
	if err != nil {
		return err
	}
	s.HavePrivPassphrase, err = encoding.Is[bool](j0["havePrivPassphrase"])
	if err != nil {
		return err
	}
	err = encoding.In("privPassphrase", j0)
	if err != nil {
		return err
	}
	s.PrivPassphrase, err = encoding.Is[string](j0["privPassphrase"])
	if err != nil {
		return err
	}
	return nil
}

func (a *AuxInfo) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["fullname"] = a.Fullname
	j0["telephone"] = a.Telephone
	j0["eMail"] = a.EMail
	return j0
}

func (a *AuxInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("fullname", j0)
	if err != nil {
		return err
	}
	a.Fullname, err = encoding.Is[string](j0["fullname"])
	if err != nil {
		return err
	}
	err = encoding.In("telephone", j0)
	if err != nil {
		return err
	}
	a.Telephone, err = encoding.Is[string](j0["telephone"])
	if err != nil {
		return err
	}
	err = encoding.In("eMail", j0)
	if err != nil {
		return err
	}
	a.EMail, err = encoding.Is[string](j0["eMail"])
	if err != nil {
		return err
	}
	return nil
}

func (p *Preferences) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["temperatureUnit"] = p.TemperatureUnit
	j0["lengthUnit"] = p.LengthUnit
	j0["pressureUnit"] = p.PressureUnit
	return j0
}

func (p *Preferences) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("temperatureUnit", j0)
	if err != nil {
		return err
	}
	p.TemperatureUnit, err = encoding.AsEnum[TemperatureEnum](j0["temperatureUnit"])
	if err != nil {
		return err
	}
	err = encoding.In("lengthUnit", j0)
	if err != nil {
		return err
	}
	p.LengthUnit, err = encoding.AsEnum[LengthEnum](j0["lengthUnit"])
	if err != nil {
		return err
	}
	err = encoding.In("pressureUnit", j0)
	if err != nil {
		return err
	}
	p.PressureUnit, err = encoding.AsEnum[PressureEnum](j0["pressureUnit"])
	if err != nil {
		return err
	}
	return nil
}

func (u *Info) Encode() map[string]any {
	j0 := make(map[string]any, 9)
	j0["enabled"] = u.Enabled
	j0["locked"] = u.Locked
	j0["blocked"] = u.Blocked
	j0["needPasswordChange"] = u.NeedPasswordChange
	j0["auxInfo"] = u.AuxInfo.Encode()
	j0["snmpV3Settings"] = u.SnmpV3Settings.Encode()
	j0["sshPublicKey"] = u.SshPublicKey
	j0["preferences"] = u.Preferences.Encode()
	j0["roleIds"] = encoding.NonNilSlice(u.RoleIds)
	return j0
}

func (u *Info) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("enabled", j0)
	if err != nil {
		return err
	}
	u.Enabled, err = encoding.Is[bool](j0["enabled"])
	if err != nil {
		return err
	}
	err = encoding.In("locked", j0)
	if err != nil {
		return err
	}
	u.Locked, err = encoding.Is[bool](j0["locked"])
	if err != nil {
		return err
	}
	err = encoding.In("blocked", j0)
	if err != nil {
		return err
	}
	u.Blocked, err = encoding.Is[bool](j0["blocked"])
	if err != nil {
		return err
	}
	err = encoding.In("needPasswordChange", j0)
	if err != nil {
		return err
	}
	u.NeedPasswordChange, err = encoding.Is[bool](j0["needPasswordChange"])
	if err != nil {
		return err
	}
	err = encoding.In("auxInfo", j0)
	if err != nil {
		return err
	}
	err = u.AuxInfo.Decode(j0["auxInfo"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("snmpV3Settings", j0)
	if err != nil {
		return err
	}
	err = u.SnmpV3Settings.Decode(j0["snmpV3Settings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("sshPublicKey", j0)
	if err != nil {
		return err
	}
	u.SshPublicKey, err = encoding.Is[string](j0["sshPublicKey"])
	if err != nil {
		return err
	}
	err = encoding.In("preferences", j0)
	if err != nil {
		return err
	}
	err = u.Preferences.Decode(j0["preferences"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("roleIds", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["roleIds"])
	if err != nil {
		return err
	}
	u.RoleIds = make([]int32, 0, len(s1))
	for _, a1 := range s1 {
		var e1 int32
		e1, err = encoding.AsInt32(a1)
		if err != nil {
			return err
		}
		u.RoleIds = append(u.RoleIds, e1)
	}
	return nil
}

func (u *Capabilities) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["canSetPassword"] = u.CanSetPassword
	j0["canSetPreferences"] = u.CanSetPreferences
	return j0
}

func (u *Capabilities) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("canSetPassword", j0)
	if err != nil {
		return err
	}
	u.CanSetPassword, err = encoding.Is[bool](j0["canSetPassword"])
	if err != nil {
		return err
	}
	err = encoding.In("canSetPreferences", j0)
	if err != nil {
		return err
	}
	u.CanSetPreferences, err = encoding.Is[bool](j0["canSetPreferences"])
	if err != nil {
		return err
	}
	return nil
}
