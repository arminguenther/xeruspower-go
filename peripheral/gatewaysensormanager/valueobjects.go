// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package gatewaysensormanager

import (
	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/idl/event"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40211/peripheral/modbuscfg"
	"github.com/arminguenther/xeruspower-go/v40211/sensors/numericsensor"
	"github.com/arminguenther/xeruspower-go/v40211/sensors/sensor"
)

func init() {
	valobj.Register(func() ConfigurationChangedEvent { return &_ConfigurationChangedEvent{} })
	valobj.Register(func() DeviceFeedback { return &_DeviceFeedback{} })
	valobj.Register(func() Feedback { return &_Feedback{} })
	valobj.Register(func() FeedbackChangedEvent { return &_FeedbackChangedEvent{} })
	valobj.Register(func() InterpretationRule { return &_InterpretationRule{} })
	valobj.Register(func() InterpretationRuleCatchAll { return &_InterpretationRuleCatchAll{} })
	valobj.Register(func() InterpretationRuleEnum { return &_InterpretationRuleEnum{} })
	valobj.Register(func() InterpretationRuleIEEE754INF { return &_InterpretationRuleIEEE754INF{} })
	valobj.Register(func() InterpretationRuleIEEE754NAN { return &_InterpretationRuleIEEE754NAN{} })
	valobj.Register(func() InterpretationRuleInvertable { return &_InterpretationRuleInvertable{} })
	valobj.Register(func() InterpretationRuleModbusException { return &_InterpretationRuleModbusException{} })
	valobj.Register(func() InterpretationRuleModbusSpecificError { return &_InterpretationRuleModbusSpecificError{} })
	valobj.Register(func() InterpretationRuleModbusSystemError { return &_InterpretationRuleModbusSystemError{} })
	valobj.Register(func() InterpretationRuleRAW { return &_InterpretationRuleRAW{} })
	valobj.Register(func() InterpretationRuleRangeRAW { return &_InterpretationRuleRangeRAW{} })
	valobj.Register(func() ModbusSensor { return &_ModbusSensor{} })
	valobj.Register(func() ModbusValueEncoding16 { return &_ModbusValueEncoding16{} })
	valobj.Register(func() ModbusValueEncoding32 { return &_ModbusValueEncoding32{} })
	valobj.Register(func() ModbusValueEncoding48 { return &_ModbusValueEncoding48{} })
	valobj.Register(func() ModbusValueEncoding64 { return &_ModbusValueEncoding64{} })
	valobj.Register(func() ModbusValueEncoding8 { return &_ModbusValueEncoding8{} })
	valobj.Register(func() ModbusValueEncodingBit { return &_ModbusValueEncodingBit{} })
	valobj.Register(func() NumericSensorClass { return &_NumericSensorClass{} })
	valobj.Register(func() NumericValueEncoding { return &_NumericValueEncoding{} })
	valobj.Register(func() RemoteDevice { return &_RemoteDevice{} })
	valobj.Register(func() RemoteModbusDevice { return &_RemoteModbusDevice{} })
	valobj.Register(func() RemoteModbusRTUDevice { return &_RemoteModbusRTUDevice{} })
	valobj.Register(func() RemoteModbusTCPDevice { return &_RemoteModbusTCPDevice{} })
	valobj.Register(func() RemoteSnmpDevice { return &_RemoteSnmpDevice{} })
	valobj.Register(func() RemoteSnmpV1V2Device { return &_RemoteSnmpV1V2Device{} })
	valobj.Register(func() RemoteSnmpV3Device { return &_RemoteSnmpV3Device{} })
	valobj.Register(func() Sensor { return &_Sensor{} })
	valobj.Register(func() SensorClass { return &_SensorClass{} })
	valobj.Register(func() SensorFeedback { return &_SensorFeedback{} })
	valobj.Register(func() SnmpSensor { return &_SnmpSensor{} })
	valobj.Register(func() StateSensorClass { return &_StateSensorClass{} })
	valobj.Register(func() SwitchSensorClass { return &_SwitchSensorClass{} })
	valobj.Register(func() ValueEncoding { return &_ValueEncoding{} })
}

type _SensorClass struct {
	idl.ValueObject
	classId string
}

// NewSensorClass returns a new SensorClass value object.
func NewSensorClass(classId string) SensorClass {
	return &_SensorClass{idl.NewValueObject(), classId}
}

