// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package gatewaysensormanager

import (
	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/idl/event"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40510/peripheral/modbuscfg"
	"github.com/arminguenther/xeruspower-go/v40510/sensors/numericsensor"
	"github.com/arminguenther/xeruspower-go/v40510/sensors/sensor"
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
		Name:  "peripheral.GatewaySensorManager_3_0_0.SensorClass",
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
		Name:  "peripheral.GatewaySensorManager_3_0_0.NumericSensorClass",
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
		Name:  "peripheral.GatewaySensorManager_3_0_0.StateSensorClass",
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
		Name:  "peripheral.GatewaySensorManager_3_0_0.SwitchSensorClass",
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
		Name:  "peripheral.GatewaySensorManager_3_0_0.RemoteDevice",
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
	detectionIdentifiers     map[int32]string
	unitId                   int32
	unsupportedFunctionCodes []byte
}

// NewRemoteModbusDevice returns a new RemoteModbusDevice value object.
func NewRemoteModbusDevice(deviceId string, disabled bool, name string, timeoutMs int32, retry int32, detectionIdentifiers map[int32]string, unitId int32, unsupportedFunctionCodes []byte) RemoteModbusDevice {
	return &_RemoteModbusDevice{NewRemoteDevice(deviceId, disabled, name, timeoutMs, retry), detectionIdentifiers, unitId, unsupportedFunctionCodes}
}

func (r *_RemoteModbusDevice) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.RemoteModbusDevice",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (r *_RemoteModbusDevice) DetectionIdentifiers() map[int32]string {
	return r.detectionIdentifiers
}

func (r *_RemoteModbusDevice) UnitId() int32 {
	return r.unitId
}

func (r *_RemoteModbusDevice) UnsupportedFunctionCodes() []byte {
	return r.unsupportedFunctionCodes
}

func (r *_RemoteModbusDevice) isRemoteModbusDevice() {}

type _RemoteModbusRTUDevice struct {
	RemoteModbusDevice
	busInterface             string
	busSettings              modbuscfg.SerialSettings
	interframeDelayDeciChars int32
}

// NewRemoteModbusRTUDevice returns a new RemoteModbusRTUDevice value object.
func NewRemoteModbusRTUDevice(deviceId string, disabled bool, name string, timeoutMs int32, retry int32, detectionIdentifiers map[int32]string, unitId int32, unsupportedFunctionCodes []byte, busInterface string, busSettings modbuscfg.SerialSettings, interframeDelayDeciChars int32) RemoteModbusRTUDevice {
	return &_RemoteModbusRTUDevice{NewRemoteModbusDevice(deviceId, disabled, name, timeoutMs, retry, detectionIdentifiers, unitId, unsupportedFunctionCodes), busInterface, busSettings, interframeDelayDeciChars}
}

func (r *_RemoteModbusRTUDevice) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.RemoteModbusRTUDevice",
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
func NewRemoteModbusTCPDevice(deviceId string, disabled bool, name string, timeoutMs int32, retry int32, detectionIdentifiers map[int32]string, unitId int32, unsupportedFunctionCodes []byte, ipAddress string, tcpPort int32) RemoteModbusTCPDevice {
	return &_RemoteModbusTCPDevice{NewRemoteModbusDevice(deviceId, disabled, name, timeoutMs, retry, detectionIdentifiers, unitId, unsupportedFunctionCodes), ipAddress, tcpPort}
}

func (r *_RemoteModbusTCPDevice) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.RemoteModbusTCPDevice",
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
		Name:  "peripheral.GatewaySensorManager_3_0_0.RemoteSnmpDevice",
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
		Name:  "peripheral.GatewaySensorManager_3_0_0.RemoteSnmpV1V2Device",
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
		Name:  "peripheral.GatewaySensorManager_3_0_0.RemoteSnmpV3Device",
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
	accessType     AccessType
}

// NewInterpretationRule returns a new InterpretationRule value object.
func NewInterpretationRule(interpretation Interpretation, ignoreTimeout int32, accessType AccessType) InterpretationRule {
	return &_InterpretationRule{idl.NewValueObject(), interpretation, ignoreTimeout, accessType}
}

func (i *_InterpretationRule) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.InterpretationRule",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRule) Interpretation() Interpretation {
	return i.interpretation
}

func (i *_InterpretationRule) IgnoreTimeout() int32 {
	return i.ignoreTimeout
}

func (i *_InterpretationRule) AccessType() AccessType {
	return i.accessType
}

func (i *_InterpretationRule) isInterpretationRule() {}

type _InterpretationRuleInvertable struct {
	InterpretationRule
	invertCondition bool
}

