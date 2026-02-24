// Copyright 2026 Raritan Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package firmwareupdatestatus

import (
	"github.com/arminguenther/xeruspower-go/idl"
	"github.com/arminguenther/xeruspower-go/internal/encoding"
)

func (u *UpdateStatus) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["state"] = u.State
	j0["elapsed"] = u.Elapsed
	j0["estimated"] = u.Estimated
	j0["error_message"] = u.Error_message
	return j0
}

func (u *UpdateStatus) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("state", j0)
	if err != nil {
		return err
	}
	u.State, err = encoding.Is[string](j0["state"])
	if err != nil {
		return err
	}
	err = encoding.In("elapsed", j0)
	if err != nil {
		return err
	}
	u.Elapsed, err = encoding.AsInt32(j0["elapsed"])
	if err != nil {
		return err
	}
	err = encoding.In("estimated", j0)
	if err != nil {
		return err
	}
	u.Estimated, err = encoding.AsInt32(j0["estimated"])
	if err != nil {
		return err
	}
	err = encoding.In("error_message", j0)
	if err != nil {
		return err
	}
	u.Error_message, err = encoding.Is[string](j0["error_message"])
	if err != nil {
		return err
	}
	return nil
}
