// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package pdu

import (
	"context"

	"github.com/arminguenther/xeruspower-go/hmi/internalbeeper"
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
	"github.com/arminguenther/xeruspower-go/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/pdumodel/controller"
	"github.com/arminguenther/xeruspower-go/pdumodel/inlet"
	"github.com/arminguenther/xeruspower-go/pdumodel/nameplate"
	"github.com/arminguenther/xeruspower-go/pdumodel/outlet"
	"github.com/arminguenther/xeruspower-go/pdumodel/overcurrentprotector"
	"github.com/arminguenther/xeruspower-go/pdumodel/transferswitch"
	"github.com/arminguenther/xeruspower-go/peripheral/peripheraldevicemanager"
	"github.com/arminguenther/xeruspower-go/portsmodel/port"
	"github.com/arminguenther/xeruspower-go/sensors/alertedsensormanager"
	"github.com/arminguenther/xeruspower-go/sensors/sensorlogger"
)

func init() {
	object.Register(NewPdu)
}

type _Pdu struct {
	idl.Object
}

// NewPdu returns a new Pdu interface for the object at given RID.
func NewPdu(rid string, caller idl.Caller) Pdu {
	return &_Pdu{idl.NewObject(rid, caller)}
}

func (p *_Pdu) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "pdumodel.Pdu",
		Major: 6, Submajor: 2, Minor: 2,
	}
}

