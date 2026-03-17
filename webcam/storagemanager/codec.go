// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package storagemanager

import (
	"github.com/arminguenther/xeruspower-go/v40020/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40020/idl"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40020/internal/encoding/valobj"
	"github.com/arminguenther/xeruspower-go/v40020/webcam/webcam"
)

func (w *WebcamStorageInfo) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["webcam"] = object.ToMap(w.Webcam)
	j0["newestIndex"] = w.NewestIndex
	j0["oldestIndex"] = w.OldestIndex
	j0["count"] = w.Count
	return j0
}

func (w *WebcamStorageInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("webcam", j0)
	if err != nil {
		return err
	}
	w.Webcam, err = object.As[webcam.Webcam](j0["webcam"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newestIndex", j0)
	if err != nil {
		return err
	}
	w.NewestIndex, err = encoding.AsInt64(j0["newestIndex"])
	if err != nil {
		return err
	}
	err = encoding.In("oldestIndex", j0)
	if err != nil {
		return err
	}
	w.OldestIndex, err = encoding.AsInt64(j0["oldestIndex"])
	if err != nil {
		return err
	}
	err = encoding.In("count", j0)
	if err != nil {
		return err
	}
	w.Count, err = encoding.AsInt32(j0["count"])
	if err != nil {
		return err
	}
	return nil
}

func (s *StorageInformation) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["status"] = s.Status
	j0["capacity"] = s.Capacity
	j0["used"] = s.Used
	s1 := make([]any, 0, len(s.WebcamStorageInfo))
	for _, e1 := range s.WebcamStorageInfo {
		s1 = append(s1, e1.Encode())
	}
	j0["webcamStorageInfo"] = s1
	return j0
}

func (s *StorageInformation) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("status", j0)
	if err != nil {
		return err
	}
	s.Status, err = encoding.AsEnum[StorageStatus](j0["status"])
	if err != nil {
		return err
	}
	err = encoding.In("capacity", j0)
	if err != nil {
		return err
	}
	s.Capacity, err = encoding.AsInt32(j0["capacity"])
	if err != nil {
		return err
	}
	err = encoding.In("used", j0)
	if err != nil {
		return err
	}
	s.Used, err = encoding.AsInt32(j0["used"])
	if err != nil {
		return err
	}
	err = encoding.In("webcamStorageInfo", j0)
	if err != nil {
		return err
	}
	var s1 []any
	s1, err = encoding.Is[[]any](j0["webcamStorageInfo"])
	if err != nil {
		return err
	}
	s.WebcamStorageInfo = make([]WebcamStorageInfo, 0, len(s1))
	for _, a1 := range s1 {
		var e1 WebcamStorageInfo
		err = e1.Decode(a1, caller)
		if err != nil {
			return err
		}
		s.WebcamStorageInfo = append(s.WebcamStorageInfo, e1)
	}
	return nil
}

func (s *StorageSettings) Encode() map[string]any {
	j0 := make(map[string]any, 5)
	j0["type"] = s.Type
	j0["capacity"] = s.Capacity
	j0["server"] = s.Server
	j0["username"] = s.Username
	j0["password"] = s.Password
	return j0
}

func (s *StorageSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("type", j0)
	if err != nil {
		return err
	}
	s.Type, err = encoding.AsEnum[StorageType](j0["type"])
	if err != nil {
		return err
	}
	err = encoding.In("capacity", j0)
	if err != nil {
		return err
	}
	s.Capacity, err = encoding.AsInt32(j0["capacity"])
	if err != nil {
		return err
	}
	err = encoding.In("server", j0)
	if err != nil {
		return err
	}
	s.Server, err = encoding.Is[string](j0["server"])
	if err != nil {
		return err
	}
	err = encoding.In("username", j0)
	if err != nil {
		return err
	}
	s.Username, err = encoding.Is[string](j0["username"])
	if err != nil {
		return err
	}
	err = encoding.In("password", j0)
	if err != nil {
		return err
	}
	s.Password, err = encoding.Is[string](j0["password"])
	if err != nil {
		return err
	}
	return nil
}

func (s *StorageMetaData) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["index"] = s.Index
	j0["webcam"] = object.ToMap(s.Webcam)
	return j0
}

func (s *StorageMetaData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("index", j0)
	if err != nil {
		return err
	}
	s.Index, err = encoding.AsInt64(j0["index"])
	if err != nil {
		return err
	}
	err = encoding.In("webcam", j0)
	if err != nil {
		return err
	}
	s.Webcam, err = object.As[webcam.Webcam](j0["webcam"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (i *ImageStorageMetaData) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["imageMeta"] = i.ImageMeta.Encode()
	j0["fileSize"] = i.FileSize
	j0["storageMeta"] = i.StorageMeta.Encode()
	return j0
}

func (i *ImageStorageMetaData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("imageMeta", j0)
	if err != nil {
		return err
	}
	err = i.ImageMeta.Decode(j0["imageMeta"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("fileSize", j0)
	if err != nil {
		return err
	}
	i.FileSize, err = encoding.AsInt32(j0["fileSize"])
	if err != nil {
		return err
	}
	err = encoding.In("storageMeta", j0)
	if err != nil {
		return err
	}
	err = i.StorageMeta.Decode(j0["storageMeta"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *StorageImage) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["image"] = s.Image.Encode()
	j0["metaData"] = s.MetaData.Encode()
	return j0
}

func (s *StorageImage) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("image", j0)
	if err != nil {
		return err
	}
	err = s.Image.Decode(j0["image"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("metaData", j0)
	if err != nil {
		return err
	}
	err = s.MetaData.Decode(j0["metaData"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (a *Activity) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["webcam"] = object.ToMap(a.Webcam)
	j0["interval"] = a.Interval
	j0["count"] = a.Count
	j0["done"] = a.Done
	return j0
}

func (a *Activity) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("webcam", j0)
	if err != nil {
		return err
	}
	a.Webcam, err = object.As[webcam.Webcam](j0["webcam"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("interval", j0)
	if err != nil {
		return err
	}
	a.Interval, err = encoding.AsInt32(j0["interval"])
	if err != nil {
		return err
	}
	err = encoding.In("count", j0)
	if err != nil {
		return err
	}
	a.Count, err = encoding.AsInt32(j0["count"])
	if err != nil {
		return err
	}
	err = encoding.In("done", j0)
	if err != nil {
		return err
	}
	a.Done, err = encoding.AsInt32(j0["done"])
	if err != nil {
		return err
	}
	return nil
}

func (i *_ImageUploadStartedEvent) Decode(value map[string]any, caller idl.Caller) error {
	i.UserEvent = valobj.For[userevent.UserEvent]()
	err := i.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("webcam", value)
	if err != nil {
		return err
	}
	i.webcam, err = object.As[webcam.Webcam](value["webcam"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("folderUrl", value)
	if err != nil {
		return err
	}
	i.folderUrl, err = encoding.Is[string](value["folderUrl"])
	if err != nil {
		return err
	}
	return nil
}
