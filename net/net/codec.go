// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package net

import (
	"github.com/arminguenther/xeruspower-go/v40411/idl"
	"github.com/arminguenther/xeruspower-go/v40411/idl/event"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40411/internal/encoding/valobj"
)

func (i *IpAddrCidr) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["addr"] = i.Addr
	j0["prefixLen"] = i.PrefixLen
	return j0
}

func (i *IpAddrCidr) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("addr", j0)
	if err != nil {
		return err
	}
	i.Addr, err = encoding.Is[string](j0["addr"])
	if err != nil {
		return err
	}
	err = encoding.In("prefixLen", j0)
	if err != nil {
		return err
	}
	i.PrefixLen, err = encoding.AsInt32(j0["prefixLen"])
	if err != nil {
		return err
	}
	return nil
}

func (i *IpRoute) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["destNetAddrCidr"] = i.DestNetAddrCidr.Encode()
	j0["nextHopAddr"] = i.NextHopAddr
	j0["ifName"] = i.IfName
	return j0
}

func (i *IpRoute) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("destNetAddrCidr", j0)
	if err != nil {
		return err
	}
	err = i.DestNetAddrCidr.Decode(j0["destNetAddrCidr"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("nextHopAddr", j0)
	if err != nil {
		return err
	}
	i.NextHopAddr, err = encoding.Is[string](j0["nextHopAddr"])
	if err != nil {
		return err
	}
	err = encoding.In("ifName", j0)
	if err != nil {
		return err
	}
	i.IfName, err = encoding.Is[string](j0["ifName"])
	if err != nil {
		return err
	}
	return nil
}

func (p *PortForwardingSettings) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["enabled"] = p.Enabled
	j0["role"] = p.Role
	j0["primaryUnitDownstreamIfName"] = p.PrimaryUnitDownstreamIfName
	return j0
}

func (p *PortForwardingSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("enabled", j0)
	if err != nil {
		return err
	}
	p.Enabled, err = encoding.Is[bool](j0["enabled"])
	if err != nil {
		return err
	}
	err = encoding.In("role", j0)
	if err != nil {
		return err
	}
	p.Role, err = encoding.AsEnum[PortForwardingRole](j0["role"])
	if err != nil {
		return err
	}
	err = encoding.In("primaryUnitDownstreamIfName", j0)
	if err != nil {
		return err
	}
	p.PrimaryUnitDownstreamIfName, err = encoding.Is[string](j0["primaryUnitDownstreamIfName"])
	if err != nil {
		return err
	}
	return nil
}

func (d *DnsSettings) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["serverAddrs"] = encoding.NonNilSlice(d.ServerAddrs)
	j0["searchSuffixes"] = encoding.NonNilSlice(d.SearchSuffixes)
	j0["resolverPrefersIPv6"] = d.ResolverPrefersIPv6
	return j0
}

func (d *DnsSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("serverAddrs", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["serverAddrs"])
	if err != nil {
		return err
	}
	d.ServerAddrs = make([]string, 0, len(s1))
	for _, a1 := range s1 {
		var e1 string
		e1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		d.ServerAddrs = append(d.ServerAddrs, e1)
	}
	err = encoding.In("searchSuffixes", j0)
	if err != nil {
		return err
	}
	var s2 []any
	s2, err = encoding.Is[[]any](j0["searchSuffixes"])
	if err != nil {
		return err
	}
	d.SearchSuffixes = make([]string, 0, len(s2))
	for _, a2 := range s2 {
		var e2 string
		e2, err = encoding.Is[string](a2)
		if err != nil {
			return err
		}
		d.SearchSuffixes = append(d.SearchSuffixes, e2)
	}
	err = encoding.In("resolverPrefersIPv6", j0)
	if err != nil {
		return err
	}
	d.ResolverPrefersIPv6, err = encoding.Is[bool](j0["resolverPrefersIPv6"])
	if err != nil {
		return err
	}
	return nil
}

func (i *IpRoutingSettings) Encode() map[string]any {
	j0 := make(map[string]any, 1)
	s1 := make([]any, 0, len(i.StaticRoutes))
	for _, e1 := range i.StaticRoutes {
		s1 = append(s1, e1.Encode())
	}
	j0["staticRoutes"] = s1
	return j0
}

func (i *IpRoutingSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("staticRoutes", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["staticRoutes"])
	if err != nil {
		return err
	}
	i.StaticRoutes = make([]IpRoute, 0, len(s1))
	for _, a1 := range s1 {
		var e1 IpRoute
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		i.StaticRoutes = append(i.StaticRoutes, e1)
	}
	return nil
}

func (r *RoutingSettings) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["ipv4"] = r.Ipv4.Encode()
	j0["ipv6"] = r.Ipv6.Encode()
	return j0
}

