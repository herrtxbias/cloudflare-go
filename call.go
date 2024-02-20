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

// CallService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewCallService] method instead.
type CallService struct {
	Options []option.RequestOption
}

// NewCallService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCallService(opts ...option.RequestOption) (r *CallService) {
	r = &CallService{}
	r.Options = opts
	return
}

// Creates a new Cloudflare calls app. An app is an unique enviroment where each
// Session can access all Tracks within the app.
func (r *CallService) New(ctx context.Context, accountID string, body CallNewParams, opts ...option.RequestOption) (res *CallNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CallNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/calls/apps", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all apps in the Cloudflare account
func (r *CallService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]CallListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CallListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/calls/apps", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an app from Cloudflare Calls
func (r *CallService) Delete(ctx context.Context, accountID string, appID string, opts ...option.RequestOption) (res *CallDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CallDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/calls/apps/%s", accountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches details for a single Calls app.
func (r *CallService) Get(ctx context.Context, accountID string, appID string, opts ...option.RequestOption) (res *CallGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CallGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/calls/apps/%s", accountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edit details for a single app.
func (r *CallService) Replace(ctx context.Context, accountID string, appID string, body CallReplaceParams, opts ...option.RequestOption) (res *CallReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CallReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%s/calls/apps/%s", accountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CallNewResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// Bearer token to use the Calls API.
	Secret string `json:"secret"`
	// A Cloudflare-generated unique identifier for a item.
	Uid  string              `json:"uid"`
	JSON callNewResponseJSON `json:"-"`
}

// callNewResponseJSON contains the JSON metadata for the struct [CallNewResponse]
type callNewResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CallListResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	Uid  string               `json:"uid"`
	JSON callListResponseJSON `json:"-"`
}

// callListResponseJSON contains the JSON metadata for the struct
// [CallListResponse]
type callListResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CallDeleteResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	Uid  string                 `json:"uid"`
	JSON callDeleteResponseJSON `json:"-"`
}

// callDeleteResponseJSON contains the JSON metadata for the struct
// [CallDeleteResponse]
type callDeleteResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CallGetResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	Uid  string              `json:"uid"`
	JSON callGetResponseJSON `json:"-"`
}

// callGetResponseJSON contains the JSON metadata for the struct [CallGetResponse]
type callGetResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CallReplaceResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	Uid  string                  `json:"uid"`
	JSON callReplaceResponseJSON `json:"-"`
}

// callReplaceResponseJSON contains the JSON metadata for the struct
// [CallReplaceResponse]
type callReplaceResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CallNewParams struct {
	// A short description of Calls app, not shown to end users.
	Name param.Field[string] `json:"name"`
}

func (r CallNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CallNewResponseEnvelope struct {
	Errors   []CallNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CallNewResponseEnvelopeMessages `json:"messages,required"`
	Result   CallNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CallNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    callNewResponseEnvelopeJSON    `json:"-"`
}

// callNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CallNewResponseEnvelope]
type callNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CallNewResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    callNewResponseEnvelopeErrorsJSON `json:"-"`
}

// callNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CallNewResponseEnvelopeErrors]
type callNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CallNewResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    callNewResponseEnvelopeMessagesJSON `json:"-"`
}

// callNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [CallNewResponseEnvelopeMessages]
type callNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CallNewResponseEnvelopeSuccess bool

const (
	CallNewResponseEnvelopeSuccessTrue CallNewResponseEnvelopeSuccess = true
)

type CallListResponseEnvelope struct {
	Errors   []CallListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CallListResponseEnvelopeMessages `json:"messages,required"`
	Result   []CallListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success CallListResponseEnvelopeSuccess `json:"success,required"`
	JSON    callListResponseEnvelopeJSON    `json:"-"`
}

// callListResponseEnvelopeJSON contains the JSON metadata for the struct
// [CallListResponseEnvelope]
type callListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CallListResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    callListResponseEnvelopeErrorsJSON `json:"-"`
}

// callListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CallListResponseEnvelopeErrors]
type callListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CallListResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    callListResponseEnvelopeMessagesJSON `json:"-"`
}

// callListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [CallListResponseEnvelopeMessages]
type callListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CallListResponseEnvelopeSuccess bool

const (
	CallListResponseEnvelopeSuccessTrue CallListResponseEnvelopeSuccess = true
)

type CallDeleteResponseEnvelope struct {
	Errors   []CallDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CallDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CallDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CallDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    callDeleteResponseEnvelopeJSON    `json:"-"`
}

// callDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [CallDeleteResponseEnvelope]
type callDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CallDeleteResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    callDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// callDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CallDeleteResponseEnvelopeErrors]
type callDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CallDeleteResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    callDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// callDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [CallDeleteResponseEnvelopeMessages]
type callDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CallDeleteResponseEnvelopeSuccess bool

const (
	CallDeleteResponseEnvelopeSuccessTrue CallDeleteResponseEnvelopeSuccess = true
)

type CallGetResponseEnvelope struct {
	Errors   []CallGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CallGetResponseEnvelopeMessages `json:"messages,required"`
	Result   CallGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CallGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    callGetResponseEnvelopeJSON    `json:"-"`
}

// callGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CallGetResponseEnvelope]
type callGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CallGetResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    callGetResponseEnvelopeErrorsJSON `json:"-"`
}

// callGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CallGetResponseEnvelopeErrors]
type callGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CallGetResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    callGetResponseEnvelopeMessagesJSON `json:"-"`
}

// callGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [CallGetResponseEnvelopeMessages]
type callGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CallGetResponseEnvelopeSuccess bool

const (
	CallGetResponseEnvelopeSuccessTrue CallGetResponseEnvelopeSuccess = true
)

type CallReplaceParams struct {
	// A short description of Calls app, not shown to end users.
	Name param.Field[string] `json:"name"`
}

func (r CallReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CallReplaceResponseEnvelope struct {
	Errors   []CallReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CallReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   CallReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CallReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    callReplaceResponseEnvelopeJSON    `json:"-"`
}

// callReplaceResponseEnvelopeJSON contains the JSON metadata for the struct
// [CallReplaceResponseEnvelope]
type callReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CallReplaceResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    callReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// callReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CallReplaceResponseEnvelopeErrors]
type callReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CallReplaceResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    callReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// callReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CallReplaceResponseEnvelopeMessages]
type callReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CallReplaceResponseEnvelopeSuccess bool

const (
	CallReplaceResponseEnvelopeSuccessTrue CallReplaceResponseEnvelopeSuccess = true
)