func (s *_SensorClass) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.SensorClass",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_SensorClass) ClassId() string {
	return s.classId
}

func (s *_SensorClass) isSensorClass() {}

type _NumericSensorClass struct {
	SensorClass
	metadata               numericsensor.MetaData
	defaultThresholds      numericsensor.Thresholds
	preferCommonThresholds bool
}

// NewNumericSensorClass returns a new NumericSensorClass value object.
func NewNumericSensorClass(classId string, metadata numericsensor.MetaData, defaultThresholds numericsensor.Thresholds, preferCommonThresholds bool) NumericSensorClass {
	return &_NumericSensorClass{NewSensorClass(classId), metadata, defaultThresholds, preferCommonThresholds}
}

func (n *_NumericSensorClass) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.NumericSensorClass",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (n *_NumericSensorClass) Metadata() numericsensor.MetaData {
	return n.metadata
}

func (n *_NumericSensorClass) DefaultThresholds() numericsensor.Thresholds {
	return n.defaultThresholds
}

func (n *_NumericSensorClass) PreferCommonThresholds() bool {
	return n.preferCommonThresholds
}

func (n *_NumericSensorClass) isNumericSensorClass() {}

type _StateSensorClass struct {
	SensorClass
	type_ sensor.TypeSpec
}

// NewStateSensorClass returns a new StateSensorClass value object.
func NewStateSensorClass(classId string, type_ sensor.TypeSpec) StateSensorClass {
	return &_StateSensorClass{NewSensorClass(classId), type_}
}

func (s *_StateSensorClass) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.StateSensorClass",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_StateSensorClass) Type() sensor.TypeSpec {
	return s.type_
}

func (s *_StateSensorClass) isStateSensorClass() {}

type _SwitchSensorClass struct {
	StateSensorClass
}

// NewSwitchSensorClass returns a new SwitchSensorClass value object.
func NewSwitchSensorClass(classId string, type_ sensor.TypeSpec) SwitchSensorClass {
	return &_SwitchSensorClass{NewStateSensorClass(classId, type_)}
}

func (s *_SwitchSensorClass) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.SwitchSensorClass",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_SwitchSensorClass) isSwitchSensorClass() {}

type _RemoteDevice struct {
	idl.ValueObject
	deviceId  string
	disabled  bool
	name      string
	timeoutMs int32
	retry     int32
}

// NewRemoteDevice returns a new RemoteDevice value object.
func NewRemoteDevice(deviceId string, disabled bool, name string, timeoutMs int32, retry int32) RemoteDevice {
	return &_RemoteDevice{idl.NewValueObject(), deviceId, disabled, name, timeoutMs, retry}
}

func (r *_RemoteDevice) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.RemoteDevice",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RemoteDevice) DeviceId() string {
	return r.deviceId
}

func (r *_RemoteDevice) Disabled() bool {
	return r.disabled
}

func (r *_RemoteDevice) Name() string {
	return r.name
}

func (r *_RemoteDevice) TimeoutMs() int32 {
	return r.timeoutMs
}

func (r *_RemoteDevice) Retry() int32 {
	return r.retry
}

func (r *_RemoteDevice) isRemoteDevice() {}

type _RemoteModbusDevice struct {
	RemoteDevice
	detectionIdentifiers map[int32]string
	unitId               int32
}

// NewRemoteModbusDevice returns a new RemoteModbusDevice value object.
func NewRemoteModbusDevice(deviceId string, disabled bool, name string, timeoutMs int32, retry int32, detectionIdentifiers map[int32]string, unitId int32) RemoteModbusDevice {
	return &_RemoteModbusDevice{NewRemoteDevice(deviceId, disabled, name, timeoutMs, retry), detectionIdentifiers, unitId}
}

func (r *_RemoteModbusDevice) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.RemoteModbusDevice",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RemoteModbusDevice) DetectionIdentifiers() map[int32]string {
	return r.detectionIdentifiers
}

func (r *_RemoteModbusDevice) UnitId() int32 {
	return r.unitId
}

func (r *_RemoteModbusDevice) isRemoteModbusDevice() {}

type _RemoteModbusRTUDevice struct {
	RemoteModbusDevice
	busInterface             string
	busSettings              modbuscfg.SerialSettings
	interframeDelayDeciChars int32
}

