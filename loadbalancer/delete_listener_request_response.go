// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// DeleteListenerRequest wrapper for the DeleteListener operation
type DeleteListenerRequest struct {

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the load balancer associated with the listener to delete.
	LoadBalancerID *string `mandatory:"true" contributesTo:"path" name:"loadBalancerId"`

	// The name of the listener to delete.
	// Example: `My listener`
	ListenerName *string `mandatory:"true" contributesTo:"path" name:"listenerName"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`
}

func (request DeleteListenerRequest) String() string {
	return common.PointerString(request)
}

// DeleteListenerResponse wrapper for the DeleteListener operation
type DeleteListenerResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the work request.
	OpcWorkRequestID *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response DeleteListenerResponse) String() string {
	return common.PointerString(response)
}
