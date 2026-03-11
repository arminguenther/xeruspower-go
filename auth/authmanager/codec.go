// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package authmanager

import (
	"github.com/arminguenther/xeruspower-go/v40410/idl"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding"
)

func (p *Policy) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["type"] = p.Type
	j0["disallowLocalAuthAdminRemoteAccess"] = p.DisallowLocalAuthAdminRemoteAccess
	j0["useLocalIfRemoteFailed"] = p.UseLocalIfRemoteFailed
	return j0
}

func (p *Policy) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	p.Type, err = encoding.AsEnum[Type](j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("disallowLocalAuthAdminRemoteAccess", j0)
	if err != nil {
		return err
	}
	p.DisallowLocalAuthAdminRemoteAccess, err = encoding.Is[bool](j0["disallowLocalAuthAdminRemoteAccess"])
	if err != nil {
		return err
	}
	err = encoding.In("useLocalIfRemoteFailed", j0)
	if err != nil {
		return err
	}
	p.UseLocalIfRemoteFailed, err = encoding.Is[bool](j0["useLocalIfRemoteFailed"])
	if err != nil {
		return err
	}
	return nil
}
