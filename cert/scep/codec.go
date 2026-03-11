// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package scep

import (
	"github.com/arminguenther/xeruspower-go/v40412/idl"
	"github.com/arminguenther/xeruspower-go/v40412/idl/event"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding"
	"github.com/arminguenther/xeruspower-go/v40412/internal/encoding/valobj"
)

func (e *EnrollmentData) Encode() map[string]any {
	j0 := make(map[string]any, 2)
	j0["challenge"] = e.Challenge
	j0["certInfo"] = e.CertInfo.Encode()
	return j0
}

func (e *EnrollmentData) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("challenge", j0)
	if err != nil {
		return err
	}
	e.Challenge, err = encoding.Is[string](j0["challenge"])
	if err != nil {
		return err
	}
	err = encoding.In("certInfo", j0)
	if err != nil {
		return err
	}
	err = e.CertInfo.Decode(j0["certInfo"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *Settings) Encode() map[string]any {
	j0 := make(map[string]any, 6)
	j0["caCertFingerprint"] = s.CaCertFingerprint
	j0["url"] = s.Url
	j0["renewDaysBeforeExpiry"] = s.RenewDaysBeforeExpiry
	j0["pendingWaitTimeSeconds"] = s.PendingWaitTimeSeconds
	j0["pendingRetries"] = s.PendingRetries
	j0["skipVerificationOfIssuedCertChain"] = s.SkipVerificationOfIssuedCertChain
	return j0
}

func (s *Settings) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("caCertFingerprint", j0)
	if err != nil {
		return err
	}
	s.CaCertFingerprint, err = encoding.Is[string](j0["caCertFingerprint"])
	if err != nil {
		return err
	}
	err = encoding.In("url", j0)
	if err != nil {
		return err
	}
	s.Url, err = encoding.Is[string](j0["url"])
	if err != nil {
		return err
	}
	err = encoding.In("renewDaysBeforeExpiry", j0)
	if err != nil {
		return err
	}
	s.RenewDaysBeforeExpiry, err = encoding.AsInt32(j0["renewDaysBeforeExpiry"])
	if err != nil {
		return err
	}
	err = encoding.In("pendingWaitTimeSeconds", j0)
	if err != nil {
		return err
	}
	s.PendingWaitTimeSeconds, err = encoding.AsInt32(j0["pendingWaitTimeSeconds"])
	if err != nil {
		return err
	}
	err = encoding.In("pendingRetries", j0)
	if err != nil {
		return err
	}
	s.PendingRetries, err = encoding.AsInt32(j0["pendingRetries"])
	if err != nil {
		return err
	}
	err = encoding.In("skipVerificationOfIssuedCertChain", j0)
	if err != nil {
		return err
	}
	s.SkipVerificationOfIssuedCertChain, err = encoding.Is[bool](j0["skipVerificationOfIssuedCertChain"])
	if err != nil {
		return err
	}
	return nil
}

func (e *EnrollmentRunInfo) Encode() map[string]any {
	j0 := make(map[string]any, 4)
	j0["status"] = e.Status
	j0["startedAt"] = e.StartedAt.Unix()
	j0["endedAt"] = e.EndedAt.Unix()
	j0["failureReason"] = e.FailureReason
	return j0
}

func (e *EnrollmentRunInfo) Decode(v any, caller idl.Caller) error {
	j0, err := encoding.Is[map[string]any](v)
	if err != nil {
		return err
	}
	err = encoding.In("status", j0)
	if err != nil {
		return err
	}
	e.Status, err = encoding.AsEnum[EnrollmentRunStatus](j0["status"])
	if err != nil {
		return err
	}
	err = encoding.In("startedAt", j0)
	if err != nil {
		return err
	}
	e.StartedAt, err = encoding.AsTime(j0["startedAt"])
	if err != nil {
		return err
	}
	err = encoding.In("endedAt", j0)
	if err != nil {
		return err
	}
	e.EndedAt, err = encoding.AsTime(j0["endedAt"])
	if err != nil {
		return err
	}
	err = encoding.In("failureReason", j0)
	if err != nil {
		return err
	}
	e.FailureReason, err = encoding.Is[string](j0["failureReason"])
	if err != nil {
		return err
	}
	return nil
}

func (s *Status) Encode() map[string]any {
	j0 := make(map[string]any, 3)
	j0["enabled"] = s.Enabled
	j0["lastSuccessfulEnrollment"] = s.LastSuccessfulEnrollment.Unix()
	j0["lastEnrollmentRunInfo"] = s.LastEnrollmentRunInfo.Encode()
	return j0
}

func (s *Status) Decode(v any, caller idl.Caller) error {
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
	err = encoding.In("lastSuccessfulEnrollment", j0)
	if err != nil {
		return err
	}
	s.LastSuccessfulEnrollment, err = encoding.AsTime(j0["lastSuccessfulEnrollment"])
	if err != nil {
		return err
	}
	err = encoding.In("lastEnrollmentRunInfo", j0)
	if err != nil {
		return err
	}
	err = s.LastEnrollmentRunInfo.Decode(j0["lastEnrollmentRunInfo"], caller)
	if err != nil {
		return err
	}
	return nil
}

func (s *_StatusChangedEvent) Decode(value map[string]any, caller idl.Caller) error {
	s.Event = valobj.For[event.Event]()
	err := s.Event.Decode(value, caller)
	if err != nil {
		return err
	}
	err = encoding.In("newStatus", value)
	if err != nil {
		return err
	}
	err = s.newStatus.Decode(value["newStatus"], caller)
	if err != nil {
		return err
	}
	return nil
}
