// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarHTTPService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRadarHTTPService] method instead.
type RadarHTTPService struct {
	Options    []option.RequestOption
	TLSVersion *RadarHTTPTLSVersionService
}

// NewRadarHTTPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarHTTPService(opts ...option.RequestOption) (r *RadarHTTPService) {
	r = &RadarHTTPService{}
	r.Options = opts
	r.TLSVersion = NewRadarHTTPTLSVersionService(opts...)
	return
}
