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

// AccessOrganizationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessOrganizationService] method
// instead.
type AccessOrganizationService struct {
	Options []option.RequestOption
}

// NewAccessOrganizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessOrganizationService(opts ...option.RequestOption) (r *AccessOrganizationService) {
	r = &AccessOrganizationService{}
	r.Options = opts
	return
}

// Sets up a Zero Trust organization for your account or zone.
func (r *AccessOrganizationService) New(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessOrganizationNewParams, opts ...option.RequestOption) (res *AccessOrganizationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessOrganizationNewResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the configuration for your Zero Trust organization.
func (r *AccessOrganizationService) List(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *AccessOrganizationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessOrganizationListResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the configuration for your Zero Trust organization.
func (r *AccessOrganizationService) Replace(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessOrganizationReplaceParams, opts ...option.RequestOption) (res *AccessOrganizationReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessOrganizationReplaceResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Revokes a user's access across all applications.
func (r *AccessOrganizationService) RevokeUsers(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessOrganizationRevokeUsersParams, opts ...option.RequestOption) (res *AccessOrganizationRevokeUsersResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessOrganizationRevokeUsersResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/organizations/revoke_user", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessOrganizationNewResponse struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                     `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                `json:"created_at" format:"date-time"`
	CustomPages            AccessOrganizationNewResponseCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                     `json:"is_ui_read_only"`
	LoginDesign  AccessOrganizationNewResponseLoginDesign `json:"login_design"`
	// The name of your Zero Trust organization.
	Name string `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration string `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason string    `json:"ui_read_only_toggle_reason"`
	UpdatedAt              time.Time `json:"updated_at" format:"date-time"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime string `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WarpAuthSessionDuration string                            `json:"warp_auth_session_duration"`
	JSON                    accessOrganizationNewResponseJSON `json:"-"`
}

// accessOrganizationNewResponseJSON contains the JSON metadata for the struct
// [AccessOrganizationNewResponse]
type accessOrganizationNewResponseJSON struct {
	AllowAuthenticateViaWarp       apijson.Field
	AuthDomain                     apijson.Field
	AutoRedirectToIdentity         apijson.Field
	CreatedAt                      apijson.Field
	CustomPages                    apijson.Field
	IsUiReadOnly                   apijson.Field
	LoginDesign                    apijson.Field
	Name                           apijson.Field
	SessionDuration                apijson.Field
	UiReadOnlyToggleReason         apijson.Field
	UpdatedAt                      apijson.Field
	UserSeatExpirationInactiveTime apijson.Field
	WarpAuthSessionDuration        apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccessOrganizationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationNewResponseCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                       `json:"identity_denied"`
	JSON           accessOrganizationNewResponseCustomPagesJSON `json:"-"`
}

// accessOrganizationNewResponseCustomPagesJSON contains the JSON metadata for the
// struct [AccessOrganizationNewResponseCustomPages]
type accessOrganizationNewResponseCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessOrganizationNewResponseCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationNewResponseLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                       `json:"text_color"`
	JSON      accessOrganizationNewResponseLoginDesignJSON `json:"-"`
}

// accessOrganizationNewResponseLoginDesignJSON contains the JSON metadata for the
// struct [AccessOrganizationNewResponseLoginDesign]
type accessOrganizationNewResponseLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessOrganizationNewResponseLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationListResponse struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                      `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                 `json:"created_at" format:"date-time"`
	CustomPages            AccessOrganizationListResponseCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                      `json:"is_ui_read_only"`
	LoginDesign  AccessOrganizationListResponseLoginDesign `json:"login_design"`
	// The name of your Zero Trust organization.
	Name string `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration string `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason string    `json:"ui_read_only_toggle_reason"`
	UpdatedAt              time.Time `json:"updated_at" format:"date-time"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime string `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WarpAuthSessionDuration string                             `json:"warp_auth_session_duration"`
	JSON                    accessOrganizationListResponseJSON `json:"-"`
}

// accessOrganizationListResponseJSON contains the JSON metadata for the struct
// [AccessOrganizationListResponse]
type accessOrganizationListResponseJSON struct {
	AllowAuthenticateViaWarp       apijson.Field
	AuthDomain                     apijson.Field
	AutoRedirectToIdentity         apijson.Field
	CreatedAt                      apijson.Field
	CustomPages                    apijson.Field
	IsUiReadOnly                   apijson.Field
	LoginDesign                    apijson.Field
	Name                           apijson.Field
	SessionDuration                apijson.Field
	UiReadOnlyToggleReason         apijson.Field
	UpdatedAt                      apijson.Field
	UserSeatExpirationInactiveTime apijson.Field
	WarpAuthSessionDuration        apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccessOrganizationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationListResponseCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                        `json:"identity_denied"`
	JSON           accessOrganizationListResponseCustomPagesJSON `json:"-"`
}

// accessOrganizationListResponseCustomPagesJSON contains the JSON metadata for the
// struct [AccessOrganizationListResponseCustomPages]
type accessOrganizationListResponseCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessOrganizationListResponseCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationListResponseLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                        `json:"text_color"`
	JSON      accessOrganizationListResponseLoginDesignJSON `json:"-"`
}

// accessOrganizationListResponseLoginDesignJSON contains the JSON metadata for the
// struct [AccessOrganizationListResponseLoginDesign]
type accessOrganizationListResponseLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessOrganizationListResponseLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationReplaceResponse struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                         `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                    `json:"created_at" format:"date-time"`
	CustomPages            AccessOrganizationReplaceResponseCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                         `json:"is_ui_read_only"`
	LoginDesign  AccessOrganizationReplaceResponseLoginDesign `json:"login_design"`
	// The name of your Zero Trust organization.
	Name string `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration string `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason string    `json:"ui_read_only_toggle_reason"`
	UpdatedAt              time.Time `json:"updated_at" format:"date-time"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime string `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WarpAuthSessionDuration string                                `json:"warp_auth_session_duration"`
	JSON                    accessOrganizationReplaceResponseJSON `json:"-"`
}

// accessOrganizationReplaceResponseJSON contains the JSON metadata for the struct
// [AccessOrganizationReplaceResponse]
type accessOrganizationReplaceResponseJSON struct {
	AllowAuthenticateViaWarp       apijson.Field
	AuthDomain                     apijson.Field
	AutoRedirectToIdentity         apijson.Field
	CreatedAt                      apijson.Field
	CustomPages                    apijson.Field
	IsUiReadOnly                   apijson.Field
	LoginDesign                    apijson.Field
	Name                           apijson.Field
	SessionDuration                apijson.Field
	UiReadOnlyToggleReason         apijson.Field
	UpdatedAt                      apijson.Field
	UserSeatExpirationInactiveTime apijson.Field
	WarpAuthSessionDuration        apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccessOrganizationReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationReplaceResponseCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                           `json:"identity_denied"`
	JSON           accessOrganizationReplaceResponseCustomPagesJSON `json:"-"`
}

// accessOrganizationReplaceResponseCustomPagesJSON contains the JSON metadata for
// the struct [AccessOrganizationReplaceResponseCustomPages]
type accessOrganizationReplaceResponseCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessOrganizationReplaceResponseCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationReplaceResponseLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                           `json:"text_color"`
	JSON      accessOrganizationReplaceResponseLoginDesignJSON `json:"-"`
}

