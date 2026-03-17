// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package alarmmanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
)

func init() {
	object.Register(NewAlarmManager)
}

type _AlarmManager struct {
	idl.Object
}

// NewAlarmManager returns a new AlarmManager interface for the object at given RID.
func NewAlarmManager(rid string, caller idl.Caller) AlarmManager {
	return &_AlarmManager{idl.NewObject(rid, caller)}
}

func (a *_AlarmManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "event.AlarmManager",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (a *_AlarmManager) AcknowledgeAlarm(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := a.Caller().Call(ctx, a.RID(), "acknowledgeAlarm", map[string]any{
		"alarmId": in0,
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
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (a *_AlarmManager) ListAlarms(ctx context.Context) ([]Alarm, error) {
	var ret []Alarm
	val, err := a.Caller().Call(ctx, a.RID(), "listAlarms", nil)
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
	ret = make([]Alarm, 0, len(s0))
	for _, a0 := range s0 {
		var e0 Alarm
		err = e0.Decode(a0, a.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}