func (p *_Pdu) GetNameplate(ctx context.Context) (nameplate.Nameplate, error) {
	var ret nameplate.Nameplate
	val, err := p.Caller().Call(ctx, p.RID(), "getNameplate", nil)
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
	err = ret.Decode(res["_ret_"], p.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (p *_Pdu) GetMetaData(ctx context.Context) (MetaData, error) {
	var ret MetaData
	val, err := p.Caller().Call(ctx, p.RID(), "getMetaData", nil)
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
	err = ret.Decode(res["_ret_"], p.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (p *_Pdu) GetSensors(ctx context.Context) (Sensors, error) {
	var ret Sensors
	val, err := p.Caller().Call(ctx, p.RID(), "getSensors", nil)
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
	err = ret.Decode(res["_ret_"], p.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (p *_Pdu) GetSensorLogger(ctx context.Context) (sensorlogger.Logger, error) {
	var ret sensorlogger.Logger
	val, err := p.Caller().Call(ctx, p.RID(), "getSensorLogger", nil)
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
	ret, err = object.As[sensorlogger.Logger](res["_ret_"], p.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (p *_Pdu) GetAlertedSensorManager(ctx context.Context) (alertedsensormanager.AlertedSensorManager, error) {
	var ret alertedsensormanager.AlertedSensorManager
	val, err := p.Caller().Call(ctx, p.RID(), "getAlertedSensorManager", nil)
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
	ret, err = object.As[alertedsensormanager.AlertedSensorManager](res["_ret_"], p.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (p *_Pdu) GetControllers(ctx context.Context) ([]controller.Controller, error) {
	var ret []controller.Controller
	val, err := p.Caller().Call(ctx, p.RID(), "getControllers", nil)
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
	ret = make([]controller.Controller, 0, len(s0))
	for _, a0 := range s0 {
		var e0 controller.Controller
		e0, err = object.As[controller.Controller](a0, p.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (p *_Pdu) GetOutlets(ctx context.Context) ([]outlet.Outlet, error) {
	var ret []outlet.Outlet
	val, err := p.Caller().Call(ctx, p.RID(), "getOutlets", nil)
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
	ret = make([]outlet.Outlet, 0, len(s0))
	for _, a0 := range s0 {
		var e0 outlet.Outlet
		e0, err = object.As[outlet.Outlet](a0, p.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (p *_Pdu) GetOverCurrentProtectors(ctx context.Context) ([]overcurrentprotector.OverCurrentProtector, error) {
	var ret []overcurrentprotector.OverCurrentProtector
	val, err := p.Caller().Call(ctx, p.RID(), "getOverCurrentProtectors", nil)
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
	ret = make([]overcurrentprotector.OverCurrentProtector, 0, len(s0))
	for _, a0 := range s0 {
		var e0 overcurrentprotector.OverCurrentProtector
		e0, err = object.As[overcurrentprotector.OverCurrentProtector](a0, p.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (p *_Pdu) GetInlets(ctx context.Context) ([]inlet.Inlet, error) {
	var ret []inlet.Inlet
	val, err := p.Caller().Call(ctx, p.RID(), "getInlets", nil)
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
	ret = make([]inlet.Inlet, 0, len(s0))
	for _, a0 := range s0 {
		var e0 inlet.Inlet
		e0, err = object.As[inlet.Inlet](a0, p.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (p *_Pdu) GetTransferSwitches(ctx context.Context) ([]transferswitch.TransferSwitch, error) {
	var ret []transferswitch.TransferSwitch
	val, err := p.Caller().Call(ctx, p.RID(), "getTransferSwitches", nil)
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
	ret = make([]transferswitch.TransferSwitch, 0, len(s0))
	for _, a0 := range s0 {
		var e0 transferswitch.TransferSwitch
		e0, err = object.As[transferswitch.TransferSwitch](a0, p.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (p *_Pdu) GetPeripheralDeviceManager(ctx context.Context) (peripheraldevicemanager.DeviceManager, error) {
	var ret peripheraldevicemanager.DeviceManager
	val, err := p.Caller().Call(ctx, p.RID(), "getPeripheralDeviceManager", nil)
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
	ret, err = object.As[peripheraldevicemanager.DeviceManager](res["_ret_"], p.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (p *_Pdu) GetBeeper(ctx context.Context) (internalbeeper.InternalBeeper, error) {
	var ret internalbeeper.InternalBeeper
	val, err := p.Caller().Call(ctx, p.RID(), "getBeeper", nil)
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
	ret, err = object.As[internalbeeper.InternalBeeper](res["_ret_"], p.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (p *_Pdu) GetSettings(ctx context.Context) (Settings, error) {
	var ret Settings
	val, err := p.Caller().Call(ctx, p.RID(), "getSettings", nil)
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
	err = ret.Decode(res["_ret_"], p.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (p *_Pdu) IsLoadSheddingActive(ctx context.Context) (bool, error) {
	var ret bool
	val, err := p.Caller().Call(ctx, p.RID(), "isLoadSheddingActive", nil)
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

func (p *_Pdu) SetSettings(ctx context.Context, in0 Settings) (int32, error) {
	var ret int32
	val, err := p.Caller().Call(ctx, p.RID(), "setSettings", map[string]any{
		"settings": in0.Encode(),
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

func (p *_Pdu) SetLoadSheddingActive(ctx context.Context, in0 bool) error {
	_, err := p.Caller().Call(ctx, p.RID(), "setLoadSheddingActive", map[string]any{
		"active": in0,
	})
	return err
}

func (p *_Pdu) GetFeaturePorts(ctx context.Context) ([]port.Port, error) {
	var ret []port.Port
	val, err := p.Caller().Call(ctx, p.RID(), "getFeaturePorts", nil)
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
	ret = make([]port.Port, 0, len(s0))
	for _, a0 := range s0 {
		var e0 port.Port
		e0, err = object.As[port.Port](a0, p.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (p *_Pdu) GetSensorPorts(ctx context.Context) ([]port.Port, error) {
	var ret []port.Port
	val, err := p.Caller().Call(ctx, p.RID(), "getSensorPorts", nil)
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
	ret = make([]port.Port, 0, len(s0))
	for _, a0 := range s0 {
		var e0 port.Port
		e0, err = object.As[port.Port](a0, p.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (p *_Pdu) GetRemoteHubPorts(ctx context.Context) ([]port.Port, error) {
	var ret []port.Port
	val, err := p.Caller().Call(ctx, p.RID(), "getRemoteHubPorts", nil)
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
	ret = make([]port.Port, 0, len(s0))
	for _, a0 := range s0 {
		var e0 port.Port
		e0, err = object.As[port.Port](a0, p.Caller())
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (p *_Pdu) EnterRS485ConfigModeAndAssignCtrlBoardAddress(ctx context.Context, in0 int32) (int32, error) {
	var ret int32
	val, err := p.Caller().Call(ctx, p.RID(), "enterRS485ConfigModeAndAssignCtrlBoardAddress", map[string]any{
		"addr": in0,
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

func (p *_Pdu) EnterRS485ConfigModeAndAssignSCBoardAddress(ctx context.Context, in0 int32, in1 int32) (int32, error) {
	var ret int32
	val, err := p.Caller().Call(ctx, p.RID(), "enterRS485ConfigModeAndAssignSCBoardAddress", map[string]any{
		"deviceId": in0,
		"addr":     in1,
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

func (p *_Pdu) LeaveRS485ConfigMode(ctx context.Context) (int32, error) {
	var ret int32
	val, err := p.Caller().Call(ctx, p.RID(), "leaveRS485ConfigMode", nil)
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

func (p *_Pdu) SetAllOutletPowerStates(ctx context.Context, in0 outlet.PowerState) (int32, error) {
	var ret int32
	val, err := p.Caller().Call(ctx, p.RID(), "setAllOutletPowerStates", map[string]any{
		"pstate": in0,
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

func (p *_Pdu) SetMultipleOutletPowerStates(ctx context.Context, in0 []int32, in1 outlet.PowerState, in2 bool) (int32, error) {
	var ret int32
	val, err := p.Caller().Call(ctx, p.RID(), "setMultipleOutletPowerStates", map[string]any{
		"outletNumbers":   encoding.NonNilSlice(in0),
		"state":           in1,
		"respectSequence": in2,
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

func (p *_Pdu) CycleAllOutletPowerStates(ctx context.Context) (int32, error) {
	var ret int32
	val, err := p.Caller().Call(ctx, p.RID(), "cycleAllOutletPowerStates", nil)
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

func (p *_Pdu) CycleMultipleOutletPowerStates(ctx context.Context, in0 []int32, in1 bool) (int32, error) {
	var ret int32
	val, err := p.Caller().Call(ctx, p.RID(), "cycleMultipleOutletPowerStates", map[string]any{
		"outletNumbers":   encoding.NonNilSlice(in0),
		"respectSequence": in1,
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

func (p *_Pdu) GetStatistic(ctx context.Context) (Statistic, error) {
	var ret Statistic
	val, err := p.Caller().Call(ctx, p.RID(), "getStatistic", nil)
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
	err = ret.Decode(res["_ret_"], p.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (p *_Pdu) GetOutletSequenceState(ctx context.Context) (OutletSequenceState, error) {
	var ret OutletSequenceState
	val, err := p.Caller().Call(ctx, p.RID(), "getOutletSequenceState", nil)
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
	err = ret.Decode(res["_ret_"], p.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (p *_Pdu) CancelOutletSequence(ctx context.Context) error {
	_, err := p.Caller().Call(ctx, p.RID(), "cancelOutletSequence", nil)
	return err
}