// NewRemoteModbusRTUDevice returns a new RemoteModbusRTUDevice value object.
func NewRemoteModbusRTUDevice(deviceId string, disabled bool, name string, timeoutMs int32, retry int32, detectionIdentifiers map[int32]string, unitId int32, busInterface string, busSettings modbuscfg.SerialSettings, interframeDelayDeciChars int32) RemoteModbusRTUDevice {
	return &_RemoteModbusRTUDevice{NewRemoteModbusDevice(deviceId, disabled, name, timeoutMs, retry, detectionIdentifiers, unitId), busInterface, busSettings, interframeDelayDeciChars}
}

func (r *_RemoteModbusRTUDevice) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.RemoteModbusRTUDevice",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RemoteModbusRTUDevice) BusInterface() string {
	return r.busInterface
}

func (r *_RemoteModbusRTUDevice) BusSettings() modbuscfg.SerialSettings {
	return r.busSettings
}

func (r *_RemoteModbusRTUDevice) InterframeDelayDeciChars() int32 {
	return r.interframeDelayDeciChars
}

func (r *_RemoteModbusRTUDevice) isRemoteModbusRTUDevice() {}

type _RemoteModbusTCPDevice struct {
	RemoteModbusDevice
	ipAddress string
	tcpPort   int32
}

// NewRemoteModbusTCPDevice returns a new RemoteModbusTCPDevice value object.
func NewRemoteModbusTCPDevice(deviceId string, disabled bool, name string, timeoutMs int32, retry int32, detectionIdentifiers map[int32]string, unitId int32, ipAddress string, tcpPort int32) RemoteModbusTCPDevice {
	return &_RemoteModbusTCPDevice{NewRemoteModbusDevice(deviceId, disabled, name, timeoutMs, retry, detectionIdentifiers, unitId), ipAddress, tcpPort}
}

func (r *_RemoteModbusTCPDevice) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.RemoteModbusTCPDevice",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RemoteModbusTCPDevice) IpAddress() string {
	return r.ipAddress
}

func (r *_RemoteModbusTCPDevice) TcpPort() int32 {
	return r.tcpPort
}

func (r *_RemoteModbusTCPDevice) isRemoteModbusTCPDevice() {}

type _RemoteSnmpDevice struct {
	RemoteDevice
	host string
}

// NewRemoteSnmpDevice returns a new RemoteSnmpDevice value object.
func NewRemoteSnmpDevice(deviceId string, disabled bool, name string, timeoutMs int32, retry int32, host string) RemoteSnmpDevice {
	return &_RemoteSnmpDevice{NewRemoteDevice(deviceId, disabled, name, timeoutMs, retry), host}
}

func (r *_RemoteSnmpDevice) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.RemoteSnmpDevice",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RemoteSnmpDevice) Host() string {
	return r.host
}

func (r *_RemoteSnmpDevice) isRemoteSnmpDevice() {}

type _RemoteSnmpV1V2Device struct {
	RemoteSnmpDevice
	community string
}

// NewRemoteSnmpV1V2Device returns a new RemoteSnmpV1V2Device value object.
func NewRemoteSnmpV1V2Device(deviceId string, disabled bool, name string, timeoutMs int32, retry int32, host string, community string) RemoteSnmpV1V2Device {
	return &_RemoteSnmpV1V2Device{NewRemoteSnmpDevice(deviceId, disabled, name, timeoutMs, retry, host), community}
}

func (r *_RemoteSnmpV1V2Device) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.RemoteSnmpV1V2Device",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RemoteSnmpV1V2Device) Community() string {
	return r.community
}

func (r *_RemoteSnmpV1V2Device) isRemoteSnmpV1V2Device() {}

type _RemoteSnmpV3Device struct {
	RemoteSnmpDevice
	user              string
	level             SnmpSecurityLevel
	authProtocol      SnmpAuthProtocol
	authPassphrase    string
	privacyProtocol   SnmpPrivProtocol
	privacyPassphrase string
}

