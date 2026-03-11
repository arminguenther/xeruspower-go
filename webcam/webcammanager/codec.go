// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package webcammanager

import (
	"github.com/arminguenther/xeruspower-go/v40410/idl"
	"github.com/arminguenther/xeruspower-go/v40410/idl/event"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40410/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40410/webcam/webcam"
)

func (w *_WebcamEvent) Decode(value map[string]any, caller idl.Caller) error {
	w.Event = valobj.For[event.Event]()
	err := w.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("webcam", value)
	if err != nil {
		return err
	}
	w.webcam, err = object.As[webcam.Webcam](value["webcam"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("information", value)
	if err != nil {
		return err
	}
	err = w.information.Decode(value["information"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("name", value)
	if err != nil {
		return err
	}
	w.name, err = encoding.Is[string](value["name"])
	if err != nil {
		return err
	}
	return nil
}

func (w *_WebcamAddedEvent) Decode(value map[string]any, caller idl.Caller) error {
	w.WebcamEvent = valobj.For[WebcamEvent]()
	err := w.WebcamEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}

func (w *_WebcamRemovedEvent) Decode(value map[string]any, caller idl.Caller) error {
	w.WebcamEvent = valobj.For[WebcamEvent]()
	err := w.WebcamEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	return nil
}
