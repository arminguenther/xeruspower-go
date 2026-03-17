// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package firmware

import (
	"github.com/arminguenther/xeruspower-go/v40211/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/idl/event"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding/valobj"
)

func (u *UpdateHistoryEntry) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["timestamp"] = u.Timestamp.Unix()
	j0["oldVersion"] = u.OldVersion
	j0["imageVersion"] = u.ImageVersion
	j0["imageMD5"] = u.ImageMD5
	j0["status"] = u.Status
	return j0
}

func (u *UpdateHistoryEntry) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("timestamp", j0)
	if err != nil {
		return err
	}
	u.Timestamp, err = encoding.AsTime(j0["timestamp"])
	if err != nil {
		return err
	}
	err = encoding.In("oldVersion", j0)
	if err != nil {
		return err
	}
	u.OldVersion, err = encoding.Is[string](j0["oldVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("imageVersion", j0)
	if err != nil {
		return err
	}
	u.ImageVersion, err = encoding.Is[string](j0["imageVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("imageMD5", j0)
	if err != nil {
		return err
	}
	u.ImageMD5, err = encoding.Is[string](j0["imageMD5"])
	if err != nil {
		return err
	}
	err = encoding.In("status", j0)
	if err != nil {
		return err
	}
	u.Status, err = encoding.AsEnum[UpdateHistoryStatus](j0["status"])
	if err != nil {
		return err
	}
	return nil
}

func (i *ImageStatus) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["state"] = i.State
	j0["error_message"] = i.Error_message
	j0["time_started"] = i.Time_started.Unix()
	j0["size_total"] = i.Size_total
	j0["size_done"] = i.Size_done
	return j0
}

func (i *ImageStatus) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("state", j0)
	if err != nil {
		return err
	}
	i.State, err = encoding.AsEnum[ImageState](j0["state"])
	if err != nil {
		return err
	}
	err = encoding.In("error_message", j0)
	if err != nil {
		return err
	}
	i.Error_message, err = encoding.Is[string](j0["error_message"])
	if err != nil {
		return err
	}
	err = encoding.In("time_started", j0)
	if err != nil {
		return err
	}
	i.Time_started, err = encoding.AsTime(j0["time_started"])
	if err != nil {
		return err
	}
	err = encoding.In("size_total", j0)
	if err != nil {
		return err
	}
	i.Size_total, err = encoding.AsInt32(j0["size_total"])
	if err != nil {
		return err
	}
	err = encoding.In("size_done", j0)
	if err != nil {
		return err
	}
	i.Size_done, err = encoding.AsInt32(j0["size_done"])
	if err != nil {
		return err
	}
	return nil
}

func (i *ImageInfo) Encode() map[string]any {
	j0 := make(map[string]any, 17)
	j0["valid"] = i.Valid
	j0["version"] = i.Version
	j0["min_required_version"] = i.Min_required_version
	j0["min_downgrade_version"] = i.Min_downgrade_version
	j0["product"] = i.Product
	j0["platform"] = i.Platform
	j0["oem"] = i.Oem
	j0["hwid_whitelist"] = i.Hwid_whitelist
	j0["hwid_blacklist"] = i.Hwid_blacklist
	j0["compatible"] = i.Compatible
	j0["signature_present"] = i.Signature_present
	j0["signed_by"] = i.Signed_by
	j0["signature_good"] = i.Signature_good
	j0["certified_by"] = i.Certified_by
	j0["certificate_good"] = i.Certificate_good
	j0["model_list_present"] = i.Model_list_present
	j0["model_supported"] = i.Model_supported
	return j0
}

func (i *ImageInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("valid", j0)
	if err != nil {
		return err
	}
	i.Valid, err = encoding.Is[bool](j0["valid"])
	if err != nil {
		return err
	}
	err = encoding.In("version", j0)
	if err != nil {
		return err
	}
	i.Version, err = encoding.Is[string](j0["version"])
	if err != nil {
		return err
	}
	err = encoding.In("min_required_version", j0)
	if err != nil {
		return err
	}
	i.Min_required_version, err = encoding.Is[string](j0["min_required_version"])
	if err != nil {
		return err
	}
	err = encoding.In("min_downgrade_version", j0)
	if err != nil {
		return err
	}
	i.Min_downgrade_version, err = encoding.Is[string](j0["min_downgrade_version"])
	if err != nil {
		return err
	}
	err = encoding.In("product", j0)
	if err != nil {
		return err
	}
	i.Product, err = encoding.Is[string](j0["product"])
	if err != nil {
		return err
	}
	err = encoding.In("platform", j0)
	if err != nil {
		return err
	}
	i.Platform, err = encoding.Is[string](j0["platform"])
	if err != nil {
		return err
	}
	err = encoding.In("oem", j0)
	if err != nil {
		return err
	}
	i.Oem, err = encoding.Is[string](j0["oem"])
	if err != nil {
		return err
	}
	err = encoding.In("hwid_whitelist", j0)
	if err != nil {
		return err
	}
	i.Hwid_whitelist, err = encoding.Is[string](j0["hwid_whitelist"])
	if err != nil {
		return err
	}
	err = encoding.In("hwid_blacklist", j0)
	if err != nil {
		return err
	}
	i.Hwid_blacklist, err = encoding.Is[string](j0["hwid_blacklist"])
	if err != nil {
		return err
	}
	err = encoding.In("compatible", j0)
	if err != nil {
		return err
	}
	i.Compatible, err = encoding.Is[bool](j0["compatible"])
	if err != nil {
		return err
	}
	err = encoding.In("signature_present", j0)
	if err != nil {
		return err
	}
	i.Signature_present, err = encoding.Is[bool](j0["signature_present"])
	if err != nil {
		return err
	}
	err = encoding.In("signed_by", j0)
	if err != nil {
		return err
	}
	i.Signed_by, err = encoding.Is[string](j0["signed_by"])
	if err != nil {
		return err
	}
	err = encoding.In("signature_good", j0)
	if err != nil {
		return err
	}
	i.Signature_good, err = encoding.Is[bool](j0["signature_good"])
	if err != nil {
		return err
	}
	err = encoding.In("certified_by", j0)
	if err != nil {
		return err
	}
	i.Certified_by, err = encoding.Is[string](j0["certified_by"])
	if err != nil {
		return err
	}
	err = encoding.In("certificate_good", j0)
	if err != nil {
		return err
	}
	i.Certificate_good, err = encoding.Is[bool](j0["certificate_good"])
	if err != nil {
		return err
	}
	err = encoding.In("model_list_present", j0)
	if err != nil {
		return err
	}
	i.Model_list_present, err = encoding.Is[bool](j0["model_list_present"])
	if err != nil {
		return err
	}
	err = encoding.In("model_supported", j0)
	if err != nil {
		return err
	}
	i.Model_supported, err = encoding.Is[bool](j0["model_supported"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_SystemStartupEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.Event = valobj.For[event.Event]()
	err := s.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *_SystemShutdownEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.UserEvent = valobj.For[userevent.UserEvent]()
	err := s.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (f *_ValidationFailedEvent) Decode(value map[string]any, caller idl.Caller) error {
	f.UserEvent = valobj.For[userevent.UserEvent]()
	err := f.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (f *_UpdateEvent) Decode(value map[string]any, caller idl.Caller) error {
	f.UserEvent = valobj.For[userevent.UserEvent]()
	err := f.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("oldVersion", value)
	if err != nil {
		return err
	}
	f.oldVersion, err = encoding.Is[string](value["oldVersion"])
	if err != nil {
		return err
	}
	err = encoding.In("newVersion", value)
	if err != nil {
		return err
	}
	f.newVersion, err = encoding.Is[string](value["newVersion"])
	if err != nil {
		return err
	}
	return nil
}

func (f *_UpdateStartedEvent) Decode(value map[string]any, caller idl.Caller) error {
	f.UpdateEvent = valobj.For[UpdateEvent]()
	err := f.UpdateEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (f *_UpdateCompletedEvent) Decode(value map[string]any, caller idl.Caller) error {
	f.UpdateEvent = valobj.For[UpdateEvent]()
	err := f.UpdateEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (f *_UpdateFailedEvent) Decode(value map[string]any, caller idl.Caller) error {
	f.UpdateEvent = valobj.For[UpdateEvent]()
	err := f.UpdateEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}