// NewRemoteSnmpV3Device returns a new RemoteSnmpV3Device value object.
func NewRemoteSnmpV3Device(deviceId string, disabled bool, name string, timeoutMs int32, retry int32, host string, user string, level SnmpSecurityLevel, authProtocol SnmpAuthProtocol, authPassphrase string, privacyProtocol SnmpPrivProtocol, privacyPassphrase string) RemoteSnmpV3Device {
	return &_RemoteSnmpV3Device{NewRemoteSnmpDevice(deviceId, disabled, name, timeoutMs, retry, host), user, level, authProtocol, authPassphrase, privacyProtocol, privacyPassphrase}
}

func (r *_RemoteSnmpV3Device) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.RemoteSnmpV3Device",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RemoteSnmpV3Device) User() string {
	return r.user
}

func (r *_RemoteSnmpV3Device) Level() SnmpSecurityLevel {
	return r.level
}

func (r *_RemoteSnmpV3Device) AuthProtocol() SnmpAuthProtocol {
	return r.authProtocol
}

func (r *_RemoteSnmpV3Device) AuthPassphrase() string {
	return r.authPassphrase
}

func (r *_RemoteSnmpV3Device) PrivacyProtocol() SnmpPrivProtocol {
	return r.privacyProtocol
}

func (r *_RemoteSnmpV3Device) PrivacyPassphrase() string {
	return r.privacyPassphrase
}

func (r *_RemoteSnmpV3Device) isRemoteSnmpV3Device() {}

type _InterpretationRule struct {
	idl.ValueObject
	interpretation Interpretation
	ignoreTimeout  int32
}

// NewInterpretationRule returns a new InterpretationRule value object.
func NewInterpretationRule(interpretation Interpretation, ignoreTimeout int32) InterpretationRule {
	return &_InterpretationRule{idl.NewValueObject(), interpretation, ignoreTimeout}
}

func (i *_InterpretationRule) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.InterpretationRule",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRule) Interpretation() Interpretation {
	return i.interpretation
}

func (i *_InterpretationRule) IgnoreTimeout() int32 {
	return i.ignoreTimeout
}

func (i *_InterpretationRule) isInterpretationRule() {}

type _InterpretationRuleInvertable struct {
	InterpretationRule
	invertCondition bool
}

// NewInterpretationRuleInvertable returns a new InterpretationRuleInvertable value object.
func NewInterpretationRuleInvertable(interpretation Interpretation, ignoreTimeout int32, invertCondition bool) InterpretationRuleInvertable {
	return &_InterpretationRuleInvertable{NewInterpretationRule(interpretation, ignoreTimeout), invertCondition}
}

func (i *_InterpretationRuleInvertable) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.InterpretationRuleInvertable",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRuleInvertable) InvertCondition() bool {
	return i.invertCondition
}

func (i *_InterpretationRuleInvertable) isInterpretationRuleInvertable() {}

type _InterpretationRuleModbusException struct {
	InterpretationRuleInvertable
	exceptions []int32
}

// NewInterpretationRuleModbusException returns a new InterpretationRuleModbusException value object.
func NewInterpretationRuleModbusException(interpretation Interpretation, ignoreTimeout int32, invertCondition bool, exceptions []int32) InterpretationRuleModbusException {
	return &_InterpretationRuleModbusException{NewInterpretationRuleInvertable(interpretation, ignoreTimeout, invertCondition), exceptions}
}

func (i *_InterpretationRuleModbusException) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.InterpretationRuleModbusException",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRuleModbusException) Exceptions() []int32 {
	return i.exceptions
}

func (i *_InterpretationRuleModbusException) isInterpretationRuleModbusException() {}

type _InterpretationRuleModbusSystemError struct {
	InterpretationRuleInvertable
	errnos []int32
}

// NewInterpretationRuleModbusSystemError returns a new InterpretationRuleModbusSystemError value object.
func NewInterpretationRuleModbusSystemError(interpretation Interpretation, ignoreTimeout int32, invertCondition bool, errnos []int32) InterpretationRuleModbusSystemError {
	return &_InterpretationRuleModbusSystemError{NewInterpretationRuleInvertable(interpretation, ignoreTimeout, invertCondition), errnos}
}

func (i *_InterpretationRuleModbusSystemError) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.InterpretationRuleModbusSystemError",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRuleModbusSystemError) Errnos() []int32 {
	return i.errnos
}

func (i *_InterpretationRuleModbusSystemError) isInterpretationRuleModbusSystemError() {}

