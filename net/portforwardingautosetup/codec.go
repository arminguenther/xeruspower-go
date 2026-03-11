// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package portforwardingautosetup

import (
	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/idl/event"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/valobj"
)

func (p *ExpansionUnit) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["state"] = p.State
	j0["errMsg"] = p.ErrMsg
	return j0
}

func (p *ExpansionUnit) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("state", j0)
	if err != nil {
		return err
	}
	p.State, err = encoding.AsEnum[ExpansionUnitState](j0["state"])
	if err != nil {
		return err
	}
	err = encoding.In("errMsg", j0)
	if err != nil {
		return err
	}
	p.ErrMsg, err = encoding.Is[string](j0["errMsg"])
	if err != nil {
		return err
	}
	return nil
}

func (p *Status) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["startTime"] = p.StartTime.Unix()
	j0["remainingExpansionUnits"] = p.RemainingExpansionUnits
	j0["runningState"] = p.RunningState
	s1 := make([]any, 0, len(p.ExpansionUnits))
	for k1, v1 := range p.ExpansionUnits {
		s1 = append(s1, map[string]any{"key": k1, "value": v1.Encode()})
	}
	j0["expansionUnits"] = s1
	return j0
}

func (p *Status) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("startTime", j0)
	if err != nil {
		return err
	}
	p.StartTime, err = encoding.AsTime(j0["startTime"])
	if err != nil {
		return err
	}
	err = encoding.In("remainingExpansionUnits", j0)
	if err != nil {
		return err
	}
	p.RemainingExpansionUnits, err = encoding.AsInt32(j0["remainingExpansionUnits"])
	if err != nil {
		return err
	}
	err = encoding.In("runningState", j0)
	if err != nil {
		return err
	}
	p.RunningState, err = encoding.AsEnum[RunningState](j0["runningState"])
	if err != nil {
		return err
	}
	err = encoding.In("expansionUnits", j0)
	if err != nil {
		return err
	}
	i1, e1, l1 := encoding.AsMapItems(j0["expansionUnits"])
	p.ExpansionUnits = make(map[int32]ExpansionUnit, l1)
	for a1, b1 := range i1 {
		var k1 int32
		k1, err = encoding.AsInt32(a1)
		if err != nil {
			return err
		}
		var v1 ExpansionUnit
		err = v1.Decode(b1, caller)
		if err != nil {
			return err
		}
		p.ExpansionUnits[k1] = v1
	}
	err = e1()
	if err != nil {
		return err
	}
	return nil
}

func (p *_StatusChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	p.Event = valobj.For[event.Event]()
	err := p.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("newStatus", value)
	if err != nil {
		return err
	}
	err = p.newStatus.Decode(value["newStatus"], caller)
	if err != nil {
		return err
	}
	return nil
}
