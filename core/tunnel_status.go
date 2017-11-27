// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TunnelStatus. Specific connection details for an IPSec tunnel.
type TunnelStatus struct {

	// The IP address of Oracle's VPN headend.
	// Example: `129.146.17.50`
	IpAddress *string `mandatory:"true" json:"ipAddress,omitempty"`

	// The tunnel's current state.
	LifecycleState TunnelStatusLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the IPSec connection was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`

	// When the state of the tunnel last changed, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeStateModified *common.SDKTime `mandatory:"false" json:"timeStateModified,omitempty"`
}

func (model TunnelStatus) String() string {
	return common.PointerString(model)
}

type TunnelStatusLifecycleStateEnum string

const (
	TUNNEL_STATUS_LIFECYCLE_STATE_UP                   TunnelStatusLifecycleStateEnum = "UP"
	TUNNEL_STATUS_LIFECYCLE_STATE_DOWN                 TunnelStatusLifecycleStateEnum = "DOWN"
	TUNNEL_STATUS_LIFECYCLE_STATE_DOWN_FOR_MAINTENANCE TunnelStatusLifecycleStateEnum = "DOWN_FOR_MAINTENANCE"
	TUNNEL_STATUS_LIFECYCLE_STATE_UNKNOWN              TunnelStatusLifecycleStateEnum = "UNKNOWN"
)

var mapping_tunnelstatus_lifecycleState = map[string]TunnelStatusLifecycleStateEnum{
	"UP":                   TUNNEL_STATUS_LIFECYCLE_STATE_UP,
	"DOWN":                 TUNNEL_STATUS_LIFECYCLE_STATE_DOWN,
	"DOWN_FOR_MAINTENANCE": TUNNEL_STATUS_LIFECYCLE_STATE_DOWN_FOR_MAINTENANCE,
	"UNKNOWN":              TUNNEL_STATUS_LIFECYCLE_STATE_UNKNOWN,
}

func GetTunnelStatusLifecycleStateEnumValues() []TunnelStatusLifecycleStateEnum {
	values := make([]TunnelStatusLifecycleStateEnum, 0)
	for _, v := range mapping_tunnelstatus_lifecycleState {
		if v != TUNNEL_STATUS_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
