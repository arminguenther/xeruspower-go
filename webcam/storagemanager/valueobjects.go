// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package storagemanager

import (
	"github.com/arminguenther/xeruspower-go/v40413/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40413/idl"
	"github.com/arminguenther/xeruspower-go/v40413/internal/encoding/valobj"
	webcam_ "github.com/arminguenther/xeruspower-go/v40413/webcam/webcam"
)

func init() {
	valobj.Register(func() ImageUploadStartedEvent { return &_ImageUploadStartedEvent{} })
}

type _ImageUploadStartedEvent struct {
	userevent.UserEvent
	webcam    webcam_.Webcam
	folderUrl string
}

func (i *_ImageUploadStartedEvent) TypeCode() idl.TypeCode {
	return idl.TypeCode{
		Name:  "webcam.StorageManager_1_0_3.ImageUploadStartedEvent",
		Major: 1, Submajor: 0, Minor: 0,
	}
}

func (i *_ImageUploadStartedEvent) Webcam() webcam_.Webcam {
	return i.webcam
}

func (i *_ImageUploadStartedEvent) FolderUrl() string {
	return i.folderUrl
}

func (i *_ImageUploadStartedEvent) isImageUploadStartedEvent() {}
