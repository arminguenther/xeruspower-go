// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package luaservice

import (
	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding"
)

func (s *ScriptState) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["execState"] = s.ExecState
	j0["exitType"] = s.ExitType
	j0["exitStatus"] = s.ExitStatus
	return j0
}

func (s *ScriptState) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("execState", j0)
	if err != nil {
		return err
	}
	s.ExecState, err = encoding.AsEnum[ScriptStateExecState](j0["execState"])
	if err != nil {
		return err
	}
	err = encoding.In("exitType", j0)
	if err != nil {
		return err
	}
	s.ExitType, err = encoding.AsEnum[ScriptStateExitType](j0["exitType"])
	if err != nil {
		return err
	}
	err = encoding.In("exitStatus", j0)
	if err != nil {
		return err
	}
	s.ExitStatus, err = encoding.AsInt32(j0["exitStatus"])
	if err != nil {
		return err
	}
	return nil
}

func (s *ScriptOptions) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	s1 := make([]any, 0, len(s.DefaultArgs))
	for k1, v1 := range s.DefaultArgs {
		s1 = append(s1, map[string]any{"key": k1, "value": v1})
	}
	j0["defaultArgs"] = s1
	j0["autoStart"] = s.AutoStart
	j0["autoRestart"] = s.AutoRestart
	return j0
}

func (s *ScriptOptions) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("defaultArgs", j0)
	if err != nil {
		return err
	}
	i1, e1, l1 := encoding.AsMapItems(j0["defaultArgs"])
	s.DefaultArgs = make(map[string]string, l1)
	for a1, b1 := range i1 {
		var k1 string
		k1, err = encoding.Is[string](a1)
		if err != nil {
			return err
		}
		var v1 string
		v1, err = encoding.Is[string](b1)
		if err != nil {
			return err
		}
		s.DefaultArgs[k1] = v1
	}
	err = e1()
	if err != nil {
		return err
	}
	err = encoding.In("autoStart", j0)
	if err != nil {
		return err
	}
	s.AutoStart, err = encoding.Is[bool](j0["autoStart"])
	if err != nil {
		return err
	}
	err = encoding.In("autoRestart", j0)
	if err != nil {
		return err
	}
	s.AutoRestart, err = encoding.Is[bool](j0["autoRestart"])
	if err != nil {
		return err
	}
	return nil
}

func (e *Environment) Encode() map[string]any {
	j0 := make(map[string]any, 9)
	j0["maxScriptMemoryGrowth"] = e.MaxScriptMemoryGrowth
	j0["maxAmountOfScripts"] = e.MaxAmountOfScripts
	j0["amountOfScripts"] = e.AmountOfScripts
	j0["maxScriptSize"] = e.MaxScriptSize
	j0["maxAllScriptSize"] = e.MaxAllScriptSize
	j0["allScriptSize"] = e.AllScriptSize
	j0["outputBufferSize"] = e.OutputBufferSize
	j0["restartInterval"] = e.RestartInterval
	j0["autoStartDelay"] = e.AutoStartDelay
	return j0
}

func (e *Environment) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("maxScriptMemoryGrowth", j0)
	if err != nil {
		return err
	}
	e.MaxScriptMemoryGrowth, err = encoding.AsInt32(j0["maxScriptMemoryGrowth"])
	if err != nil {
		return err
	}
	err = encoding.In("maxAmountOfScripts", j0)
	if err != nil {
		return err
	}
	e.MaxAmountOfScripts, err = encoding.AsInt32(j0["maxAmountOfScripts"])
	if err != nil {
		return err
	}
	err = encoding.In("amountOfScripts", j0)
	if err != nil {
		return err
	}
	e.AmountOfScripts, err = encoding.AsInt32(j0["amountOfScripts"])
	if err != nil {
		return err
	}
	err = encoding.In("maxScriptSize", j0)
	if err != nil {
		return err
	}
	e.MaxScriptSize, err = encoding.AsInt32(j0["maxScriptSize"])
	if err != nil {
		return err
	}
	err = encoding.In("maxAllScriptSize", j0)
	if err != nil {
		return err
	}
	e.MaxAllScriptSize, err = encoding.AsInt32(j0["maxAllScriptSize"])
	if err != nil {
		return err
	}
	err = encoding.In("allScriptSize", j0)
	if err != nil {
		return err
	}
	e.AllScriptSize, err = encoding.AsInt32(j0["allScriptSize"])
	if err != nil {
		return err
	}
	err = encoding.In("outputBufferSize", j0)
	if err != nil {
		return err
	}
	e.OutputBufferSize, err = encoding.AsInt32(j0["outputBufferSize"])
	if err != nil {
		return err
	}
	err = encoding.In("restartInterval", j0)
	if err != nil {
		return err
	}
	e.RestartInterval, err = encoding.AsInt32(j0["restartInterval"])
	if err != nil {
		return err
	}
	err = encoding.In("autoStartDelay", j0)
	if err != nil {
		return err
	}
	e.AutoStartDelay, err = encoding.AsInt32(j0["autoStartDelay"])
	if err != nil {
		return err
	}
	return nil
}
