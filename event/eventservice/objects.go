// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package eventservice

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40032/idl"
	"github.com/arminguenther/xeruspower-go/v40032/idl/event"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40032/internal/encoding/valobj"
)

func init() {
	object.Register(NewChannel)
	object.Register(NewConsumer)
	object.Register(NewService)
}

type _Consumer struct {
	idl.Object
}

// NewConsumer returns a new Consumer interface for the object at given RID.
func NewConsumer(rid string, caller idl.Caller) Consumer {
	return &_Consumer{idl.NewObject(rid, caller)}
}

func (c *_Consumer) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "event.Consumer",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (c *_Consumer) PushEvents(ctx context.Context, in0 []event.Event) error {
	s0 := make([]any, 0, len(in0))
	for _, e0 := range in0 {
		s0 = append(s0, valobj.ToMap(e0))
	}
	_, err := c.Caller().Call(ctx, c.RID(), "pushEvents", map[string]any{
		"events": s0,
	})
	return err
}

type _Channel struct {
	idl.Object
}

// NewChannel returns a new Channel interface for the object at given RID.
func NewChannel(rid string, caller idl.Caller) Channel {
	return &_Channel{idl.NewObject(rid, caller)}
}

func (c *_Channel) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "event.Channel",
		Major: 1, Submajor: 0, Minor: 1,
	}
}

func (c *_Channel) DemandEventType(ctx context.Context, in0 idl.TypeCode) error {
	_, err := c.Caller().Call(ctx, c.RID(), "demandEventType", map[string]any{
		"type": in0.String(),
	})
	return err
}

func (c *_Channel) CancelEventType(ctx context.Context, in0 idl.TypeCode) error {
	_, err := c.Caller().Call(ctx, c.RID(), "cancelEventType", map[string]any{
		"type": in0.String(),
	})
	return err
}

func (c *_Channel) DemandEventTypes(ctx context.Context, in0 []idl.TypeCode) error {
	s0 := make([]any, 0, len(in0))
	for _, e0 := range in0 {
		s0 = append(s0, e0.String())
	}
	_, err := c.Caller().Call(ctx, c.RID(), "demandEventTypes", map[string]any{
		"types": s0,
	})
	return err
}

func (c *_Channel) CancelEventTypes(ctx context.Context, in0 []idl.TypeCode) error {
	s0 := make([]any, 0, len(in0))
	for _, e0 := range in0 {
		s0 = append(s0, e0.String())
	}
	_, err := c.Caller().Call(ctx, c.RID(), "cancelEventTypes", map[string]any{
		"types": s0,
	})
	return err
}

func (c *_Channel) DemandEvent(ctx context.Context, in0 idl.TypeCode, in1 idl.Object) error {
	_, err := c.Caller().Call(ctx, c.RID(), "demandEvent", map[string]any{
		"type": in0.String(),
		"src":  object.ToMap(in1),
	})
	return err
}

func (c *_Channel) CancelEvent(ctx context.Context, in0 idl.TypeCode, in1 idl.Object) error {
	_, err := c.Caller().Call(ctx, c.RID(), "cancelEvent", map[string]any{
		"type": in0.String(),
		"src":  object.ToMap(in1),
	})
	return err
}

func (c *_Channel) DemandEvents(ctx context.Context, in0 []ChannelEventSelect) error {
	s0 := make([]any, 0, len(in0))
	for _, e0 := range in0 {
		s0 = append(s0, e0.Encode())
	}
	_, err := c.Caller().Call(ctx, c.RID(), "demandEvents", map[string]any{
		"events": s0,
	})
	return err
}

func (c *_Channel) CancelEvents(ctx context.Context, in0 []ChannelEventSelect) error {
	s0 := make([]any, 0, len(in0))
	for _, e0 := range in0 {
		s0 = append(s0, e0.Encode())
	}
	_, err := c.Caller().Call(ctx, c.RID(), "cancelEvents", map[string]any{
		"events": s0,
	})
	return err
}

func (c *_Channel) Subscribe(ctx context.Context, in0 Consumer) error {
	_, err := c.Caller().Call(ctx, c.RID(), "subscribe", map[string]any{
		"consumer": object.ToMap(in0),
	})
	return err
}

func (c *_Channel) Unsubscribe(ctx context.Context, in0 Consumer) (int32, error) {
	var ret int32
	val, err := c.Caller().Call(ctx, c.RID(), "unsubscribe", map[string]any{
		"consumer": object.ToMap(in0),
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

func (c *_Channel) PollEvents(ctx context.Context) (bool, []event.Event, error) {
	var ret bool
	var out0 []event.Event
	val, err := c.Caller().Call(ctx, c.RID(), "pollEvents", nil)
	if err != nil {
		return ret, out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, out0, err
	}
	ret, err = encoding.Is[bool](res["_ret_"])
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("events", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["events"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]event.Event, 0, len(s0))
	for _, a0 := range s0 {
		var e0 event.Event
		e0, err = valobj.As[event.Event](a0, c.Caller())
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}

func (c *_Channel) PollEventsNb(ctx context.Context) (bool, []event.Event, error) {
	var ret bool
	var out0 []event.Event
	val, err := c.Caller().Call(ctx, c.RID(), "pollEventsNb", nil)
	if err != nil {
		return ret, out0, err
	}
	res, err := encoding.Is[map[string]any](val)
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("_ret_", res)
	if err != nil {
		return ret, out0, err
	}
	ret, err = encoding.Is[bool](res["_ret_"])
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("events", res)
	if err != nil {
		return ret, out0, err
	}
	var s0 []any
	s0, err = encoding.Is[[]any](res["events"])
	if err != nil {
		return ret, out0, err
	}
	out0 = make([]event.Event, 0, len(s0))
	for _, a0 := range s0 {
		var e0 event.Event
		e0, err = valobj.As[event.Event](a0, c.Caller())
		if err != nil {
			return ret, out0, err
		}
		out0 = append(out0, e0)
	}
	return ret, out0, nil
}

type _Service struct {
	idl.Object
}

// NewService returns a new Service interface for the object at given RID.
func NewService(rid string, caller idl.Caller) Service {
	return &_Service{idl.NewObject(rid, caller)}
}

func (s *_Service) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "event.Service",
		Major: 1, Submajor: 0, Minor: 1,
	}
}

func (s *_Service) CreateChannel(ctx context.Context) (Channel, error) {
	var ret Channel
	val, err := s.Caller().Call(ctx, s.RID(), "createChannel", nil)
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
	ret, err = object.As[Channel](res["_ret_"], s.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (s *_Service) DestroyChannel(ctx context.Context, in0 Channel) (int32, error) {
	var ret int32
	val, err := s.Caller().Call(ctx, s.RID(), "destroyChannel", map[string]any{
		"channel": object.ToMap(in0),
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

func (s *_Service) PushEvent(ctx context.Context, in0 event.Event) error {
	_, err := s.Caller().Call(ctx, s.RID(), "pushEvent", map[string]any{
		"event": valobj.ToMap(in0),
	})
	return err
}

func (s *_Service) PushEvents(ctx context.Context, in0 []event.Event) error {
	s0 := make([]any, 0, len(in0))
	for _, e0 := range in0 {
		s0 = append(s0, valobj.ToMap(e0))
	}
	_, err := s.Caller().Call(ctx, s.RID(), "pushEvents", map[string]any{
		"events": s0,
	})
	return err
}
