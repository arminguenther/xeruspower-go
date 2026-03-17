// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package cascademanager

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding/object"
)

func init() {
	object.Register(NewCascadeManager)
}

type _CascadeManager struct {
	idl.Object
}

// NewCascadeManager returns a new CascadeManager interface for the object at given RID.
func NewCascadeManager(rid string, caller idl.Caller) CascadeManager {
	return &_CascadeManager{idl.NewObject(rid, caller)}
}

func (c *_CascadeManager) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "cascading.CascadeManager",
		Major: 2, Submajor: 0, Minor: 2,
	}
}

func (c *_CascadeManager) GetPrimaryUnitSettings(ctx context.Context) (PrimaryUnitSettings, error) {
	var ret PrimaryUnitSettings
	val, err := c.Caller().Call(ctx, c.RID(), "getPrimaryUnitSettings", nil)
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
	err = ret.Decode(res["_ret_"], c.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (c *_CascadeManager) SetPrimaryUnitSettings(ctx context.Context, in0 PrimaryUnitSettings) (int32, error) {
	var ret int32
	val, err := c.Caller().Call(ctx, c.RID(), "setPrimaryUnitSettings", map[string]any{
		"primaryUnitSettings": in0.Encode(),
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

func (c *_CascadeManager) GetStatus(ctx context.Context) (Status, error) {
	var ret Status
	val, err := c.Caller().Call(ctx, c.RID(), "getStatus", nil)
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
	err = ret.Decode(res["_ret_"], c.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (c *_CascadeManager) GetLinkPortStatus(ctx context.Context) (LinkPortStatus, error) {
	var ret LinkPortStatus
	val, err := c.Caller().Call(ctx, c.RID(), "getLinkPortStatus", nil)
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
	err = ret.Decode(res["_ret_"], c.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func (c *_CascadeManager) AddLinkUnit(ctx context.Context, in0 int32, in1 string, in2 string, in3 string, in4 string) (int32, error) {
	var ret int32
	val, err := c.Caller().Call(ctx, c.RID(), "addLinkUnit", map[string]any{
		"linkId":      in0,
		"host":        in1,
		"login":       in2,
		"password":    in3,
		"newPassword": in4,
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

func (c *_CascadeManager) AddLinkUnit2(ctx context.Context, in0 int32, in1 string, in2 string, in3 string, in4 string, in5 bool) (int32, error) {
	var ret int32
	val, err := c.Caller().Call(ctx, c.RID(), "addLinkUnit2", map[string]any{
		"linkId":                   in0,
		"host":                     in1,
		"login":                    in2,
		"password":                 in3,
		"newPassword":              in4,
		"disableStrongPasswordReq": in5,
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

func (c *_CascadeManager) ReleaseLinkUnit(ctx context.Context, in0 int32) (int32, error) {
	var ret int32
	val, err := c.Caller().Call(ctx, c.RID(), "releaseLinkUnit", map[string]any{
		"linkId": in0,
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

func (c *_CascadeManager) RequestLink(ctx context.Context, in0 string) (int32, error) {
	var ret int32
	val, err := c.Caller().Call(ctx, c.RID(), "requestLink", map[string]any{
		"token": in0,
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

func (c *_CascadeManager) FinalizeLink(ctx context.Context, in0 string) error {
	_, err := c.Caller().Call(ctx, c.RID(), "finalizeLink", map[string]any{
		"token": in0,
	})
	return err
}

func (c *_CascadeManager) Unlink(ctx context.Context) error {
	_, err := c.Caller().Call(ctx, c.RID(), "unlink", nil)
	return err
}

func (c *_CascadeManager) GetSupportedRoles(ctx context.Context) ([]Role, error) {
	var ret []Role
	val, err := c.Caller().Call(ctx, c.RID(), "getSupportedRoles", nil)
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
	ret = make([]Role, 0, len(s0))
	for _, a0 := range s0 {
		var e0 Role
		e0, err = encoding.AsEnum[Role](a0)
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (c *_CascadeManager) GetSupportedLinkUnitTypes(ctx context.Context) ([]LinkUnitType, error) {
	var ret []LinkUnitType
	val, err := c.Caller().Call(ctx, c.RID(), "getSupportedLinkUnitTypes", nil)
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
	ret = make([]LinkUnitType, 0, len(s0))
	for _, a0 := range s0 {
		var e0 LinkUnitType
		e0, err = encoding.AsEnum[LinkUnitType](a0)
		if err != nil {
			return ret, err
		}
		ret = append(ret, e0)
	}
	return ret, nil
}

func (c *_CascadeManager) AddCascadeLinkUnit(ctx context.Context, in0 int32, in1 int32, in2 string, in3 string, in4 bool) (int32, error) {
	var ret int32
	val, err := c.Caller().Call(ctx, c.RID(), "addCascadeLinkUnit", map[string]any{
		"linkId":            in0,
		"nodeIndex":         in1,
		"login":             in2,
		"password":          in3,
		"positionDependent": in4,
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

func (c *_CascadeManager) AddLinkPortLinkUnit(ctx context.Context) (int32, error) {
	var ret int32
	val, err := c.Caller().Call(ctx, c.RID(), "addLinkPortLinkUnit", nil)
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

func (c *_CascadeManager) AddSecureSerialLinkUnit(ctx context.Context, in0 int32, in1 string) (int32, error) {
	var ret int32
	val, err := c.Caller().Call(ctx, c.RID(), "addSecureSerialLinkUnit", map[string]any{
		"linkId":     in0,
		"installKey": in1,
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
