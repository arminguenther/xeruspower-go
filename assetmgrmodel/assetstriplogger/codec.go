// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package assetstriplogger

import (
	"github.com/arminguenther/xeruspower-go/v40032/assetmgrmodel/assetstrip"
	"github.com/arminguenther/xeruspower-go/v40032/idl"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding"
)

func (i *Info) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["capacity"] = i.Capacity
	j0["oldestRecord"] = i.OldestRecord
	j0["newestRecord"] = i.NewestRecord
	j0["totalEventCount"] = i.TotalEventCount
	return j0
}

func (i *Info) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("capacity", j0)
	if err != nil {
		return err
	}
	i.Capacity, err = encoding.AsInt32(j0["capacity"])
	if err != nil {
		return err
	}
	err = encoding.In("oldestRecord", j0)
	if err != nil {
		return err
	}
	i.OldestRecord, err = encoding.AsInt32(j0["oldestRecord"])
	if err != nil {
		return err
	}
	err = encoding.In("newestRecord", j0)
	if err != nil {
		return err
	}
	i.NewestRecord, err = encoding.AsInt32(j0["newestRecord"])
	if err != nil {
		return err
	}
	err = encoding.In("totalEventCount", j0)
	if err != nil {
		return err
	}
	i.TotalEventCount, err = encoding.AsInt32(j0["totalEventCount"])
	if err != nil {
		return err
	}
	return nil
}

func (r *Record) Encode() map[string]any {
	j0 := make(map[string]any, 9)
	j0["timestamp"] = r.Timestamp.Unix()
	j0["type"] = r.Type
	j0["assetStripNumber"] = r.AssetStripNumber
	j0["rackUnitNumber"] = r.RackUnitNumber
	j0["rackUnitPosition"] = r.RackUnitPosition
	j0["slotNumber"] = r.SlotNumber
	j0["tagId"] = r.TagId
	j0["parentBladeId"] = r.ParentBladeId
	j0["state"] = r.State
	return j0
}

func (r *Record) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("timestamp", j0)
	if err != nil {
		return err
	}
	r.Timestamp, err = encoding.AsTime(j0["timestamp"])
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	r.Type, err = encoding.AsEnum[RecordType](j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("assetStripNumber", j0)
	if err != nil {
		return err
	}
	r.AssetStripNumber, err = encoding.AsInt32(j0["assetStripNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("rackUnitNumber", j0)
	if err != nil {
		return err
	}
	r.RackUnitNumber, err = encoding.AsInt32(j0["rackUnitNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("rackUnitPosition", j0)
	if err != nil {
		return err
	}
	r.RackUnitPosition, err = encoding.AsInt32(j0["rackUnitPosition"])
	if err != nil {
		return err
	}
	err = encoding.In("slotNumber", j0)
	if err != nil {
		return err
	}
	r.SlotNumber, err = encoding.AsInt32(j0["slotNumber"])
	if err != nil {
		return err
	}
	err = encoding.In("tagId", j0)
	if err != nil {
		return err
	}
	r.TagId, err = encoding.Is[string](j0["tagId"])
	if err != nil {
		return err
	}
	err = encoding.In("parentBladeId", j0)
	if err != nil {
		return err
	}
	r.ParentBladeId, err = encoding.Is[string](j0["parentBladeId"])
	if err != nil {
		return err
	}
	err = encoding.In("state", j0)
	if err != nil {
		return err
	}
	r.State, err = encoding.AsEnum[assetstrip.State](j0["state"])
	if err != nil {
		return err
	}
	return nil
}
