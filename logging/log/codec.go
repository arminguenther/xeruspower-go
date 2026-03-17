// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package log

import (
	"github.com/arminguenther/xeruspower-go/v40220/idl"
	"github.com/arminguenther/xeruspower-go/v40220/internal/encoding"
)

func (l *Info) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["creationTime"] = l.CreationTime
	j0["idFirst"] = l.IdFirst
	j0["idNext"] = l.IdNext
	return j0
}

func (l *Info) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("creationTime", j0)
	if err != nil {
		return err
	}
	l.CreationTime, err = encoding.AsInt64(j0["creationTime"])
	if err != nil {
		return err
	}
	err = encoding.In("idFirst", j0)
	if err != nil {
		return err
	}
	l.IdFirst, err = encoding.AsInt32(j0["idFirst"])
	if err != nil {
		return err
	}
	err = encoding.In("idNext", j0)
	if err != nil {
		return err
	}
	l.IdNext, err = encoding.AsInt32(j0["idNext"])
	if err != nil {
		return err
	}
	return nil
}

func (l *Entry) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["id"] = l.Id
	j0["timestamp"] = l.Timestamp.Unix()
	j0["eventClass"] = l.EventClass
	j0["message"] = l.Message
	return j0
}

func (l *Entry) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("id", j0)
	if err != nil {
		return err
	}
	l.Id, err = encoding.AsInt32(j0["id"])
	if err != nil {
		return err
	}
	err = encoding.In("timestamp", j0)
	if err != nil {
		return err
	}
	l.Timestamp, err = encoding.AsTime(j0["timestamp"])
	if err != nil {
		return err
	}
	err = encoding.In("eventClass", j0)
	if err != nil {
		return err
	}
	l.EventClass, err = encoding.Is[string](j0["eventClass"])
	if err != nil {
		return err
	}
	err = encoding.In("message", j0)
	if err != nil {
		return err
	}
	l.Message, err = encoding.Is[string](j0["message"])
	if err != nil {
		return err
	}
	return nil
}

func (l *Chunk) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["logCreationTime"] = l.LogCreationTime
	j0["idFirst"] = l.IdFirst
	j0["allEntryCnt"] = l.AllEntryCnt
	s1 := make([]any, 0, len(l.SelEntries))
	for _, e1 := range l.SelEntries {
		s1 = append(s1, e1.Encode())
	}
	j0["selEntries"] = s1
	return j0
}

func (l *Chunk) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("logCreationTime", j0)
	if err != nil {
		return err
	}
	l.LogCreationTime, err = encoding.AsInt64(j0["logCreationTime"])
	if err != nil {
		return err
	}
	err = encoding.In("idFirst", j0)
	if err != nil {
		return err
	}
	l.IdFirst, err = encoding.AsInt32(j0["idFirst"])
	if err != nil {
		return err
	}
	err = encoding.In("allEntryCnt", j0)
	if err != nil {
		return err
	}
	l.AllEntryCnt, err = encoding.AsInt32(j0["allEntryCnt"])
	if err != nil {
		return err
	}
	err = encoding.In("selEntries", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["selEntries"])
	if err != nil {
		return err
	}
	l.SelEntries = make([]Entry, 0, len(s1))
	for _, a1 := range s1 {
		var e1 Entry
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		l.SelEntries = append(l.SelEntries, e1)
	}
	return nil
}
