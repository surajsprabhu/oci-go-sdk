// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import "net/http"

// ListCustomerSecretKeysRequest wrapper for the ListCustomerSecretKeys operation
type ListCustomerSecretKeysRequest struct {

	// The OCID of the user.
	UserID *string `mandatory:"true" contributesTo:"path" name:"userId"`
}

// ListCustomerSecretKeysResponse wrapper for the ListCustomerSecretKeys operation
type ListCustomerSecretKeysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []CustomerSecretKeySummary instance
	Items []CustomerSecretKeySummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}