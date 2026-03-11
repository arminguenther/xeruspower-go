// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package zigbeedevice

import (
	"github.com/arminguenther/xeruspower-go/v40510/idl"
	"github.com/arminguenther/xeruspower-go/v40510/idl/event"
	"github.com/arminguenther/xeruspower-go/v40510/internal/encoding/valobj"
)

func init() {
	valobj.Register(func() ClusterDataEvent { return &_ClusterDataEvent{} })
}

type _ClusterDataEvent struct {
	event.Event
	value ClusterValue
}

func (c *_ClusterDataEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "zigbee.ZigbeeDevice.ClusterDataEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_ClusterDataEvent) Value() ClusterValue {
	return c.value
}

func (c *_ClusterDataEvent) isClusterDataEvent() {}