func (r *RoutingSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("ipv4", j0)
	if err != nil {
		return err
	}
	err = r.Ipv4.Decode(j0["ipv4"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("ipv6", j0)
	if err != nil {
		return err
	}
	err = r.Ipv6.Decode(j0["ipv6"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommonSettings) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["dns"] = c.Dns.Encode()
	j0["routing"] = c.Routing.Encode()
	j0["portForwarding"] = c.PortForwarding.Encode()
	return j0
}

func (c *CommonSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("dns", j0)
	if err != nil {
		return err
	}
	err = c.Dns.Decode(j0["dns"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("routing", j0)
	if err != nil {
		return err
	}
	err = c.Routing.Decode(j0["routing"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("portForwarding", j0)
	if err != nil {
		return err
	}
	err = c.PortForwarding.Decode(j0["portForwarding"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (p *PortForwardingPrimaryUnitAddrInfo) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["ifName"] = p.IfName
	j0["addr"] = p.Addr
	return j0
}

func (p *PortForwardingPrimaryUnitAddrInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("ifName", j0)
	if err != nil {
		return err
	}
	p.IfName, err = encoding.Is[string](j0["ifName"])
	if err != nil {
		return err
	}
	err = encoding.In("addr", j0)
	if err != nil {
		return err
	}
	p.Addr, err = encoding.Is[string](j0["addr"])
	if err != nil {
		return err
	}
	return nil
}

func (p *PortForwardingInfo) Encode() map[string]any {
	j0 := make(map[string]any, 8)
	j0["enabled"] = p.Enabled
	j0["nodeIndexValid"] = p.NodeIndexValid
	j0["nodeIndex"] = p.NodeIndex
	j0["expansionUnitConnected"] = p.ExpansionUnitConnected
	j0["primaryUnitDownstreamIfName"] = p.PrimaryUnitDownstreamIfName
	s1 := make([]any, 0, len(p.PrimaryUnitIPv4AddrInfos))
	for _, e1 := range p.PrimaryUnitIPv4AddrInfos {
		s1 = append(s1, e1.Encode())
	}
	j0["primaryUnitIPv4AddrInfos"] = s1
	s2 := make([]any, 0, len(p.PrimaryUnitIPv6AddrInfos))
	for _, e2 := range p.PrimaryUnitIPv6AddrInfos {
		s2 = append(s2, e2.Encode())
	}
	j0["primaryUnitIPv6AddrInfos"] = s2
	j0["linkLocalIPv6Address"] = p.LinkLocalIPv6Address
	return j0
}

func (p *PortForwardingInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("enabled", j0)
	if err != nil {
		return err
	}
	p.Enabled, err = encoding.Is[bool](j0["enabled"])
	if err != nil {
		return err
	}
	err = encoding.In("nodeIndexValid", j0)
	if err != nil {
		return err
	}
	p.NodeIndexValid, err = encoding.Is[bool](j0["nodeIndexValid"])
	if err != nil {
		return err
	}
	err = encoding.In("nodeIndex", j0)
	if err != nil {
		return err
	}
	p.NodeIndex, err = encoding.AsInt32(j0["nodeIndex"])
	if err != nil {
		return err
	}
	err = encoding.In("expansionUnitConnected", j0)
	if err != nil {
		return err
	}
	p.ExpansionUnitConnected, err = encoding.Is[bool](j0["expansionUnitConnected"])
	if err != nil {
		return err
	}
	err = encoding.In("primaryUnitDownstreamIfName", j0)
	if err != nil {
		return err
	}
	p.PrimaryUnitDownstreamIfName, err = encoding.Is[string](j0["primaryUnitDownstreamIfName"])
	if err != nil {
		return err
	}
	err = encoding.In("primaryUnitIPv4AddrInfos", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["primaryUnitIPv4AddrInfos"])
	if err != nil {
		return err
	}
	p.PrimaryUnitIPv4AddrInfos = make([]PortForwardingPrimaryUnitAddrInfo, 0, len(s1))
	for _, a1 := range s1 {
		var e1 PortForwardingPrimaryUnitAddrInfo
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		p.PrimaryUnitIPv4AddrInfos = append(p.PrimaryUnitIPv4AddrInfos, e1)
	}
	err = encoding.In("primaryUnitIPv6AddrInfos", j0)
	if err != nil {
		return err
	}
	var s2 []any
	s2, err = encoding.Is[[]any](j0["primaryUnitIPv6AddrInfos"])
	if err != nil {
		return err
	}
	p.PrimaryUnitIPv6AddrInfos = make([]PortForwardingPrimaryUnitAddrInfo, 0, len(s2))
	for _, a2 := range s2 {
		var e2 PortForwardingPrimaryUnitAddrInfo
		err = e2.Decode(a2, caller)
		if err != nil {
			return err
		}
		p.PrimaryUnitIPv6AddrInfos = append(p.PrimaryUnitIPv6AddrInfos, e2)
	}
	err = encoding.In("linkLocalIPv6Address", j0)
	if err != nil {
		return err
	}
	p.LinkLocalIPv6Address, err = encoding.Is[string](j0["linkLocalIPv6Address"])
	if err != nil {
		return err
	}
	return nil
}

func (d *DnsInfo) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["serverAddrs"] = encoding.NonNilSlice(d.ServerAddrs)
	j0["searchSuffixes"] = encoding.NonNilSlice(d.SearchSuffixes)
	j0["resolverPrefersIPv6"] = d.ResolverPrefersIPv6
	return j0
}

func (d *DnsInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("serverAddrs", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["serverAddrs"])
	if err != nil {
		return err
	}
	d.ServerAddrs = make([]string, 0, len(s1))
	for _, a1 := range s1 {
		var e1 string
		e1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		d.ServerAddrs = append(d.ServerAddrs, e1)
	}
	err = encoding.In("searchSuffixes", j0)
	if err != nil {
		return err
	}
	var s2 []any
	s2, err = encoding.Is[[]any](j0["searchSuffixes"])
	if err != nil {
		return err
	}
	d.SearchSuffixes = make([]string, 0, len(s2))
	for _, a2 := range s2 {
		var e2 string
		e2, err = encoding.Is[string](a2)
		if err != nil {
			return err
		}
		d.SearchSuffixes = append(d.SearchSuffixes, e2)
	}
	err = encoding.In("resolverPrefersIPv6", j0)
	if err != nil {
		return err
	}
	d.ResolverPrefersIPv6, err = encoding.Is[bool](j0["resolverPrefersIPv6"])
	if err != nil {
		return err
	}
	return nil
}

func (r *RoutingInfo) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	s1 := make([]any, 0, len(r.Ipv4Routes))
	for _, e1 := range r.Ipv4Routes {
		s1 = append(s1, e1.Encode())
	}
	j0["ipv4Routes"] = s1
	s2 := make([]any, 0, len(r.Ipv6Routes))
	for _, e2 := range r.Ipv6Routes {
		s2 = append(s2, e2.Encode())
	}
	j0["ipv6Routes"] = s2
	return j0
}

func (r *RoutingInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("ipv4Routes", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["ipv4Routes"])
	if err != nil {
		return err
	}
	r.Ipv4Routes = make([]IpRoute, 0, len(s1))
	for _, a1 := range s1 {
		var e1 IpRoute
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		r.Ipv4Routes = append(r.Ipv4Routes, e1)
	}
	err = encoding.In("ipv6Routes", j0)
	if err != nil {
		return err
	}
	var s2 []any
	s2, err = encoding.Is[[]any](j0["ipv6Routes"])
	if err != nil {
		return err
	}
	r.Ipv6Routes = make([]IpRoute, 0, len(s2))
	for _, a2 := range s2 {
		var e2 IpRoute
		err = e2.Decode(a2, caller)
		if err != nil {
			return err
		}
		r.Ipv6Routes = append(r.Ipv6Routes, e2)
	}
	return nil
}

func (c *CommonInfo) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["dns"] = c.Dns.Encode()
	j0["routing"] = c.Routing.Encode()
	j0["portForwarding"] = c.PortForwarding.Encode()
	return j0
}

func (c *CommonInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("dns", j0)
	if err != nil {
		return err
	}
	err = c.Dns.Decode(j0["dns"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("routing", j0)
	if err != nil {
		return err
	}
	err = c.Routing.Decode(j0["routing"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("portForwarding", j0)
	if err != nil {
		return err
	}
	err = c.PortForwarding.Decode(j0["portForwarding"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (i *InterfaceIpSettings) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["enabled"] = i.Enabled
	j0["configMethod"] = i.ConfigMethod
	j0["staticAddrCidr"] = i.StaticAddrCidr.Encode()
	j0["staticDefaultGatewayAddr"] = i.StaticDefaultGatewayAddr
	j0["dhcpPreferredHostname"] = i.DhcpPreferredHostname
	j0["stablePrivacyEnabled"] = i.StablePrivacyEnabled
	return j0
}

func (i *InterfaceIpSettings) Decode(v any, caller idl.Caller) error {
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
	err = encoding.In("configMethod", j0)
	if err != nil {
		return err
	}
	i.ConfigMethod, err = encoding.AsEnum[IpConfigMethod](j0["configMethod"])
	if err != nil {
		return err
	}
	err = encoding.In("staticAddrCidr", j0)
	if err != nil {
		return err
	}
	err = i.StaticAddrCidr.Decode(j0["staticAddrCidr"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("staticDefaultGatewayAddr", j0)
	if err != nil {
		return err
	}
	i.StaticDefaultGatewayAddr, err = encoding.Is[string](j0["staticDefaultGatewayAddr"])
	if err != nil {
		return err
	}
	err = encoding.In("dhcpPreferredHostname", j0)
	if err != nil {
		return err
	}
	i.DhcpPreferredHostname, err = encoding.Is[string](j0["dhcpPreferredHostname"])
	if err != nil {
		return err
	}
	err = encoding.In("stablePrivacyEnabled", j0)
	if err != nil {
		return err
	}
	i.StablePrivacyEnabled, err = encoding.Is[bool](j0["stablePrivacyEnabled"])
	if err != nil {
		return err
	}
	return nil
}

func (i *InterfaceIPv4Info) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["enabled"] = i.Enabled
	j0["configMethod"] = i.ConfigMethod
	s1 := make([]any, 0, len(i.AddrsCidr))
	for _, e1 := range i.AddrsCidr {
		s1 = append(s1, e1.Encode())
	}
	j0["addrsCidr"] = s1
	j0["dhcpServerAddr"] = i.DhcpServerAddr
	j0["dhcpPreferredHostname"] = i.DhcpPreferredHostname
	return j0
}

func (i *InterfaceIPv4Info) Decode(v any, caller idl.Caller) error {
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
	err = encoding.In("configMethod", j0)
	if err != nil {
		return err
	}
	i.ConfigMethod, err = encoding.AsEnum[IpConfigMethod](j0["configMethod"])
	if err != nil {
		return err
	}
	err = encoding.In("addrsCidr", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["addrsCidr"])
	if err != nil {
		return err
	}
	i.AddrsCidr = make([]IpAddrCidr, 0, len(s1))
	for _, a1 := range s1 {
		var e1 IpAddrCidr
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		i.AddrsCidr = append(i.AddrsCidr, e1)
	}
	err = encoding.In("dhcpServerAddr", j0)
	if err != nil {
		return err
	}
	i.DhcpServerAddr, err = encoding.Is[string](j0["dhcpServerAddr"])
	if err != nil {
		return err
	}
	err = encoding.In("dhcpPreferredHostname", j0)
	if err != nil {
		return err
	}
	i.DhcpPreferredHostname, err = encoding.Is[string](j0["dhcpPreferredHostname"])
	if err != nil {
		return err
	}
	return nil
}

func (i *InterfaceIPv6Info) Encode() map[string]any {
	j0 := make(map[string]any, 7)
	j0["enabled"] = i.Enabled
	j0["configMethod"] = i.ConfigMethod
	s1 := make([]any, 0, len(i.AddrsCidr))
	for _, e1 := range i.AddrsCidr {
		s1 = append(s1, e1.Encode())
	}
	j0["addrsCidr"] = s1
	j0["dhcpServerId"] = i.DhcpServerId
	j0["dhcpPreferredHostname"] = i.DhcpPreferredHostname
	j0["raManaged"] = i.RaManaged
	j0["raOtherConf"] = i.RaOtherConf
	return j0
}

func (i *InterfaceIPv6Info) Decode(v any, caller idl.Caller) error {
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
	err = encoding.In("configMethod", j0)
	if err != nil {
		return err
	}
	i.ConfigMethod, err = encoding.AsEnum[IpConfigMethod](j0["configMethod"])
	if err != nil {
		return err
	}
	err = encoding.In("addrsCidr", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["addrsCidr"])
	if err != nil {
		return err
	}
	i.AddrsCidr = make([]IpAddrCidr, 0, len(s1))
	for _, a1 := range s1 {
		var e1 IpAddrCidr
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		i.AddrsCidr = append(i.AddrsCidr, e1)
	}
	err = encoding.In("dhcpServerId", j0)
	if err != nil {
		return err
	}
	i.DhcpServerId, err = encoding.Is[string](j0["dhcpServerId"])
	if err != nil {
		return err
	}
	err = encoding.In("dhcpPreferredHostname", j0)
	if err != nil {
		return err
	}
	i.DhcpPreferredHostname, err = encoding.Is[string](j0["dhcpPreferredHostname"])
	if err != nil {
		return err
	}
	err = encoding.In("raManaged", j0)
	if err != nil {
		return err
	}
	i.RaManaged, err = encoding.Is[bool](j0["raManaged"])
	if err != nil {
		return err
	}
	err = encoding.In("raOtherConf", j0)
	if err != nil {
		return err
	}
	i.RaOtherConf, err = encoding.Is[bool](j0["raOtherConf"])
	if err != nil {
		return err
	}
	return nil
}

func (i *InterfaceSettings) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["enabled"] = i.Enabled
	j0["ipv4"] = i.Ipv4.Encode()
	j0["ipv6"] = i.Ipv6.Encode()
	return j0
}

func (i *InterfaceSettings) Decode(v any, caller idl.Caller) error {
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
	err = encoding.In("ipv4", j0)
	if err != nil {
		return err
	}
	err = i.Ipv4.Decode(j0["ipv4"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("ipv6", j0)
	if err != nil {
		return err
	}
	err = i.Ipv6.Decode(j0["ipv6"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (i *InterfaceInfo) Encode() map[string]any {
	j0 := make(map[string]any, 10)
	j0["name"] = i.Name
	j0["label"] = i.Label
	j0["type"] = i.Type
	j0["enabled"] = i.Enabled
	j0["controllingIfName"] = i.ControllingIfName
	j0["state"] = i.State
	j0["macAddr"] = i.MacAddr
	j0["mtu"] = i.Mtu
	j0["ipv4"] = i.Ipv4.Encode()
	j0["ipv6"] = i.Ipv6.Encode()
	return j0
}

func (i *InterfaceInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	i.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("label", j0)
	if err != nil {
		return err
	}
	i.Label, err = encoding.Is[string](j0["label"])
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	i.Type, err = encoding.AsEnum[InterfaceType](j0["type"])
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
	err = encoding.In("controllingIfName", j0)
	if err != nil {
		return err
	}
	i.ControllingIfName, err = encoding.Is[string](j0["controllingIfName"])
	if err != nil {
		return err
	}
	err = encoding.In("state", j0)
	if err != nil {
		return err
	}
	i.State, err = encoding.AsEnum[InterfaceOpState](j0["state"])
	if err != nil {
		return err
	}
	err = encoding.In("macAddr", j0)
	if err != nil {
		return err
	}
	i.MacAddr, err = encoding.Is[string](j0["macAddr"])
	if err != nil {
		return err
	}
	err = encoding.In("mtu", j0)
	if err != nil {
		return err
	}
	i.Mtu, err = encoding.AsInt32(j0["mtu"])
	if err != nil {
		return err
	}
	err = encoding.In("ipv4", j0)
	if err != nil {
		return err
	}
	err = i.Ipv4.Decode(j0["ipv4"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("ipv6", j0)
	if err != nil {
		return err
	}
	err = i.Ipv6.Decode(j0["ipv6"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (e *EapAuthSettings) Encode() map[string]any {
	j0 := make(map[string]any, 14)
	j0["identity"] = e.Identity
	j0["password"] = e.Password
	j0["clearPassword"] = e.ClearPassword
	j0["clientCertChain"] = e.ClientCertChain
	j0["clientPrivKey"] = e.ClientPrivKey
	j0["clearClientPrivKey"] = e.ClearClientPrivKey
	j0["clientPrivKeyPassword"] = e.ClientPrivKeyPassword
	j0["outerMethod"] = e.OuterMethod
	j0["innerMethod"] = e.InnerMethod
	j0["caCertChain"] = e.CaCertChain
	j0["forceTrustedCert"] = e.ForceTrustedCert
	j0["allowOffTimeRangeCerts"] = e.AllowOffTimeRangeCerts
	j0["allowNotYetValidCertsIfTimeBeforeBuild"] = e.AllowNotYetValidCertsIfTimeBeforeBuild
	j0["authServerName"] = e.AuthServerName
	return j0
}

func (e *EapAuthSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("identity", j0)
	if err != nil {
		return err
	}
	e.Identity, err = encoding.Is[string](j0["identity"])
	if err != nil {
		return err
	}
	err = encoding.In("password", j0)
	if err != nil {
		return err
	}
	e.Password, err = encoding.Is[string](j0["password"])
	if err != nil {
		return err
	}
	err = encoding.In("clearPassword", j0)
	if err != nil {
		return err
	}
	e.ClearPassword, err = encoding.Is[bool](j0["clearPassword"])
	if err != nil {
		return err
	}
	err = encoding.In("clientCertChain", j0)
	if err != nil {
		return err
	}
	e.ClientCertChain, err = encoding.Is[string](j0["clientCertChain"])
	if err != nil {
		return err
	}
	err = encoding.In("clientPrivKey", j0)
	if err != nil {
		return err
	}
	e.ClientPrivKey, err = encoding.Is[string](j0["clientPrivKey"])
	if err != nil {
		return err
	}
	err = encoding.In("clearClientPrivKey", j0)
	if err != nil {
		return err
	}
	e.ClearClientPrivKey, err = encoding.Is[bool](j0["clearClientPrivKey"])
	if err != nil {
		return err
	}
	err = encoding.In("clientPrivKeyPassword", j0)
	if err != nil {
		return err
	}
	e.ClientPrivKeyPassword, err = encoding.Is[string](j0["clientPrivKeyPassword"])
	if err != nil {
		return err
	}
	err = encoding.In("outerMethod", j0)
	if err != nil {
		return err
	}
	e.OuterMethod, err = encoding.AsEnum[EapOuterAuthMethod](j0["outerMethod"])
	if err != nil {
		return err
	}
	err = encoding.In("innerMethod", j0)
	if err != nil {
		return err
	}
	e.InnerMethod, err = encoding.AsEnum[EapInnerAuthMethod](j0["innerMethod"])
	if err != nil {
		return err
	}
	err = encoding.In("caCertChain", j0)
	if err != nil {
		return err
	}
	e.CaCertChain, err = encoding.Is[string](j0["caCertChain"])
	if err != nil {
		return err
	}
	err = encoding.In("forceTrustedCert", j0)
	if err != nil {
		return err
	}
	e.ForceTrustedCert, err = encoding.Is[bool](j0["forceTrustedCert"])
	if err != nil {
		return err
	}
	err = encoding.In("allowOffTimeRangeCerts", j0)
	if err != nil {
		return err
	}
	e.AllowOffTimeRangeCerts, err = encoding.Is[bool](j0["allowOffTimeRangeCerts"])
	if err != nil {
		return err
	}
	err = encoding.In("allowNotYetValidCertsIfTimeBeforeBuild", j0)
	if err != nil {
		return err
	}
	e.AllowNotYetValidCertsIfTimeBeforeBuild, err = encoding.Is[bool](j0["allowNotYetValidCertsIfTimeBeforeBuild"])
	if err != nil {
		return err
	}
	err = encoding.In("authServerName", j0)
	if err != nil {
		return err
	}
	e.AuthServerName, err = encoding.Is[string](j0["authServerName"])
	if err != nil {
		return err
	}
	return nil
}

func (e *EthLinkMode) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["speed"] = e.Speed
	j0["duplexMode"] = e.DuplexMode
	return j0
}

func (e *EthLinkMode) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("speed", j0)
	if err != nil {
		return err
	}
	e.Speed, err = encoding.AsEnum[EthSpeed](j0["speed"])
	if err != nil {
		return err
	}
	err = encoding.In("duplexMode", j0)
	if err != nil {
		return err
	}
	e.DuplexMode, err = encoding.AsEnum[EthDuplexMode](j0["duplexMode"])
	if err != nil {
		return err
	}
	return nil
}

func (e *EthSettings) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["linkMode"] = e.LinkMode.Encode()
	j0["mtu"] = e.Mtu
	j0["authType"] = e.AuthType
	j0["eap"] = e.Eap.Encode()
	j0["lldpEnabled"] = e.LldpEnabled
	return j0
}

func (e *EthSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("linkMode", j0)
	if err != nil {
		return err
	}
	err = e.LinkMode.Decode(j0["linkMode"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("mtu", j0)
	if err != nil {
		return err
	}
	e.Mtu, err = encoding.AsInt32(j0["mtu"])
	if err != nil {
		return err
	}
	err = encoding.In("authType", j0)
	if err != nil {
		return err
	}
	e.AuthType, err = encoding.AsEnum[EthAuthType](j0["authType"])
	if err != nil {
		return err
	}
	err = encoding.In("eap", j0)
	if err != nil {
		return err
	}
	err = e.Eap.Decode(j0["eap"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("lldpEnabled", j0)
	if err != nil {
		return err
	}
	e.LldpEnabled, err = encoding.Is[bool](j0["lldpEnabled"])
	if err != nil {
		return err
	}
	return nil
}

func (e *EthInfo) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["eapStatus"] = e.EapStatus
	j0["linkMode"] = e.LinkMode.Encode()
	j0["linkModeValid"] = e.LinkModeValid
	j0["autonegEnabled"] = e.AutonegEnabled
	j0["linkDetected"] = e.LinkDetected
	s1 := make([]any, 0, len(e.SupportedLinkModes))
	for _, e1 := range e.SupportedLinkModes {
		s1 = append(s1, e1.Encode())
	}
	j0["supportedLinkModes"] = s1
	return j0
}

func (e *EthInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("eapStatus", j0)
	if err != nil {
		return err
	}
	e.EapStatus, err = encoding.AsEnum[EapStatus](j0["eapStatus"])
	if err != nil {
		return err
	}
	err = encoding.In("linkMode", j0)
	if err != nil {
		return err
	}
	err = e.LinkMode.Decode(j0["linkMode"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("linkModeValid", j0)
	if err != nil {
		return err
	}
	e.LinkModeValid, err = encoding.Is[bool](j0["linkModeValid"])
	if err != nil {
		return err
	}
	err = encoding.In("autonegEnabled", j0)
	if err != nil {
		return err
	}
	e.AutonegEnabled, err = encoding.Is[bool](j0["autonegEnabled"])
	if err != nil {
		return err
	}
	err = encoding.In("linkDetected", j0)
	if err != nil {
		return err
	}
	e.LinkDetected, err = encoding.Is[bool](j0["linkDetected"])
	if err != nil {
		return err
	}
	err = encoding.In("supportedLinkModes", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["supportedLinkModes"])
	if err != nil {
		return err
	}
	e.SupportedLinkModes = make([]EthLinkMode, 0, len(s1))
	for _, a1 := range s1 {
		var e1 EthLinkMode
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		e.SupportedLinkModes = append(e.SupportedLinkModes, e1)
	}
	return nil
}

func (w *WlanSettings) Encode() map[string]any {
	j0 := make(map[string]any, 9)
	j0["enableHT"] = w.EnableHT
	j0["ssid"] = w.Ssid
	j0["bssid"] = w.Bssid
	j0["mtu"] = w.Mtu
	j0["secProtocol"] = w.SecProtocol
	j0["authType"] = w.AuthType
	j0["psk"] = w.Psk
	j0["clearPsk"] = w.ClearPsk
	j0["eap"] = w.Eap.Encode()
	return j0
}

func (w *WlanSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("enableHT", j0)
	if err != nil {
		return err
	}
	w.EnableHT, err = encoding.Is[bool](j0["enableHT"])
	if err != nil {
		return err
	}
	err = encoding.In("ssid", j0)
	if err != nil {
		return err
	}
	w.Ssid, err = encoding.Is[string](j0["ssid"])
	if err != nil {
		return err
	}
	err = encoding.In("bssid", j0)
	if err != nil {
		return err
	}
	w.Bssid, err = encoding.Is[string](j0["bssid"])
	if err != nil {
		return err
	}
	err = encoding.In("mtu", j0)
	if err != nil {
		return err
	}
	w.Mtu, err = encoding.AsInt32(j0["mtu"])
	if err != nil {
		return err
	}
	err = encoding.In("secProtocol", j0)
	if err != nil {
		return err
	}
	w.SecProtocol, err = encoding.AsEnum[WlanSecProtocol](j0["secProtocol"])
	if err != nil {
		return err
	}
	err = encoding.In("authType", j0)
	if err != nil {
		return err
	}
	w.AuthType, err = encoding.AsEnum[WlanAuthType](j0["authType"])
	if err != nil {
		return err
	}
	err = encoding.In("psk", j0)
	if err != nil {
		return err
	}
	w.Psk, err = encoding.Is[string](j0["psk"])
	if err != nil {
		return err
	}
	err = encoding.In("clearPsk", j0)
	if err != nil {
		return err
	}
	w.ClearPsk, err = encoding.Is[bool](j0["clearPsk"])
	if err != nil {
		return err
	}
	err = encoding.In("eap", j0)
	if err != nil {
		return err
	}
	err = w.Eap.Decode(j0["eap"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (w *WlanInfo) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["associated"] = w.Associated
	j0["ssid"] = w.Ssid
	j0["bssid"] = w.Bssid
	j0["channel"] = w.Channel
	j0["channelWidth"] = w.ChannelWidth
	return j0
}

func (w *WlanInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("associated", j0)
	if err != nil {
		return err
	}
	w.Associated, err = encoding.Is[bool](j0["associated"])
	if err != nil {
		return err
	}
	err = encoding.In("ssid", j0)
	if err != nil {
		return err
	}
	w.Ssid, err = encoding.Is[string](j0["ssid"])
	if err != nil {
		return err
	}
	err = encoding.In("bssid", j0)
	if err != nil {
		return err
	}
	w.Bssid, err = encoding.Is[string](j0["bssid"])
	if err != nil {
		return err
	}
	err = encoding.In("channel", j0)
	if err != nil {
		return err
	}
	w.Channel, err = encoding.AsInt32(j0["channel"])
	if err != nil {
		return err
	}
	err = encoding.In("channelWidth", j0)
	if err != nil {
		return err
	}
	w.ChannelWidth, err = encoding.AsEnum[WlanChannelWidth](j0["channelWidth"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["common"] = s.Common.Encode()
	s1 := make([]any, 0, len(s.IfMap))
	for k1, v1 := range s.IfMap {
		s1 = append(s1, map[string]any{"key": k1, "value": v1.Encode()})
	}
	j0["ifMap"] = s1
	s2 := make([]any, 0, len(s.EthMap))
	for k2, v2 := range s.EthMap {
		s2 = append(s2, map[string]any{"key": k2, "value": v2.Encode()})
	}
	j0["ethMap"] = s2
	s3 := make([]any, 0, len(s.WlanMap))
	for k3, v3 := range s.WlanMap {
		s3 = append(s3, map[string]any{"key": k3, "value": v3.Encode()})
	}
	j0["wlanMap"] = s3
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("common", j0)
	if err != nil {
		return err
	}
	err = s.Common.Decode(j0["common"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("ifMap", j0)
	if err != nil {
		return err
	}
	i1, e1, l1 := encoding.AsMapItems(j0["ifMap"])
	s.IfMap = make(map[string]InterfaceSettings, l1)
	for a1, b1 := range i1 {
		var k1 string
		k1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		var v1 InterfaceSettings
		err = v1.Decode(b1, caller)
		if err != nil {
			return err
		}
		s.IfMap[k1] = v1
	}
	err = e1()
	if err != nil {
		return err
	}
	err = encoding.In("ethMap", j0)
	if err != nil {
		return err
	}
	i2, e2, l2 := encoding.AsMapItems(j0["ethMap"])
	s.EthMap = make(map[string]EthSettings, l2)
	for a2, b2 := range i2 {
		var k2 string
		k2, err = encoding.Is[string](a2)
		if err != nil {
			return err
		}
		var v2 EthSettings
		err = v2.Decode(b2, caller)
		if err != nil {
			return err
		}
		s.EthMap[k2] = v2
	}
	err = e2()
	if err != nil {
		return err
	}
	err = encoding.In("wlanMap", j0)
	if err != nil {
		return err
	}
	i3, e3, l3 := encoding.AsMapItems(j0["wlanMap"])
	s.WlanMap = make(map[string]WlanSettings, l3)
	for a3, b3 := range i3 {
		var k3 string
		k3, err = encoding.Is[string](a3)
		if err != nil {
			return err
		}
		var v3 WlanSettings
		err = v3.Decode(b3, caller)
		if err != nil {
			return err
		}
		s.WlanMap[k3] = v3
	}
	err = e3()
	if err != nil {
		return err
	}
	return nil
}

func (i *Info) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["common"] = i.Common.Encode()
	s1 := make([]any, 0, len(i.IfMap))
	for k1, v1 := range i.IfMap {
		s1 = append(s1, map[string]any{"key": k1, "value": v1.Encode()})
	}
	j0["ifMap"] = s1
	s2 := make([]any, 0, len(i.EthMap))
	for k2, v2 := range i.EthMap {
		s2 = append(s2, map[string]any{"key": k2, "value": v2.Encode()})
	}
	j0["ethMap"] = s2
	s3 := make([]any, 0, len(i.WlanMap))
	for k3, v3 := range i.WlanMap {
		s3 = append(s3, map[string]any{"key": k3, "value": v3.Encode()})
	}
	j0["wlanMap"] = s3
	return j0
}

func (i *Info) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("common", j0)
	if err != nil {
		return err
	}
	err = i.Common.Decode(j0["common"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("ifMap", j0)
	if err != nil {
		return err
	}
	i1, e1, l1 := encoding.AsMapItems(j0["ifMap"])
	i.IfMap = make(map[string]InterfaceInfo, l1)
	for a1, b1 := range i1 {
		var k1 string
		k1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		var v1 InterfaceInfo
		err = v1.Decode(b1, caller)
		if err != nil {
			return err
		}
		i.IfMap[k1] = v1
	}
	err = e1()
	if err != nil {
		return err
	}
	err = encoding.In("ethMap", j0)
	if err != nil {
		return err
	}
	i2, e2, l2 := encoding.AsMapItems(j0["ethMap"])
	i.EthMap = make(map[string]EthInfo, l2)
	for a2, b2 := range i2 {
		var k2 string
		k2, err = encoding.Is[string](a2)
		if err != nil {
			return err
		}
		var v2 EthInfo
		err = v2.Decode(b2, caller)
		if err != nil {
			return err
		}
		i.EthMap[k2] = v2
	}
	err = e2()
	if err != nil {
		return err
	}
	err = encoding.In("wlanMap", j0)
	if err != nil {
		return err
	}
	i3, e3, l3 := encoding.AsMapItems(j0["wlanMap"])
	i.WlanMap = make(map[string]WlanInfo, l3)
	for a3, b3 := range i3 {
		var k3 string
		k3, err = encoding.Is[string](a3)
		if err != nil {
			return err
		}
		var v3 WlanInfo
		err = v3.Decode(b3, caller)
		if err != nil {
			return err
		}
		i.WlanMap[k3] = v3
	}
	err = e3()
	if err != nil {
		return err
	}
	return nil
}

func (p *PortForwardingProtocolMapping) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["appProtoId"] = p.AppProtoId
	j0["appProtoName"] = p.AppProtoName
	j0["transportProtoName"] = p.TransportProtoName
	return j0
}

func (p *PortForwardingProtocolMapping) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("appProtoId", j0)
	if err != nil {
		return err
	}
	p.AppProtoId, err = encoding.AsInt32(j0["appProtoId"])
	if err != nil {
		return err
	}
	err = encoding.In("appProtoName", j0)
	if err != nil {
		return err
	}
	p.AppProtoName, err = encoding.Is[string](j0["appProtoName"])
	if err != nil {
		return err
	}
	err = encoding.In("transportProtoName", j0)
	if err != nil {
		return err
	}
	p.TransportProtoName, err = encoding.Is[string](j0["transportProtoName"])
	if err != nil {
		return err
	}
	return nil
}

func (c *_CommonInfoChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.Event = valobj.For[event.Event]()
	err := c.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("commonInfo", value)
	if err != nil {
		return err
	}
	err = c.commonInfo.Decode(value["commonInfo"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (i *_InterfaceInfoChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	i.Event = valobj.For[event.Event]()
	err := i.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("ifInfo", value)
	if err != nil {
		return err
	}
	err = i.ifInfo.Decode(value["ifInfo"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (e *_EthInfoChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	e.Event = valobj.For[event.Event]()
	err := e.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("ifName", value)
	if err != nil {
		return err
	}
	e.ifName, err = encoding.Is[string](value["ifName"])
	if err != nil {
		return err
	}
	err = encoding.In("ifLabel", value)
	if err != nil {
		return err
	}
	e.ifLabel, err = encoding.Is[string](value["ifLabel"])
	if err != nil {
		return err
	}
	err = encoding.In("ethInfo", value)
	if err != nil {
		return err
	}
	err = e.ethInfo.Decode(value["ethInfo"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (w *_WlanInfoChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	w.Event = valobj.For[event.Event]()
	err := w.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("ifName", value)
	if err != nil {
		return err
	}
	w.ifName, err = encoding.Is[string](value["ifName"])
	if err != nil {
		return err
	}
	err = encoding.In("wlanInfo", value)
	if err != nil {
		return err
	}
	err = w.wlanInfo.Decode(value["wlanInfo"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (l *_LinkStateChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	l.Event = valobj.For[event.Event]()
	err := l.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("ifName", value)
	if err != nil {
		return err
	}
	l.ifName, err = encoding.Is[string](value["ifName"])
	if err != nil {
		return err
	}
	err = encoding.In("ifLabel", value)
	if err != nil {
		return err
	}
	l.ifLabel, err = encoding.Is[string](value["ifLabel"])
	if err != nil {
		return err
	}
	err = encoding.In("ifType", value)
	if err != nil {
		return err
	}
	l.ifType, err = encoding.AsEnum[InterfaceType](value["ifType"])
	if err != nil {
		return err
	}
	err = encoding.In("ifState", value)
	if err != nil {
		return err
	}
	l.ifState, err = encoding.AsEnum[InterfaceOpState](value["ifState"])
	if err != nil {
		return err
	}
	return nil
}

func (p *_PortForwardingExpansionUnitPresenceStateChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	p.Event = valobj.For[event.Event]()
	err := p.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("expansionUnitPresent", value)
	if err != nil {
		return err
	}
	p.expansionUnitPresent, err = encoding.Is[bool](value["expansionUnitPresent"])
	if err != nil {
		return err
	}
	return nil
}