// NewInterpretationRuleInvertable returns a new InterpretationRuleInvertable value object.
func NewInterpretationRuleInvertable(interpretation Interpretation, ignoreTimeout int32, accessType AccessType, invertCondition bool) InterpretationRuleInvertable {
	return &_InterpretationRuleInvertable{NewInterpretationRule(interpretation, ignoreTimeout, accessType), invertCondition}
}

func (i *_InterpretationRuleInvertable) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.InterpretationRuleInvertable",
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
func NewInterpretationRuleModbusException(interpretation Interpretation, ignoreTimeout int32, accessType AccessType, invertCondition bool, exceptions []int32) InterpretationRuleModbusException {
	return &_InterpretationRuleModbusException{NewInterpretationRuleInvertable(interpretation, ignoreTimeout, accessType, invertCondition), exceptions}
}

func (i *_InterpretationRuleModbusException) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.InterpretationRuleModbusException",
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
func NewInterpretationRuleModbusSystemError(interpretation Interpretation, ignoreTimeout int32, accessType AccessType, invertCondition bool, errnos []int32) InterpretationRuleModbusSystemError {
	return &_InterpretationRuleModbusSystemError{NewInterpretationRuleInvertable(interpretation, ignoreTimeout, accessType, invertCondition), errnos}
}

func (i *_InterpretationRuleModbusSystemError) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.InterpretationRuleModbusSystemError",
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
func NewInterpretationRuleModbusSpecificError(interpretation Interpretation, ignoreTimeout int32, accessType AccessType, invertCondition bool, errors []modbuscfg.SpecificModbusErrors) InterpretationRuleModbusSpecificError {
	return &_InterpretationRuleModbusSpecificError{NewInterpretationRuleInvertable(interpretation, ignoreTimeout, accessType, invertCondition), errors}
}

func (i *_InterpretationRuleModbusSpecificError) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.InterpretationRuleModbusSpecificError",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRuleModbusSpecificError) Errors() []modbuscfg.SpecificModbusErrors {
	return i.errors
}

func (i *_InterpretationRuleModbusSpecificError) isInterpretationRuleModbusSpecificError() {}

type _InterpretationRuleRAW struct {
	InterpretationRuleInvertable
	value string
	mask  string
}

// NewInterpretationRuleRAW returns a new InterpretationRuleRAW value object.
func NewInterpretationRuleRAW(interpretation Interpretation, ignoreTimeout int32, accessType AccessType, invertCondition bool, value string, mask string) InterpretationRuleRAW {
	return &_InterpretationRuleRAW{NewInterpretationRuleInvertable(interpretation, ignoreTimeout, accessType, invertCondition), value, mask}
}

func (i *_InterpretationRuleRAW) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.InterpretationRuleRAW",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRuleRAW) Value() string {
	return i.value
}

func (i *_InterpretationRuleRAW) Mask() string {
	return i.mask
}

func (i *_InterpretationRuleRAW) isInterpretationRuleRAW() {}

type _InterpretationRuleRangeRAW struct {
	InterpretationRuleInvertable
	min  string
	max  string
	mask string
}

// NewInterpretationRuleRangeRAW returns a new InterpretationRuleRangeRAW value object.
func NewInterpretationRuleRangeRAW(interpretation Interpretation, ignoreTimeout int32, accessType AccessType, invertCondition bool, min string, max string, mask string) InterpretationRuleRangeRAW {
	return &_InterpretationRuleRangeRAW{NewInterpretationRuleInvertable(interpretation, ignoreTimeout, accessType, invertCondition), min, max, mask}
}

func (i *_InterpretationRuleRangeRAW) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.InterpretationRuleRangeRAW",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRuleRangeRAW) Min() string {
	return i.min
}

func (i *_InterpretationRuleRangeRAW) Max() string {
	return i.max
}

func (i *_InterpretationRuleRangeRAW) Mask() string {
	return i.mask
}

func (i *_InterpretationRuleRangeRAW) isInterpretationRuleRangeRAW() {}

type _InterpretationRuleEnum struct {
	InterpretationRuleInvertable
	enumValues []string
}

// NewInterpretationRuleEnum returns a new InterpretationRuleEnum value object.
func NewInterpretationRuleEnum(interpretation Interpretation, ignoreTimeout int32, accessType AccessType, invertCondition bool, enumValues []string) InterpretationRuleEnum {
	return &_InterpretationRuleEnum{NewInterpretationRuleInvertable(interpretation, ignoreTimeout, accessType, invertCondition), enumValues}
}

func (i *_InterpretationRuleEnum) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.InterpretationRuleEnum",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRuleEnum) EnumValues() []string {
	return i.enumValues
}

func (i *_InterpretationRuleEnum) isInterpretationRuleEnum() {}

