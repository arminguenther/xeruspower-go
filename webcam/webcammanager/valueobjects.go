// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package webcammanager

import (
	"github.com/arminguenther/xeruspower-go/v40100/idl"
	"github.com/arminguenther/xeruspower-go/v40100/idl/event"
	"github.com/arminguenther/xeruspower-go/v40100/internal/encoding/valobj"
	webcam_ "github.com/arminguenther/xeruspower-go/v40100/webcam/webcam"
)

func init() {
	valobj.Register(func() WebcamAddedEvent { return &_WebcamAddedEvent{} })
	valobj.Register(func() WebcamEvent { return &_WebcamEvent{} })
	valobj.Register(func() WebcamRemovedEvent { return &_WebcamRemovedEvent{} })
}

type _WebcamEvent struct {
	event.Event
	webcam      webcam_.Webcam
	information webcam_.Information
	name        string
}

func (w *_WebcamEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "webcam.WebcamEvent",
		Major: 1, Submajor: 0, Minor: 1,
	}
}

func (w *_WebcamEvent) Webcam() webcam_.Webcam {
	return w.webcam
}

func (w *_WebcamEvent) Information() webcam_.Information {
	return w.information
}

func (w *_WebcamEvent) Name() string {
	return w.name
}

func (w *_WebcamEvent) isWebcamEvent() {}

type _WebcamAddedEvent struct {
	WebcamEvent
}

func (w *_WebcamAddedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "webcam.WebcamAddedEvent",
		Major: 1, Submajor: 0, Minor: 1,
	}
}

func (w *_WebcamAddedEvent) isWebcamAddedEvent() {}

type _WebcamRemovedEvent struct {
	WebcamEvent
}

func (w *_WebcamRemovedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "webcam.WebcamRemovedEvent",
		Major: 1, Submajor: 0, Minor: 1,
	}
}

func (w *_WebcamRemovedEvent) isWebcamRemovedEvent() {}
