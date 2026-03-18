// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package transferswitch

import (
	"github.com/arminguenther/xeruspower-go/v40033/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40033/idl"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40033/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40033/sensors/numericsensor"
	"github.com/arminguenther/xeruspower-go/v40033/sensors/statesensor"
)

func (m *MetaData) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["label"] = m.Label
	j0["namePlate"] = m.NamePlate.Encode()
	j0["rating"] = m.Rating.Encode()
	j0["type"] = m.Type
	j0["sourceCount"] = m.SourceCount
	return j0
}

func (m *MetaData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("label", j0)
	if err != nil {
		return err
	}
	m.Label, err = encoding.Is[string](j0["label"])
	if err != nil {
		return err
	}
	err = encoding.In("namePlate", j0)
	if err != nil {
		return err
	}
	err = m.NamePlate.Decode(j0["namePlate"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("rating", j0)
	if err != nil {
		return err
	}
	err = m.Rating.Decode(j0["rating"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	m.Type, err = encoding.AsEnum[Type](j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("sourceCount", j0)
	if err != nil {
		return err
	}
	m.SourceCount, err = encoding.AsInt32(j0["sourceCount"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Sensors) Encode() map[string]any {
	j0 := make(map[string]any, 7)
	j0["selectedSource"] = object.ToMap(s.SelectedSource)
	j0["operationalState"] = object.ToMap(s.OperationalState)
	j0["sourceVoltagePhaseSyncAngle"] = object.ToMap(s.SourceVoltagePhaseSyncAngle)
	j0["overloadAlarm"] = object.ToMap(s.OverloadAlarm)
	j0["phaseSyncAlarm"] = object.ToMap(s.PhaseSyncAlarm)
	j0["switchFault"] = object.ToMap(s.SwitchFault)
	j0["selectedBypassSource"] = object.ToMap(s.SelectedBypassSource)
	return j0
}

func (s *Sensors) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("selectedSource", j0)
	if err != nil {
		return err
	}
	s.SelectedSource, err = object.As[statesensor.StateSensor](j0["selectedSource"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("operationalState", j0)
	if err != nil {
		return err
	}
	s.OperationalState, err = object.As[statesensor.StateSensor](j0["operationalState"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("sourceVoltagePhaseSyncAngle", j0)
	if err != nil {
		return err
	}
	s.SourceVoltagePhaseSyncAngle, err = object.As[numericsensor.NumericSensor](j0["sourceVoltagePhaseSyncAngle"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("overloadAlarm", j0)
	if err != nil {
		return err
	}
	s.OverloadAlarm, err = object.As[statesensor.StateSensor](j0["overloadAlarm"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("phaseSyncAlarm", j0)
	if err != nil {
		return err
	}
	s.PhaseSyncAlarm, err = object.As[statesensor.StateSensor](j0["phaseSyncAlarm"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("switchFault", j0)
	if err != nil {
		return err
	}
	s.SwitchFault, err = object.As[statesensor.StateSensor](j0["switchFault"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("selectedBypassSource", j0)
	if err != nil {
		return err
	}
	s.SelectedBypassSource, err = object.As[statesensor.StateSensor](j0["selectedBypassSource"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 7)
	j0["name"] = s.Name
	j0["preferredSource"] = s.PreferredSource
	j0["autoRetransfer"] = s.AutoRetransfer
	j0["noAutoRetransferIfPhaseFault"] = s.NoAutoRetransferIfPhaseFault
	j0["autoRetransferWaitTime"] = s.AutoRetransferWaitTime
	j0["manualTransferEnabled"] = s.ManualTransferEnabled
	j0["phaseSyncSensorEnabled"] = s.PhaseSyncSensorEnabled
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("name", j0)
	if err != nil {
		return err
	}
	s.Name, err = encoding.Is[string](j0["name"])
	if err != nil {
		return err
	}
	err = encoding.In("preferredSource", j0)
	if err != nil {
		return err
	}
	s.PreferredSource, err = encoding.AsInt32(j0["preferredSource"])
	if err != nil {
		return err
	}
	err = encoding.In("autoRetransfer", j0)
	if err != nil {
		return err
	}
	s.AutoRetransfer, err = encoding.Is[bool](j0["autoRetransfer"])
	if err != nil {
		return err
	}
	err = encoding.In("noAutoRetransferIfPhaseFault", j0)
	if err != nil {
		return err
	}
	s.NoAutoRetransferIfPhaseFault, err = encoding.Is[bool](j0["noAutoRetransferIfPhaseFault"])
	if err != nil {
		return err
	}
	err = encoding.In("autoRetransferWaitTime", j0)
	if err != nil {
		return err
	}
	s.AutoRetransferWaitTime, err = encoding.AsInt32(j0["autoRetransferWaitTime"])
	if err != nil {
		return err
	}
	err = encoding.In("manualTransferEnabled", j0)
	if err != nil {
		return err
	}
	s.ManualTransferEnabled, err = encoding.Is[bool](j0["manualTransferEnabled"])
	if err != nil {
		return err
	}
	err = encoding.In("phaseSyncSensorEnabled", j0)
	if err != nil {
		return err
	}
	s.PhaseSyncSensorEnabled, err = encoding.Is[bool](j0["phaseSyncSensorEnabled"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Statistics) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["transferCount"] = s.TransferCount
	j0["powerFailDetectTime"] = s.PowerFailDetectTime
	j0["relayOpenTime"] = s.RelayOpenTime
	j0["totalTransferTime"] = s.TotalTransferTime
	return j0
}

func (s *Statistics) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("transferCount", j0)
	if err != nil {
		return err
	}
	s.TransferCount, err = encoding.AsInt32(j0["transferCount"])
	if err != nil {
		return err
	}
	err = encoding.In("powerFailDetectTime", j0)
	if err != nil {
		return err
	}
	s.PowerFailDetectTime, err = encoding.AsInt32(j0["powerFailDetectTime"])
	if err != nil {
		return err
	}
	err = encoding.In("relayOpenTime", j0)
	if err != nil {
		return err
	}
	s.RelayOpenTime, err = encoding.AsInt32(j0["relayOpenTime"])
	if err != nil {
		return err
	}
	err = encoding.In("totalTransferTime", j0)
	if err != nil {
		return err
	}
	s.TotalTransferTime, err = encoding.AsInt32(j0["totalTransferTime"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_SettingsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.UserEvent = valobj.For[userevent.UserEvent]()
	err := s.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldSettings", value)
	if err != nil {
		return err
	}
	err = s.oldSettings.Decode(value["oldSettings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newSettings", value)
	if err != nil {
		return err
	}
	err = s.newSettings.Decode(value["newSettings"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (t *TransferLogEntry) Encode() map[string]any {
	j0 := make(map[string]any, 7)
	j0["timestamp"] = t.Timestamp.Unix()
	j0["oldInlet"] = t.OldInlet
	j0["newInlet"] = t.NewInlet
	j0["reason"] = t.Reason
	j0["waveform"] = t.Waveform.Encode()
	j0["statistics"] = t.Statistics.Encode()
	j0["switchFault"] = t.SwitchFault
	return j0
}

func (t *TransferLogEntry) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("timestamp", j0)
	if err != nil {
		return err
	}
	t.Timestamp, err = encoding.AsTime(j0["timestamp"])
	if err != nil {
		return err
	}
	err = encoding.In("oldInlet", j0)
	if err != nil {
		return err
	}
	t.OldInlet, err = encoding.AsInt32(j0["oldInlet"])
	if err != nil {
		return err
	}
	err = encoding.In("newInlet", j0)
	if err != nil {
		return err
	}
	t.NewInlet, err = encoding.AsInt32(j0["newInlet"])
	if err != nil {
		return err
	}
	err = encoding.In("reason", j0)
	if err != nil {
		return err
	}
	t.Reason, err = encoding.AsEnum[TransferReason](j0["reason"])
	if err != nil {
		return err
	}
	err = encoding.In("waveform", j0)
	if err != nil {
		return err
	}
	err = t.Waveform.Decode(j0["waveform"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("statistics", j0)
	if err != nil {
		return err
	}
	err = t.Statistics.Decode(j0["statistics"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("switchFault", j0)
	if err != nil {
		return err
	}
	t.SwitchFault, err = encoding.AsInt32(j0["switchFault"])
	if err != nil {
		return err
	}
	return nil
}