type _InterpretationRuleModbusSpecificError struct {
	InterpretationRuleInvertable
	errors []modbuscfg.SpecificModbusErrors
}

// NewInterpretationRuleModbusSpecificError returns a new InterpretationRuleModbusSpecificError value object.
func NewInterpretationRuleModbusSpecificError(interpretation Interpretation, ignoreTimeout int32, invertCondition bool, errors []modbuscfg.SpecificModbusErrors) InterpretationRuleModbusSpecificError {
	return &_InterpretationRuleModbusSpecificError{NewInterpretationRuleInvertable(interpretation, ignoreTimeout, invertCondition), errors}
}

func (i *_InterpretationRuleModbusSpecificError) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.InterpretationRuleModbusSpecificError",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRuleModbusSpecificError) Errors() []modbuscfg.SpecificModbusErrors {
	return i.errors
}

func (i *_InterpretationRuleModbusSpecificError) isInterpretationRuleModbusSpecificError() {}

type _InterpretationRuleRAW struct {
	InterpretationRuleInvertable
	value int64
	mask  int64
}

// NewInterpretationRuleRAW returns a new InterpretationRuleRAW value object.
func NewInterpretationRuleRAW(interpretation Interpretation, ignoreTimeout int32, invertCondition bool, value int64, mask int64) InterpretationRuleRAW {
	return &_InterpretationRuleRAW{NewInterpretationRuleInvertable(interpretation, ignoreTimeout, invertCondition), value, mask}
}

func (i *_InterpretationRuleRAW) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.InterpretationRuleRAW",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRuleRAW) Value() int64 {
	return i.value
}

func (i *_InterpretationRuleRAW) Mask() int64 {
	return i.mask
}

func (i *_InterpretationRuleRAW) isInterpretationRuleRAW() {}

type _InterpretationRuleRangeRAW struct {
	InterpretationRuleInvertable
	min  int64
	max  int64
	mask int64
}

// NewInterpretationRuleRangeRAW returns a new InterpretationRuleRangeRAW value object.
func NewInterpretationRuleRangeRAW(interpretation Interpretation, ignoreTimeout int32, invertCondition bool, min int64, max int64, mask int64) InterpretationRuleRangeRAW {
	return &_InterpretationRuleRangeRAW{NewInterpretationRuleInvertable(interpretation, ignoreTimeout, invertCondition), min, max, mask}
}

func (i *_InterpretationRuleRangeRAW) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.InterpretationRuleRangeRAW",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRuleRangeRAW) Min() int64 {
	return i.min
}

func (i *_InterpretationRuleRangeRAW) Max() int64 {
	return i.max
}

func (i *_InterpretationRuleRangeRAW) Mask() int64 {
	return i.mask
}

func (i *_InterpretationRuleRangeRAW) isInterpretationRuleRangeRAW() {}

type _InterpretationRuleEnum struct {
	InterpretationRuleInvertable
	enumValues []int64
}

// NewInterpretationRuleEnum returns a new InterpretationRuleEnum value object.
func NewInterpretationRuleEnum(interpretation Interpretation, ignoreTimeout int32, invertCondition bool, enumValues []int64) InterpretationRuleEnum {
	return &_InterpretationRuleEnum{NewInterpretationRuleInvertable(interpretation, ignoreTimeout, invertCondition), enumValues}
}

func (i *_InterpretationRuleEnum) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.InterpretationRuleEnum",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRuleEnum) EnumValues() []int64 {
	return i.enumValues
}

func (i *_InterpretationRuleEnum) isInterpretationRuleEnum() {}

type _InterpretationRuleIEEE754INF struct {
	InterpretationRuleInvertable
}

// NewInterpretationRuleIEEE754INF returns a new InterpretationRuleIEEE754INF value object.
func NewInterpretationRuleIEEE754INF(interpretation Interpretation, ignoreTimeout int32, invertCondition bool) InterpretationRuleIEEE754INF {
	return &_InterpretationRuleIEEE754INF{NewInterpretationRuleInvertable(interpretation, ignoreTimeout, invertCondition)}
}

func (i *_InterpretationRuleIEEE754INF) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.InterpretationRuleIEEE754INF",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRuleIEEE754INF) isInterpretationRuleIEEE754INF() {}

type _InterpretationRuleIEEE754NAN struct {
	InterpretationRuleInvertable
}

