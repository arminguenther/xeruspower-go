// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package gatewaysensormanager

import (
	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/idl/event"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40220/peripheral/modbuscfg"
)

func (s *_SensorClass) Encode() map[string]any {
	ret := s.ValueObject.Encode()
	ret["classId"] = s.classId
	return ret
}

func (s *_SensorClass) Decode(value map[string]any, caller idl.Caller) error {
	s.ValueObject = valobj.For[idl.ValueObject]()
	var err error
	err = encoding.In("classId", value)
	if err != nil {
		return err
	}
	s.classId, err = encoding.Is[string](value["classId"])
	if err != nil {
		return err
	}
	return nil
}

func (n *_NumericSensorClass) Encode() map[string]any {
	ret := n.SensorClass.Encode()
	ret["metadata"] = n.metadata.Encode()
	ret["defaultThresholds"] = n.defaultThresholds.Encode()
	ret["preferCommonThresholds"] = n.preferCommonThresholds
	return ret
}

func (n *_NumericSensorClass) Decode(value map[string]any, caller idl.Caller) error {
	n.SensorClass = valobj.For[SensorClass]()
	err := n.SensorClass.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("metadata", value)
	if err != nil {
		return err
	}
	err = n.metadata.Decode(value["metadata"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("defaultThresholds", value)
	if err != nil {
		return err
	}
	err = n.defaultThresholds.Decode(value["defaultThresholds"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("preferCommonThresholds", value)
	if err != nil {
		return err
	}
	n.preferCommonThresholds, err = encoding.Is[bool](value["preferCommonThresholds"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_StateSensorClass) Encode() map[string]any {
	ret := s.SensorClass.Encode()
	ret["type"] = s.type_.Encode()
	return ret
}

func (s *_StateSensorClass) Decode(value map[string]any, caller idl.Caller) error {
	s.SensorClass = valobj.For[SensorClass]()
	err := s.SensorClass.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("type", value)
	if err != nil {
		return err
	}
	err = s.type_.Decode(value["type"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *_SwitchSensorClass) Encode() map[string]any {
	ret := s.StateSensorClass.Encode()
	return ret
}

func (s *_SwitchSensorClass) Decode(value map[string]any, caller idl.Caller) error {
	s.StateSensorClass = valobj.For[StateSensorClass]()
	err := s.StateSensorClass.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (r *_RemoteDevice) Encode() map[string]any {
	ret := r.ValueObject.Encode()
	ret["deviceId"] = r.deviceId
	ret["disabled"] = r.disabled
	ret["name"] = r.name
	ret["timeoutMs"] = r.timeoutMs
	ret["retry"] = r.retry
	return ret
}

func (r *_RemoteDevice) Decode(value map[string]any, caller idl.Caller) error {
	r.ValueObject = valobj.For[idl.ValueObject]()
	var err error
	err = encoding.In("deviceId", value)
	if err != nil {
		return err
	}
	r.deviceId, err = encoding.Is[string](value["deviceId"])
	if err != nil {
		return err
	}
	err = encoding.In("disabled", value)
	if err != nil {
		return err
	}
	r.disabled, err = encoding.Is[bool](value["disabled"])
	if err != nil {
		return err
	}
	err = encoding.In("name", value)
	if err != nil {
		return err
	}
	r.name, err = encoding.Is[string](value["name"])
	if err != nil {
		return err
	}
	err = encoding.In("timeoutMs", value)
	if err != nil {
		return err
	}
	r.timeoutMs, err = encoding.AsInt32(value["timeoutMs"])
	if err != nil {
		return err
	}
	err = encoding.In("retry", value)
	if err != nil {
		return err
	}
	r.retry, err = encoding.AsInt32(value["retry"])
	if err != nil {
		return err
	}
	return nil
}

func (r *_RemoteModbusDevice) Encode() map[string]any {
	ret := r.RemoteDevice.Encode()
	s0 := make([]any, 0, len(r.detectionIdentifiers))
	for k0, v0 := range r.detectionIdentifiers {
		s0 = append(s0, map[string]any{"key": k0, "value": v0})
	}
	ret["detectionIdentifiers"] = s0
	ret["unitId"] = r.unitId
	return ret
}

func (r *_RemoteModbusDevice) Decode(value map[string]any, caller idl.Caller) error {
	r.RemoteDevice = valobj.For[RemoteDevice]()
	err := r.RemoteDevice.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("detectionIdentifiers", value)
	if err != nil {
		return err
	}
	i0, e0, l0 := encoding.AsMapItems(value["detectionIdentifiers"])
	r.detectionIdentifiers = make(map[int32]string, l0)
	for a0, b0 := range i0 {
		var k0 int32
		k0, err = encoding.AsInt32(a0)
		if err != nil {
			return err
		}
		var v0 string
		v0, err = encoding.Is[string](b0)
		if err != nil {
			return err
		}
		r.detectionIdentifiers[k0] = v0
	}
	err = e0()
	if err != nil {
		return err
	}
	err = encoding.In("unitId", value)
	if err != nil {
		return err
	}
	r.unitId, err = encoding.AsInt32(value["unitId"])
	if err != nil {
		return err
	}
	return nil
}

func (r *_RemoteModbusRTUDevice) Encode() map[string]any {
	ret := r.RemoteModbusDevice.Encode()
	ret["busInterface"] = r.busInterface
	ret["busSettings"] = r.busSettings.Encode()
	ret["interframeDelayDeciChars"] = r.interframeDelayDeciChars
	return ret
}

func (r *_RemoteModbusRTUDevice) Decode(value map[string]any, caller idl.Caller) error {
	r.RemoteModbusDevice = valobj.For[RemoteModbusDevice]()
	err := r.RemoteModbusDevice.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("busInterface", value)
	if err != nil {
		return err
	}
	r.busInterface, err = encoding.Is[string](value["busInterface"])
	if err != nil {
		return err
	}
	err = encoding.In("busSettings", value)
	if err != nil {
		return err
	}
	err = r.busSettings.Decode(value["busSettings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("interframeDelayDeciChars", value)
	if err != nil {
		return err
	}
	r.interframeDelayDeciChars, err = encoding.AsInt32(value["interframeDelayDeciChars"])
	if err != nil {
		return err
	}
	return nil
}

func (r *_RemoteModbusTCPDevice) Encode() map[string]any {
	ret := r.RemoteModbusDevice.Encode()
	ret["ipAddress"] = r.ipAddress
	ret["tcpPort"] = r.tcpPort
	return ret
}

func (r *_RemoteModbusTCPDevice) Decode(value map[string]any, caller idl.Caller) error {
	r.RemoteModbusDevice = valobj.For[RemoteModbusDevice]()
	err := r.RemoteModbusDevice.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("ipAddress", value)
	if err != nil {
		return err
	}
	r.ipAddress, err = encoding.Is[string](value["ipAddress"])
	if err != nil {
		return err
	}
	err = encoding.In("tcpPort", value)
	if err != nil {
		return err
	}
	r.tcpPort, err = encoding.AsInt32(value["tcpPort"])
	if err != nil {
		return err
	}
	return nil
}

func (r *_RemoteSnmpDevice) Encode() map[string]any {
	ret := r.RemoteDevice.Encode()
	ret["host"] = r.host
	return ret
}

func (r *_RemoteSnmpDevice) Decode(value map[string]any, caller idl.Caller) error {
	r.RemoteDevice = valobj.For[RemoteDevice]()
	err := r.RemoteDevice.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("host", value)
	if err != nil {
		return err
	}
	r.host, err = encoding.Is[string](value["host"])
	if err != nil {
		return err
	}
	return nil
}

func (r *_RemoteSnmpV1V2Device) Encode() map[string]any {
	ret := r.RemoteSnmpDevice.Encode()
	ret["community"] = r.community
	return ret
}

func (r *_RemoteSnmpV1V2Device) Decode(value map[string]any, caller idl.Caller) error {
	r.RemoteSnmpDevice = valobj.For[RemoteSnmpDevice]()
	err := r.RemoteSnmpDevice.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("community", value)
	if err != nil {
		return err
	}
	r.community, err = encoding.Is[string](value["community"])
	if err != nil {
		return err
	}
	return nil
}

func (r *_RemoteSnmpV3Device) Encode() map[string]any {
	ret := r.RemoteSnmpDevice.Encode()
	ret["user"] = r.user
	ret["level"] = r.level
	ret["authProtocol"] = r.authProtocol
	ret["authPassphrase"] = r.authPassphrase
	ret["privacyProtocol"] = r.privacyProtocol
	ret["privacyPassphrase"] = r.privacyPassphrase
	return ret
}

func (r *_RemoteSnmpV3Device) Decode(value map[string]any, caller idl.Caller) error {
	r.RemoteSnmpDevice = valobj.For[RemoteSnmpDevice]()
	err := r.RemoteSnmpDevice.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("user", value)
	if err != nil {
		return err
	}
	r.user, err = encoding.Is[string](value["user"])
	if err != nil {
		return err
	}
	err = encoding.In("level", value)
	if err != nil {
		return err
	}
	r.level, err = encoding.AsEnum[SnmpSecurityLevel](value["level"])
	if err != nil {
		return err
	}
	err = encoding.In("authProtocol", value)
	if err != nil {
		return err
	}
	r.authProtocol, err = encoding.AsEnum[SnmpAuthProtocol](value["authProtocol"])
	if err != nil {
		return err
	}
	err = encoding.In("authPassphrase", value)
	if err != nil {
		return err
	}
	r.authPassphrase, err = encoding.Is[string](value["authPassphrase"])
	if err != nil {
		return err
	}
	err = encoding.In("privacyProtocol", value)
	if err != nil {
		return err
	}
	r.privacyProtocol, err = encoding.AsEnum[SnmpPrivProtocol](value["privacyProtocol"])
	if err != nil {
		return err
	}
	err = encoding.In("privacyPassphrase", value)
	if err != nil {
		return err
	}
	r.privacyPassphrase, err = encoding.Is[string](value["privacyPassphrase"])
	if err != nil {
		return err
	}
	return nil
}

func (i *_InterpretationRule) Encode() map[string]any {
	ret := i.ValueObject.Encode()
	ret["interpretation"] = i.interpretation
	ret["ignoreTimeout"] = i.ignoreTimeout
	return ret
}

func (i *_InterpretationRule) Decode(value map[string]any, caller idl.Caller) error {
	i.ValueObject = valobj.For[idl.ValueObject]()
	var err error
	err = encoding.In("interpretation", value)
	if err != nil {
		return err
	}
	i.interpretation, err = encoding.AsEnum[Interpretation](value["interpretation"])
	if err != nil {
		return err
	}
	err = encoding.In("ignoreTimeout", value)
	if err != nil {
		return err
	}
	i.ignoreTimeout, err = encoding.AsInt32(value["ignoreTimeout"])
	if err != nil {
		return err
	}
	return nil
}

func (i *_InterpretationRuleInvertable) Encode() map[string]any {
	ret := i.InterpretationRule.Encode()
	ret["invertCondition"] = i.invertCondition
	return ret
}

func (i *_InterpretationRuleInvertable) Decode(value map[string]any, caller idl.Caller) error {
	i.InterpretationRule = valobj.For[InterpretationRule]()
	err := i.InterpretationRule.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("invertCondition", value)
	if err != nil {
		return err
	}
	i.invertCondition, err = encoding.Is[bool](value["invertCondition"])
	if err != nil {
		return err
	}
	return nil
}

func (i *_InterpretationRuleModbusException) Encode() map[string]any {
	ret := i.InterpretationRuleInvertable.Encode()
	ret["exceptions"] = encoding.NonNilSlice(i.exceptions)
	return ret
}

func (i *_InterpretationRuleModbusException) Decode(value map[string]any, caller idl.Caller) error {
	i.InterpretationRuleInvertable = valobj.For[InterpretationRuleInvertable]()
	err := i.InterpretationRuleInvertable.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("exceptions", value)
	if err != nil {
		return err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](value["exceptions"])
	if err != nil {
		return err
	}
	i.exceptions = make([]int32, 0, len(s0))
	for _, a0 := range s0 {
		var e0 int32
		e0, err = encoding.AsInt32(a0)
		if err != nil {
			return err
		}
		i.exceptions = append(i.exceptions, e0)
	}
	return nil
}

func (i *_InterpretationRuleModbusSystemError) Encode() map[string]any {
	ret := i.InterpretationRuleInvertable.Encode()
	ret["errnos"] = encoding.NonNilSlice(i.errnos)
	return ret
}

func (i *_InterpretationRuleModbusSystemError) Decode(value map[string]any, caller idl.Caller) error {
	i.InterpretationRuleInvertable = valobj.For[InterpretationRuleInvertable]()
	err := i.InterpretationRuleInvertable.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("errnos", value)
	if err != nil {
		return err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](value["errnos"])
	if err != nil {
		return err
	}
	i.errnos = make([]int32, 0, len(s0))
	for _, a0 := range s0 {
		var e0 int32
		e0, err = encoding.AsInt32(a0)
		if err != nil {
			return err
		}
		i.errnos = append(i.errnos, e0)
	}
	return nil
}

func (i *_InterpretationRuleModbusSpecificError) Encode() map[string]any {
	ret := i.InterpretationRuleInvertable.Encode()
	ret["errors"] = encoding.NonNilSlice(i.errors)
	return ret
}

func (i *_InterpretationRuleModbusSpecificError) Decode(value map[string]any, caller idl.Caller) error {
	i.InterpretationRuleInvertable = valobj.For[InterpretationRuleInvertable]()
	err := i.InterpretationRuleInvertable.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("errors", value)
	if err != nil {
		return err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](value["errors"])
	if err != nil {
		return err
	}
	i.errors = make([]modbuscfg.SpecificModbusErrors, 0, len(s0))
	for _, a0 := range s0 {
		var e0 modbuscfg.SpecificModbusErrors
		e0, err = encoding.AsEnum[modbuscfg.SpecificModbusErrors](a0)
		if err != nil {
			return err
		}
		i.errors = append(i.errors, e0)
	}
	return nil
}

func (i *_InterpretationRuleRAW) Encode() map[string]any {
	ret := i.InterpretationRuleInvertable.Encode()
	ret["value"] = i.value
	ret["mask"] = i.mask
	return ret
}

func (i *_InterpretationRuleRAW) Decode(value map[string]any, caller idl.Caller) error {
	i.InterpretationRuleInvertable = valobj.For[InterpretationRuleInvertable]()
	err := i.InterpretationRuleInvertable.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("value", value)
	if err != nil {
		return err
	}
	i.value, err = encoding.AsInt64(value["value"])
	if err != nil {
		return err
	}
	err = encoding.In("mask", value)
	if err != nil {
		return err
	}
	i.mask, err = encoding.AsInt64(value["mask"])
	if err != nil {
		return err
	}
	return nil
}

func (i *_InterpretationRuleRangeRAW) Encode() map[string]any {
	ret := i.InterpretationRuleInvertable.Encode()
	ret["min"] = i.min
	ret["max"] = i.max
	ret["mask"] = i.mask
	return ret
}

func (i *_InterpretationRuleRangeRAW) Decode(value map[string]any, caller idl.Caller) error {
	i.InterpretationRuleInvertable = valobj.For[InterpretationRuleInvertable]()
	err := i.InterpretationRuleInvertable.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("min", value)
	if err != nil {
		return err
	}
	i.min, err = encoding.AsInt64(value["min"])
	if err != nil {
		return err
	}
	err = encoding.In("max", value)
	if err != nil {
		return err
	}
	i.max, err = encoding.AsInt64(value["max"])
	if err != nil {
		return err
	}
	err = encoding.In("mask", value)
	if err != nil {
		return err
	}
	i.mask, err = encoding.AsInt64(value["mask"])
	if err != nil {
		return err
	}
	return nil
}

func (i *_InterpretationRuleEnum) Encode() map[string]any {
	ret := i.InterpretationRuleInvertable.Encode()
	ret["enumValues"] = encoding.NonNilSlice(i.enumValues)
	return ret
}

func (i *_InterpretationRuleEnum) Decode(value map[string]any, caller idl.Caller) error {
	i.InterpretationRuleInvertable = valobj.For[InterpretationRuleInvertable]()
	err := i.InterpretationRuleInvertable.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("enumValues", value)
	if err != nil {
		return err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](value["enumValues"])
	if err != nil {
		return err
	}
	i.enumValues = make([]int64, 0, len(s0))
	for _, a0 := range s0 {
		var e0 int64
		e0, err = encoding.AsInt64(a0)
		if err != nil {
			return err
		}
		i.enumValues = append(i.enumValues, e0)
	}
	return nil
}

func (i *_InterpretationRuleIEEE754INF) Encode() map[string]any {
	ret := i.InterpretationRuleInvertable.Encode()
	return ret
}

func (i *_InterpretationRuleIEEE754INF) Decode(value map[string]any, caller idl.Caller) error {
	i.InterpretationRuleInvertable = valobj.For[InterpretationRuleInvertable]()
	err := i.InterpretationRuleInvertable.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (i *_InterpretationRuleIEEE754NAN) Encode() map[string]any {
	ret := i.InterpretationRuleInvertable.Encode()
	return ret
}

func (i *_InterpretationRuleIEEE754NAN) Decode(value map[string]any, caller idl.Caller) error {
	i.InterpretationRuleInvertable = valobj.For[InterpretationRuleInvertable]()
	err := i.InterpretationRuleInvertable.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (i *_InterpretationRuleCatchAll) Encode() map[string]any {
	ret := i.InterpretationRule.Encode()
	return ret
}

func (i *_InterpretationRuleCatchAll) Decode(value map[string]any, caller idl.Caller) error {
	i.InterpretationRule = valobj.For[InterpretationRule]()
	err := i.InterpretationRule.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (v *_ValueEncoding) Encode() map[string]any {
	ret := v.ValueObject.Encode()
	ret["encodingId"] = v.encodingId
	ret["type"] = v.type_
	ret["invertState"] = v.invertState
	s0 := make([]any, 0, len(v.interpretationRules))
	for _, e0 := range v.interpretationRules {
		s0 = append(s0, valobj.ToMap(e0))
	}
	ret["interpretationRules"] = s0
	ret["minAccessInterval"] = v.minAccessInterval
	return ret
}

func (v *_ValueEncoding) Decode(value map[string]any, caller idl.Caller) error {
	v.ValueObject = valobj.For[idl.ValueObject]()
	var err error
	err = encoding.In("encodingId", value)
	if err != nil {
		return err
	}
	v.encodingId, err = encoding.Is[string](value["encodingId"])
	if err != nil {
		return err
	}
	err = encoding.In("type", value)
	if err != nil {
		return err
	}
	v.type_, err = encoding.AsEnum[EncodingType](value["type"])
	if err != nil {
		return err
	}
	err = encoding.In("invertState", value)
	if err != nil {
		return err
	}
	v.invertState, err = encoding.Is[bool](value["invertState"])
	if err != nil {
		return err
	}
	err = encoding.In("interpretationRules", value)
	if err != nil {
		return err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](value["interpretationRules"])
	if err != nil {
		return err
	}
	v.interpretationRules = make([]InterpretationRule, 0, len(s0))
	for _, a0 := range s0 {
		var e0 InterpretationRule
		e0, err = valobj.As[InterpretationRule](a0, caller)
		if err != nil {
			return err
		}
		v.interpretationRules = append(v.interpretationRules, e0)
	}
	err = encoding.In("minAccessInterval", value)
	if err != nil {
		return err
	}
	v.minAccessInterval, err = encoding.AsInt32(value["minAccessInterval"])
	if err != nil {
		return err
	}
	return nil
}

func (n *_NumericValueEncoding) Encode() map[string]any {
	ret := n.ValueEncoding.Encode()
	ret["scalingFactor"] = n.scalingFactor
	ret["offset"] = n.offset
	return ret
}

func (n *_NumericValueEncoding) Decode(value map[string]any, caller idl.Caller) error {
	n.ValueEncoding = valobj.For[ValueEncoding]()
	err := n.ValueEncoding.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("scalingFactor", value)
	if err != nil {
		return err
	}
	n.scalingFactor, err = encoding.AsFloat32(value["scalingFactor"])
	if err != nil {
		return err
	}
	err = encoding.In("offset", value)
	if err != nil {
		return err
	}
	n.offset, err = encoding.AsFloat32(value["offset"])
	if err != nil {
		return err
	}
	return nil
}

func (m *_ModbusValueEncodingBit) Encode() map[string]any {
	ret := m.ValueEncoding.Encode()
	return ret
}

func (m *_ModbusValueEncodingBit) Decode(value map[string]any, caller idl.Caller) error {
	m.ValueEncoding = valobj.For[ValueEncoding]()
	err := m.ValueEncoding.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (m *_ModbusValueEncoding8) Encode() map[string]any {
	ret := m.NumericValueEncoding.Encode()
	ret["byteSwap"] = m.byteSwap
	ret["mask"] = m.mask
	ret["start"] = m.start
	ret["width"] = m.width
	return ret
}

func (m *_ModbusValueEncoding8) Decode(value map[string]any, caller idl.Caller) error {
	m.NumericValueEncoding = valobj.For[NumericValueEncoding]()
	err := m.NumericValueEncoding.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("byteSwap", value)
	if err != nil {
		return err
	}
	m.byteSwap, err = encoding.Is[bool](value["byteSwap"])
	if err != nil {
		return err
	}
	err = encoding.In("mask", value)
	if err != nil {
		return err
	}
	m.mask, err = encoding.AsInt64(value["mask"])
	if err != nil {
		return err
	}
	err = encoding.In("start", value)
	if err != nil {
		return err
	}
	m.start, err = encoding.AsInt32(value["start"])
	if err != nil {
		return err
	}
	err = encoding.In("width", value)
	if err != nil {
		return err
	}
	m.width, err = encoding.AsInt32(value["width"])
	if err != nil {
		return err
	}
	return nil
}

func (m *_ModbusValueEncoding16) Encode() map[string]any {
	ret := m.ModbusValueEncoding8.Encode()
	return ret
}

func (m *_ModbusValueEncoding16) Decode(value map[string]any, caller idl.Caller) error {
	m.ModbusValueEncoding8 = valobj.For[ModbusValueEncoding8]()
	err := m.ModbusValueEncoding8.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (m *_ModbusValueEncoding32) Encode() map[string]any {
	ret := m.ModbusValueEncoding16.Encode()
	ret["endianness"] = m.endianness
	return ret
}

func (m *_ModbusValueEncoding32) Decode(value map[string]any, caller idl.Caller) error {
	m.ModbusValueEncoding16 = valobj.For[ModbusValueEncoding16]()
	err := m.ModbusValueEncoding16.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("endianness", value)
	if err != nil {
		return err
	}
	m.endianness, err = encoding.AsEnum[ModbusEndianness](value["endianness"])
	if err != nil {
		return err
	}
	return nil
}

func (m *_ModbusValueEncoding48) Encode() map[string]any {
	ret := m.ModbusValueEncoding32.Encode()
	return ret
}

func (m *_ModbusValueEncoding48) Decode(value map[string]any, caller idl.Caller) error {
	m.ModbusValueEncoding32 = valobj.For[ModbusValueEncoding32]()
	err := m.ModbusValueEncoding32.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (m *_ModbusValueEncoding64) Encode() map[string]any {
	ret := m.ModbusValueEncoding32.Encode()
	return ret
}

func (m *_ModbusValueEncoding64) Decode(value map[string]any, caller idl.Caller) error {
	m.ModbusValueEncoding32 = valobj.For[ModbusValueEncoding32]()
	err := m.ModbusValueEncoding32.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *_Sensor) Encode() map[string]any {
	ret := s.ValueObject.Encode()
	ret["sensorId"] = s.sensorId
	ret["disabled"] = s.disabled
	ret["deviceId"] = s.deviceId
	ret["classId"] = s.classId
	ret["encodingId"] = s.encodingId
	ret["defaultName"] = s.defaultName
	return ret
}

func (s *_Sensor) Decode(value map[string]any, caller idl.Caller) error {
	s.ValueObject = valobj.For[idl.ValueObject]()
	var err error
	err = encoding.In("sensorId", value)
	if err != nil {
		return err
	}
	s.sensorId, err = encoding.Is[string](value["sensorId"])
	if err != nil {
		return err
	}
	err = encoding.In("disabled", value)
	if err != nil {
		return err
	}
	s.disabled, err = encoding.Is[bool](value["disabled"])
	if err != nil {
		return err
	}
	err = encoding.In("deviceId", value)
	if err != nil {
		return err
	}
	s.deviceId, err = encoding.Is[string](value["deviceId"])
	if err != nil {
		return err
	}
	err = encoding.In("classId", value)
	if err != nil {
		return err
	}
	s.classId, err = encoding.Is[string](value["classId"])
	if err != nil {
		return err
	}
	err = encoding.In("encodingId", value)
	if err != nil {
		return err
	}
	s.encodingId, err = encoding.Is[string](value["encodingId"])
	if err != nil {
		return err
	}
	err = encoding.In("defaultName", value)
	if err != nil {
		return err
	}
	s.defaultName, err = encoding.Is[string](value["defaultName"])
	if err != nil {
		return err
	}
	return nil
}

func (m *_ModbusSensor) Encode() map[string]any {
	ret := m.Sensor.Encode()
	ret["function"] = m.function
	ret["regAddr"] = m.regAddr
	return ret
}

func (m *_ModbusSensor) Decode(value map[string]any, caller idl.Caller) error {
	m.Sensor = valobj.For[Sensor]()
	err := m.Sensor.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("function", value)
	if err != nil {
		return err
	}
	m.function, err = encoding.AsEnum[modbuscfg.ModbusFunction](value["function"])
	if err != nil {
		return err
	}
	err = encoding.In("regAddr", value)
	if err != nil {
		return err
	}
	m.regAddr, err = encoding.AsInt32(value["regAddr"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_SnmpSensor) Encode() map[string]any {
	ret := s.Sensor.Encode()
	ret["oid"] = s.oid
	return ret
}

func (s *_SnmpSensor) Decode(value map[string]any, caller idl.Caller) error {
	s.Sensor = valobj.For[Sensor]()
	err := s.Sensor.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oid", value)
	if err != nil {
		return err
	}
	s.oid, err = encoding.Is[string](value["oid"])
	if err != nil {
		return err
	}
	return nil
}

func (c *ConfigurationPackage) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["disabled"] = c.Disabled
	j0["name"] = c.Name
	s1 := make([]any, 0, len(c.Classes))
	for _, e1 := range c.Classes {
		s1 = append(s1, valobj.ToMap(e1))
	}
	j0["classes"] = s1
	s2 := make([]any, 0, len(c.Devices))
	for _, e2 := range c.Devices {
		s2 = append(s2, valobj.ToMap(e2))
	}
	j0["devices"] = s2
	s3 := make([]any, 0, len(c.Encodings))
	for _, e3 := range c.Encodings {
		s3 = append(s3, valobj.ToMap(e3))
	}
	j0["encodings"] = s3
	s4 := make([]any, 0, len(c.Sensors))
	for _, e4 := range c.Sensors {
		s4 = append(s4, valobj.ToMap(e4))
	}
	j0["sensors"] = s4
	return j0
}

func (c *ConfigurationPackage) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("disabled", j0)
	if err != nil {
		return err
	}
	c.Disabled, err = encoding.Is[bool](j0["disabled"])
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	c.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("classes", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["classes"])
	if err != nil {
		return err
	}
	c.Classes = make([]SensorClass, 0, len(s1))
	for _, a1 := range s1 {
		var e1 SensorClass
		e1, err = valobj.As[SensorClass](a1, caller)
		if err != nil {
			return err
		}
		c.Classes = append(c.Classes, e1)
	}
	err = encoding.In("devices", j0)
	if err != nil {
		return err
	}
	var s2 []any
	s2, err = encoding.Is[[]any](j0["devices"])
	if err != nil {
		return err
	}
	c.Devices = make([]RemoteDevice, 0, len(s2))
	for _, a2 := range s2 {
		var e2 RemoteDevice
		e2, err = valobj.As[RemoteDevice](a2, caller)
		if err != nil {
			return err
		}
		c.Devices = append(c.Devices, e2)
	}
	err = encoding.In("encodings", j0)
	if err != nil {
		return err
	}
	var s3 []any
	s3, err = encoding.Is[[]any](j0["encodings"])
	if err != nil {
		return err
	}
	c.Encodings = make([]ValueEncoding, 0, len(s3))
	for _, a3 := range s3 {
		var e3 ValueEncoding
		e3, err = valobj.As[ValueEncoding](a3, caller)
		if err != nil {
			return err
		}
		c.Encodings = append(c.Encodings, e3)
	}
	err = encoding.In("sensors", j0)
	if err != nil {
		return err
	}
	var s4 []any
	s4, err = encoding.Is[[]any](j0["sensors"])
	if err != nil {
		return err
	}
	c.Sensors = make([]Sensor, 0, len(s4))
	for _, a4 := range s4 {
		var e4 Sensor
		e4, err = valobj.As[Sensor](a4, caller)
		if err != nil {
			return err
		}
		c.Sensors = append(c.Sensors, e4)
	}
	return nil
}

func (c *_ConfigurationChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.Event = valobj.For[event.Event]()
	err := c.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("configuration", value)
	if err != nil {
		return err
	}
	i0, e0, l0 := encoding.AsMapItems(value["configuration"])
	c.configuration = make(map[string]ConfigurationPackage, l0)
	for a0, b0 := range i0 {
		var k0 string
		k0, err = encoding.Is[string](a0)
		if err != nil {
			return err
		}
		var v0 ConfigurationPackage
		err = v0.Decode(b0, caller)
		if err != nil {
			return err
		}
		c.configuration[k0] = v0
	}
	err = e0()
	if err != nil {
		return err
	}
	return nil
}

func (f *FeedbackObject) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["key"] = f.Key
	j0["value"] = f.Value
	j0["stateTansitionTo"] = f.StateTansitionTo
	return j0
}

func (f *FeedbackObject) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("key", j0)
	if err != nil {
		return err
	}
	f.Key, err = encoding.Is[string](j0["key"])
	if err != nil {
		return err
	}
	err = encoding.In("value", j0)
	if err != nil {
		return err
	}
	f.Value, err = encoding.Is[string](j0["value"])
	if err != nil {
		return err
	}
	err = encoding.In("stateTansitionTo", j0)
	if err != nil {
		return err
	}
	f.StateTansitionTo, err = encoding.AsEnum[FeedbackObjectFeedbackState](j0["stateTansitionTo"])
	if err != nil {
		return err
	}
	return nil
}

func (f *_Feedback) Decode(value map[string]any, caller idl.Caller) error {
	f.ValueObject = valobj.For[idl.ValueObject]()
	var err error
	err = encoding.In("currentState", value)
	if err != nil {
		return err
	}
	f.currentState, err = encoding.AsEnum[FeedbackObjectFeedbackState](value["currentState"])
	if err != nil {
		return err
	}
	err = encoding.In("infos", value)
	if err != nil {
		return err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](value["infos"])
	if err != nil {
		return err
	}
	f.infos = make([]FeedbackObject, 0, len(s0))
	for _, a0 := range s0 {
		var e0 FeedbackObject
		err = e0.Decode(a0, caller)
		if err != nil {
			return err
		}
		f.infos = append(f.infos, e0)
	}
	return nil
}

func (d *_DeviceFeedback) Decode(value map[string]any, caller idl.Caller) error {
	d.Feedback = valobj.For[Feedback]()
	err := d.Feedback.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("packageId", value)
	if err != nil {
		return err
	}
	d.packageId, err = encoding.Is[string](value["packageId"])
	if err != nil {
		return err
	}
	err = encoding.In("deviceId", value)
	if err != nil {
		return err
	}
	d.deviceId, err = encoding.Is[string](value["deviceId"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_SensorFeedback) Decode(value map[string]any, caller idl.Caller) error {
	s.Feedback = valobj.For[Feedback]()
	err := s.Feedback.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("packageId", value)
	if err != nil {
		return err
	}
	s.packageId, err = encoding.Is[string](value["packageId"])
	if err != nil {
		return err
	}
	err = encoding.In("deviceId", value)
	if err != nil {
		return err
	}
	s.deviceId, err = encoding.Is[string](value["deviceId"])
	if err != nil {
		return err
	}
	err = encoding.In("sensorId", value)
	if err != nil {
		return err
	}
	s.sensorId, err = encoding.Is[string](value["sensorId"])
	if err != nil {
		return err
	}
	return nil
}

func (f *_FeedbackChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	f.Event = valobj.For[event.Event]()
	err := f.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("feedback", value)
	if err != nil {
		return err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](value["feedback"])
	if err != nil {
		return err
	}
	f.feedback = make([]Feedback, 0, len(s0))
	for _, a0 := range s0 {
		var e0 Feedback
		e0, err = valobj.As[Feedback](a0, caller)
		if err != nil {
			return err
		}
		f.feedback = append(f.feedback, e0)
	}
	return nil
}
