// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package sessionmanager

import (
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
)

func (s *Session) Encode() map[string]any {
	j0 := make(map[string]any, 8)
	j0["sessionId"] = s.SessionId
	j0["username"] = s.Username
	j0["remoteIp"] = s.RemoteIp
	j0["clientType"] = s.ClientType
	j0["creationTime"] = s.CreationTime.Unix()
	j0["timeout"] = s.Timeout
	j0["idle"] = s.Idle
	j0["userIdle"] = s.UserIdle
	return j0
}

func (s *Session) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("sessionId", j0)
	if err != nil {
		return err
	}
	s.SessionId, err = encoding.AsInt32(j0["sessionId"])
	if err != nil {
		return err
	}
	err = encoding.In("username", j0)
	if err != nil {
		return err
	}
	s.Username, err = encoding.Is[string](j0["username"])
	if err != nil {
		return err
	}
	err = encoding.In("remoteIp", j0)
	if err != nil {
		return err
	}
	s.RemoteIp, err = encoding.Is[string](j0["remoteIp"])
	if err != nil {
		return err
	}
	err = encoding.In("clientType", j0)
	if err != nil {
		return err
	}
	s.ClientType, err = encoding.Is[string](j0["clientType"])
	if err != nil {
		return err
	}
	err = encoding.In("creationTime", j0)
	if err != nil {
		return err
	}
	s.CreationTime, err = encoding.AsTime(j0["creationTime"])
	if err != nil {
		return err
	}
	err = encoding.In("timeout", j0)
	if err != nil {
		return err
	}
	s.Timeout, err = encoding.AsInt32(j0["timeout"])
	if err != nil {
		return err
	}
	err = encoding.In("idle", j0)
	if err != nil {
		return err
	}
	s.Idle, err = encoding.AsInt32(j0["idle"])
	if err != nil {
		return err
	}
	err = encoding.In("userIdle", j0)
	if err != nil {
		return err
	}
	s.UserIdle, err = encoding.AsInt32(j0["userIdle"])
	if err != nil {
		return err
	}
	return nil
}

func (h *HistoryEntry) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["creationTime"] = h.CreationTime.Unix()
	j0["remoteIp"] = h.RemoteIp
	j0["clientType"] = h.ClientType
	return j0
}

func (h *HistoryEntry) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("creationTime", j0)
	if err != nil {
		return err
	}
	h.CreationTime, err = encoding.AsTime(j0["creationTime"])
	if err != nil {
		return err
	}
	err = encoding.In("remoteIp", j0)
	if err != nil {
		return err
	}
	h.RemoteIp, err = encoding.Is[string](j0["remoteIp"])
	if err != nil {
		return err
	}
	err = encoding.In("clientType", j0)
	if err != nil {
		return err
	}
	h.ClientType, err = encoding.Is[string](j0["clientType"])
	if err != nil {
		return err
	}
	return nil
}
