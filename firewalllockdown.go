// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// FirewallLockdownService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFirewallLockdownService] method
// instead.
type FirewallLockdownService struct {
	Options []option.RequestOption
}

// NewFirewallLockdownService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFirewallLockdownService(opts ...option.RequestOption) (r *FirewallLockdownService) {
	r = &FirewallLockdownService{}
	r.Options = opts
	return
}

// Creates a new Zone Lockdown rule.
func (r *FirewallLockdownService) New(ctx context.Context, zoneIdentifier string, body FirewallLockdownNewParams, opts ...option.RequestOption) (res *FirewallLockdownNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallLockdownNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Zone Lockdown rules. You can filter the results using several optional
// parameters.
func (r *FirewallLockdownService) List(ctx context.Context, zoneIdentifier string, query FirewallLockdownListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[FirewallLockdownListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", zoneIdentifier)
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

// Fetches Zone Lockdown rules. You can filter the results using several optional
// parameters.
func (r *FirewallLockdownService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query FirewallLockdownListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[FirewallLockdownListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Deletes an existing Zone Lockdown rule.
func (r *FirewallLockdownService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *FirewallLockdownDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallLockdownDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a Zone Lockdown rule.
func (r *FirewallLockdownService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *FirewallLockdownGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallLockdownGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Zone Lockdown rule.
func (r *FirewallLockdownService) Replace(ctx context.Context, zoneIdentifier string, id string, body FirewallLockdownReplaceParams, opts ...option.RequestOption) (res *FirewallLockdownReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallLockdownReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FirewallLockdownNewResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations FirewallLockdownNewResponseConfigurations `json:"configurations,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// An informative summary of the rule.
	Description string `json:"description,required"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// When true, indicates that the rule is currently paused.
	Paused bool `json:"paused,required"`
	// The URLs to include in the rule definition. You can use wildcards. Each entered
	// URL will be escaped before use, which means you can only use simple wildcard
	// patterns.
	URLs []string                        `json:"urls,required"`
	JSON firewallLockdownNewResponseJSON `json:"-"`
}

// firewallLockdownNewResponseJSON contains the JSON metadata for the struct
// [FirewallLockdownNewResponse]
type firewallLockdownNewResponseJSON struct {
	ID             apijson.Field
	Configurations apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	ModifiedOn     apijson.Field
	Paused         apijson.Field
	URLs           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FirewallLockdownNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasIPConfiguration] or
// [FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasCidrConfiguration].
type FirewallLockdownNewResponseConfigurations interface {
	implementsFirewallLockdownNewResponseConfigurations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallLockdownNewResponseConfigurations)(nil)).Elem(), "")
}

type FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                       `json:"value"`
	JSON  firewallLockdownNewResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON `json:"-"`
}

// firewallLockdownNewResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasIPConfiguration]
type firewallLockdownNewResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasIPConfiguration) implementsFirewallLockdownNewResponseConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget string

const (
	FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasIPConfigurationTargetIP FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget = "ip"
)

type FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                         `json:"value"`
	JSON  firewallLockdownNewResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON `json:"-"`
}

// firewallLockdownNewResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasCidrConfiguration]
type firewallLockdownNewResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasCidrConfiguration) implementsFirewallLockdownNewResponseConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget string

const (
	FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasCidrConfigurationTargetIPRange FirewallLockdownNewResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget = "ip_range"
)

type FirewallLockdownListResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations FirewallLockdownListResponseConfigurations `json:"configurations,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// An informative summary of the rule.
	Description string `json:"description,required"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// When true, indicates that the rule is currently paused.
	Paused bool `json:"paused,required"`
	// The URLs to include in the rule definition. You can use wildcards. Each entered
	// URL will be escaped before use, which means you can only use simple wildcard
	// patterns.
	URLs []string                         `json:"urls,required"`
	JSON firewallLockdownListResponseJSON `json:"-"`
}

// firewallLockdownListResponseJSON contains the JSON metadata for the struct
// [FirewallLockdownListResponse]
type firewallLockdownListResponseJSON struct {
	ID             apijson.Field
	Configurations apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	ModifiedOn     apijson.Field
	Paused         apijson.Field
	URLs           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FirewallLockdownListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [FirewallLockdownListResponseConfigurationsLegacyJhsSchemasIPConfiguration] or
// [FirewallLockdownListResponseConfigurationsLegacyJhsSchemasCidrConfiguration].
type FirewallLockdownListResponseConfigurations interface {
	implementsFirewallLockdownListResponseConfigurations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallLockdownListResponseConfigurations)(nil)).Elem(), "")
}