// NewInterpretationRuleIEEE754NAN returns a new InterpretationRuleIEEE754NAN value object.
func NewInterpretationRuleIEEE754NAN(interpretation Interpretation, ignoreTimeout int32, invertCondition bool) InterpretationRuleIEEE754NAN {
	return &_InterpretationRuleIEEE754NAN{NewInterpretationRuleInvertable(interpretation, ignoreTimeout, invertCondition)}
}

func (i *_InterpretationRuleIEEE754NAN) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.InterpretationRuleIEEE754NAN",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRuleIEEE754NAN) isInterpretationRuleIEEE754NAN() {}

type _InterpretationRuleCatchAll struct {
	InterpretationRule
}

// NewInterpretationRuleCatchAll returns a new InterpretationRuleCatchAll value object.
func NewInterpretationRuleCatchAll(interpretation Interpretation, ignoreTimeout int32) InterpretationRuleCatchAll {
	return &_InterpretationRuleCatchAll{NewInterpretationRule(interpretation, ignoreTimeout)}
}

func (i *_InterpretationRuleCatchAll) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.InterpretationRuleCatchAll",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRuleCatchAll) isInterpretationRuleCatchAll() {}

type _ValueEncoding struct {
	idl.ValueObject
	encodingId          string
	type_               EncodingType
	invertState         bool
	interpretationRules []InterpretationRule
	minAccessInterval   int32
}

// NewValueEncoding returns a new ValueEncoding value object.
func NewValueEncoding(encodingId string, type_ EncodingType, invertState bool, interpretationRules []InterpretationRule, minAccessInterval int32) ValueEncoding {
	return &_ValueEncoding{idl.NewValueObject(), encodingId, type_, invertState, interpretationRules, minAccessInterval}
}

func (v *_ValueEncoding) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.ValueEncoding",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (v *_ValueEncoding) EncodingId() string {
	return v.encodingId
}

func (v *_ValueEncoding) Type() EncodingType {
	return v.type_
}

func (v *_ValueEncoding) InvertState() bool {
	return v.invertState
}

func (v *_ValueEncoding) InterpretationRules() []InterpretationRule {
	return v.interpretationRules
}

func (v *_ValueEncoding) MinAccessInterval() int32 {
	return v.minAccessInterval
}

func (v *_ValueEncoding) isValueEncoding() {}

type _NumericValueEncoding struct {
	ValueEncoding
	scalingFactor float32
	offset        float32
}

// NewNumericValueEncoding returns a new NumericValueEncoding value object.
func NewNumericValueEncoding(encodingId string, type_ EncodingType, invertState bool, interpretationRules []InterpretationRule, minAccessInterval int32, scalingFactor float32, offset float32) NumericValueEncoding {
	return &_NumericValueEncoding{NewValueEncoding(encodingId, type_, invertState, interpretationRules, minAccessInterval), scalingFactor, offset}
}

func (n *_NumericValueEncoding) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.NumericValueEncoding",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (n *_NumericValueEncoding) ScalingFactor() float32 {
	return n.scalingFactor
}

func (n *_NumericValueEncoding) Offset() float32 {
	return n.offset
}

func (n *_NumericValueEncoding) isNumericValueEncoding() {}

type _ModbusValueEncodingBit struct {
	ValueEncoding
}

// NewModbusValueEncodingBit returns a new ModbusValueEncodingBit value object.
func NewModbusValueEncodingBit(encodingId string, type_ EncodingType, invertState bool, interpretationRules []InterpretationRule, minAccessInterval int32) ModbusValueEncodingBit {
	return &_ModbusValueEncodingBit{NewValueEncoding(encodingId, type_, invertState, interpretationRules, minAccessInterval)}
}

func (m *_ModbusValueEncodingBit) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.ModbusValueEncodingBit",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_ModbusValueEncodingBit) isModbusValueEncodingBit() {}

type _ModbusValueEncoding8 struct {
	NumericValueEncoding
	byteSwap bool
	mask     int64
	start    int32
	width    int32
}

// NewModbusValueEncoding8 returns a new ModbusValueEncoding8 value object.
func NewModbusValueEncoding8(encodingId string, type_ EncodingType, invertState bool, interpretationRules []InterpretationRule, minAccessInterval int32, scalingFactor float32, offset float32, byteSwap bool, mask int64, start int32, width int32) ModbusValueEncoding8 {
	return &_ModbusValueEncoding8{NewNumericValueEncoding(encodingId, type_, invertState, interpretationRules, minAccessInterval, scalingFactor, offset), byteSwap, mask, start, width}
}

