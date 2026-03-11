// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package usermanager

import (
	"github.com/arminguenther/xeruspower-go/v40510/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/valobj"
)

func (a *Account) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["name"] = a.Name
	j0["info"] = a.Info.Encode()
	return j0
}

func (a *Account) Decode(v any, caller idl.Caller) error {
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
	err = encoding.In("info", j0)
	if err != nil {
		return err
	}
	err = a.Info.Decode(j0["info"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (a *_AccountEvent) Decode(value map[string]any, caller idl.Caller) error {
	a.UserEvent = valobj.For[userevent.UserEvent]()
	err := a.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("username", value)
	if err != nil {
		return err
	}
	a.username, err = encoding.Is[string](value["username"])
	if err != nil {
		return err
	}
	return nil
}

func (a *_AccountAdded) Decode(value map[string]any, caller idl.Caller) error {
	a.AccountEvent = valobj.For[AccountEvent]()
	err := a.AccountEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (a *_AccountRenamed) Decode(value map[string]any, caller idl.Caller) error {
	a.AccountEvent = valobj.For[AccountEvent]()
	err := a.AccountEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("newUsername", value)
	if err != nil {
		return err
	}
	a.newUsername, err = encoding.Is[string](value["newUsername"])
	if err != nil {
		return err
	}
	return nil
}

func (a *_AccountRemoved) Decode(value map[string]any, caller idl.Caller) error {
	a.AccountEvent = valobj.For[AccountEvent]()
	err := a.AccountEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (p *_PasswordChanged) Decode(value map[string]any, caller idl.Caller) error {
	p.AccountEvent = valobj.For[AccountEvent]()
	err := p.AccountEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (a *_AccountChanged) Decode(value map[string]any, caller idl.Caller) error {
	a.AccountEvent = valobj.For[AccountEvent]()
	err := a.AccountEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}
