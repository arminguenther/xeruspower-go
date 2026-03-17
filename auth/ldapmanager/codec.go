// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package ldapmanager

import (
	"github.com/arminguenther/xeruspower-go/v40200/idl"
	"github.com/arminguenther/xeruspower-go/v40200/internal/encoding"
)

func (s *ServerSettings) Encode() map[string]any {
	j0 := make(map[string]any, 23)
	j0["id"] = s.Id
	j0["server"] = s.Server
	j0["adoptSettingsId"] = s.AdoptSettingsId
	j0["type"] = s.Type
	j0["secProto"] = s.SecProto
	j0["port"] = s.Port
	j0["sslPort"] = s.SslPort
	j0["forceTrustedCert"] = s.ForceTrustedCert
	j0["allowOffTimeRangeCerts"] = s.AllowOffTimeRangeCerts
	j0["certificate"] = s.Certificate
	j0["adsDomain"] = s.AdsDomain
	j0["useAnonymousBind"] = s.UseAnonymousBind
	j0["bindDN"] = s.BindDN
	j0["bindPwd"] = s.BindPwd
	j0["searchBaseDN"] = s.SearchBaseDN
	j0["loginNameAttr"] = s.LoginNameAttr
	j0["userEntryObjClass"] = s.UserEntryObjClass
	j0["userSearchFilter"] = s.UserSearchFilter
	j0["groupInfoInUserEntry"] = s.GroupInfoInUserEntry
	j0["supportNestedGroups"] = s.SupportNestedGroups
	j0["groupMemberAttr"] = s.GroupMemberAttr
	j0["groupEntryObjClass"] = s.GroupEntryObjClass
	j0["groupSearchFilter"] = s.GroupSearchFilter
	return j0
}

func (s *ServerSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("id", j0)
	if err != nil {
		return err
	}
	s.Id, err = encoding.Is[string](j0["id"])
	if err != nil {
		return err
	}
	err = encoding.In("server", j0)
	if err != nil {
		return err
	}
	s.Server, err = encoding.Is[string](j0["server"])
	if err != nil {
		return err
	}
	err = encoding.In("adoptSettingsId", j0)
	if err != nil {
		return err
	}
	s.AdoptSettingsId, err = encoding.Is[string](j0["adoptSettingsId"])
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	s.Type, err = encoding.AsEnum[ServerType](j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("secProto", j0)
	if err != nil {
		return err
	}
	s.SecProto, err = encoding.AsEnum[SecurityProtocol](j0["secProto"])
	if err != nil {
		return err
	}
	err = encoding.In("port", j0)
	if err != nil {
		return err
	}
	s.Port, err = encoding.AsInt32(j0["port"])
	if err != nil {
		return err
	}
	err = encoding.In("sslPort", j0)
	if err != nil {
		return err
	}
	s.SslPort, err = encoding.AsInt32(j0["sslPort"])
	if err != nil {
		return err
	}
	err = encoding.In("forceTrustedCert", j0)
	if err != nil {
		return err
	}
	s.ForceTrustedCert, err = encoding.Is[bool](j0["forceTrustedCert"])
	if err != nil {
		return err
	}
	err = encoding.In("allowOffTimeRangeCerts", j0)
	if err != nil {
		return err
	}
	s.AllowOffTimeRangeCerts, err = encoding.Is[bool](j0["allowOffTimeRangeCerts"])
	if err != nil {
		return err
	}
	err = encoding.In("certificate", j0)
	if err != nil {
		return err
	}
	s.Certificate, err = encoding.Is[string](j0["certificate"])
	if err != nil {
		return err
	}
	err = encoding.In("adsDomain", j0)
	if err != nil {
		return err
	}
	s.AdsDomain, err = encoding.Is[string](j0["adsDomain"])
	if err != nil {
		return err
	}
	err = encoding.In("useAnonymousBind", j0)
	if err != nil {
		return err
	}
	s.UseAnonymousBind, err = encoding.Is[bool](j0["useAnonymousBind"])
	if err != nil {
		return err
	}
	err = encoding.In("bindDN", j0)
	if err != nil {
		return err
	}
	s.BindDN, err = encoding.Is[string](j0["bindDN"])
	if err != nil {
		return err
	}
	err = encoding.In("bindPwd", j0)
	if err != nil {
		return err
	}
	s.BindPwd, err = encoding.Is[string](j0["bindPwd"])
	if err != nil {
		return err
	}
	err = encoding.In("searchBaseDN", j0)
	if err != nil {
		return err
	}
	s.SearchBaseDN, err = encoding.Is[string](j0["searchBaseDN"])
	if err != nil {
		return err
	}
	err = encoding.In("loginNameAttr", j0)
	if err != nil {
		return err
	}
	s.LoginNameAttr, err = encoding.Is[string](j0["loginNameAttr"])
	if err != nil {
		return err
	}
	err = encoding.In("userEntryObjClass", j0)
	if err != nil {
		return err
	}
	s.UserEntryObjClass, err = encoding.Is[string](j0["userEntryObjClass"])
	if err != nil {
		return err
	}
	err = encoding.In("userSearchFilter", j0)
	if err != nil {
		return err
	}
	s.UserSearchFilter, err = encoding.Is[string](j0["userSearchFilter"])
	if err != nil {
		return err
	}
	err = encoding.In("groupInfoInUserEntry", j0)
	if err != nil {
		return err
	}
	s.GroupInfoInUserEntry, err = encoding.Is[bool](j0["groupInfoInUserEntry"])
	if err != nil {
		return err
	}
	err = encoding.In("supportNestedGroups", j0)
	if err != nil {
		return err
	}
	s.SupportNestedGroups, err = encoding.Is[bool](j0["supportNestedGroups"])
	if err != nil {
		return err
	}
	err = encoding.In("groupMemberAttr", j0)
	if err != nil {
		return err
	}
	s.GroupMemberAttr, err = encoding.Is[string](j0["groupMemberAttr"])
	if err != nil {
		return err
	}
	err = encoding.In("groupEntryObjClass", j0)
	if err != nil {
		return err
	}
	s.GroupEntryObjClass, err = encoding.Is[string](j0["groupEntryObjClass"])
	if err != nil {
		return err
	}
	err = encoding.In("groupSearchFilter", j0)
	if err != nil {
		return err
	}
	s.GroupSearchFilter, err = encoding.Is[string](j0["groupSearchFilter"])
	if err != nil {
		return err
	}
	return nil
}
