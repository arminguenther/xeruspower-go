// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package serversslcert

import (
	"github.com/arminguenther/xeruspower-go/v40000/idl"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding"
)

func (c *CommonAttributes) Encode() map[string]any {
	j0 := make(map[string]any, 7)
	j0["country"] = c.Country
	j0["stateOrProvince"] = c.StateOrProvince
	j0["locality"] = c.Locality
	j0["organization"] = c.Organization
	j0["organizationalUnit"] = c.OrganizationalUnit
	j0["commonName"] = c.CommonName
	j0["emailAddress"] = c.EmailAddress
	return j0
}

func (c *CommonAttributes) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("country", j0)
	if err != nil {
		return err
	}
	c.Country, err = encoding.Is[string](j0["country"])
	if err != nil {
		return err
	}
	err = encoding.In("stateOrProvince", j0)
	if err != nil {
		return err
	}
	c.StateOrProvince, err = encoding.Is[string](j0["stateOrProvince"])
	if err != nil {
		return err
	}
	err = encoding.In("locality", j0)
	if err != nil {
		return err
	}
	c.Locality, err = encoding.Is[string](j0["locality"])
	if err != nil {
		return err
	}
	err = encoding.In("organization", j0)
	if err != nil {
		return err
	}
	c.Organization, err = encoding.Is[string](j0["organization"])
	if err != nil {
		return err
	}
	err = encoding.In("organizationalUnit", j0)
	if err != nil {
		return err
	}
	c.OrganizationalUnit, err = encoding.Is[string](j0["organizationalUnit"])
	if err != nil {
		return err
	}
	err = encoding.In("commonName", j0)
	if err != nil {
		return err
	}
	c.CommonName, err = encoding.Is[string](j0["commonName"])
	if err != nil {
		return err
	}
	err = encoding.In("emailAddress", j0)
	if err != nil {
		return err
	}
	c.EmailAddress, err = encoding.Is[string](j0["emailAddress"])
	if err != nil {
		return err
	}
	return nil
}

func (r *ReqInfo) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["subject"] = r.Subject.Encode()
	j0["names"] = encoding.NonNilSlice(r.Names)
	j0["keyType"] = r.KeyType
	j0["ellipticCurve"] = r.EllipticCurve
	j0["rsaKeyLength"] = r.RsaKeyLength
	return j0
}

