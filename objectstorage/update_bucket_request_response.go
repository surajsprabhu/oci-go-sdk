// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// UpdateBucketRequest wrapper for the UpdateBucket operation
type UpdateBucketRequest struct {

	// The top-level namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the bucket.
	// Example: `my-new-bucket1`
	BucketName *string `mandatory:"true" contributesTo:"path" name:"bucketName"`

	// Request object for updating a bucket.
	UpdateBucketDetails `contributesTo:"body"`

	// The entity tag to match. For creating and committing a multipart upload to an object, this is the entity tag of the target object.
	// For uploading a part, this is the entity tag of the target part.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The client request ID for tracing.
	OpcClientRequestID *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`
}

func (request UpdateBucketRequest) String() string {
	return common.PointerString(request)
}

// UpdateBucketResponse wrapper for the UpdateBucket operation
type UpdateBucketResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Bucket instance
	Bucket `presentIn:"body"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestID *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, please provide this request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`

	// The entity tag for the updated bucket.
	ETag *string `presentIn:"header" name:"etag"`
}

func (response UpdateBucketResponse) String() string {
	return common.PointerString(response)
}
