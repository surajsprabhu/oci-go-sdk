// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// Request wrapper for the GetIdentityProvider operation
type GetIdentityProviderRequest struct {

	// The OCID of the identity provider.
	IdentityProviderID string
}

// Response wrapper for the GetIdentityProvider operation
type GetIdentityProviderResponse struct {

	// The IdentityProvider instance
	IdentityProvider

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string

	// For optimistic concurrency control. See `if-match`.
	Etag string
}