func (r *ReqInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("subject", j0)
	if err != nil {
		return err
	}
	err = r.Subject.Decode(j0["subject"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("names", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["names"])
	if err != nil {
		return err
	}
	r.Names = make([]string, 0, len(s1))
	for _, a1 := range s1 {
		var e1 string
		e1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		r.Names = append(r.Names, e1)
	}
	err = encoding.In("keyType", j0)
	if err != nil {
		return err
	}
	r.KeyType, err = encoding.AsEnum[KeyType](j0["keyType"])
	if err != nil {
		return err
	}
	err = encoding.In("ellipticCurve", j0)
	if err != nil {
		return err
	}
	r.EllipticCurve, err = encoding.AsEnum[EllipticCurve](j0["ellipticCurve"])
	if err != nil {
		return err
	}
	err = encoding.In("rsaKeyLength", j0)
	if err != nil {
		return err
	}
	r.RsaKeyLength, err = encoding.AsInt32(j0["rsaKeyLength"])
	if err != nil {
		return err
	}
	return nil
}

func (c *CertInfo) Encode() map[string]any {
	j0 := make(map[string]any, 9)
	j0["subject"] = c.Subject.Encode()
	j0["issuer"] = c.Issuer.Encode()
	j0["names"] = encoding.NonNilSlice(c.Names)
	j0["invalidBefore"] = c.InvalidBefore
	j0["invalidAfter"] = c.InvalidAfter
	j0["serialNumber"] = c.SerialNumber
	j0["keyType"] = c.KeyType
	j0["ellipticCurve"] = c.EllipticCurve
	j0["rsaKeyLength"] = c.RsaKeyLength
	return j0
}

func (c *CertInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("subject", j0)
	if err != nil {
		return err
	}
	err = c.Subject.Decode(j0["subject"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("issuer", j0)
	if err != nil {
		return err
	}
	err = c.Issuer.Decode(j0["issuer"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("names", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["names"])
	if err != nil {
		return err
	}
	c.Names = make([]string, 0, len(s1))
	for _, a1 := range s1 {
		var e1 string
		e1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		c.Names = append(c.Names, e1)
	}
	err = encoding.In("invalidBefore", j0)
	if err != nil {
		return err
	}
	c.InvalidBefore, err = encoding.Is[string](j0["invalidBefore"])
	if err != nil {
		return err
	}
	err = encoding.In("invalidAfter", j0)
	if err != nil {
		return err
	}
	c.InvalidAfter, err = encoding.Is[string](j0["invalidAfter"])
	if err != nil {
		return err
	}
	err = encoding.In("serialNumber", j0)
	if err != nil {
		return err
	}
	c.SerialNumber, err = encoding.Is[string](j0["serialNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("keyType", j0)
	if err != nil {
		return err
	}
	c.KeyType, err = encoding.AsEnum[KeyType](j0["keyType"])
	if err != nil {
		return err
	}
	err = encoding.In("ellipticCurve", j0)
	if err != nil {
		return err
	}
	c.EllipticCurve, err = encoding.AsEnum[EllipticCurve](j0["ellipticCurve"])
	if err != nil {
		return err
	}
	err = encoding.In("rsaKeyLength", j0)
	if err != nil {
		return err
	}
	c.RsaKeyLength, err = encoding.AsInt32(j0["rsaKeyLength"])
	if err != nil {
		return err
	}
	return nil
}

func (i *Info) Encode() map[string]any {
	j0 := make(map[string]any, 8)
	j0["havePendingReq"] = i.HavePendingReq
	j0["havePendingCert"] = i.HavePendingCert
	j0["pendingReqInfo"] = i.PendingReqInfo.Encode()
	j0["pendingCertInfo"] = i.PendingCertInfo.Encode()
	s1 := make([]any, 0, len(i.PendingCertChainInfos))
	for _, e1 := range i.PendingCertChainInfos {
		s1 = append(s1, e1.Encode())
	}
	j0["pendingCertChainInfos"] = s1
	j0["activeCertInfo"] = i.ActiveCertInfo.Encode()
	s2 := make([]any, 0, len(i.ActiveCertChainInfos))
	for _, e2 := range i.ActiveCertChainInfos {
		s2 = append(s2, e2.Encode())
	}
	j0["activeCertChainInfos"] = s2
	j0["maxSignDays"] = i.MaxSignDays
	return j0
}

func (i *Info) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("havePendingReq", j0)
	if err != nil {
		return err
	}
	i.HavePendingReq, err = encoding.Is[bool](j0["havePendingReq"])
	if err != nil {
		return err
	}
	err = encoding.In("havePendingCert", j0)
	if err != nil {
		return err
	}
	i.HavePendingCert, err = encoding.Is[bool](j0["havePendingCert"])
	if err != nil {
		return err
	}
	err = encoding.In("pendingReqInfo", j0)
	if err != nil {
		return err
	}
	err = i.PendingReqInfo.Decode(j0["pendingReqInfo"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("pendingCertInfo", j0)
	if err != nil {
		return err
	}
	err = i.PendingCertInfo.Decode(j0["pendingCertInfo"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("pendingCertChainInfos", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["pendingCertChainInfos"])
	if err != nil {
		return err
	}
	i.PendingCertChainInfos = make([]CertInfo, 0, len(s1))
	for _, a1 := range s1 {
		var e1 CertInfo
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		i.PendingCertChainInfos = append(i.PendingCertChainInfos, e1)
	}
	err = encoding.In("activeCertInfo", j0)
	if err != nil {
		return err
	}
	err = i.ActiveCertInfo.Decode(j0["activeCertInfo"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("activeCertChainInfos", j0)
	if err != nil {
		return err
	}
	var s2 []any
	s2, err = encoding.Is[[]any](j0["activeCertChainInfos"])
	if err != nil {
		return err
	}
	i.ActiveCertChainInfos = make([]CertInfo, 0, len(s2))
	for _, a2 := range s2 {
		var e2 CertInfo
		err = e2.Decode(a2, caller)
		if err != nil {
			return err
		}
		i.ActiveCertChainInfos = append(i.ActiveCertChainInfos, e2)
	}
	err = encoding.In("maxSignDays", j0)
	if err != nil {
		return err
	}
	i.MaxSignDays, err = encoding.AsInt32(j0["maxSignDays"])
	if err != nil {
		return err
	}
	return nil
}
