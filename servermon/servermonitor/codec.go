// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package servermonitor

import (
	"github.com/arminguenther/xeruspower-go/v40211/event/userevent"
	"github.com/arminguenther/xeruspower-go/v40211/idl"
	"github.com/arminguenther/xeruspower-go/v40211/idl/event"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding/object"
	"github.com/arminguenther/xeruspower-go/v40211/internal/encoding/valobj"
)

func (s *ServerPowerSettings) Encode() map[string]any {
	j0 := make(map[string]any, 9)
	j0["enabled"] = s.Enabled
	j0["target"] = object.ToMap(s.Target)
	j0["powerCheck"] = s.PowerCheck
	j0["powerThreshold"] = s.PowerThreshold
	j0["timeout"] = s.Timeout
	j0["shutdownCmd"] = s.ShutdownCmd
	j0["username"] = s.Username
	j0["password"] = s.Password
	j0["sshPort"] = s.SshPort
	return j0
}

func (s *ServerPowerSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("enabled", j0)
	if err != nil {
		return err
	}
	s.Enabled, err = encoding.Is[bool](j0["enabled"])
	if err != nil {
		return err
	}
	err = encoding.In("target", j0)
	if err != nil {
		return err
	}
	s.Target, err = object.As[idl.Object](j0["target"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("powerCheck", j0)
	if err != nil {
		return err
	}
	s.PowerCheck, err = encoding.AsEnum[ServerPowerCheckMethod](j0["powerCheck"])
	if err != nil {
		return err
	}
	err = encoding.In("powerThreshold", j0)
	if err != nil {
		return err
	}
	s.PowerThreshold, err = encoding.AsFloat64(j0["powerThreshold"])
	if err != nil {
		return err
	}
	err = encoding.In("timeout", j0)
	if err != nil {
		return err
	}
	s.Timeout, err = encoding.AsInt32(j0["timeout"])
	if err != nil {
		return err
	}
	err = encoding.In("shutdownCmd", j0)
	if err != nil {
		return err
	}
	s.ShutdownCmd, err = encoding.Is[string](j0["shutdownCmd"])
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
	err = encoding.In("sshPort", j0)
	if err != nil {
		return err
	}
	s.SshPort, err = encoding.AsInt32(j0["sshPort"])
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerSettings) Encode() map[string]any {
	j0 := make(map[string]any, 9)
	j0["host"] = s.Host
	j0["enabled"] = s.Enabled
	j0["pingInterval"] = s.PingInterval
	j0["retryInterval"] = s.RetryInterval
	j0["activationCount"] = s.ActivationCount
	j0["failureCount"] = s.FailureCount
	j0["resumeDelay"] = s.ResumeDelay
	j0["resumeCount"] = s.ResumeCount
	j0["powerSettings"] = s.PowerSettings.Encode()
	return j0
}

func (s *ServerSettings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("host", j0)
	if err != nil {
		return err
	}
	s.Host, err = encoding.Is[string](j0["host"])
	if err != nil {
		return err
	}
	err = encoding.In("enabled", j0)
	if err != nil {
		return err
	}
	s.Enabled, err = encoding.Is[bool](j0["enabled"])
	if err != nil {
		return err
	}
	err = encoding.In("pingInterval", j0)
	if err != nil {
		return err
	}
	s.PingInterval, err = encoding.AsInt32(j0["pingInterval"])
	if err != nil {
		return err
	}
	err = encoding.In("retryInterval", j0)
	if err != nil {
		return err
	}
	s.RetryInterval, err = encoding.AsInt32(j0["retryInterval"])
	if err != nil {
		return err
	}
	err = encoding.In("activationCount", j0)
	if err != nil {
		return err
	}
	s.ActivationCount, err = encoding.AsInt32(j0["activationCount"])
	if err != nil {
		return err
	}
	err = encoding.In("failureCount", j0)
	if err != nil {
		return err
	}
	s.FailureCount, err = encoding.AsInt32(j0["failureCount"])
	if err != nil {
		return err
	}
	err = encoding.In("resumeDelay", j0)
	if err != nil {
		return err
	}
	s.ResumeDelay, err = encoding.AsInt32(j0["resumeDelay"])
	if err != nil {
		return err
	}
	err = encoding.In("resumeCount", j0)
	if err != nil {
		return err
	}
	s.ResumeCount, err = encoding.AsInt32(j0["resumeCount"])
	if err != nil {
		return err
	}
	err = encoding.In("powerSettings", j0)
	if err != nil {
		return err
	}
	err = s.PowerSettings.Decode(j0["powerSettings"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerStatus) Encode() map[string]any {
	j0 := make(map[string]any, 9)
	j0["powerState"] = s.PowerState
	j0["lastPowerControlResult"] = s.LastPowerControlResult
	j0["reachable"] = s.Reachable
	j0["lastRequest"] = s.LastRequest.Unix()
	j0["lastResponse"] = s.LastResponse.Unix()
	j0["requests"] = s.Requests
	j0["responses"] = s.Responses
	j0["failures"] = s.Failures
	j0["resumes"] = s.Resumes
	return j0
}

func (s *ServerStatus) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("powerState", j0)
	if err != nil {
		return err
	}
	s.PowerState, err = encoding.AsEnum[ServerPowerState](j0["powerState"])
	if err != nil {
		return err
	}
	err = encoding.In("lastPowerControlResult", j0)
	if err != nil {
		return err
	}
	s.LastPowerControlResult, err = encoding.AsEnum[ServerPowerControlResult](j0["lastPowerControlResult"])
	if err != nil {
		return err
	}
	err = encoding.In("reachable", j0)
	if err != nil {
		return err
	}
	s.Reachable, err = encoding.AsEnum[ServerReachability](j0["reachable"])
	if err != nil {
		return err
	}
	err = encoding.In("lastRequest", j0)
	if err != nil {
		return err
	}
	s.LastRequest, err = encoding.AsTime(j0["lastRequest"])
	if err != nil {
		return err
	}
	err = encoding.In("lastResponse", j0)
	if err != nil {
		return err
	}
	s.LastResponse, err = encoding.AsTime(j0["lastResponse"])
	if err != nil {
		return err
	}
	err = encoding.In("requests", j0)
	if err != nil {
		return err
	}
	s.Requests, err = encoding.AsInt32(j0["requests"])
	if err != nil {
		return err
	}
	err = encoding.In("responses", j0)
	if err != nil {
		return err
	}
	s.Responses, err = encoding.AsInt32(j0["responses"])
	if err != nil {
		return err
	}
	err = encoding.In("failures", j0)
	if err != nil {
		return err
	}
	s.Failures, err = encoding.AsInt32(j0["failures"])
	if err != nil {
		return err
	}
	err = encoding.In("resumes", j0)
	if err != nil {
		return err
	}
	s.Resumes, err = encoding.AsInt32(j0["resumes"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_ServerPowerStateEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.Event = valobj.For[event.Event]()
	err := s.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("id", value)
	if err != nil {
		return err
	}
	s.id, err = encoding.AsInt32(value["id"])
	if err != nil {
		return err
	}
	err = encoding.In("host", value)
	if err != nil {
		return err
	}
	s.host, err = encoding.Is[string](value["host"])
	if err != nil {
		return err
	}
	err = encoding.In("oldPowerState", value)
	if err != nil {
		return err
	}
	s.oldPowerState, err = encoding.AsEnum[ServerPowerState](value["oldPowerState"])
	if err != nil {
		return err
	}
	err = encoding.In("newPowerState", value)
	if err != nil {
		return err
	}
	s.newPowerState, err = encoding.AsEnum[ServerPowerState](value["newPowerState"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_ServerPowerControlInitiatedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.UserEvent = valobj.For[userevent.UserEvent]()
	err := s.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("id", value)
	if err != nil {
		return err
	}
	s.id, err = encoding.AsInt32(value["id"])
	if err != nil {
		return err
	}
	err = encoding.In("host", value)
	if err != nil {
		return err
	}
	s.host, err = encoding.Is[string](value["host"])
	if err != nil {
		return err
	}
	err = encoding.In("on", value)
	if err != nil {
		return err
	}
	s.on, err = encoding.Is[bool](value["on"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_ServerPowerControlCompletedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.Event = valobj.For[event.Event]()
	err := s.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("id", value)
	if err != nil {
		return err
	}
	s.id, err = encoding.AsInt32(value["id"])
	if err != nil {
		return err
	}
	err = encoding.In("host", value)
	if err != nil {
		return err
	}
	s.host, err = encoding.Is[string](value["host"])
	if err != nil {
		return err
	}
	err = encoding.In("result", value)
	if err != nil {
		return err
	}
	s.result, err = encoding.AsEnum[ServerPowerControlResult](value["result"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_ServerReachabilityEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.Event = valobj.For[event.Event]()
	err := s.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("id", value)
	if err != nil {
		return err
	}
	s.id, err = encoding.AsInt32(value["id"])
	if err != nil {
		return err
	}
	err = encoding.In("host", value)
	if err != nil {
		return err
	}
	s.host, err = encoding.Is[string](value["host"])
	if err != nil {
		return err
	}
	err = encoding.In("reachable", value)
	if err != nil {
		return err
	}
	s.reachable, err = encoding.AsEnum[ServerReachability](value["reachable"])
	if err != nil {
		return err
	}
	return nil
}

func (s *_ServerAddedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.UserEvent = valobj.For[userevent.UserEvent]()
	err := s.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("id", value)
	if err != nil {
		return err
	}
	s.id, err = encoding.AsInt32(value["id"])
	if err != nil {
		return err
	}
	err = encoding.In("settings", value)
	if err != nil {
		return err
	}
	err = s.settings.Decode(value["settings"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *_ServerSettingsChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.UserEvent = valobj.For[userevent.UserEvent]()
	err := s.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("id", value)
	if err != nil {
		return err
	}
	s.id, err = encoding.AsInt32(value["id"])
	if err != nil {
		return err
	}
	err = encoding.In("oldSettings", value)
	if err != nil {
		return err
	}
	err = s.oldSettings.Decode(value["oldSettings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("newSettings", value)
	if err != nil {
		return err
	}
	err = s.newSettings.Decode(value["newSettings"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *_ServerDeletedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.UserEvent = valobj.For[userevent.UserEvent]()
	err := s.UserEvent.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("id", value)
	if err != nil {
		return err
	}
	s.id, err = encoding.AsInt32(value["id"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["settings"] = s.Settings.Encode()
	j0["status"] = s.Status.Encode()
	return j0
}

func (s *Server) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("settings", j0)
	if err != nil {
		return err
	}
	err = s.Settings.Decode(j0["settings"], caller)
	if err != nil {
		return err
	}
	err = encoding.In("status", j0)
	if err != nil {
		return err
	}
	err = s.Status.Decode(j0["status"], caller)
	if err != nil {
		return err
	}
	return nil
}