func (m *_ModbusValueEncoding8) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.ModbusValueEncoding8",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_ModbusValueEncoding8) ByteSwap() bool {
	return m.byteSwap
}

func (m *_ModbusValueEncoding8) Mask() int64 {
	return m.mask
}

func (m *_ModbusValueEncoding8) Start() int32 {
	return m.start
}

func (m *_ModbusValueEncoding8) Width() int32 {
	return m.width
}

func (m *_ModbusValueEncoding8) isModbusValueEncoding8() {}

type _ModbusValueEncoding16 struct {
	ModbusValueEncoding8
}

// NewModbusValueEncoding16 returns a new ModbusValueEncoding16 value object.
func NewModbusValueEncoding16(encodingId string, type_ EncodingType, invertState bool, interpretationRules []InterpretationRule, minAccessInterval int32, scalingFactor float32, offset float32, byteSwap bool, mask int64, start int32, width int32) ModbusValueEncoding16 {
	return &_ModbusValueEncoding16{NewModbusValueEncoding8(encodingId, type_, invertState, interpretationRules, minAccessInterval, scalingFactor, offset, byteSwap, mask, start, width)}
}

func (m *_ModbusValueEncoding16) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.ModbusValueEncoding16",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_ModbusValueEncoding16) isModbusValueEncoding16() {}

type _ModbusValueEncoding32 struct {
	ModbusValueEncoding16
	endianness ModbusEndianness
}

// NewModbusValueEncoding32 returns a new ModbusValueEncoding32 value object.
func NewModbusValueEncoding32(encodingId string, type_ EncodingType, invertState bool, interpretationRules []InterpretationRule, minAccessInterval int32, scalingFactor float32, offset float32, byteSwap bool, mask int64, start int32, width int32, endianness ModbusEndianness) ModbusValueEncoding32 {
	return &_ModbusValueEncoding32{NewModbusValueEncoding16(encodingId, type_, invertState, interpretationRules, minAccessInterval, scalingFactor, offset, byteSwap, mask, start, width), endianness}
}

func (m *_ModbusValueEncoding32) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.ModbusValueEncoding32",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_ModbusValueEncoding32) Endianness() ModbusEndianness {
	return m.endianness
}

func (m *_ModbusValueEncoding32) isModbusValueEncoding32() {}

type _ModbusValueEncoding48 struct {
	ModbusValueEncoding32
}

// NewModbusValueEncoding48 returns a new ModbusValueEncoding48 value object.
func NewModbusValueEncoding48(encodingId string, type_ EncodingType, invertState bool, interpretationRules []InterpretationRule, minAccessInterval int32, scalingFactor float32, offset float32, byteSwap bool, mask int64, start int32, width int32, endianness ModbusEndianness) ModbusValueEncoding48 {
	return &_ModbusValueEncoding48{NewModbusValueEncoding32(encodingId, type_, invertState, interpretationRules, minAccessInterval, scalingFactor, offset, byteSwap, mask, start, width, endianness)}
}

func (m *_ModbusValueEncoding48) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.ModbusValueEncoding48",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_ModbusValueEncoding48) isModbusValueEncoding48() {}

type _ModbusValueEncoding64 struct {
	ModbusValueEncoding32
}

// NewModbusValueEncoding64 returns a new ModbusValueEncoding64 value object.
func NewModbusValueEncoding64(encodingId string, type_ EncodingType, invertState bool, interpretationRules []InterpretationRule, minAccessInterval int32, scalingFactor float32, offset float32, byteSwap bool, mask int64, start int32, width int32, endianness ModbusEndianness) ModbusValueEncoding64 {
	return &_ModbusValueEncoding64{NewModbusValueEncoding32(encodingId, type_, invertState, interpretationRules, minAccessInterval, scalingFactor, offset, byteSwap, mask, start, width, endianness)}
}

func (m *_ModbusValueEncoding64) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.ModbusValueEncoding64",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_ModbusValueEncoding64) isModbusValueEncoding64() {}

