// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package webcamchannel

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40410/idl"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40410/webcam/webcam"
)

func init() {
	object.Register(NewChannel)
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
		Name:  "webcam.Channel",
		Major: 1, Submajor: 0, Minor: 2,
	}
}

func (c *_Channel) GetClientType(ctx context.Context) (string, error) {
	var ret string
	val, err := c.Caller().Call(ctx, c.RID(), "getClientType", nil)
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
	ret, err = encoding.Is[string](res["_ret_"])
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (c *_Channel) GetWebcam(ctx context.Context) (webcam.Webcam, error) {
	var ret webcam.Webcam
	val, err := c.Caller().Call(ctx, c.RID(), "getWebcam", nil)
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
	ret, err = object.As[webcam.Webcam](res["_ret_"], c.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (c *_Channel) IsAvailable(ctx context.Context) (bool, error) {
	var ret bool
	val, err := c.Caller().Call(ctx, c.RID(), "isAvailable", nil)
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

func (c *_Channel) Release(ctx context.Context) error {
	_, err := c.Caller().Call(ctx, c.RID(), "release", nil)
	return err
}

func (c *_Channel) CaptureImage(ctx context.Context) (int32, webcam.Image, error) {
	var ret int32
	var out0 webcam.Image
	val, err := c.Caller().Call(ctx, c.RID(), "captureImage", nil)
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
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("image", res)
	if err != nil {
		return ret, out0, err
	}
	err = out0.Decode(res["image"], c.Caller())
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}

func (c *_Channel) TriggerCapture(ctx context.Context) (int32, string, error) {
	var ret int32
	var out0 string
	val, err := c.Caller().Call(ctx, c.RID(), "triggerCapture", nil)
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
	ret, err = encoding.AsInt32(res["_ret_"])
	if err != nil {
		return ret, out0, err
	}
	err = encoding.In("captureToken", res)
	if err != nil {
		return ret, out0, err
	}
	out0, err = encoding.Is[string](res["captureToken"])
	if err != nil {
		return ret, out0, err
	}
	return ret, out0, nil
}
