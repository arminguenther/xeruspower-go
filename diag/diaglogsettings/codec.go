// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package diaglogsettings

import (
	"github.com/arminguenther/xeruspower-go/v40000/idl"
	"github.com/arminguenther/xeruspower-go/v40000/internal/encoding"
)

func (l *LogLevelEntry) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["ctxName"] = l.CtxName
	j0["logLevel"] = l.LogLevel
	return j0
}

func (l *LogLevelEntry) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("ctxName", j0)
	if err != nil {
		return err
	}
	l.CtxName, err = encoding.Is[string](j0["ctxName"])
	if err != nil {
		return err
	}
	err = encoding.In("logLevel", j0)
	if err != nil {
		return err
	}
	l.LogLevel, err = encoding.AsEnum[LogLevel](j0["logLevel"])
	if err != nil {
		return err
	}
	return nil
}