type _Sensor struct {
	idl.ValueObject
	sensorId    string
	disabled    bool
	deviceId    string
	classId     string
	encodingId  string
	defaultName string
}

// NewSensor returns a new Sensor value object.
func NewSensor(sensorId string, disabled bool, deviceId string, classId string, encodingId string, defaultName string) Sensor {
	return &_Sensor{idl.NewValueObject(), sensorId, disabled, deviceId, classId, encodingId, defaultName}
}

func (s *_Sensor) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.Sensor",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_Sensor) SensorId() string {
	return s.sensorId
}

func (s *_Sensor) Disabled() bool {
	return s.disabled
}

func (s *_Sensor) DeviceId() string {
	return s.deviceId
}

func (s *_Sensor) ClassId() string {
	return s.classId
}

func (s *_Sensor) EncodingId() string {
	return s.encodingId
}

func (s *_Sensor) DefaultName() string {
	return s.defaultName
}

func (s *_Sensor) isSensor() {}

type _ModbusSensor struct {
	Sensor
	function modbuscfg.ModbusFunction
	regAddr  int32
}

// NewModbusSensor returns a new ModbusSensor value object.
func NewModbusSensor(sensorId string, disabled bool, deviceId string, classId string, encodingId string, defaultName string, function modbuscfg.ModbusFunction, regAddr int32) ModbusSensor {
	return &_ModbusSensor{NewSensor(sensorId, disabled, deviceId, classId, encodingId, defaultName), function, regAddr}
}

func (m *_ModbusSensor) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.ModbusSensor",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_ModbusSensor) Function() modbuscfg.ModbusFunction {
	return m.function
}

func (m *_ModbusSensor) RegAddr() int32 {
	return m.regAddr
}

func (m *_ModbusSensor) isModbusSensor() {}

type _SnmpSensor struct {
	Sensor
	oid string
}

// NewSnmpSensor returns a new SnmpSensor value object.
func NewSnmpSensor(sensorId string, disabled bool, deviceId string, classId string, encodingId string, defaultName string, oid string) SnmpSensor {
	return &_SnmpSensor{NewSensor(sensorId, disabled, deviceId, classId, encodingId, defaultName), oid}
}

func (s *_SnmpSensor) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.SnmpSensor",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_SnmpSensor) Oid() string {
	return s.oid
}

func (s *_SnmpSensor) isSnmpSensor() {}

type _ConfigurationChangedEvent struct {
	event.Event
	configuration map[string]ConfigurationPackage
}

func (c *_ConfigurationChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.ConfigurationChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_ConfigurationChangedEvent) Configuration() map[string]ConfigurationPackage {
	return c.configuration
}

func (c *_ConfigurationChangedEvent) isConfigurationChangedEvent() {}

type _Feedback struct {
	idl.ValueObject
	currentState FeedbackObjectFeedbackState
	infos        []FeedbackObject
}

func (f *_Feedback) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.Feedback",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (f *_Feedback) CurrentState() FeedbackObjectFeedbackState {
	return f.currentState
}

func (f *_Feedback) Infos() []FeedbackObject {
	return f.infos
}

func (f *_Feedback) isFeedback() {}

type _DeviceFeedback struct {
	Feedback
	packageId string
	deviceId  string
}

func (d *_DeviceFeedback) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.DeviceFeedback",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (d *_DeviceFeedback) PackageId() string {
	return d.packageId
}

func (d *_DeviceFeedback) DeviceId() string {
	return d.deviceId
}

func (d *_DeviceFeedback) isDeviceFeedback() {}

type _SensorFeedback struct {
	Feedback
	packageId string
	deviceId  string
	sensorId  string
}

func (s *_SensorFeedback) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.SensorFeedback",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (s *_SensorFeedback) PackageId() string {
	return s.packageId
}

func (s *_SensorFeedback) DeviceId() string {
	return s.deviceId
}

func (s *_SensorFeedback) SensorId() string {
	return s.sensorId
}

func (s *_SensorFeedback) isSensorFeedback() {}

type _FeedbackChangedEvent struct {
	event.Event
	feedback []Feedback
}

func (f *_FeedbackChangedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_2_0_1.FeedbackChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (f *_FeedbackChangedEvent) Feedback() []Feedback {
	return f.feedback
}

func (f *_FeedbackChangedEvent) isFeedbackChangedEvent() {}
