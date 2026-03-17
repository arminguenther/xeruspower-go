// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package nameservice

import (
	"context"

	"github.com/arminguenther/xeruspower-go/v40100/idl"
	"github.com/arminguenther/xeruspower-go/v40100/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40100/internal/encoding/object"
)

func init() {
	object.Register(NewNameService)
}

type _NameService struct {
	idl.Object
}

// NewNameService returns a new NameService interface for the object at given RID.
func NewNameService(rid string, caller idl.Caller) NameService {
	return &_NameService{idl.NewObject(rid, caller)}
}

func (n *_NameService) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "jsonrpc.NameService",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (n *_NameService) Lookup(ctx context.Context, in0 string) (idl.Object, error) {
	var ret idl.Object
	val, err := n.Caller().Call(ctx, n.RID(), "lookup", map[string]any{
		"name": in0,
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
	ret, err = object.As[idl.Object](res["_ret_"], n.Caller())
	if err != nil {
		return ret, err
	}
	return ret, nil
}