type _InterpretationRuleIEEE754INF struct {
	InterpretationRuleInvertable
}

// NewInterpretationRuleIEEE754INF returns a new InterpretationRuleIEEE754INF value object.
func NewInterpretationRuleIEEE754INF(interpretation Interpretation, ignoreTimeout int32, accessType AccessType, invertCondition bool) InterpretationRuleIEEE754INF {
	return &_InterpretationRuleIEEE754INF{NewInterpretationRuleInvertable(interpretation, ignoreTimeout, accessType, invertCondition)}
}

func (i *_InterpretationRuleIEEE754INF) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.InterpretationRuleIEEE754INF",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRuleIEEE754INF) isInterpretationRuleIEEE754INF() {}

type _InterpretationRuleIEEE754NAN struct {
	InterpretationRuleInvertable
}

// NewInterpretationRuleIEEE754NAN returns a new InterpretationRuleIEEE754NAN value object.
func NewInterpretationRuleIEEE754NAN(interpretation Interpretation, ignoreTimeout int32, accessType AccessType, invertCondition bool) InterpretationRuleIEEE754NAN {
	return &_InterpretationRuleIEEE754NAN{NewInterpretationRuleInvertable(interpretation, ignoreTimeout, accessType, invertCondition)}
}

func (i *_InterpretationRuleIEEE754NAN) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.InterpretationRuleIEEE754NAN",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_InterpretationRuleIEEE754NAN) isInterpretationRuleIEEE754NAN() {}

type _InterpretationRuleCatchAll struct {
	InterpretationRule
}

// NewInterpretationRuleCatchAll returns a new InterpretationRuleCatchAll value object.
func NewInterpretationRuleCatchAll(interpretation Interpretation, ignoreTimeout int32, accessType AccessType) InterpretationRuleCatchAll {
	return &_InterpretationRuleCatchAll{NewInterpretationRule(interpretation, ignoreTimeout, accessType)}
}

func (i *_InterpretationRuleCatchAll) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.InterpretationRuleCatchAll",
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
		Name:  "peripheral.GatewaySensorManager_3_0_0.ValueEncoding",
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
		Name:  "peripheral.GatewaySensorManager_3_0_0.NumericValueEncoding",
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
		Name:  "peripheral.GatewaySensorManager_3_0_0.ModbusValueEncodingBit",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_ModbusValueEncodingBit) isModbusValueEncodingBit() {}

type _ModbusValueEncoding8 struct {
	NumericValueEncoding
	byteSwap bool
	mask     string
	start    int32
	width    int32
}

// NewModbusValueEncoding8 returns a new ModbusValueEncoding8 value object.
func NewModbusValueEncoding8(encodingId string, type_ EncodingType, invertState bool, interpretationRules []InterpretationRule, minAccessInterval int32, scalingFactor float32, offset float32, byteSwap bool, mask string, start int32, width int32) ModbusValueEncoding8 {
	return &_ModbusValueEncoding8{NewNumericValueEncoding(encodingId, type_, invertState, interpretationRules, minAccessInterval, scalingFactor, offset), byteSwap, mask, start, width}
}

func (m *_ModbusValueEncoding8) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.ModbusValueEncoding8",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_ModbusValueEncoding8) ByteSwap() bool {
	return m.byteSwap
}

