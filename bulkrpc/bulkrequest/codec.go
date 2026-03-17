// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package bulkrequest

import (
	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding"
)

func (r *Request) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["rid"] = r.Rid
	j0["json"] = r.Json
	return j0
}

func (r *Request) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("rid", j0)
	if err != nil {
		return err
	}
	r.Rid, err = encoding.Is[string](j0["rid"])
	if err != nil {
		return err
	}
	err = encoding.In("json", j0)
	if err != nil {
		return err
	}
	r.Json, err = encoding.Is[string](j0["json"])
	if err != nil {
		return err
	}
	return nil
}

func (r *Response) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["json"] = r.Json
	j0["statcode"] = r.Statcode
	return j0
}

func (r *Response) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("json", j0)
	if err != nil {
		return err
	}
	r.Json, err = encoding.Is[string](j0["json"])
	if err != nil {
		return err
	}
	err = encoding.In("statcode", j0)
	if err != nil {
		return err
	}
	r.Statcode, err = encoding.AsInt32(j0["statcode"])
	if err != nil {
		return err
	}
	return nil
}

func (r *BulkRequestRequest) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["rid"] = r.Rid
	j0["json"] = r.Json.Encode()
	return j0
}

func (r *BulkRequestRequest) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("rid", j0)
	if err != nil {
		return err
	}
	r.Rid, err = encoding.Is[string](j0["rid"])
	if err != nil {
		return err
	}
	err = encoding.In("json", j0)
	if err != nil {
		return err
	}
	err = r.Json.Decode(j0["json"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (r *BulkRequestResponse) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["json"] = r.Json.Encode()
	j0["statcode"] = r.Statcode
	return j0
}

func (r *BulkRequestResponse) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("json", j0)
	if err != nil {
		return err
	}
	err = r.Json.Decode(j0["json"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("statcode", j0)
	if err != nil {
		return err
	}
	r.Statcode, err = encoding.AsInt32(j0["statcode"])
	if err != nil {
		return err
	}
	return nil
}
