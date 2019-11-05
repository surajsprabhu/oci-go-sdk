// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestLogEntry Log entries related to a specific work request.
type WorkRequestLogEntry struct {

	// The description of an action that occurred.
	Message *string `mandatory:"true" json:"message"`

	// The date and time the log entry occurred.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`
}

func (m WorkRequestLogEntry) String() string {
	return common.PointerString(m)
}