// accessOrganizationReplaceResponseLoginDesignJSON contains the JSON metadata for
// the struct [AccessOrganizationReplaceResponseLoginDesign]
type accessOrganizationReplaceResponseLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessOrganizationReplaceResponseLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationRevokeUsersResponse bool

const (
	AccessOrganizationRevokeUsersResponseTrue  AccessOrganizationRevokeUsersResponse = true
	AccessOrganizationRevokeUsersResponseFalse AccessOrganizationRevokeUsersResponse = false
)

type AccessOrganizationNewParams struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain,required"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name,required"`
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp param.Field[bool] `json:"allow_authenticate_via_warp"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]                                   `json:"is_ui_read_only"`
	LoginDesign  param.Field[AccessOrganizationNewParamsLoginDesign] `json:"login_design"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason param.Field[string] `json:"ui_read_only_toggle_reason"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime param.Field[string] `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WarpAuthSessionDuration param.Field[string] `json:"warp_auth_session_duration"`
}

func (r AccessOrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessOrganizationNewParamsLoginDesign struct {
	// The background color on your login page.
	BackgroundColor param.Field[string] `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText param.Field[string] `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText param.Field[string] `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath param.Field[string] `json:"logo_path"`
	// The text color on your login page.
	TextColor param.Field[string] `json:"text_color"`
}

func (r AccessOrganizationNewParamsLoginDesign) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessOrganizationNewResponseEnvelope struct {
	Errors   []AccessOrganizationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessOrganizationNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessOrganizationNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessOrganizationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessOrganizationNewResponseEnvelopeJSON    `json:"-"`
}

// accessOrganizationNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessOrganizationNewResponseEnvelope]
type accessOrganizationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationNewResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accessOrganizationNewResponseEnvelopeErrorsJSON `json:"-"`
}

// accessOrganizationNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessOrganizationNewResponseEnvelopeErrors]
type accessOrganizationNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationNewResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accessOrganizationNewResponseEnvelopeMessagesJSON `json:"-"`
}

// accessOrganizationNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessOrganizationNewResponseEnvelopeMessages]
type accessOrganizationNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessOrganizationNewResponseEnvelopeSuccess bool

const (
	AccessOrganizationNewResponseEnvelopeSuccessTrue AccessOrganizationNewResponseEnvelopeSuccess = true
)

type AccessOrganizationListResponseEnvelope struct {
	Errors   []AccessOrganizationListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessOrganizationListResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessOrganizationListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessOrganizationListResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessOrganizationListResponseEnvelopeJSON    `json:"-"`
}

// accessOrganizationListResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessOrganizationListResponseEnvelope]
type accessOrganizationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationListResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accessOrganizationListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessOrganizationListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessOrganizationListResponseEnvelopeErrors]
type accessOrganizationListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationListResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accessOrganizationListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessOrganizationListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessOrganizationListResponseEnvelopeMessages]
type accessOrganizationListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessOrganizationListResponseEnvelopeSuccess bool

const (
	AccessOrganizationListResponseEnvelopeSuccessTrue AccessOrganizationListResponseEnvelopeSuccess = true
)

type AccessOrganizationReplaceParams struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity param.Field[bool]                                       `json:"auto_redirect_to_identity"`
	CustomPages            param.Field[AccessOrganizationReplaceParamsCustomPages] `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]                                       `json:"is_ui_read_only"`
	LoginDesign  param.Field[AccessOrganizationReplaceParamsLoginDesign] `json:"login_design"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason param.Field[string] `json:"ui_read_only_toggle_reason"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime param.Field[string] `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WarpAuthSessionDuration param.Field[string] `json:"warp_auth_session_duration"`
}

func (r AccessOrganizationReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessOrganizationReplaceParamsCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden param.Field[string] `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied param.Field[string] `json:"identity_denied"`
}

func (r AccessOrganizationReplaceParamsCustomPages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessOrganizationReplaceParamsLoginDesign struct {
	// The background color on your login page.
	BackgroundColor param.Field[string] `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText param.Field[string] `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText param.Field[string] `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath param.Field[string] `json:"logo_path"`
	// The text color on your login page.
	TextColor param.Field[string] `json:"text_color"`
}

func (r AccessOrganizationReplaceParamsLoginDesign) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessOrganizationReplaceResponseEnvelope struct {
	Errors   []AccessOrganizationReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessOrganizationReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessOrganizationReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessOrganizationReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessOrganizationReplaceResponseEnvelopeJSON    `json:"-"`
}

// accessOrganizationReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessOrganizationReplaceResponseEnvelope]
type accessOrganizationReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationReplaceResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accessOrganizationReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// accessOrganizationReplaceResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessOrganizationReplaceResponseEnvelopeErrors]
type accessOrganizationReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationReplaceResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accessOrganizationReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// accessOrganizationReplaceResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessOrganizationReplaceResponseEnvelopeMessages]
type accessOrganizationReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessOrganizationReplaceResponseEnvelopeSuccess bool

const (
	AccessOrganizationReplaceResponseEnvelopeSuccessTrue AccessOrganizationReplaceResponseEnvelopeSuccess = true
)

type AccessOrganizationRevokeUsersParams struct {
	// The email of the user to revoke.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessOrganizationRevokeUsersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessOrganizationRevokeUsersResponseEnvelope struct {
	Result  AccessOrganizationRevokeUsersResponse                `json:"result"`
	Success AccessOrganizationRevokeUsersResponseEnvelopeSuccess `json:"success"`
	JSON    accessOrganizationRevokeUsersResponseEnvelopeJSON    `json:"-"`
}

// accessOrganizationRevokeUsersResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessOrganizationRevokeUsersResponseEnvelope]
type accessOrganizationRevokeUsersResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationRevokeUsersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationRevokeUsersResponseEnvelopeSuccess bool

const (
	AccessOrganizationRevokeUsersResponseEnvelopeSuccessTrue  AccessOrganizationRevokeUsersResponseEnvelopeSuccess = true
	AccessOrganizationRevokeUsersResponseEnvelopeSuccessFalse AccessOrganizationRevokeUsersResponseEnvelopeSuccess = false
)
