// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ClientCertificateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewClientCertificateService] method
// instead.
type ClientCertificateService struct {
	Options []option.RequestOption
}

// NewClientCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewClientCertificateService(opts ...option.RequestOption) (r *ClientCertificateService) {
	r = &ClientCertificateService{}
	r.Options = opts
	return
}

// Create a new API Shield mTLS Client Certificate
func (r *ClientCertificateService) New(ctx context.Context, zoneID string, body ClientCertificateNewParams, opts ...option.RequestOption) (res *ClientCertificateNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ClientCertificateNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/client_certificates", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// If a API Shield mTLS Client Certificate is in a pending_revocation state, you
// may reactivate it with this endpoint.
func (r *ClientCertificateService) Update(ctx context.Context, zoneID string, clientCertificateID string, opts ...option.RequestOption) (res *ClientCertificateUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ClientCertificateUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/client_certificates/%s", zoneID, clientCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all of your Zone's API Shield mTLS Client Certificates by Status and/or
// using Pagination
func (r *ClientCertificateService) List(ctx context.Context, zoneID string, query ClientCertificateListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[ClientCertificateListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/client_certificates", zoneID)
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

// List all of your Zone's API Shield mTLS Client Certificates by Status and/or
// using Pagination
func (r *ClientCertificateService) ListAutoPaging(ctx context.Context, zoneID string, query ClientCertificateListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[ClientCertificateListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneID, query, opts...))
}

// Set a API Shield mTLS Client Certificate to pending_revocation status for
// processing to revoked status.
func (r *ClientCertificateService) Delete(ctx context.Context, zoneID string, clientCertificateID string, opts ...option.RequestOption) (res *ClientCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ClientCertificateDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/client_certificates/%s", zoneID, clientCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Details for a single mTLS API Shield Client Certificate
func (r *ClientCertificateService) Get(ctx context.Context, zoneID string, clientCertificateID string, opts ...option.RequestOption) (res *ClientCertificateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ClientCertificateGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/client_certificates/%s", zoneID, clientCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ClientCertificateNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ClientCertificateNewResponseCertificateAuthority `json:"certificate_authority"`
	// Common Name of the Client Certificate
	CommonName string `json:"common_name"`
	// Country, provided by the CSR
	Country string `json:"country"`
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr string `json:"csr"`
	// Date that the Client Certificate expires
	ExpiresOn string `json:"expires_on"`
	// Unique identifier of the Client Certificate
	FingerprintSha256 string `json:"fingerprint_sha256"`
	// Date that the Client Certificate was issued by the Certificate Authority
	IssuedOn string `json:"issued_on"`
	// Location, provided by the CSR
	Location string `json:"location"`
	// Organization, provided by the CSR
	Organization string `json:"organization"`
	// Organizational Unit, provided by the CSR
	OrganizationalUnit string `json:"organizational_unit"`
	// The serial number on the created Client Certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the Client Certificate..
	Signature string `json:"signature"`
	// Subject Key Identifier
	Ski string `json:"ski"`
	// State, provided by the CSR
	State string `json:"state"`
	// Client Certificates may be active or revoked, and the pending_reactivation or
	// pending_revocation represent in-progress asynchronous transitions
	Status ClientCertificateNewResponseStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                            `json:"validity_days"`
	JSON         clientCertificateNewResponseJSON `json:"-"`
}

// clientCertificateNewResponseJSON contains the JSON metadata for the struct
// [ClientCertificateNewResponse]
type clientCertificateNewResponseJSON struct {
	ID                   apijson.Field
	Certificate          apijson.Field
	CertificateAuthority apijson.Field
	CommonName           apijson.Field
	Country              apijson.Field
	Csr                  apijson.Field
	ExpiresOn            apijson.Field
	FingerprintSha256    apijson.Field
	IssuedOn             apijson.Field
	Location             apijson.Field
	Organization         apijson.Field
	OrganizationalUnit   apijson.Field
	SerialNumber         apijson.Field
	Signature            apijson.Field
	Ski                  apijson.Field
	State                apijson.Field
	Status               apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ClientCertificateNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority used to issue the Client Certificate
type ClientCertificateNewResponseCertificateAuthority struct {
	ID   string                                               `json:"id"`
	Name string                                               `json:"name"`
	JSON clientCertificateNewResponseCertificateAuthorityJSON `json:"-"`
}

// clientCertificateNewResponseCertificateAuthorityJSON contains the JSON metadata
// for the struct [ClientCertificateNewResponseCertificateAuthority]
type clientCertificateNewResponseCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateNewResponseCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ClientCertificateNewResponseStatus string

const (
	ClientCertificateNewResponseStatusActive              ClientCertificateNewResponseStatus = "active"
	ClientCertificateNewResponseStatusPendingReactivation ClientCertificateNewResponseStatus = "pending_reactivation"
	ClientCertificateNewResponseStatusPendingRevocation   ClientCertificateNewResponseStatus = "pending_revocation"
	ClientCertificateNewResponseStatusRevoked             ClientCertificateNewResponseStatus = "revoked"
)

type ClientCertificateUpdateResponse struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ClientCertificateUpdateResponseCertificateAuthority `json:"certificate_authority"`
	// Common Name of the Client Certificate
	CommonName string `json:"common_name"`
	// Country, provided by the CSR
	Country string `json:"country"`
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr string `json:"csr"`
	// Date that the Client Certificate expires
	ExpiresOn string `json:"expires_on"`
	// Unique identifier of the Client Certificate
	FingerprintSha256 string `json:"fingerprint_sha256"`
	// Date that the Client Certificate was issued by the Certificate Authority
	IssuedOn string `json:"issued_on"`
	// Location, provided by the CSR
	Location string `json:"location"`
	// Organization, provided by the CSR
	Organization string `json:"organization"`
	// Organizational Unit, provided by the CSR
	OrganizationalUnit string `json:"organizational_unit"`
	// The serial number on the created Client Certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the Client Certificate..
	Signature string `json:"signature"`
	// Subject Key Identifier
	Ski string `json:"ski"`
	// State, provided by the CSR
	State string `json:"state"`
	// Client Certificates may be active or revoked, and the pending_reactivation or
	// pending_revocation represent in-progress asynchronous transitions
	Status ClientCertificateUpdateResponseStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                               `json:"validity_days"`
	JSON         clientCertificateUpdateResponseJSON `json:"-"`
}

// clientCertificateUpdateResponseJSON contains the JSON metadata for the struct
// [ClientCertificateUpdateResponse]
type clientCertificateUpdateResponseJSON struct {
	ID                   apijson.Field
	Certificate          apijson.Field
	CertificateAuthority apijson.Field
	CommonName           apijson.Field
	Country              apijson.Field
	Csr                  apijson.Field
	ExpiresOn            apijson.Field
	FingerprintSha256    apijson.Field
	IssuedOn             apijson.Field
	Location             apijson.Field
	Organization         apijson.Field
	OrganizationalUnit   apijson.Field
	SerialNumber         apijson.Field
	Signature            apijson.Field
	Ski                  apijson.Field
	State                apijson.Field
	Status               apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ClientCertificateUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority used to issue the Client Certificate
type ClientCertificateUpdateResponseCertificateAuthority struct {
	ID   string                                                  `json:"id"`
	Name string                                                  `json:"name"`
	JSON clientCertificateUpdateResponseCertificateAuthorityJSON `json:"-"`
}

// clientCertificateUpdateResponseCertificateAuthorityJSON contains the JSON
// metadata for the struct [ClientCertificateUpdateResponseCertificateAuthority]
type clientCertificateUpdateResponseCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateUpdateResponseCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ClientCertificateUpdateResponseStatus string

const (
	ClientCertificateUpdateResponseStatusActive              ClientCertificateUpdateResponseStatus = "active"
	ClientCertificateUpdateResponseStatusPendingReactivation ClientCertificateUpdateResponseStatus = "pending_reactivation"
	ClientCertificateUpdateResponseStatusPendingRevocation   ClientCertificateUpdateResponseStatus = "pending_revocation"
	ClientCertificateUpdateResponseStatusRevoked             ClientCertificateUpdateResponseStatus = "revoked"
)

type ClientCertificateListResponse struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ClientCertificateListResponseCertificateAuthority `json:"certificate_authority"`
	// Common Name of the Client Certificate
	CommonName string `json:"common_name"`
	// Country, provided by the CSR
	Country string `json:"country"`
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr string `json:"csr"`
	// Date that the Client Certificate expires
	ExpiresOn string `json:"expires_on"`
	// Unique identifier of the Client Certificate
	FingerprintSha256 string `json:"fingerprint_sha256"`
	// Date that the Client Certificate was issued by the Certificate Authority
	IssuedOn string `json:"issued_on"`
	// Location, provided by the CSR
	Location string `json:"location"`
	// Organization, provided by the CSR
	Organization string `json:"organization"`
	// Organizational Unit, provided by the CSR
	OrganizationalUnit string `json:"organizational_unit"`
	// The serial number on the created Client Certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the Client Certificate..
	Signature string `json:"signature"`
	// Subject Key Identifier
	Ski string `json:"ski"`
	// State, provided by the CSR
	State string `json:"state"`
	// Client Certificates may be active or revoked, and the pending_reactivation or
	// pending_revocation represent in-progress asynchronous transitions
	Status ClientCertificateListResponseStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                             `json:"validity_days"`
	JSON         clientCertificateListResponseJSON `json:"-"`
}

// clientCertificateListResponseJSON contains the JSON metadata for the struct
// [ClientCertificateListResponse]
type clientCertificateListResponseJSON struct {
	ID                   apijson.Field
	Certificate          apijson.Field
	CertificateAuthority apijson.Field
	CommonName           apijson.Field
	Country              apijson.Field
	Csr                  apijson.Field
	ExpiresOn            apijson.Field
	FingerprintSha256    apijson.Field
	IssuedOn             apijson.Field
	Location             apijson.Field
	Organization         apijson.Field
	OrganizationalUnit   apijson.Field
	SerialNumber         apijson.Field
	Signature            apijson.Field
	Ski                  apijson.Field
	State                apijson.Field
	Status               apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ClientCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority used to issue the Client Certificate
type ClientCertificateListResponseCertificateAuthority struct {
	ID   string                                                `json:"id"`
	Name string                                                `json:"name"`
	JSON clientCertificateListResponseCertificateAuthorityJSON `json:"-"`
}

// clientCertificateListResponseCertificateAuthorityJSON contains the JSON metadata
// for the struct [ClientCertificateListResponseCertificateAuthority]
type clientCertificateListResponseCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateListResponseCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ClientCertificateListResponseStatus string

const (
	ClientCertificateListResponseStatusActive              ClientCertificateListResponseStatus = "active"
	ClientCertificateListResponseStatusPendingReactivation ClientCertificateListResponseStatus = "pending_reactivation"
	ClientCertificateListResponseStatusPendingRevocation   ClientCertificateListResponseStatus = "pending_revocation"
	ClientCertificateListResponseStatusRevoked             ClientCertificateListResponseStatus = "revoked"
)

type ClientCertificateDeleteResponse struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ClientCertificateDeleteResponseCertificateAuthority `json:"certificate_authority"`
	// Common Name of the Client Certificate
	CommonName string `json:"common_name"`
	// Country, provided by the CSR
	Country string `json:"country"`
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr string `json:"csr"`
	// Date that the Client Certificate expires
	ExpiresOn string `json:"expires_on"`
	// Unique identifier of the Client Certificate
	FingerprintSha256 string `json:"fingerprint_sha256"`
	// Date that the Client Certificate was issued by the Certificate Authority
	IssuedOn string `json:"issued_on"`
	// Location, provided by the CSR
	Location string `json:"location"`
	// Organization, provided by the CSR
	Organization string `json:"organization"`
	// Organizational Unit, provided by the CSR
	OrganizationalUnit string `json:"organizational_unit"`
	// The serial number on the created Client Certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the Client Certificate..
	Signature string `json:"signature"`
	// Subject Key Identifier
	Ski string `json:"ski"`
	// State, provided by the CSR
	State string `json:"state"`
	// Client Certificates may be active or revoked, and the pending_reactivation or
	// pending_revocation represent in-progress asynchronous transitions
	Status ClientCertificateDeleteResponseStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                               `json:"validity_days"`
	JSON         clientCertificateDeleteResponseJSON `json:"-"`
}

// clientCertificateDeleteResponseJSON contains the JSON metadata for the struct
// [ClientCertificateDeleteResponse]
type clientCertificateDeleteResponseJSON struct {
	ID                   apijson.Field
	Certificate          apijson.Field
	CertificateAuthority apijson.Field
	CommonName           apijson.Field
	Country              apijson.Field
	Csr                  apijson.Field
	ExpiresOn            apijson.Field
	FingerprintSha256    apijson.Field
	IssuedOn             apijson.Field
	Location             apijson.Field
	Organization         apijson.Field
	OrganizationalUnit   apijson.Field
	SerialNumber         apijson.Field
	Signature            apijson.Field
	Ski                  apijson.Field
	State                apijson.Field
	Status               apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ClientCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority used to issue the Client Certificate
type ClientCertificateDeleteResponseCertificateAuthority struct {
	ID   string                                                  `json:"id"`
	Name string                                                  `json:"name"`
	JSON clientCertificateDeleteResponseCertificateAuthorityJSON `json:"-"`
}

// clientCertificateDeleteResponseCertificateAuthorityJSON contains the JSON
// metadata for the struct [ClientCertificateDeleteResponseCertificateAuthority]
type clientCertificateDeleteResponseCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateDeleteResponseCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ClientCertificateDeleteResponseStatus string

const (
	ClientCertificateDeleteResponseStatusActive              ClientCertificateDeleteResponseStatus = "active"
	ClientCertificateDeleteResponseStatusPendingReactivation ClientCertificateDeleteResponseStatus = "pending_reactivation"
	ClientCertificateDeleteResponseStatusPendingRevocation   ClientCertificateDeleteResponseStatus = "pending_revocation"
	ClientCertificateDeleteResponseStatusRevoked             ClientCertificateDeleteResponseStatus = "revoked"
)

type ClientCertificateGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ClientCertificateGetResponseCertificateAuthority `json:"certificate_authority"`
	// Common Name of the Client Certificate
	CommonName string `json:"common_name"`
	// Country, provided by the CSR
	Country string `json:"country"`
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr string `json:"csr"`
	// Date that the Client Certificate expires
	ExpiresOn string `json:"expires_on"`
	// Unique identifier of the Client Certificate
	FingerprintSha256 string `json:"fingerprint_sha256"`
	// Date that the Client Certificate was issued by the Certificate Authority
	IssuedOn string `json:"issued_on"`
	// Location, provided by the CSR
	Location string `json:"location"`
	// Organization, provided by the CSR
	Organization string `json:"organization"`
	// Organizational Unit, provided by the CSR
	OrganizationalUnit string `json:"organizational_unit"`
	// The serial number on the created Client Certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the Client Certificate..
	Signature string `json:"signature"`
	// Subject Key Identifier
	Ski string `json:"ski"`
	// State, provided by the CSR
	State string `json:"state"`
	// Client Certificates may be active or revoked, and the pending_reactivation or
	// pending_revocation represent in-progress asynchronous transitions
	Status ClientCertificateGetResponseStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                            `json:"validity_days"`
	JSON         clientCertificateGetResponseJSON `json:"-"`
}

// clientCertificateGetResponseJSON contains the JSON metadata for the struct
// [ClientCertificateGetResponse]
type clientCertificateGetResponseJSON struct {
	ID                   apijson.Field
	Certificate          apijson.Field
	CertificateAuthority apijson.Field
	CommonName           apijson.Field
	Country              apijson.Field
	Csr                  apijson.Field
	ExpiresOn            apijson.Field
	FingerprintSha256    apijson.Field
	IssuedOn             apijson.Field
	Location             apijson.Field
	Organization         apijson.Field
	OrganizationalUnit   apijson.Field
	SerialNumber         apijson.Field
	Signature            apijson.Field
	Ski                  apijson.Field
	State                apijson.Field
	Status               apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ClientCertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate Authority used to issue the Client Certificate
type ClientCertificateGetResponseCertificateAuthority struct {
	ID   string                                               `json:"id"`
	Name string                                               `json:"name"`
	JSON clientCertificateGetResponseCertificateAuthorityJSON `json:"-"`
}

// clientCertificateGetResponseCertificateAuthorityJSON contains the JSON metadata
// for the struct [ClientCertificateGetResponseCertificateAuthority]
type clientCertificateGetResponseCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateGetResponseCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ClientCertificateGetResponseStatus string

const (
	ClientCertificateGetResponseStatusActive              ClientCertificateGetResponseStatus = "active"
	ClientCertificateGetResponseStatusPendingReactivation ClientCertificateGetResponseStatus = "pending_reactivation"
	ClientCertificateGetResponseStatusPendingRevocation   ClientCertificateGetResponseStatus = "pending_revocation"
	ClientCertificateGetResponseStatusRevoked             ClientCertificateGetResponseStatus = "revoked"
)

type ClientCertificateNewParams struct {
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr param.Field[string] `json:"csr,required"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays param.Field[int64] `json:"validity_days,required"`
}

func (r ClientCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ClientCertificateNewResponseEnvelope struct {
	Errors   []ClientCertificateNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ClientCertificateNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ClientCertificateNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ClientCertificateNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    clientCertificateNewResponseEnvelopeJSON    `json:"-"`
}

// clientCertificateNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ClientCertificateNewResponseEnvelope]
type clientCertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClientCertificateNewResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    clientCertificateNewResponseEnvelopeErrorsJSON `json:"-"`
}

// clientCertificateNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ClientCertificateNewResponseEnvelopeErrors]
type clientCertificateNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClientCertificateNewResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    clientCertificateNewResponseEnvelopeMessagesJSON `json:"-"`
}

// clientCertificateNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ClientCertificateNewResponseEnvelopeMessages]
type clientCertificateNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ClientCertificateNewResponseEnvelopeSuccess bool

const (
	ClientCertificateNewResponseEnvelopeSuccessTrue ClientCertificateNewResponseEnvelopeSuccess = true
)

type ClientCertificateUpdateResponseEnvelope struct {
	Errors   []ClientCertificateUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ClientCertificateUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ClientCertificateUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ClientCertificateUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    clientCertificateUpdateResponseEnvelopeJSON    `json:"-"`
}

// clientCertificateUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ClientCertificateUpdateResponseEnvelope]
type clientCertificateUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClientCertificateUpdateResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    clientCertificateUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// clientCertificateUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ClientCertificateUpdateResponseEnvelopeErrors]
type clientCertificateUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClientCertificateUpdateResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    clientCertificateUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// clientCertificateUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ClientCertificateUpdateResponseEnvelopeMessages]
type clientCertificateUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ClientCertificateUpdateResponseEnvelopeSuccess bool

const (
	ClientCertificateUpdateResponseEnvelopeSuccessTrue ClientCertificateUpdateResponseEnvelopeSuccess = true
)

type ClientCertificateListParams struct {
	// Limit to the number of records returned.
	Limit param.Field[int64] `query:"limit"`
	// Offset the results
	Offset param.Field[int64] `query:"offset"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of records per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Client Certitifcate Status to filter results by.
	Status param.Field[ClientCertificateListParamsStatus] `query:"status"`
}

// URLQuery serializes [ClientCertificateListParams]'s query parameters as
// `url.Values`.
func (r ClientCertificateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Client Certitifcate Status to filter results by.
type ClientCertificateListParamsStatus string

const (
	ClientCertificateListParamsStatusAll                 ClientCertificateListParamsStatus = "all"
	ClientCertificateListParamsStatusActive              ClientCertificateListParamsStatus = "active"
	ClientCertificateListParamsStatusPendingReactivation ClientCertificateListParamsStatus = "pending_reactivation"
	ClientCertificateListParamsStatusPendingRevocation   ClientCertificateListParamsStatus = "pending_revocation"
	ClientCertificateListParamsStatusRevoked             ClientCertificateListParamsStatus = "revoked"
)

type ClientCertificateDeleteResponseEnvelope struct {
	Errors   []ClientCertificateDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ClientCertificateDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ClientCertificateDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ClientCertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    clientCertificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// clientCertificateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ClientCertificateDeleteResponseEnvelope]
type clientCertificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClientCertificateDeleteResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    clientCertificateDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// clientCertificateDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ClientCertificateDeleteResponseEnvelopeErrors]
type clientCertificateDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClientCertificateDeleteResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    clientCertificateDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// clientCertificateDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ClientCertificateDeleteResponseEnvelopeMessages]
type clientCertificateDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ClientCertificateDeleteResponseEnvelopeSuccess bool

const (
	ClientCertificateDeleteResponseEnvelopeSuccessTrue ClientCertificateDeleteResponseEnvelopeSuccess = true
)

type ClientCertificateGetResponseEnvelope struct {
	Errors   []ClientCertificateGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ClientCertificateGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ClientCertificateGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ClientCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    clientCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// clientCertificateGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ClientCertificateGetResponseEnvelope]
type clientCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClientCertificateGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    clientCertificateGetResponseEnvelopeErrorsJSON `json:"-"`
}

// clientCertificateGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ClientCertificateGetResponseEnvelopeErrors]
type clientCertificateGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ClientCertificateGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    clientCertificateGetResponseEnvelopeMessagesJSON `json:"-"`
}

// clientCertificateGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ClientCertificateGetResponseEnvelopeMessages]
type clientCertificateGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ClientCertificateGetResponseEnvelopeSuccess bool

const (
	ClientCertificateGetResponseEnvelopeSuccessTrue ClientCertificateGetResponseEnvelopeSuccess = true
)
