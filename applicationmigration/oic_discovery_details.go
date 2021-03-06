// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// OicDiscoveryDetails Specifies the credentials to access the source OIC instance
type OicDiscoveryDetails struct {

	// The OIC instance admin user
	ServiceInstanceUser *string `mandatory:"true" json:"serviceInstanceUser"`

	// The OIC instance admin password
	ServiceInstancePassword *string `mandatory:"true" json:"serviceInstancePassword"`
}

func (m OicDiscoveryDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OicDiscoveryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOicDiscoveryDetails OicDiscoveryDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeOicDiscoveryDetails
	}{
		"OIC",
		(MarshalTypeOicDiscoveryDetails)(m),
	}

	return json.Marshal(&s)
}
