// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package smtp

import (
	"github.com/arminguenther/xeruspower-go/v40300/idl"
	"github.com/arminguenther/xeruspower-go/v40300/internal/encoding"
)

func (c *Configuration) Encode() map[string]any {
	j0 := make(map[string]any, 11)
	j0["host"] = c.Host
	j0["port"] = c.Port
	j0["useTls"] = c.UseTls
	j0["allowOffTimeRangeCerts"] = c.AllowOffTimeRangeCerts
	j0["caCertChain"] = c.CaCertChain
	j0["sender"] = c.Sender
	j0["useAuth"] = c.UseAuth
	j0["username"] = c.Username
	j0["password"] = c.Password
	j0["retryCount"] = c.RetryCount
	j0["retryInterval"] = c.RetryInterval
	return j0
}

func (c *Configuration) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("host", j0)
	if err != nil {
		return err
	}
	c.Host, err = encoding.Is[string](j0["host"])
	if err != nil {
		return err
	}
	err = encoding.In("port", j0)
	if err != nil {
		return err
	}
	c.Port, err = encoding.AsInt32(j0["port"])
	if err != nil {
		return err
	}
	err = encoding.In("useTls", j0)
	if err != nil {
		return err
	}
	c.UseTls, err = encoding.Is[bool](j0["useTls"])
	if err != nil {
		return err
	}
	err = encoding.In("allowOffTimeRangeCerts", j0)
	if err != nil {
		return err
	}
	c.AllowOffTimeRangeCerts, err = encoding.Is[bool](j0["allowOffTimeRangeCerts"])
	if err != nil {
		return err
	}
	err = encoding.In("caCertChain", j0)
	if err != nil {
		return err
	}
	c.CaCertChain, err = encoding.Is[string](j0["caCertChain"])
	if err != nil {
		return err
	}
	err = encoding.In("sender", j0)
	if err != nil {
		return err
	}
	c.Sender, err = encoding.Is[string](j0["sender"])
	if err != nil {
		return err
	}
	err = encoding.In("useAuth", j0)
	if err != nil {
		return err
	}
	c.UseAuth, err = encoding.Is[bool](j0["useAuth"])
	if err != nil {
		return err
	}
	err = encoding.In("username", j0)
	if err != nil {
		return err
	}
	c.Username, err = encoding.Is[string](j0["username"])
	if err != nil {
		return err
	}
	err = encoding.In("password", j0)
	if err != nil {
		return err
	}
	c.Password, err = encoding.Is[string](j0["password"])
	if err != nil {
		return err
	}
	err = encoding.In("retryCount", j0)
	if err != nil {
		return err
	}
	c.RetryCount, err = encoding.AsInt32(j0["retryCount"])
	if err != nil {
		return err
	}
	err = encoding.In("retryInterval", j0)
	if err != nil {
		return err
	}
	c.RetryInterval, err = encoding.AsInt32(j0["retryInterval"])
	if err != nil {
		return err
	}
	return nil
}

func (t *TestResult) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["status"] = t.Status
	j0["message"] = t.Message
	return j0
}

func (t *TestResult) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("status", j0)
	if err != nil {
		return err
	}
	t.Status, err = encoding.AsInt32(j0["status"])
	if err != nil {
		return err
	}
	err = encoding.In("message", j0)
	if err != nil {
		return err
	}
	t.Message, err = encoding.Is[string](j0["message"])
	if err != nil {
		return err
	}
	return nil
}
