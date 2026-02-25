// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package sessionmanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
)

func init() {
	object.Register(NewSessionManager)
}

type _SessionManager struct {
	idl.Object
}

// NewSessionManager returns a new SessionManager interface for the object at given RID.
func NewSessionManager(rid string, caller idl.Caller) SessionManager {
	return &_SessionManager{idl.NewObject(rid, caller)}
}

func (s *_SessionManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "session.SessionManager",
		Major: 2, Submajor: 0, Minor: 0,
	}
}

func (s *_SessionManager) NewSession(ctx context.Context) (int32, Session, string, error) {
	var ret int32
	var out0 Session
	var out1 string
	val, err := s.Caller().Call(ctx, s.RID(), "newSession", nil)
	if err != nil {
		return ret, out0, out1, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, out0, out1, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, out0, out1, err
	}
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, out0, out1, err
	}
	err = encoding.In("session", res)
	if err != nil {
		return ret, out0, out1, err
	}
	err = out0.Decode(res["session"], s.Caller())
	if err != nil {
		return ret, out0, out1, err
	}
	err = encoding.In("token", res)
	if err != nil {
		return ret, out0, out1, err
	}
	out1, err = encoding.Is[string](res["token"])
	if err != nil {
		return ret, out0, out1, err
	}
	return ret, out0, out1, nil
}

func (s *_SessionManager) GetCurrentSession(ctx context.Context) (Session, error) {
	var ret Session
	val, err := s.Caller().Call(ctx, s.RID(), "getCurrentSession", nil)
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	err = ret.Decode(res["_ret_"], s.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_SessionManager) GetSessions(ctx context.Context) ([]Session, error) {
	var ret []Session
	val, err := s.Caller().Call(ctx, s.RID(), "getSessions", nil)
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["_ret_"])
	if err != nil {
		return ret, err
	}
	ret = make([]Session, 0, len(s0))
	for _, a0 := range s0 {
		var e0 Session
		err = e0.Decode(a0, s.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (s *_SessionManager) CloseSession(ctx context.Context, in0 int32, in1 CloseReason) error {
	_, err := s.Caller().Call(ctx, s.RID(), "closeSession", map[string]any{
		"sessionId": in0,
		"reason":    in1,
	})
	return err
}

func (s *_SessionManager) CloseCurrentSession(ctx context.Context, in0 CloseReason) error {
	_, err := s.Caller().Call(ctx, s.RID(), "closeCurrentSession", map[string]any{
		"reason": in0,
	})
	return err
}

func (s *_SessionManager) TouchCurrentSession(ctx context.Context, in0 bool) error {
	_, err := s.Caller().Call(ctx, s.RID(), "touchCurrentSession", map[string]any{
		"userActivity": in0,
	})
	return err
}

func (s *_SessionManager) GetSessionHistory(ctx context.Context) ([]HistoryEntry, error) {
	var ret []HistoryEntry
	val, err := s.Caller().Call(ctx, s.RID(), "getSessionHistory", nil)
	if err != nil {
		return ret, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["_ret_"])
	if err != nil {
		return ret, err
	}
	ret = make([]HistoryEntry, 0, len(s0))
	for _, a0 := range s0 {
		var e0 HistoryEntry
		err = e0.Decode(a0, s.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}
