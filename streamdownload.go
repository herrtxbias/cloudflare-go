// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// StreamDownloadService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStreamDownloadService] method
// instead.
type StreamDownloadService struct {
	Options []option.RequestOption
}

// NewStreamDownloadService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStreamDownloadService(opts ...option.RequestOption) (r *StreamDownloadService) {
	r = &StreamDownloadService{}
	r.Options = opts
	return
}

// Creates a download for a video when a video is ready to view.
func (r *StreamDownloadService) New(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *StreamDownloadNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamDownloadNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the downloads created for a video.
func (r *StreamDownloadService) List(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *StreamDownloadListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamDownloadListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete the downloads for a video.
func (r *StreamDownloadService) Delete(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *StreamDownloadDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamDownloadDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [StreamDownloadNewResponseUnknown] or [shared.UnionString].
type StreamDownloadNewResponse interface {
	ImplementsStreamDownloadNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamDownloadNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [StreamDownloadListResponseUnknown] or [shared.UnionString].
type StreamDownloadListResponse interface {
	ImplementsStreamDownloadListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamDownloadListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [StreamDownloadDeleteResponseUnknown] or
// [shared.UnionString].
type StreamDownloadDeleteResponse interface {
	ImplementsStreamDownloadDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamDownloadDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type StreamDownloadNewResponseEnvelope struct {
	Errors   []StreamDownloadNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamDownloadNewResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamDownloadNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamDownloadNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamDownloadNewResponseEnvelopeJSON    `json:"-"`
}

// streamDownloadNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamDownloadNewResponseEnvelope]
type streamDownloadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamDownloadNewResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    streamDownloadNewResponseEnvelopeErrorsJSON `json:"-"`
}

// streamDownloadNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamDownloadNewResponseEnvelopeErrors]
type streamDownloadNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamDownloadNewResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    streamDownloadNewResponseEnvelopeMessagesJSON `json:"-"`
}

// streamDownloadNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [StreamDownloadNewResponseEnvelopeMessages]
type streamDownloadNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamDownloadNewResponseEnvelopeSuccess bool

const (
	StreamDownloadNewResponseEnvelopeSuccessTrue StreamDownloadNewResponseEnvelopeSuccess = true
)

type StreamDownloadListResponseEnvelope struct {
	Errors   []StreamDownloadListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamDownloadListResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamDownloadListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamDownloadListResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamDownloadListResponseEnvelopeJSON    `json:"-"`
}

// streamDownloadListResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamDownloadListResponseEnvelope]
type streamDownloadListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamDownloadListResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    streamDownloadListResponseEnvelopeErrorsJSON `json:"-"`
}

// streamDownloadListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamDownloadListResponseEnvelopeErrors]
type streamDownloadListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamDownloadListResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    streamDownloadListResponseEnvelopeMessagesJSON `json:"-"`
}

// streamDownloadListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamDownloadListResponseEnvelopeMessages]
type streamDownloadListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamDownloadListResponseEnvelopeSuccess bool

const (
	StreamDownloadListResponseEnvelopeSuccessTrue StreamDownloadListResponseEnvelopeSuccess = true
)

type StreamDownloadDeleteResponseEnvelope struct {
	Errors   []StreamDownloadDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamDownloadDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamDownloadDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamDownloadDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamDownloadDeleteResponseEnvelopeJSON    `json:"-"`
}

// streamDownloadDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamDownloadDeleteResponseEnvelope]
type streamDownloadDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamDownloadDeleteResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    streamDownloadDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// streamDownloadDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StreamDownloadDeleteResponseEnvelopeErrors]
type streamDownloadDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamDownloadDeleteResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    streamDownloadDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// streamDownloadDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamDownloadDeleteResponseEnvelopeMessages]
type streamDownloadDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDownloadDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamDownloadDeleteResponseEnvelopeSuccess bool

const (
	StreamDownloadDeleteResponseEnvelopeSuccessTrue StreamDownloadDeleteResponseEnvelopeSuccess = true
)
