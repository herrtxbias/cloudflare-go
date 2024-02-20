// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// HealthcheckService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewHealthcheckService] method
// instead.
type HealthcheckService struct {
	Options  []option.RequestOption
	Previews *HealthcheckPreviewService
}

// NewHealthcheckService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHealthcheckService(opts ...option.RequestOption) (r *HealthcheckService) {
	r = &HealthcheckService{}
	r.Options = opts
	r.Previews = NewHealthcheckPreviewService(opts...)
	return
}

// Create a new health check.
func (r *HealthcheckService) New(ctx context.Context, zoneIdentifier string, body HealthcheckNewParams, opts ...option.RequestOption) (res *HealthcheckNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured health checks.
func (r *HealthcheckService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]HealthcheckListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckListResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a health check.
func (r *HealthcheckService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *HealthcheckDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured health check.
func (r *HealthcheckService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *HealthcheckGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a configured health check.
func (r *HealthcheckService) Replace(ctx context.Context, zoneIdentifier string, identifier string, body HealthcheckReplaceParams, opts ...option.RequestOption) (res *HealthcheckReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HealthcheckReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HealthcheckNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []HealthcheckNewResponseCheckRegion `json:"check_regions,nullable"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails int64 `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses int64     `json:"consecutive_successes"`
	CreatedOn            time.Time `json:"created_on" format:"date-time"`
	// A human-readable description of the health check.
	Description string `json:"description"`
	// The current failure reason if status is unhealthy.
	FailureReason string `json:"failure_reason"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig HealthcheckNewResponseHTTPConfig `json:"http_config,nullable"`
	// The interval between each health check. Shorter intervals may give quicker
	// notifications if the origin status changes, but will increase load on the origin
	// as we check from multiple locations.
	Interval   int64     `json:"interval"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name string `json:"name"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries int64 `json:"retries"`
	// The current status of the origin server according to the health check.
	Status HealthcheckNewResponseStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig HealthcheckNewResponseTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                     `json:"type"`
	JSON healthcheckNewResponseJSON `json:"-"`
}

// healthcheckNewResponseJSON contains the JSON metadata for the struct
// [HealthcheckNewResponse]
type healthcheckNewResponseJSON struct {
	ID                   apijson.Field
	Address              apijson.Field
	CheckRegions         apijson.Field
	ConsecutiveFails     apijson.Field
	ConsecutiveSuccesses apijson.Field
	CreatedOn            apijson.Field
	Description          apijson.Field
	FailureReason        apijson.Field
	HTTPConfig           apijson.Field
	Interval             apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	Retries              apijson.Field
	Status               apijson.Field
	Suspended            apijson.Field
	TcpConfig            apijson.Field
	Timeout              apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HealthcheckNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckNewResponseCheckRegion string

const (
	HealthcheckNewResponseCheckRegionWnam       HealthcheckNewResponseCheckRegion = "WNAM"
	HealthcheckNewResponseCheckRegionEnam       HealthcheckNewResponseCheckRegion = "ENAM"
	HealthcheckNewResponseCheckRegionWeu        HealthcheckNewResponseCheckRegion = "WEU"
	HealthcheckNewResponseCheckRegionEeu        HealthcheckNewResponseCheckRegion = "EEU"
	HealthcheckNewResponseCheckRegionNsam       HealthcheckNewResponseCheckRegion = "NSAM"
	HealthcheckNewResponseCheckRegionSsam       HealthcheckNewResponseCheckRegion = "SSAM"
	HealthcheckNewResponseCheckRegionOc         HealthcheckNewResponseCheckRegion = "OC"
	HealthcheckNewResponseCheckRegionMe         HealthcheckNewResponseCheckRegion = "ME"
	HealthcheckNewResponseCheckRegionNaf        HealthcheckNewResponseCheckRegion = "NAF"
	HealthcheckNewResponseCheckRegionSaf        HealthcheckNewResponseCheckRegion = "SAF"
	HealthcheckNewResponseCheckRegionIn         HealthcheckNewResponseCheckRegion = "IN"
	HealthcheckNewResponseCheckRegionSeas       HealthcheckNewResponseCheckRegion = "SEAS"
	HealthcheckNewResponseCheckRegionNeas       HealthcheckNewResponseCheckRegion = "NEAS"
	HealthcheckNewResponseCheckRegionAllRegions HealthcheckNewResponseCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckNewResponseHTTPConfig struct {
	// Do not validate the certificate when the health check uses HTTPS.
	AllowInsecure bool `json:"allow_insecure"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy.
	ExpectedBody string `json:"expected_body"`
	// The expected HTTP response codes (e.g. "200") or code ranges (e.g. "2xx" for all
	// codes starting with 2) of the health check.
	ExpectedCodes []string `json:"expected_codes,nullable"`
	// Follow redirects if the origin returns a 3xx status code.
	FollowRedirects bool `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden.
	Header interface{} `json:"header,nullable"`
	// The HTTP method to use for the health check.
	Method HealthcheckNewResponseHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                `json:"port"`
	JSON healthcheckNewResponseHTTPConfigJSON `json:"-"`
}

// healthcheckNewResponseHTTPConfigJSON contains the JSON metadata for the struct
// [HealthcheckNewResponseHTTPConfig]
type healthcheckNewResponseHTTPConfigJSON struct {
	AllowInsecure   apijson.Field
	ExpectedBody    apijson.Field
	ExpectedCodes   apijson.Field
	FollowRedirects apijson.Field
	Header          apijson.Field
	Method          apijson.Field
	Path            apijson.Field
	Port            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *HealthcheckNewResponseHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method to use for the health check.
type HealthcheckNewResponseHTTPConfigMethod string

const (
	HealthcheckNewResponseHTTPConfigMethodGet  HealthcheckNewResponseHTTPConfigMethod = "GET"
	HealthcheckNewResponseHTTPConfigMethodHead HealthcheckNewResponseHTTPConfigMethod = "HEAD"
)

// The current status of the origin server according to the health check.
type HealthcheckNewResponseStatus string

const (
	HealthcheckNewResponseStatusUnknown   HealthcheckNewResponseStatus = "unknown"
	HealthcheckNewResponseStatusHealthy   HealthcheckNewResponseStatus = "healthy"
	HealthcheckNewResponseStatusUnhealthy HealthcheckNewResponseStatus = "unhealthy"
	HealthcheckNewResponseStatusSuspended HealthcheckNewResponseStatus = "suspended"
)

// Parameters specific to TCP health check.
type HealthcheckNewResponseTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method HealthcheckNewResponseTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                               `json:"port"`
	JSON healthcheckNewResponseTcpConfigJSON `json:"-"`
}

// healthcheckNewResponseTcpConfigJSON contains the JSON metadata for the struct
// [HealthcheckNewResponseTcpConfig]
type healthcheckNewResponseTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckNewResponseTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The TCP connection method to use for the health check.
type HealthcheckNewResponseTcpConfigMethod string

const (
	HealthcheckNewResponseTcpConfigMethodConnectionEstablished HealthcheckNewResponseTcpConfigMethod = "connection_established"
)

type HealthcheckListResponse struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []HealthcheckListResponseCheckRegion `json:"check_regions,nullable"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails int64 `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses int64     `json:"consecutive_successes"`
	CreatedOn            time.Time `json:"created_on" format:"date-time"`
	// A human-readable description of the health check.
	Description string `json:"description"`
	// The current failure reason if status is unhealthy.
	FailureReason string `json:"failure_reason"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig HealthcheckListResponseHTTPConfig `json:"http_config,nullable"`
	// The interval between each health check. Shorter intervals may give quicker
	// notifications if the origin status changes, but will increase load on the origin
	// as we check from multiple locations.
	Interval   int64     `json:"interval"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name string `json:"name"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries int64 `json:"retries"`
	// The current status of the origin server according to the health check.
	Status HealthcheckListResponseStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig HealthcheckListResponseTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                      `json:"type"`
	JSON healthcheckListResponseJSON `json:"-"`
}

// healthcheckListResponseJSON contains the JSON metadata for the struct
// [HealthcheckListResponse]
type healthcheckListResponseJSON struct {
	ID                   apijson.Field
	Address              apijson.Field
	CheckRegions         apijson.Field
	ConsecutiveFails     apijson.Field
	ConsecutiveSuccesses apijson.Field
	CreatedOn            apijson.Field
	Description          apijson.Field
	FailureReason        apijson.Field
	HTTPConfig           apijson.Field
	Interval             apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	Retries              apijson.Field
	Status               apijson.Field
	Suspended            apijson.Field
	TcpConfig            apijson.Field
	Timeout              apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HealthcheckListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckListResponseCheckRegion string

const (
	HealthcheckListResponseCheckRegionWnam       HealthcheckListResponseCheckRegion = "WNAM"
	HealthcheckListResponseCheckRegionEnam       HealthcheckListResponseCheckRegion = "ENAM"
	HealthcheckListResponseCheckRegionWeu        HealthcheckListResponseCheckRegion = "WEU"
	HealthcheckListResponseCheckRegionEeu        HealthcheckListResponseCheckRegion = "EEU"
	HealthcheckListResponseCheckRegionNsam       HealthcheckListResponseCheckRegion = "NSAM"
	HealthcheckListResponseCheckRegionSsam       HealthcheckListResponseCheckRegion = "SSAM"
	HealthcheckListResponseCheckRegionOc         HealthcheckListResponseCheckRegion = "OC"
	HealthcheckListResponseCheckRegionMe         HealthcheckListResponseCheckRegion = "ME"
	HealthcheckListResponseCheckRegionNaf        HealthcheckListResponseCheckRegion = "NAF"
	HealthcheckListResponseCheckRegionSaf        HealthcheckListResponseCheckRegion = "SAF"
	HealthcheckListResponseCheckRegionIn         HealthcheckListResponseCheckRegion = "IN"
	HealthcheckListResponseCheckRegionSeas       HealthcheckListResponseCheckRegion = "SEAS"
	HealthcheckListResponseCheckRegionNeas       HealthcheckListResponseCheckRegion = "NEAS"
	HealthcheckListResponseCheckRegionAllRegions HealthcheckListResponseCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckListResponseHTTPConfig struct {
	// Do not validate the certificate when the health check uses HTTPS.
	AllowInsecure bool `json:"allow_insecure"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy.
	ExpectedBody string `json:"expected_body"`
	// The expected HTTP response codes (e.g. "200") or code ranges (e.g. "2xx" for all
	// codes starting with 2) of the health check.
	ExpectedCodes []string `json:"expected_codes,nullable"`
	// Follow redirects if the origin returns a 3xx status code.
	FollowRedirects bool `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden.
	Header interface{} `json:"header,nullable"`
	// The HTTP method to use for the health check.
	Method HealthcheckListResponseHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                 `json:"port"`
	JSON healthcheckListResponseHTTPConfigJSON `json:"-"`
}

// healthcheckListResponseHTTPConfigJSON contains the JSON metadata for the struct
// [HealthcheckListResponseHTTPConfig]
type healthcheckListResponseHTTPConfigJSON struct {
	AllowInsecure   apijson.Field
	ExpectedBody    apijson.Field
	ExpectedCodes   apijson.Field
	FollowRedirects apijson.Field
	Header          apijson.Field
	Method          apijson.Field
	Path            apijson.Field
	Port            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *HealthcheckListResponseHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method to use for the health check.
type HealthcheckListResponseHTTPConfigMethod string

const (
	HealthcheckListResponseHTTPConfigMethodGet  HealthcheckListResponseHTTPConfigMethod = "GET"
	HealthcheckListResponseHTTPConfigMethodHead HealthcheckListResponseHTTPConfigMethod = "HEAD"
)

// The current status of the origin server according to the health check.
type HealthcheckListResponseStatus string

const (
	HealthcheckListResponseStatusUnknown   HealthcheckListResponseStatus = "unknown"
	HealthcheckListResponseStatusHealthy   HealthcheckListResponseStatus = "healthy"
	HealthcheckListResponseStatusUnhealthy HealthcheckListResponseStatus = "unhealthy"
	HealthcheckListResponseStatusSuspended HealthcheckListResponseStatus = "suspended"
)

// Parameters specific to TCP health check.
type HealthcheckListResponseTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method HealthcheckListResponseTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                                `json:"port"`
	JSON healthcheckListResponseTcpConfigJSON `json:"-"`
}

// healthcheckListResponseTcpConfigJSON contains the JSON metadata for the struct
// [HealthcheckListResponseTcpConfig]
type healthcheckListResponseTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckListResponseTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The TCP connection method to use for the health check.
type HealthcheckListResponseTcpConfigMethod string

const (
	HealthcheckListResponseTcpConfigMethodConnectionEstablished HealthcheckListResponseTcpConfigMethod = "connection_established"
)

type HealthcheckDeleteResponse struct {
	// Identifier
	ID   string                        `json:"id"`
	JSON healthcheckDeleteResponseJSON `json:"-"`
}

// healthcheckDeleteResponseJSON contains the JSON metadata for the struct
// [HealthcheckDeleteResponse]
type healthcheckDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []HealthcheckGetResponseCheckRegion `json:"check_regions,nullable"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails int64 `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses int64     `json:"consecutive_successes"`
	CreatedOn            time.Time `json:"created_on" format:"date-time"`
	// A human-readable description of the health check.
	Description string `json:"description"`
	// The current failure reason if status is unhealthy.
	FailureReason string `json:"failure_reason"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig HealthcheckGetResponseHTTPConfig `json:"http_config,nullable"`
	// The interval between each health check. Shorter intervals may give quicker
	// notifications if the origin status changes, but will increase load on the origin
	// as we check from multiple locations.
	Interval   int64     `json:"interval"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name string `json:"name"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries int64 `json:"retries"`
	// The current status of the origin server according to the health check.
	Status HealthcheckGetResponseStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig HealthcheckGetResponseTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                     `json:"type"`
	JSON healthcheckGetResponseJSON `json:"-"`
}

// healthcheckGetResponseJSON contains the JSON metadata for the struct
// [HealthcheckGetResponse]
type healthcheckGetResponseJSON struct {
	ID                   apijson.Field
	Address              apijson.Field
	CheckRegions         apijson.Field
	ConsecutiveFails     apijson.Field
	ConsecutiveSuccesses apijson.Field
	CreatedOn            apijson.Field
	Description          apijson.Field
	FailureReason        apijson.Field
	HTTPConfig           apijson.Field
	Interval             apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	Retries              apijson.Field
	Status               apijson.Field
	Suspended            apijson.Field
	TcpConfig            apijson.Field
	Timeout              apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HealthcheckGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckGetResponseCheckRegion string

const (
	HealthcheckGetResponseCheckRegionWnam       HealthcheckGetResponseCheckRegion = "WNAM"
	HealthcheckGetResponseCheckRegionEnam       HealthcheckGetResponseCheckRegion = "ENAM"
	HealthcheckGetResponseCheckRegionWeu        HealthcheckGetResponseCheckRegion = "WEU"
	HealthcheckGetResponseCheckRegionEeu        HealthcheckGetResponseCheckRegion = "EEU"
	HealthcheckGetResponseCheckRegionNsam       HealthcheckGetResponseCheckRegion = "NSAM"
	HealthcheckGetResponseCheckRegionSsam       HealthcheckGetResponseCheckRegion = "SSAM"
	HealthcheckGetResponseCheckRegionOc         HealthcheckGetResponseCheckRegion = "OC"
	HealthcheckGetResponseCheckRegionMe         HealthcheckGetResponseCheckRegion = "ME"
	HealthcheckGetResponseCheckRegionNaf        HealthcheckGetResponseCheckRegion = "NAF"
	HealthcheckGetResponseCheckRegionSaf        HealthcheckGetResponseCheckRegion = "SAF"
	HealthcheckGetResponseCheckRegionIn         HealthcheckGetResponseCheckRegion = "IN"
	HealthcheckGetResponseCheckRegionSeas       HealthcheckGetResponseCheckRegion = "SEAS"
	HealthcheckGetResponseCheckRegionNeas       HealthcheckGetResponseCheckRegion = "NEAS"
	HealthcheckGetResponseCheckRegionAllRegions HealthcheckGetResponseCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckGetResponseHTTPConfig struct {
	// Do not validate the certificate when the health check uses HTTPS.
	AllowInsecure bool `json:"allow_insecure"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy.
	ExpectedBody string `json:"expected_body"`
	// The expected HTTP response codes (e.g. "200") or code ranges (e.g. "2xx" for all
	// codes starting with 2) of the health check.
	ExpectedCodes []string `json:"expected_codes,nullable"`
	// Follow redirects if the origin returns a 3xx status code.
	FollowRedirects bool `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden.
	Header interface{} `json:"header,nullable"`
	// The HTTP method to use for the health check.
	Method HealthcheckGetResponseHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                `json:"port"`
	JSON healthcheckGetResponseHTTPConfigJSON `json:"-"`
}

// healthcheckGetResponseHTTPConfigJSON contains the JSON metadata for the struct
// [HealthcheckGetResponseHTTPConfig]
type healthcheckGetResponseHTTPConfigJSON struct {
	AllowInsecure   apijson.Field
	ExpectedBody    apijson.Field
	ExpectedCodes   apijson.Field
	FollowRedirects apijson.Field
	Header          apijson.Field
	Method          apijson.Field
	Path            apijson.Field
	Port            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *HealthcheckGetResponseHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method to use for the health check.
type HealthcheckGetResponseHTTPConfigMethod string

const (
	HealthcheckGetResponseHTTPConfigMethodGet  HealthcheckGetResponseHTTPConfigMethod = "GET"
	HealthcheckGetResponseHTTPConfigMethodHead HealthcheckGetResponseHTTPConfigMethod = "HEAD"
)

// The current status of the origin server according to the health check.
type HealthcheckGetResponseStatus string

const (
	HealthcheckGetResponseStatusUnknown   HealthcheckGetResponseStatus = "unknown"
	HealthcheckGetResponseStatusHealthy   HealthcheckGetResponseStatus = "healthy"
	HealthcheckGetResponseStatusUnhealthy HealthcheckGetResponseStatus = "unhealthy"
	HealthcheckGetResponseStatusSuspended HealthcheckGetResponseStatus = "suspended"
)

// Parameters specific to TCP health check.
type HealthcheckGetResponseTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method HealthcheckGetResponseTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                               `json:"port"`
	JSON healthcheckGetResponseTcpConfigJSON `json:"-"`
}

// healthcheckGetResponseTcpConfigJSON contains the JSON metadata for the struct
// [HealthcheckGetResponseTcpConfig]
type healthcheckGetResponseTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckGetResponseTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The TCP connection method to use for the health check.
type HealthcheckGetResponseTcpConfigMethod string

const (
	HealthcheckGetResponseTcpConfigMethodConnectionEstablished HealthcheckGetResponseTcpConfigMethod = "connection_established"
)

type HealthcheckReplaceResponse struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []HealthcheckReplaceResponseCheckRegion `json:"check_regions,nullable"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails int64 `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses int64     `json:"consecutive_successes"`
	CreatedOn            time.Time `json:"created_on" format:"date-time"`
	// A human-readable description of the health check.
	Description string `json:"description"`
	// The current failure reason if status is unhealthy.
	FailureReason string `json:"failure_reason"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig HealthcheckReplaceResponseHTTPConfig `json:"http_config,nullable"`
	// The interval between each health check. Shorter intervals may give quicker
	// notifications if the origin status changes, but will increase load on the origin
	// as we check from multiple locations.
	Interval   int64     `json:"interval"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name string `json:"name"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries int64 `json:"retries"`
	// The current status of the origin server according to the health check.
	Status HealthcheckReplaceResponseStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig HealthcheckReplaceResponseTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string                         `json:"type"`
	JSON healthcheckReplaceResponseJSON `json:"-"`
}

// healthcheckReplaceResponseJSON contains the JSON metadata for the struct
// [HealthcheckReplaceResponse]
type healthcheckReplaceResponseJSON struct {
	ID                   apijson.Field
	Address              apijson.Field
	CheckRegions         apijson.Field
	ConsecutiveFails     apijson.Field
	ConsecutiveSuccesses apijson.Field
	CreatedOn            apijson.Field
	Description          apijson.Field
	FailureReason        apijson.Field
	HTTPConfig           apijson.Field
	Interval             apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	Retries              apijson.Field
	Status               apijson.Field
	Suspended            apijson.Field
	TcpConfig            apijson.Field
	Timeout              apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HealthcheckReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckReplaceResponseCheckRegion string

const (
	HealthcheckReplaceResponseCheckRegionWnam       HealthcheckReplaceResponseCheckRegion = "WNAM"
	HealthcheckReplaceResponseCheckRegionEnam       HealthcheckReplaceResponseCheckRegion = "ENAM"
	HealthcheckReplaceResponseCheckRegionWeu        HealthcheckReplaceResponseCheckRegion = "WEU"
	HealthcheckReplaceResponseCheckRegionEeu        HealthcheckReplaceResponseCheckRegion = "EEU"
	HealthcheckReplaceResponseCheckRegionNsam       HealthcheckReplaceResponseCheckRegion = "NSAM"
	HealthcheckReplaceResponseCheckRegionSsam       HealthcheckReplaceResponseCheckRegion = "SSAM"
	HealthcheckReplaceResponseCheckRegionOc         HealthcheckReplaceResponseCheckRegion = "OC"
	HealthcheckReplaceResponseCheckRegionMe         HealthcheckReplaceResponseCheckRegion = "ME"
	HealthcheckReplaceResponseCheckRegionNaf        HealthcheckReplaceResponseCheckRegion = "NAF"
	HealthcheckReplaceResponseCheckRegionSaf        HealthcheckReplaceResponseCheckRegion = "SAF"
	HealthcheckReplaceResponseCheckRegionIn         HealthcheckReplaceResponseCheckRegion = "IN"
	HealthcheckReplaceResponseCheckRegionSeas       HealthcheckReplaceResponseCheckRegion = "SEAS"
	HealthcheckReplaceResponseCheckRegionNeas       HealthcheckReplaceResponseCheckRegion = "NEAS"
	HealthcheckReplaceResponseCheckRegionAllRegions HealthcheckReplaceResponseCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckReplaceResponseHTTPConfig struct {
	// Do not validate the certificate when the health check uses HTTPS.
	AllowInsecure bool `json:"allow_insecure"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy.
	ExpectedBody string `json:"expected_body"`
	// The expected HTTP response codes (e.g. "200") or code ranges (e.g. "2xx" for all
	// codes starting with 2) of the health check.
	ExpectedCodes []string `json:"expected_codes,nullable"`
	// Follow redirects if the origin returns a 3xx status code.
	FollowRedirects bool `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden.
	Header interface{} `json:"header,nullable"`
	// The HTTP method to use for the health check.
	Method HealthcheckReplaceResponseHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                                    `json:"port"`
	JSON healthcheckReplaceResponseHTTPConfigJSON `json:"-"`
}

// healthcheckReplaceResponseHTTPConfigJSON contains the JSON metadata for the
// struct [HealthcheckReplaceResponseHTTPConfig]
type healthcheckReplaceResponseHTTPConfigJSON struct {
	AllowInsecure   apijson.Field
	ExpectedBody    apijson.Field
	ExpectedCodes   apijson.Field
	FollowRedirects apijson.Field
	Header          apijson.Field
	Method          apijson.Field
	Path            apijson.Field
	Port            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *HealthcheckReplaceResponseHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method to use for the health check.
type HealthcheckReplaceResponseHTTPConfigMethod string

const (
	HealthcheckReplaceResponseHTTPConfigMethodGet  HealthcheckReplaceResponseHTTPConfigMethod = "GET"
	HealthcheckReplaceResponseHTTPConfigMethodHead HealthcheckReplaceResponseHTTPConfigMethod = "HEAD"
)

// The current status of the origin server according to the health check.
type HealthcheckReplaceResponseStatus string

const (
	HealthcheckReplaceResponseStatusUnknown   HealthcheckReplaceResponseStatus = "unknown"
	HealthcheckReplaceResponseStatusHealthy   HealthcheckReplaceResponseStatus = "healthy"
	HealthcheckReplaceResponseStatusUnhealthy HealthcheckReplaceResponseStatus = "unhealthy"
	HealthcheckReplaceResponseStatusSuspended HealthcheckReplaceResponseStatus = "suspended"
)

// Parameters specific to TCP health check.
type HealthcheckReplaceResponseTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method HealthcheckReplaceResponseTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                                   `json:"port"`
	JSON healthcheckReplaceResponseTcpConfigJSON `json:"-"`
}

// healthcheckReplaceResponseTcpConfigJSON contains the JSON metadata for the
// struct [HealthcheckReplaceResponseTcpConfig]
type healthcheckReplaceResponseTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckReplaceResponseTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The TCP connection method to use for the health check.
type HealthcheckReplaceResponseTcpConfigMethod string

const (
	HealthcheckReplaceResponseTcpConfigMethodConnectionEstablished HealthcheckReplaceResponseTcpConfigMethod = "connection_established"
)

type HealthcheckNewParams struct {
	// The hostname or IP address of the origin server to run health checks on.
	Address param.Field[string] `json:"address,required"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions param.Field[[]HealthcheckNewParamsCheckRegion] `json:"check_regions"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails param.Field[int64] `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses param.Field[int64] `json:"consecutive_successes"`
	// A human-readable description of the health check.
	Description param.Field[string] `json:"description"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig param.Field[HealthcheckNewParamsHTTPConfig] `json:"http_config"`
	// The interval between each health check. Shorter intervals may give quicker
	// notifications if the origin status changes, but will increase load on the origin
	// as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// If suspended, no health checks are sent to the origin.
	Suspended param.Field[bool] `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig param.Field[HealthcheckNewParamsTcpConfig] `json:"tcp_config"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type param.Field[string] `json:"type"`
}

func (r HealthcheckNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckNewParamsCheckRegion string

const (
	HealthcheckNewParamsCheckRegionWnam       HealthcheckNewParamsCheckRegion = "WNAM"
	HealthcheckNewParamsCheckRegionEnam       HealthcheckNewParamsCheckRegion = "ENAM"
	HealthcheckNewParamsCheckRegionWeu        HealthcheckNewParamsCheckRegion = "WEU"
	HealthcheckNewParamsCheckRegionEeu        HealthcheckNewParamsCheckRegion = "EEU"
	HealthcheckNewParamsCheckRegionNsam       HealthcheckNewParamsCheckRegion = "NSAM"
	HealthcheckNewParamsCheckRegionSsam       HealthcheckNewParamsCheckRegion = "SSAM"
	HealthcheckNewParamsCheckRegionOc         HealthcheckNewParamsCheckRegion = "OC"
	HealthcheckNewParamsCheckRegionMe         HealthcheckNewParamsCheckRegion = "ME"
	HealthcheckNewParamsCheckRegionNaf        HealthcheckNewParamsCheckRegion = "NAF"
	HealthcheckNewParamsCheckRegionSaf        HealthcheckNewParamsCheckRegion = "SAF"
	HealthcheckNewParamsCheckRegionIn         HealthcheckNewParamsCheckRegion = "IN"
	HealthcheckNewParamsCheckRegionSeas       HealthcheckNewParamsCheckRegion = "SEAS"
	HealthcheckNewParamsCheckRegionNeas       HealthcheckNewParamsCheckRegion = "NEAS"
	HealthcheckNewParamsCheckRegionAllRegions HealthcheckNewParamsCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckNewParamsHTTPConfig struct {
	// Do not validate the certificate when the health check uses HTTPS.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// The expected HTTP response codes (e.g. "200") or code ranges (e.g. "2xx" for all
	// codes starting with 2) of the health check.
	ExpectedCodes param.Field[[]string] `json:"expected_codes"`
	// Follow redirects if the origin returns a 3xx status code.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden.
	Header param.Field[interface{}] `json:"header"`
	// The HTTP method to use for the health check.
	Method param.Field[HealthcheckNewParamsHTTPConfigMethod] `json:"method"`
	// The endpoint path to health check against.
	Path param.Field[string] `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port param.Field[int64] `json:"port"`
}

func (r HealthcheckNewParamsHTTPConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method to use for the health check.
type HealthcheckNewParamsHTTPConfigMethod string

const (
	HealthcheckNewParamsHTTPConfigMethodGet  HealthcheckNewParamsHTTPConfigMethod = "GET"
	HealthcheckNewParamsHTTPConfigMethodHead HealthcheckNewParamsHTTPConfigMethod = "HEAD"
)

// Parameters specific to TCP health check.
type HealthcheckNewParamsTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method param.Field[HealthcheckNewParamsTcpConfigMethod] `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port param.Field[int64] `json:"port"`
}

func (r HealthcheckNewParamsTcpConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TCP connection method to use for the health check.
type HealthcheckNewParamsTcpConfigMethod string

const (
	HealthcheckNewParamsTcpConfigMethodConnectionEstablished HealthcheckNewParamsTcpConfigMethod = "connection_established"
)

type HealthcheckNewResponseEnvelope struct {
	Errors   []HealthcheckNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckNewResponseEnvelopeMessages `json:"messages,required"`
	Result   HealthcheckNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success HealthcheckNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    healthcheckNewResponseEnvelopeJSON    `json:"-"`
}

// healthcheckNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [HealthcheckNewResponseEnvelope]
type healthcheckNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    healthcheckNewResponseEnvelopeErrorsJSON `json:"-"`
}

// healthcheckNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HealthcheckNewResponseEnvelopeErrors]
type healthcheckNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    healthcheckNewResponseEnvelopeMessagesJSON `json:"-"`
}

// healthcheckNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HealthcheckNewResponseEnvelopeMessages]
type healthcheckNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HealthcheckNewResponseEnvelopeSuccess bool

const (
	HealthcheckNewResponseEnvelopeSuccessTrue HealthcheckNewResponseEnvelopeSuccess = true
)

type HealthcheckListResponseEnvelope struct {
	Errors   []HealthcheckListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckListResponseEnvelopeMessages `json:"messages,required"`
	Result   []HealthcheckListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    HealthcheckListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo HealthcheckListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       healthcheckListResponseEnvelopeJSON       `json:"-"`
}

// healthcheckListResponseEnvelopeJSON contains the JSON metadata for the struct
// [HealthcheckListResponseEnvelope]
type healthcheckListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckListResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    healthcheckListResponseEnvelopeErrorsJSON `json:"-"`
}

// healthcheckListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HealthcheckListResponseEnvelopeErrors]
type healthcheckListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckListResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    healthcheckListResponseEnvelopeMessagesJSON `json:"-"`
}

// healthcheckListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HealthcheckListResponseEnvelopeMessages]
type healthcheckListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HealthcheckListResponseEnvelopeSuccess bool

const (
	HealthcheckListResponseEnvelopeSuccessTrue HealthcheckListResponseEnvelopeSuccess = true
)

type HealthcheckListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       healthcheckListResponseEnvelopeResultInfoJSON `json:"-"`
}

// healthcheckListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [HealthcheckListResponseEnvelopeResultInfo]
type healthcheckListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckDeleteResponseEnvelope struct {
	Errors   []HealthcheckDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   HealthcheckDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success HealthcheckDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    healthcheckDeleteResponseEnvelopeJSON    `json:"-"`
}

// healthcheckDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [HealthcheckDeleteResponseEnvelope]
type healthcheckDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckDeleteResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    healthcheckDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// healthcheckDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HealthcheckDeleteResponseEnvelopeErrors]
type healthcheckDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckDeleteResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    healthcheckDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// healthcheckDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HealthcheckDeleteResponseEnvelopeMessages]
type healthcheckDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HealthcheckDeleteResponseEnvelopeSuccess bool

const (
	HealthcheckDeleteResponseEnvelopeSuccessTrue HealthcheckDeleteResponseEnvelopeSuccess = true
)

type HealthcheckGetResponseEnvelope struct {
	Errors   []HealthcheckGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckGetResponseEnvelopeMessages `json:"messages,required"`
	Result   HealthcheckGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success HealthcheckGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    healthcheckGetResponseEnvelopeJSON    `json:"-"`
}

// healthcheckGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [HealthcheckGetResponseEnvelope]
type healthcheckGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    healthcheckGetResponseEnvelopeErrorsJSON `json:"-"`
}

// healthcheckGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HealthcheckGetResponseEnvelopeErrors]
type healthcheckGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    healthcheckGetResponseEnvelopeMessagesJSON `json:"-"`
}

// healthcheckGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [HealthcheckGetResponseEnvelopeMessages]
type healthcheckGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HealthcheckGetResponseEnvelopeSuccess bool

const (
	HealthcheckGetResponseEnvelopeSuccessTrue HealthcheckGetResponseEnvelopeSuccess = true
)

type HealthcheckReplaceParams struct {
	// The hostname or IP address of the origin server to run health checks on.
	Address param.Field[string] `json:"address,required"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions param.Field[[]HealthcheckReplaceParamsCheckRegion] `json:"check_regions"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails param.Field[int64] `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses param.Field[int64] `json:"consecutive_successes"`
	// A human-readable description of the health check.
	Description param.Field[string] `json:"description"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig param.Field[HealthcheckReplaceParamsHTTPConfig] `json:"http_config"`
	// The interval between each health check. Shorter intervals may give quicker
	// notifications if the origin status changes, but will increase load on the origin
	// as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// If suspended, no health checks are sent to the origin.
	Suspended param.Field[bool] `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig param.Field[HealthcheckReplaceParamsTcpConfig] `json:"tcp_config"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type param.Field[string] `json:"type"`
}

func (r HealthcheckReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckReplaceParamsCheckRegion string

const (
	HealthcheckReplaceParamsCheckRegionWnam       HealthcheckReplaceParamsCheckRegion = "WNAM"
	HealthcheckReplaceParamsCheckRegionEnam       HealthcheckReplaceParamsCheckRegion = "ENAM"
	HealthcheckReplaceParamsCheckRegionWeu        HealthcheckReplaceParamsCheckRegion = "WEU"
	HealthcheckReplaceParamsCheckRegionEeu        HealthcheckReplaceParamsCheckRegion = "EEU"
	HealthcheckReplaceParamsCheckRegionNsam       HealthcheckReplaceParamsCheckRegion = "NSAM"
	HealthcheckReplaceParamsCheckRegionSsam       HealthcheckReplaceParamsCheckRegion = "SSAM"
	HealthcheckReplaceParamsCheckRegionOc         HealthcheckReplaceParamsCheckRegion = "OC"
	HealthcheckReplaceParamsCheckRegionMe         HealthcheckReplaceParamsCheckRegion = "ME"
	HealthcheckReplaceParamsCheckRegionNaf        HealthcheckReplaceParamsCheckRegion = "NAF"
	HealthcheckReplaceParamsCheckRegionSaf        HealthcheckReplaceParamsCheckRegion = "SAF"
	HealthcheckReplaceParamsCheckRegionIn         HealthcheckReplaceParamsCheckRegion = "IN"
	HealthcheckReplaceParamsCheckRegionSeas       HealthcheckReplaceParamsCheckRegion = "SEAS"
	HealthcheckReplaceParamsCheckRegionNeas       HealthcheckReplaceParamsCheckRegion = "NEAS"
	HealthcheckReplaceParamsCheckRegionAllRegions HealthcheckReplaceParamsCheckRegion = "ALL_REGIONS"
)

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckReplaceParamsHTTPConfig struct {
	// Do not validate the certificate when the health check uses HTTPS.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// The expected HTTP response codes (e.g. "200") or code ranges (e.g. "2xx" for all
	// codes starting with 2) of the health check.
	ExpectedCodes param.Field[[]string] `json:"expected_codes"`
	// Follow redirects if the origin returns a 3xx status code.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden.
	Header param.Field[interface{}] `json:"header"`
	// The HTTP method to use for the health check.
	Method param.Field[HealthcheckReplaceParamsHTTPConfigMethod] `json:"method"`
	// The endpoint path to health check against.
	Path param.Field[string] `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port param.Field[int64] `json:"port"`
}

func (r HealthcheckReplaceParamsHTTPConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method to use for the health check.
type HealthcheckReplaceParamsHTTPConfigMethod string

const (
	HealthcheckReplaceParamsHTTPConfigMethodGet  HealthcheckReplaceParamsHTTPConfigMethod = "GET"
	HealthcheckReplaceParamsHTTPConfigMethodHead HealthcheckReplaceParamsHTTPConfigMethod = "HEAD"
)

// Parameters specific to TCP health check.
type HealthcheckReplaceParamsTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method param.Field[HealthcheckReplaceParamsTcpConfigMethod] `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port param.Field[int64] `json:"port"`
}

func (r HealthcheckReplaceParamsTcpConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TCP connection method to use for the health check.
type HealthcheckReplaceParamsTcpConfigMethod string

const (
	HealthcheckReplaceParamsTcpConfigMethodConnectionEstablished HealthcheckReplaceParamsTcpConfigMethod = "connection_established"
)

type HealthcheckReplaceResponseEnvelope struct {
	Errors   []HealthcheckReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HealthcheckReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   HealthcheckReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success HealthcheckReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    healthcheckReplaceResponseEnvelopeJSON    `json:"-"`
}

// healthcheckReplaceResponseEnvelopeJSON contains the JSON metadata for the struct
// [HealthcheckReplaceResponseEnvelope]
type healthcheckReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckReplaceResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    healthcheckReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// healthcheckReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [HealthcheckReplaceResponseEnvelopeErrors]
type healthcheckReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthcheckReplaceResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    healthcheckReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// healthcheckReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [HealthcheckReplaceResponseEnvelopeMessages]
type healthcheckReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HealthcheckReplaceResponseEnvelopeSuccess bool

const (
	HealthcheckReplaceResponseEnvelopeSuccessTrue HealthcheckReplaceResponseEnvelopeSuccess = true
)
