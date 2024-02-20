// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// PageProjectDomainService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPageProjectDomainService] method
// instead.
type PageProjectDomainService struct {
	Options []option.RequestOption
}

// NewPageProjectDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPageProjectDomainService(opts ...option.RequestOption) (r *PageProjectDomainService) {
	r = &PageProjectDomainService{}
	r.Options = opts
	return
}

// Add a new domain for the Pages project.
func (r *PageProjectDomainService) New(ctx context.Context, accountID string, projectName string, body PageProjectDomainNewParams, opts ...option.RequestOption) (res *PageProjectDomainNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDomainNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", accountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retry the validation status of a single domain.
func (r *PageProjectDomainService) Update(ctx context.Context, accountID string, projectName string, domainName string, opts ...option.RequestOption) (res *PageProjectDomainUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDomainUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", accountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of all domains associated with a Pages project.
func (r *PageProjectDomainService) List(ctx context.Context, accountID string, projectName string, opts ...option.RequestOption) (res *[]PageProjectDomainListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDomainListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", accountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Pages project's domain.
func (r *PageProjectDomainService) Delete(ctx context.Context, accountID string, projectName string, domainName string, opts ...option.RequestOption) (res *PageProjectDomainDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", accountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Fetch a single domain.
func (r *PageProjectDomainService) Get(ctx context.Context, accountID string, projectName string, domainName string, opts ...option.RequestOption) (res *PageProjectDomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageProjectDomainGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", accountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [PageProjectDomainNewResponseUnknown],
// [PageProjectDomainNewResponseArray] or [shared.UnionString].
type PageProjectDomainNewResponse interface {
	ImplementsPageProjectDomainNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageProjectDomainNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type PageProjectDomainNewResponseArray []interface{}

func (r PageProjectDomainNewResponseArray) ImplementsPageProjectDomainNewResponse() {}

// Union satisfied by [PageProjectDomainUpdateResponseUnknown],
// [PageProjectDomainUpdateResponseArray] or [shared.UnionString].
type PageProjectDomainUpdateResponse interface {
	ImplementsPageProjectDomainUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageProjectDomainUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type PageProjectDomainUpdateResponseArray []interface{}

func (r PageProjectDomainUpdateResponseArray) ImplementsPageProjectDomainUpdateResponse() {}

type PageProjectDomainListResponse = interface{}

type PageProjectDomainDeleteResponse = interface{}

// Union satisfied by [PageProjectDomainGetResponseUnknown],
// [PageProjectDomainGetResponseArray] or [shared.UnionString].
type PageProjectDomainGetResponse interface {
	ImplementsPageProjectDomainGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageProjectDomainGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type PageProjectDomainGetResponseArray []interface{}

func (r PageProjectDomainGetResponseArray) ImplementsPageProjectDomainGetResponse() {}

type PageProjectDomainNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r PageProjectDomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type PageProjectDomainNewResponseEnvelope struct {
	Errors   []PageProjectDomainNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDomainNewResponseEnvelopeMessages `json:"messages,required"`
	Result   PageProjectDomainNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success PageProjectDomainNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectDomainNewResponseEnvelopeJSON    `json:"-"`
}

// pageProjectDomainNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageProjectDomainNewResponseEnvelope]
type pageProjectDomainNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDomainNewResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    pageProjectDomainNewResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDomainNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [PageProjectDomainNewResponseEnvelopeErrors]
type pageProjectDomainNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDomainNewResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    pageProjectDomainNewResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDomainNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PageProjectDomainNewResponseEnvelopeMessages]
type pageProjectDomainNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectDomainNewResponseEnvelopeSuccess bool

const (
	PageProjectDomainNewResponseEnvelopeSuccessTrue PageProjectDomainNewResponseEnvelopeSuccess = true
)

type PageProjectDomainUpdateResponseEnvelope struct {
	Errors   []PageProjectDomainUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDomainUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   PageProjectDomainUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success PageProjectDomainUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectDomainUpdateResponseEnvelopeJSON    `json:"-"`
}

// pageProjectDomainUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageProjectDomainUpdateResponseEnvelope]
type pageProjectDomainUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDomainUpdateResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    pageProjectDomainUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDomainUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [PageProjectDomainUpdateResponseEnvelopeErrors]
type pageProjectDomainUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDomainUpdateResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    pageProjectDomainUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDomainUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [PageProjectDomainUpdateResponseEnvelopeMessages]
type pageProjectDomainUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectDomainUpdateResponseEnvelopeSuccess bool

const (
	PageProjectDomainUpdateResponseEnvelopeSuccessTrue PageProjectDomainUpdateResponseEnvelopeSuccess = true
)

type PageProjectDomainListResponseEnvelope struct {
	Errors   []PageProjectDomainListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDomainListResponseEnvelopeMessages `json:"messages,required"`
	Result   []PageProjectDomainListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success    PageProjectDomainListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PageProjectDomainListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       pageProjectDomainListResponseEnvelopeJSON       `json:"-"`
}

// pageProjectDomainListResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageProjectDomainListResponseEnvelope]
type pageProjectDomainListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDomainListResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    pageProjectDomainListResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDomainListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [PageProjectDomainListResponseEnvelopeErrors]
type pageProjectDomainListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDomainListResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    pageProjectDomainListResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDomainListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PageProjectDomainListResponseEnvelopeMessages]
type pageProjectDomainListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectDomainListResponseEnvelopeSuccess bool

const (
	PageProjectDomainListResponseEnvelopeSuccessTrue PageProjectDomainListResponseEnvelopeSuccess = true
)

type PageProjectDomainListResponseEnvelopeResultInfo struct {
	Count      interface{}                                         `json:"count"`
	Page       interface{}                                         `json:"page"`
	PerPage    interface{}                                         `json:"per_page"`
	TotalCount interface{}                                         `json:"total_count"`
	JSON       pageProjectDomainListResponseEnvelopeResultInfoJSON `json:"-"`
}

// pageProjectDomainListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [PageProjectDomainListResponseEnvelopeResultInfo]
type pageProjectDomainListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDomainGetResponseEnvelope struct {
	Errors   []PageProjectDomainGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageProjectDomainGetResponseEnvelopeMessages `json:"messages,required"`
	Result   PageProjectDomainGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success PageProjectDomainGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageProjectDomainGetResponseEnvelopeJSON    `json:"-"`
}

// pageProjectDomainGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageProjectDomainGetResponseEnvelope]
type pageProjectDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDomainGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    pageProjectDomainGetResponseEnvelopeErrorsJSON `json:"-"`
}

// pageProjectDomainGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [PageProjectDomainGetResponseEnvelopeErrors]
type pageProjectDomainGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageProjectDomainGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    pageProjectDomainGetResponseEnvelopeMessagesJSON `json:"-"`
}

// pageProjectDomainGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PageProjectDomainGetResponseEnvelopeMessages]
type pageProjectDomainGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageProjectDomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageProjectDomainGetResponseEnvelopeSuccess bool

const (
	PageProjectDomainGetResponseEnvelopeSuccessTrue PageProjectDomainGetResponseEnvelopeSuccess = true
)
