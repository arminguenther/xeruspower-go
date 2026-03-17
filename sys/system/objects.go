// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package system

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40100/idl"
	"github.com/arminguenther/xeruspower-go/v40100/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40100/internal/encoding/object"
)

func init() {
	object.Register(NewSystem)
}

type _System struct {
	idl.Object
}

// NewSystem returns a new System interface for the object at given RID.
func NewSystem(rid string, caller idl.Caller) System {
	return &_System{idl.NewObject(rid, caller)}
}

func (s *_System) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "sys.System",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_System) IsDaemonRunning(ctx context.Context, in0 string) (bool, error) {
	var ret bool
	val, err := s.Caller().Call(ctx, s.RID(), "isDaemonRunning", map[string]any{
		"name": in0,
	})
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
	ret, err = encoding.Is[bool](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_System) RestartDaemon(ctx context.Context, in0 string) error {
	_, err := s.Caller().Call(ctx, s.RID(), "restartDaemon", map[string]any{
		"name": in0,
	})
	return err
}
