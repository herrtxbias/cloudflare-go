// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/herrtxbias/cloudflare-go/v4/internal/apijson"
	"github.com/herrtxbias/cloudflare-go/v4/internal/param"
	"github.com/herrtxbias/cloudflare-go/v4/internal/requestconfig"
	"github.com/herrtxbias/cloudflare-go/v4/option"
	"github.com/herrtxbias/cloudflare-go/v4/shared"
)

// DeviceSettingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceSettingService] method instead.
type DeviceSettingService struct {
	Options []option.RequestOption
}

// NewDeviceSettingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceSettingService(opts ...option.RequestOption) (r *DeviceSettingService) {
	r = &DeviceSettingService{}
	r.Options = opts
	return
}

// Updates the current device settings for a Zero Trust account.
func (r *DeviceSettingService) Update(ctx context.Context, params DeviceSettingUpdateParams, opts ...option.RequestOption) (res *DeviceSettings, err error) {
	var env DeviceSettingUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/settings", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Describes the current device settings for a Zero Trust account.
func (r *DeviceSettingService) List(ctx context.Context, query DeviceSettingListParams, opts ...option.RequestOption) (res *DeviceSettings, err error) {
	var env DeviceSettingListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/settings", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patches the current device settings for a Zero Trust account.
func (r *DeviceSettingService) Edit(ctx context.Context, params DeviceSettingEditParams, opts ...option.RequestOption) (res *DeviceSettings, err error) {
	var env DeviceSettingEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/settings", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeviceSettings struct {
	// Sets the time limit, in seconds, that a user can use an override code to bypass
	// WARP.
	DisableForTime float64 `json:"disable_for_time"`
	// Enable gateway proxy filtering on TCP.
	GatewayProxyEnabled bool `json:"gateway_proxy_enabled"`
	// Enable gateway proxy filtering on UDP.
	GatewayUdpProxyEnabled bool `json:"gateway_udp_proxy_enabled"`
	// Enable installation of cloudflare managed root certificate.
	RootCertificateInstallationEnabled bool `json:"root_certificate_installation_enabled"`
	// Enable using CGNAT virtual IPv4.
	UseZtVirtualIP bool               `json:"use_zt_virtual_ip"`
	JSON           deviceSettingsJSON `json:"-"`
}

// deviceSettingsJSON contains the JSON metadata for the struct [DeviceSettings]
type deviceSettingsJSON struct {
	DisableForTime                     apijson.Field
	GatewayProxyEnabled                apijson.Field
	GatewayUdpProxyEnabled             apijson.Field
	RootCertificateInstallationEnabled apijson.Field
	UseZtVirtualIP                     apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *DeviceSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingsJSON) RawJSON() string {
	return r.raw
}

type DeviceSettingsParam struct {
	// Sets the time limit, in seconds, that a user can use an override code to bypass
	// WARP.
	DisableForTime param.Field[float64] `json:"disable_for_time"`
	// Enable gateway proxy filtering on TCP.
	GatewayProxyEnabled param.Field[bool] `json:"gateway_proxy_enabled"`
	// Enable gateway proxy filtering on UDP.
	GatewayUdpProxyEnabled param.Field[bool] `json:"gateway_udp_proxy_enabled"`
	// Enable installation of cloudflare managed root certificate.
	RootCertificateInstallationEnabled param.Field[bool] `json:"root_certificate_installation_enabled"`
	// Enable using CGNAT virtual IPv4.
	UseZtVirtualIP param.Field[bool] `json:"use_zt_virtual_ip"`
}

func (r DeviceSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceSettingUpdateParams struct {
	AccountID      param.Field[string] `path:"account_id,required"`
	DeviceSettings DeviceSettingsParam `json:"device_settings,required"`
}

func (r DeviceSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DeviceSettings)
}

type DeviceSettingUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DeviceSettings        `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceSettingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceSettingUpdateResponseEnvelopeJSON    `json:"-"`
}

// deviceSettingUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceSettingUpdateResponseEnvelope]
type deviceSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceSettingUpdateResponseEnvelopeSuccess bool

const (
	DeviceSettingUpdateResponseEnvelopeSuccessTrue DeviceSettingUpdateResponseEnvelopeSuccess = true
)

func (r DeviceSettingUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceSettingUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DeviceSettingListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceSettingListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DeviceSettings        `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceSettingListResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceSettingListResponseEnvelopeJSON    `json:"-"`
}

// deviceSettingListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceSettingListResponseEnvelope]
type deviceSettingListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceSettingListResponseEnvelopeSuccess bool

const (
	DeviceSettingListResponseEnvelopeSuccessTrue DeviceSettingListResponseEnvelopeSuccess = true
)

func (r DeviceSettingListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceSettingListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DeviceSettingEditParams struct {
	AccountID      param.Field[string] `path:"account_id,required"`
	DeviceSettings DeviceSettingsParam `json:"device_settings,required"`
}

func (r DeviceSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DeviceSettings)
}

type DeviceSettingEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DeviceSettings        `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceSettingEditResponseEnvelopeJSON    `json:"-"`
}

// deviceSettingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceSettingEditResponseEnvelope]
type deviceSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceSettingEditResponseEnvelopeSuccess bool

const (
	DeviceSettingEditResponseEnvelopeSuccessTrue DeviceSettingEditResponseEnvelopeSuccess = true
)

func (r DeviceSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
