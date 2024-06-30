// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_ca_certificates

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/cloudflare/cloudflare-go/v2/ssl"
	"github.com/tidwall/gjson"
)

// OriginCACertificateService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOriginCACertificateService] method instead.
type OriginCACertificateService struct {
	Options []option.RequestOption
}

// NewOriginCACertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOriginCACertificateService(opts ...option.RequestOption) (r *OriginCACertificateService) {
	r = &OriginCACertificateService{}
	r.Options = opts
	return
}

// Create an Origin CA certificate. Use your Origin CA Key as your User Service Key
// when calling this endpoint ([see above](#requests)).
func (r *OriginCACertificateService) New(ctx context.Context, body OriginCACertificateNewParams, opts ...option.RequestOption) (res *OriginCACertificateNewResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginCACertificateNewResponseEnvelope
	path := "certificates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all existing Origin CA certificates for a given zone. Use your Origin CA
// Key as your User Service Key when calling this endpoint
// ([see above](#requests)).
func (r *OriginCACertificateService) List(ctx context.Context, query OriginCACertificateListParams, opts ...option.RequestOption) (res *pagination.SinglePage[OriginCACertificate], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "certificates"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all existing Origin CA certificates for a given zone. Use your Origin CA
// Key as your User Service Key when calling this endpoint
// ([see above](#requests)).
func (r *OriginCACertificateService) ListAutoPaging(ctx context.Context, query OriginCACertificateListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[OriginCACertificate] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Revoke an existing Origin CA certificate by its serial number. Use your Origin
// CA Key as your User Service Key when calling this endpoint
// ([see above](#requests)).
func (r *OriginCACertificateService) Delete(ctx context.Context, certificateID string, opts ...option.RequestOption) (res *OriginCACertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginCACertificateDeleteResponseEnvelope
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("certificates/%s", certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get an existing Origin CA certificate by its serial number. Use your Origin CA
// Key as your User Service Key when calling this endpoint
// ([see above](#requests)).
func (r *OriginCACertificateService) Get(ctx context.Context, certificateID string, opts ...option.RequestOption) (res *OriginCACertificateGetResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginCACertificateGetResponseEnvelope
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("certificates/%s", certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OriginCACertificate struct {
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr string `json:"csr,required"`
	// Array of hostnames or wildcard names (e.g., \*.example.com) bound to the
	// certificate.
	Hostnames []interface{} `json:"hostnames,required"`
	// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
	// or "keyless-certificate" (for Keyless SSL servers).
	RequestType shared.CertificateRequestType `json:"request_type,required"`
	// The number of days for which the certificate should be valid.
	RequestedValidity ssl.RequestValidity `json:"requested_validity,required"`
	// Identifier
	ID string `json:"id"`
	// The Origin CA certificate. Will be newline-encoded.
	Certificate string `json:"certificate"`
	// When the certificate will expire.
	ExpiresOn time.Time               `json:"expires_on" format:"date-time"`
	JSON      originCACertificateJSON `json:"-"`
}

// originCACertificateJSON contains the JSON metadata for the struct
// [OriginCACertificate]
type originCACertificateJSON struct {
	Csr               apijson.Field
	Hostnames         apijson.Field
	RequestType       apijson.Field
	RequestedValidity apijson.Field
	ID                apijson.Field
	Certificate       apijson.Field
	ExpiresOn         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OriginCACertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [origin_ca_certificates.OriginCACertificateNewResponseUnknown] or
// [shared.UnionString].
type OriginCACertificateNewResponseUnion interface {
	ImplementsOriginCACertificatesOriginCACertificateNewResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginCACertificateNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type OriginCACertificateDeleteResponse struct {
	// Identifier
	ID   string                                `json:"id"`
	JSON originCACertificateDeleteResponseJSON `json:"-"`
}

// originCACertificateDeleteResponseJSON contains the JSON metadata for the struct
// [OriginCACertificateDeleteResponse]
type originCACertificateDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCACertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [origin_ca_certificates.OriginCACertificateGetResponseUnknown] or
// [shared.UnionString].
type OriginCACertificateGetResponseUnion interface {
	ImplementsOriginCACertificatesOriginCACertificateGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginCACertificateGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type OriginCACertificateNewParams struct {
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr param.Field[string] `json:"csr"`
	// Array of hostnames or wildcard names (e.g., \*.example.com) bound to the
	// certificate.
	Hostnames param.Field[[]interface{}] `json:"hostnames"`
	// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
	// or "keyless-certificate" (for Keyless SSL servers).
	RequestType param.Field[shared.CertificateRequestType] `json:"request_type"`
	// The number of days for which the certificate should be valid.
	RequestedValidity param.Field[ssl.RequestValidity] `json:"requested_validity"`
}

func (r OriginCACertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OriginCACertificateNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OriginCACertificateNewResponseEnvelopeSuccess `json:"success,required"`
	Result  OriginCACertificateNewResponseUnion           `json:"result"`
	JSON    originCACertificateNewResponseEnvelopeJSON    `json:"-"`
}

// originCACertificateNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginCACertificateNewResponseEnvelope]
type originCACertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCACertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginCACertificateNewResponseEnvelopeSuccess bool

const (
	OriginCACertificateNewResponseEnvelopeSuccessTrue OriginCACertificateNewResponseEnvelopeSuccess = true
)

func (r OriginCACertificateNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginCACertificateNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginCACertificateListParams struct {
	// Identifier
	ZoneID param.Field[string] `query:"zone_id"`
}

// URLQuery serializes [OriginCACertificateListParams]'s query parameters as
// `url.Values`.
func (r OriginCACertificateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type OriginCACertificateDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OriginCACertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  OriginCACertificateDeleteResponse                `json:"result"`
	JSON    originCACertificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// originCACertificateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginCACertificateDeleteResponseEnvelope]
type originCACertificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCACertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginCACertificateDeleteResponseEnvelopeSuccess bool

const (
	OriginCACertificateDeleteResponseEnvelopeSuccessTrue OriginCACertificateDeleteResponseEnvelopeSuccess = true
)

func (r OriginCACertificateDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginCACertificateDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginCACertificateGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OriginCACertificateGetResponseEnvelopeSuccess `json:"success,required"`
	Result  OriginCACertificateGetResponseUnion           `json:"result"`
	JSON    originCACertificateGetResponseEnvelopeJSON    `json:"-"`
}

// originCACertificateGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginCACertificateGetResponseEnvelope]
type originCACertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCACertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCACertificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginCACertificateGetResponseEnvelopeSuccess bool

const (
	OriginCACertificateGetResponseEnvelopeSuccessTrue OriginCACertificateGetResponseEnvelopeSuccess = true
)

func (r OriginCACertificateGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginCACertificateGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
