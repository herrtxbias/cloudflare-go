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

// StorageKvNamespaceBulkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStorageKvNamespaceBulkService]
// method instead.
type StorageKvNamespaceBulkService struct {
	Options []option.RequestOption
}

// NewStorageKvNamespaceBulkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStorageKvNamespaceBulkService(opts ...option.RequestOption) (r *StorageKvNamespaceBulkService) {
	r = &StorageKvNamespaceBulkService{}
	r.Options = opts
	return
}

// Remove multiple KV pairs from the namespace. Body should be an array of up to
// 10,000 keys to be removed.
func (r *StorageKvNamespaceBulkService) Delete(ctx context.Context, accountID string, namespaceID string, body StorageKvNamespaceBulkDeleteParams, opts ...option.RequestOption) (res *StorageKvNamespaceBulkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKvNamespaceBulkDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk", accountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Write multiple keys and values at once. Body should be an array of up to 10,000
// key-value pairs to be stored, along with optional expiration information.
// Existing values and expirations will be overwritten. If neither `expiration` nor
// `expiration_ttl` is specified, the key-value pair will never expire. If both are
// set, `expiration_ttl` is used and `expiration` is ignored. The entire request
// size must be 100 megabytes or less.
func (r *StorageKvNamespaceBulkService) Replace(ctx context.Context, accountID string, namespaceID string, body StorageKvNamespaceBulkReplaceParams, opts ...option.RequestOption) (res *StorageKvNamespaceBulkReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageKvNamespaceBulkReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk", accountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [StorageKvNamespaceBulkDeleteResponseUnknown] or
// [shared.UnionString].
type StorageKvNamespaceBulkDeleteResponse interface {
	ImplementsStorageKvNamespaceBulkDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StorageKvNamespaceBulkDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [StorageKvNamespaceBulkReplaceResponseUnknown] or
// [shared.UnionString].
type StorageKvNamespaceBulkReplaceResponse interface {
	ImplementsStorageKvNamespaceBulkReplaceResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StorageKvNamespaceBulkReplaceResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type StorageKvNamespaceBulkDeleteParams struct {
	Body param.Field[[]string] `json:"body,required"`
}

func (r StorageKvNamespaceBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type StorageKvNamespaceBulkDeleteResponseEnvelope struct {
	Errors   []StorageKvNamespaceBulkDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKvNamespaceBulkDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   StorageKvNamespaceBulkDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StorageKvNamespaceBulkDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageKvNamespaceBulkDeleteResponseEnvelopeJSON    `json:"-"`
}

// storageKvNamespaceBulkDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [StorageKvNamespaceBulkDeleteResponseEnvelope]
type storageKvNamespaceBulkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceBulkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceBulkDeleteResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    storageKvNamespaceBulkDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKvNamespaceBulkDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [StorageKvNamespaceBulkDeleteResponseEnvelopeErrors]
type storageKvNamespaceBulkDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceBulkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceBulkDeleteResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    storageKvNamespaceBulkDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKvNamespaceBulkDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [StorageKvNamespaceBulkDeleteResponseEnvelopeMessages]
type storageKvNamespaceBulkDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceBulkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKvNamespaceBulkDeleteResponseEnvelopeSuccess bool

const (
	StorageKvNamespaceBulkDeleteResponseEnvelopeSuccessTrue StorageKvNamespaceBulkDeleteResponseEnvelopeSuccess = true
)

type StorageKvNamespaceBulkReplaceParams struct {
	Body param.Field[[]StorageKvNamespaceBulkReplaceParamsBody] `json:"body,required"`
}

func (r StorageKvNamespaceBulkReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type StorageKvNamespaceBulkReplaceParamsBody struct {
	// Whether or not the server should base64 decode the value before storing it.
	// Useful for writing values that wouldn't otherwise be valid JSON strings, such as
	// images.
	Base64 param.Field[bool] `json:"base64"`
	// The time, measured in number of seconds since the UNIX epoch, at which the key
	// should expire.
	Expiration param.Field[float64] `json:"expiration"`
	// The number of seconds for which the key should be visible before it expires. At
	// least 60.
	ExpirationTTL param.Field[float64] `json:"expiration_ttl"`
	// A key's name. The name may be at most 512 bytes. All printable, non-whitespace
	// characters are valid.
	Key param.Field[string] `json:"key"`
	// Arbitrary JSON that is associated with a key.
	Metadata param.Field[interface{}] `json:"metadata"`
	// A UTF-8 encoded string to be stored, up to 25 MiB in length.
	Value param.Field[string] `json:"value"`
}

func (r StorageKvNamespaceBulkReplaceParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StorageKvNamespaceBulkReplaceResponseEnvelope struct {
	Errors   []StorageKvNamespaceBulkReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageKvNamespaceBulkReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   StorageKvNamespaceBulkReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StorageKvNamespaceBulkReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageKvNamespaceBulkReplaceResponseEnvelopeJSON    `json:"-"`
}

// storageKvNamespaceBulkReplaceResponseEnvelopeJSON contains the JSON metadata for
// the struct [StorageKvNamespaceBulkReplaceResponseEnvelope]
type storageKvNamespaceBulkReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceBulkReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceBulkReplaceResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    storageKvNamespaceBulkReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// storageKvNamespaceBulkReplaceResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [StorageKvNamespaceBulkReplaceResponseEnvelopeErrors]
type storageKvNamespaceBulkReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceBulkReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvNamespaceBulkReplaceResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    storageKvNamespaceBulkReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// storageKvNamespaceBulkReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [StorageKvNamespaceBulkReplaceResponseEnvelopeMessages]
type storageKvNamespaceBulkReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageKvNamespaceBulkReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageKvNamespaceBulkReplaceResponseEnvelopeSuccess bool

const (
	StorageKvNamespaceBulkReplaceResponseEnvelopeSuccessTrue StorageKvNamespaceBulkReplaceResponseEnvelopeSuccess = true
)
