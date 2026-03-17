// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package role

import (
	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding"
)

func (p *Privilege) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["name"] = p.Name
	j0["args"] = encoding.NonNilSlice(p.Args)
	return j0
}

func (p *Privilege) Decode(v any, caller idl.Caller) error {
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
	err = encoding.In("args", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["args"])
	if err != nil {
		return err
	}
	p.Args = make([]string, 0, len(s1))
	for _, a1 := range s1 {
		var e1 string
		e1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		p.Args = append(p.Args, e1)
	}
	return nil
}

func (i *Info) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["description"] = i.Description
	j0["locked"] = i.Locked
	s1 := make([]any, 0, len(i.Privileges))
	for _, e1 := range i.Privileges {
		s1 = append(s1, e1.Encode())
	}
	j0["privileges"] = s1
	return j0
}

func (i *Info) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("description", j0)
	if err != nil {
		return err
	}
	i.Description, err = encoding.Is[string](j0["description"])
	if err != nil {
		return err
	}
	err = encoding.In("locked", j0)
	if err != nil {
		return err
	}
	i.Locked, err = encoding.Is[bool](j0["locked"])
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
	i.Privileges = make([]Privilege, 0, len(s1))
	for _, a1 := range s1 {
		var e1 Privilege
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		i.Privileges = append(i.Privileges, e1)
	}
	return nil
}
