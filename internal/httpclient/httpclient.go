// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package httpclient

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/hc-install/version"
)

// NewHTTPClient provides a pre-configured http.Client
// e.g. with relevant User-Agent header
func NewHTTPClient() *http.Client {
	client := retryablehttp.NewClient().StandardClient()
	client.Transport = &userAgentRoundTripper{
		userAgent: fmt.Sprintf("hc-install/%s", version.Version()),
		inner:     client.Transport,
	}
	return client
}

type userAgentRoundTripper struct {
	inner     http.RoundTripper
	userAgent string
}

func (rt *userAgentRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if _, ok := req.Header["User-Agent"]; !ok {
		req.Header.Set("User-Agent", rt.userAgent)
	}
	return rt.inner.RoundTrip(req)
}