type FirewallLockdownListResponseConfigurationsLegacyJhsSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target FirewallLockdownListResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                        `json:"value"`
	JSON  firewallLockdownListResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON `json:"-"`
}

// firewallLockdownListResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallLockdownListResponseConfigurationsLegacyJhsSchemasIPConfiguration]
type firewallLockdownListResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownListResponseConfigurationsLegacyJhsSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallLockdownListResponseConfigurationsLegacyJhsSchemasIPConfiguration) implementsFirewallLockdownListResponseConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type FirewallLockdownListResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget string

const (
	FirewallLockdownListResponseConfigurationsLegacyJhsSchemasIPConfigurationTargetIP FirewallLockdownListResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget = "ip"
)

type FirewallLockdownListResponseConfigurationsLegacyJhsSchemasCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target FirewallLockdownListResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                          `json:"value"`
	JSON  firewallLockdownListResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON `json:"-"`
}

// firewallLockdownListResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallLockdownListResponseConfigurationsLegacyJhsSchemasCidrConfiguration]
type firewallLockdownListResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownListResponseConfigurationsLegacyJhsSchemasCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallLockdownListResponseConfigurationsLegacyJhsSchemasCidrConfiguration) implementsFirewallLockdownListResponseConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type FirewallLockdownListResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget string

const (
	FirewallLockdownListResponseConfigurationsLegacyJhsSchemasCidrConfigurationTargetIPRange FirewallLockdownListResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget = "ip_range"
)

type FirewallLockdownDeleteResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID   string                             `json:"id"`
	JSON firewallLockdownDeleteResponseJSON `json:"-"`
}

// firewallLockdownDeleteResponseJSON contains the JSON metadata for the struct
// [FirewallLockdownDeleteResponse]
type firewallLockdownDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownGetResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations FirewallLockdownGetResponseConfigurations `json:"configurations,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// An informative summary of the rule.
	Description string `json:"description,required"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// When true, indicates that the rule is currently paused.
	Paused bool `json:"paused,required"`
	// The URLs to include in the rule definition. You can use wildcards. Each entered
	// URL will be escaped before use, which means you can only use simple wildcard
	// patterns.
	URLs []string                        `json:"urls,required"`
	JSON firewallLockdownGetResponseJSON `json:"-"`
}

// firewallLockdownGetResponseJSON contains the JSON metadata for the struct
// [FirewallLockdownGetResponse]
type firewallLockdownGetResponseJSON struct {
	ID             apijson.Field
	Configurations apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	ModifiedOn     apijson.Field
	Paused         apijson.Field
	URLs           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FirewallLockdownGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfiguration] or
// [FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfiguration].
type FirewallLockdownGetResponseConfigurations interface {
	implementsFirewallLockdownGetResponseConfigurations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallLockdownGetResponseConfigurations)(nil)).Elem(), "")
}

type FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                       `json:"value"`
	JSON  firewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON `json:"-"`
}

// firewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfiguration]
type firewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfiguration) implementsFirewallLockdownGetResponseConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget string

const (
	FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfigurationTargetIP FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget = "ip"
)

type FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                         `json:"value"`
	JSON  firewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON `json:"-"`
}

// firewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfiguration]
type firewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfiguration) implementsFirewallLockdownGetResponseConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget string

const (
	FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfigurationTargetIPRange FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget = "ip_range"
)

type FirewallLockdownReplaceResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations FirewallLockdownReplaceResponseConfigurations `json:"configurations,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// An informative summary of the rule.
	Description string `json:"description,required"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// When true, indicates that the rule is currently paused.
	Paused bool `json:"paused,required"`
	// The URLs to include in the rule definition. You can use wildcards. Each entered
	// URL will be escaped before use, which means you can only use simple wildcard
	// patterns.
	URLs []string                            `json:"urls,required"`
	JSON firewallLockdownReplaceResponseJSON `json:"-"`
}

// firewallLockdownReplaceResponseJSON contains the JSON metadata for the struct
// [FirewallLockdownReplaceResponse]
type firewallLockdownReplaceResponseJSON struct {
	ID             apijson.Field
	Configurations apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	ModifiedOn     apijson.Field
	Paused         apijson.Field
	URLs           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FirewallLockdownReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasIPConfiguration]
// or
// [FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasCidrConfiguration].
type FirewallLockdownReplaceResponseConfigurations interface {
	implementsFirewallLockdownReplaceResponseConfigurations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallLockdownReplaceResponseConfigurations)(nil)).Elem(), "")
}

type FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                           `json:"value"`
	JSON  firewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON `json:"-"`
}

// firewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasIPConfiguration]
type firewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasIPConfiguration) implementsFirewallLockdownReplaceResponseConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget string

const (
	FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasIPConfigurationTargetIP FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget = "ip"
)

type FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                             `json:"value"`
	JSON  firewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON `json:"-"`
}

// firewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasCidrConfiguration]
type firewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasCidrConfiguration) implementsFirewallLockdownReplaceResponseConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget string

const (
	FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasCidrConfigurationTargetIPRange FirewallLockdownReplaceResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget = "ip_range"
)

type FirewallLockdownNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallLockdownNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallLockdownNewResponseEnvelope struct {
	Errors   []FirewallLockdownNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallLockdownNewResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallLockdownNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallLockdownNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallLockdownNewResponseEnvelopeJSON    `json:"-"`
}

// firewallLockdownNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallLockdownNewResponseEnvelope]
type firewallLockdownNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    firewallLockdownNewResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallLockdownNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallLockdownNewResponseEnvelopeErrors]
type firewallLockdownNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    firewallLockdownNewResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallLockdownNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FirewallLockdownNewResponseEnvelopeMessages]
type firewallLockdownNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallLockdownNewResponseEnvelopeSuccess bool

const (
	FirewallLockdownNewResponseEnvelopeSuccessTrue FirewallLockdownNewResponseEnvelopeSuccess = true
)

type FirewallLockdownListParams struct {
	// A string to search for in the description of existing rules.
	Description param.Field[string] `query:"description"`
	// A string to search for in the description of existing rules.
	DescriptionSearch param.Field[string] `query:"description_search"`
	// A single IP address to search for in existing rules.
	IP param.Field[string] `query:"ip"`
	// A single IP address range to search for in existing rules.
	IPRangeSearch param.Field[string] `query:"ip_range_search"`
	// A single IP address to search for in existing rules.
	IPSearch param.Field[string] `query:"ip_search"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
	// The priority of the rule to control the processing order. A lower number
	// indicates higher priority. If not provided, any rules with a configured priority
	// will be processed before rules without a priority.
	Priority param.Field[float64] `query:"priority"`
	// A single URI to search for in the list of URLs of existing rules.
	UriSearch param.Field[string] `query:"uri_search"`
}

// URLQuery serializes [FirewallLockdownListParams]'s query parameters as
// `url.Values`.
func (r FirewallLockdownListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallLockdownDeleteResponseEnvelope struct {
	Result FirewallLockdownDeleteResponse             `json:"result"`
	JSON   firewallLockdownDeleteResponseEnvelopeJSON `json:"-"`
}

// firewallLockdownDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallLockdownDeleteResponseEnvelope]
type firewallLockdownDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownGetResponseEnvelope struct {
	Errors   []FirewallLockdownGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallLockdownGetResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallLockdownGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallLockdownGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallLockdownGetResponseEnvelopeJSON    `json:"-"`
}

// firewallLockdownGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallLockdownGetResponseEnvelope]
type firewallLockdownGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    firewallLockdownGetResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallLockdownGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallLockdownGetResponseEnvelopeErrors]
type firewallLockdownGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    firewallLockdownGetResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallLockdownGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FirewallLockdownGetResponseEnvelopeMessages]
type firewallLockdownGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallLockdownGetResponseEnvelopeSuccess bool

const (
	FirewallLockdownGetResponseEnvelopeSuccessTrue FirewallLockdownGetResponseEnvelopeSuccess = true
)

type FirewallLockdownReplaceParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallLockdownReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallLockdownReplaceResponseEnvelope struct {
	Errors   []FirewallLockdownReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallLockdownReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallLockdownReplaceResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallLockdownReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallLockdownReplaceResponseEnvelopeJSON    `json:"-"`
}

// firewallLockdownReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallLockdownReplaceResponseEnvelope]
type firewallLockdownReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownReplaceResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    firewallLockdownReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallLockdownReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FirewallLockdownReplaceResponseEnvelopeErrors]
type firewallLockdownReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownReplaceResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    firewallLockdownReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallLockdownReplaceResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [FirewallLockdownReplaceResponseEnvelopeMessages]
type firewallLockdownReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallLockdownReplaceResponseEnvelopeSuccess bool

const (
	FirewallLockdownReplaceResponseEnvelopeSuccessTrue FirewallLockdownReplaceResponseEnvelopeSuccess = true
)