func (m *_ModbusValueEncoding8) Mask() string {
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
func NewModbusValueEncoding16(encodingId string, type_ EncodingType, invertState bool, interpretationRules []InterpretationRule, minAccessInterval int32, scalingFactor float32, offset float32, byteSwap bool, mask string, start int32, width int32) ModbusValueEncoding16 {
	return &_ModbusValueEncoding16{NewModbusValueEncoding8(encodingId, type_, invertState, interpretationRules, minAccessInterval, scalingFactor, offset, byteSwap, mask, start, width)}
}

func (m *_ModbusValueEncoding16) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.ModbusValueEncoding16",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_ModbusValueEncoding16) isModbusValueEncoding16() {}

type _ModbusValueEncoding32 struct {
	ModbusValueEncoding16
	endianness ModbusEndianness
}

// NewModbusValueEncoding32 returns a new ModbusValueEncoding32 value object.
func NewModbusValueEncoding32(encodingId string, type_ EncodingType, invertState bool, interpretationRules []InterpretationRule, minAccessInterval int32, scalingFactor float32, offset float32, byteSwap bool, mask string, start int32, width int32, endianness ModbusEndianness) ModbusValueEncoding32 {
	return &_ModbusValueEncoding32{NewModbusValueEncoding16(encodingId, type_, invertState, interpretationRules, minAccessInterval, scalingFactor, offset, byteSwap, mask, start, width), endianness}
}

func (m *_ModbusValueEncoding32) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.ModbusValueEncoding32",
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
func NewModbusValueEncoding48(encodingId string, type_ EncodingType, invertState bool, interpretationRules []InterpretationRule, minAccessInterval int32, scalingFactor float32, offset float32, byteSwap bool, mask string, start int32, width int32, endianness ModbusEndianness) ModbusValueEncoding48 {
	return &_ModbusValueEncoding48{NewModbusValueEncoding32(encodingId, type_, invertState, interpretationRules, minAccessInterval, scalingFactor, offset, byteSwap, mask, start, width, endianness)}
}

func (m *_ModbusValueEncoding48) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.ModbusValueEncoding48",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_ModbusValueEncoding48) isModbusValueEncoding48() {}

type _ModbusValueEncoding64 struct {
	ModbusValueEncoding32
}

// NewModbusValueEncoding64 returns a new ModbusValueEncoding64 value object.
func NewModbusValueEncoding64(encodingId string, type_ EncodingType, invertState bool, interpretationRules []InterpretationRule, minAccessInterval int32, scalingFactor float32, offset float32, byteSwap bool, mask string, start int32, width int32, endianness ModbusEndianness) ModbusValueEncoding64 {
	return &_ModbusValueEncoding64{NewModbusValueEncoding32(encodingId, type_, invertState, interpretationRules, minAccessInterval, scalingFactor, offset, byteSwap, mask, start, width, endianness)}
}

func (m *_ModbusValueEncoding64) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.ModbusValueEncoding64",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (m *_ModbusValueEncoding64) isModbusValueEncoding64() {}

type _Sensor struct {
	idl.ValueObject
	sensorId           string
	disabled           bool
	deviceId           string
	classId            string
	encodingId         string
	defaultName        string
	defaultDescription string
	defaultLocationX   string
	defaultLocationY   string
	defaultLocationZ   string
}

// NewSensor returns a new Sensor value object.
func NewSensor(sensorId string, disabled bool, deviceId string, classId string, encodingId string, defaultName string, defaultDescription string, defaultLocationX string, defaultLocationY string, defaultLocationZ string) Sensor {
	return &_Sensor{idl.NewValueObject(), sensorId, disabled, deviceId, classId, encodingId, defaultName, defaultDescription, defaultLocationX, defaultLocationY, defaultLocationZ}
}

func (s *_Sensor) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.Sensor",
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

func (s *_Sensor) DefaultDescription() string {
	return s.defaultDescription
}

func (s *_Sensor) DefaultLocationX() string {
	return s.defaultLocationX
}

func (s *_Sensor) DefaultLocationY() string {
	return s.defaultLocationY
}

func (s *_Sensor) DefaultLocationZ() string {
	return s.defaultLocationZ
}

func (s *_Sensor) isSensor() {}

type _ModbusSensor struct {
	Sensor
	function modbuscfg.ModbusFunction
	regAddr  int32
}

// NewModbusSensor returns a new ModbusSensor value object.
func NewModbusSensor(sensorId string, disabled bool, deviceId string, classId string, encodingId string, defaultName string, defaultDescription string, defaultLocationX string, defaultLocationY string, defaultLocationZ string, function modbuscfg.ModbusFunction, regAddr int32) ModbusSensor {
	return &_ModbusSensor{NewSensor(sensorId, disabled, deviceId, classId, encodingId, defaultName, defaultDescription, defaultLocationX, defaultLocationY, defaultLocationZ), function, regAddr}
}

func (m *_ModbusSensor) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.ModbusSensor",
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
func NewSnmpSensor(sensorId string, disabled bool, deviceId string, classId string, encodingId string, defaultName string, defaultDescription string, defaultLocationX string, defaultLocationY string, defaultLocationZ string, oid string) SnmpSensor {
	return &_SnmpSensor{NewSensor(sensorId, disabled, deviceId, classId, encodingId, defaultName, defaultDescription, defaultLocationX, defaultLocationY, defaultLocationZ), oid}
}

func (s *_SnmpSensor) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "peripheral.GatewaySensorManager_3_0_0.SnmpSensor",
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
		Name:  "peripheral.GatewaySensorManager_3_0_0.ConfigurationChangedEvent",
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
		Name:  "peripheral.GatewaySensorManager_3_0_0.Feedback",
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
		Name:  "peripheral.GatewaySensorManager_3_0_0.DeviceFeedback",
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
		Name:  "peripheral.GatewaySensorManager_3_0_0.SensorFeedback",
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
		Name:  "peripheral.GatewaySensorManager_3_0_0.FeedbackChangedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (f *_FeedbackChangedEvent) Feedback() []Feedback {
	return f.feedback
}

func (f *_FeedbackChangedEvent) isFeedbackChangedEvent() {}
