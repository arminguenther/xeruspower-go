// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package rolemanager

import (
	"github.com/arminguenther/xeruspower-go/v40220/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/valobj"
)

func (r *_RoleEvent) Decode(value map[string]any, caller idl.Caller) error {
	r.UserEvent = valobj.For[userevent.UserEvent]()
	err := r.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("rolename", value)
	if err != nil {
		return err
	}
	r.rolename, err = encoding.Is[string](value["rolename"])
	if err != nil {
		return err
	}
	return nil
}

func (r *_RoleAdded) Decode(value map[string]any, caller idl.Caller) error {
	r.RoleEvent = valobj.For[RoleEvent]()
	err := r.RoleEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (r *_RoleRemoved) Decode(value map[string]any, caller idl.Caller) error {
	r.RoleEvent = valobj.For[RoleEvent]()
	err := r.RoleEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (r *_RoleChanged) Decode(value map[string]any, caller idl.Caller) error {
	r.RoleEvent = valobj.For[RoleEvent]()
	err := r.RoleEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldSettings", value)
	if err != nil {
		return err
	}
	err = r.oldSettings.Decode(value["oldSettings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newSettings", value)
	if err != nil {
		return err
	}
	err = r.newSettings.Decode(value["newSettings"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (a *ArgumentDesc) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["name"] = a.Name
	j0["desc"] = a.Desc
	return j0
}

func (a *ArgumentDesc) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	a.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("desc", j0)
	if err != nil {
		return err
	}
	a.Desc, err = encoding.Is[string](j0["desc"])
	if err != nil {
		return err
	}
	return nil
}

func (p *PrivilegeDesc) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["name"] = p.Name
	j0["desc"] = p.Desc
	s1 := make([]any, 0, len(p.Args))
	for _, e1 := range p.Args {
		s1 = append(s1, e1.Encode())
	}
	j0["args"] = s1
	return j0
}

func (p *PrivilegeDesc) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	p.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("desc", j0)
	if err != nil {
		return err
	}
	p.Desc, err = encoding.Is[string](j0["desc"])
	if err != nil {
		return err
	}
	err = encoding.In("args", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["args"])
	if err != nil {
		return err
	}
	p.Args = make([]ArgumentDesc, 0, len(s1))
	for _, a1 := range s1 {
		var e1 ArgumentDesc
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		p.Args = append(p.Args, e1)
	}
	return nil
}

func (r *RoleAccount) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["id"] = r.Id
	j0["name"] = r.Name
	j0["info"] = r.Info.Encode()
	return j0
}

func (r *RoleAccount) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("id", j0)
	if err != nil {
		return err
	}
	r.Id, err = encoding.AsInt32(j0["id"])
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	r.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("info", j0)
	if err != nil {
		return err
	}
	err = r.Info.Decode(j0["info"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (i *Info) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	s1 := make([]any, 0, len(i.Privileges))
	for _, e1 := range i.Privileges {
		s1 = append(s1, e1.Encode())
	}
	j0["privileges"] = s1
	s2 := make([]any, 0, len(i.Roles))
	for _, e2 := range i.Roles {
		s2 = append(s2, e2.Encode())
	}
	j0["roles"] = s2
	return j0
}

func (i *Info) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("privileges", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["privileges"])
	if err != nil {
		return err
	}
	i.Privileges = make([]PrivilegeDesc, 0, len(s1))
	for _, a1 := range s1 {
		var e1 PrivilegeDesc
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		i.Privileges = append(i.Privileges, e1)
	}
	err = encoding.In("roles", j0)
	if err != nil {
		return err
	}
	var s2 []any
	s2, err = encoding.Is[[]any](j0["roles"])
	if err != nil {
		return err
	}
	i.Roles = make([]RoleAccount, 0, len(s2))
	for _, a2 := range s2 {
		var e2 RoleAccount
		err = e2.Decode(a2, caller)
		if err != nil {
			return err
		}
		i.Roles = append(i.Roles, e2)
	}
	return nil
}
