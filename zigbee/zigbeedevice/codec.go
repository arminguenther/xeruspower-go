// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package zigbeedevice

import (
	"github.com/arminguenther/xeruspower-go/v40410/idl"
	"github.com/arminguenther/xeruspower-go/v40410/idl/event"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding/valobj"
)

func (m *MetaData) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["sourceId"] = m.SourceId
	j0["clusters"] = encoding.NonNilSlice(m.Clusters)
	j0["preferredSlot"] = m.PreferredSlot
	return j0
}

func (m *MetaData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("sourceId", j0)
	if err != nil {
		return err
	}
	m.SourceId, err = encoding.AsInt32(j0["sourceId"])
	if err != nil {
		return err
	}
	err = encoding.In("clusters", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["clusters"])
	if err != nil {
		return err
	}
	m.Clusters = make([]int32, 0, len(s1))
	for _, a1 := range s1 {
		var e1 int32
		e1, err = encoding.AsInt32(a1)
		if err != nil {
			return err
		}
		m.Clusters = append(m.Clusters, e1)
	}
	err = encoding.In("preferredSlot", j0)
	if err != nil {
		return err
	}
	m.PreferredSlot, err = encoding.AsInt32(j0["preferredSlot"])
	if err != nil {
		return err
	}
	return nil
}

func (c *ClusterValue) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["id"] = c.Id
	j0["endpoint"] = c.Endpoint
	j0["attribute"] = c.Attribute
	j0["timestamp"] = c.Timestamp.Unix()
	j0["value"] = c.Value
	return j0
}

func (c *ClusterValue) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("id", j0)
	if err != nil {
		return err
	}
	c.Id, err = encoding.AsInt32(j0["id"])
	if err != nil {
		return err
	}
	err = encoding.In("endpoint", j0)
	if err != nil {
		return err
	}
	c.Endpoint, err = encoding.AsInt32(j0["endpoint"])
	if err != nil {
		return err
	}
	err = encoding.In("attribute", j0)
	if err != nil {
		return err
	}
	c.Attribute, err = encoding.AsInt32(j0["attribute"])
	if err != nil {
		return err
	}
	err = encoding.In("timestamp", j0)
	if err != nil {
		return err
	}
	c.Timestamp, err = encoding.AsTime(j0["timestamp"])
	if err != nil {
		return err
	}
	err = encoding.In("value", j0)
	if err != nil {
		return err
	}
	c.Value, err = encoding.Is[string](j0["value"])
	if err != nil {
		return err
	}
	return nil
}

func (c *_ClusterDataEvent) Decode(value map[string]any, caller idl.Caller) error {
	c.Event = valobj.For[event.Event]()
	err := c.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("value", value)
	if err != nil {
		return err
	}
	err = c.value.Decode(value["value"], caller)
	if err != nil {
		return err
	}
	return nil
}
