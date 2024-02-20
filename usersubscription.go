// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// UserSubscriptionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserSubscriptionService] method
// instead.
type UserSubscriptionService struct {
	Options []option.RequestOption
}

// NewUserSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserSubscriptionService(opts ...option.RequestOption) (r *UserSubscriptionService) {
	r = &UserSubscriptionService{}
	r.Options = opts
	return
}

// Lists all of a user's subscriptions.
func (r *UserSubscriptionService) List(ctx context.Context, opts ...option.RequestOption) (res *[]UserSubscriptionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserSubscriptionListResponseEnvelope
	path := "user/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a user's subscription.
func (r *UserSubscriptionService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *UserSubscriptionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/subscriptions/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Updates a user's subscriptions.
func (r *UserSubscriptionService) Replace(ctx context.Context, identifier string, body UserSubscriptionReplaceParams, opts ...option.RequestOption) (res *UserSubscriptionReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserSubscriptionReplaceResponseEnvelope
	path := fmt.Sprintf("user/subscriptions/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserSubscriptionListResponse struct {
	// Subscription identifier tag.
	ID  string                          `json:"id"`
	App UserSubscriptionListResponseApp `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues []UserSubscriptionListResponseComponentValue `json:"component_values"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The end of the current period and also when the next billing is due.
	CurrentPeriodEnd time.Time `json:"current_period_end" format:"date-time"`
	// When the current billing period started. May match initial_period_start if this
	// is the first period.
	CurrentPeriodStart time.Time `json:"current_period_start" format:"date-time"`
	// How often the subscription is renewed automatically.
	Frequency UserSubscriptionListResponseFrequency `json:"frequency"`
	// The price of the subscription that will be billed, in US dollars.
	Price float64 `json:"price"`
	// The rate plan applied to the subscription.
	RatePlan UserSubscriptionListResponseRatePlan `json:"rate_plan"`
	// The state that the subscription is in.
	State UserSubscriptionListResponseState `json:"state"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone UserSubscriptionListResponseZone `json:"zone"`
	JSON userSubscriptionListResponseJSON `json:"-"`
}

// userSubscriptionListResponseJSON contains the JSON metadata for the struct
// [UserSubscriptionListResponse]
type userSubscriptionListResponseJSON struct {
	ID                 apijson.Field
	App                apijson.Field
	ComponentValues    apijson.Field
	Currency           apijson.Field
	CurrentPeriodEnd   apijson.Field
	CurrentPeriodStart apijson.Field
	Frequency          apijson.Field
	Price              apijson.Field
	RatePlan           apijson.Field
	State              apijson.Field
	Zone               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserSubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionListResponseApp struct {
	// app install id.
	InstallID string                              `json:"install_id"`
	JSON      userSubscriptionListResponseAppJSON `json:"-"`
}

// userSubscriptionListResponseAppJSON contains the JSON metadata for the struct
// [UserSubscriptionListResponseApp]
type userSubscriptionListResponseAppJSON struct {
	InstallID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionListResponseApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A component value for a subscription.
type UserSubscriptionListResponseComponentValue struct {
	// The default amount assigned.
	Default float64 `json:"default"`
	// The name of the component value.
	Name string `json:"name"`
	// The unit price for the component value.
	Price float64 `json:"price"`
	// The amount of the component value assigned.
	Value float64                                        `json:"value"`
	JSON  userSubscriptionListResponseComponentValueJSON `json:"-"`
}

// userSubscriptionListResponseComponentValueJSON contains the JSON metadata for
// the struct [UserSubscriptionListResponseComponentValue]
type userSubscriptionListResponseComponentValueJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionListResponseComponentValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How often the subscription is renewed automatically.
type UserSubscriptionListResponseFrequency string

const (
	UserSubscriptionListResponseFrequencyWeekly    UserSubscriptionListResponseFrequency = "weekly"
	UserSubscriptionListResponseFrequencyMonthly   UserSubscriptionListResponseFrequency = "monthly"
	UserSubscriptionListResponseFrequencyQuarterly UserSubscriptionListResponseFrequency = "quarterly"
	UserSubscriptionListResponseFrequencyYearly    UserSubscriptionListResponseFrequency = "yearly"
)

// The rate plan applied to the subscription.
type UserSubscriptionListResponseRatePlan struct {
	// The ID of the rate plan.
	ID interface{} `json:"id"`
	// The currency applied to the rate plan subscription.
	Currency string `json:"currency"`
	// Whether this rate plan is managed externally from Cloudflare.
	ExternallyManaged bool `json:"externally_managed"`
	// Whether a rate plan is enterprise-based (or newly adopted term contract).
	IsContract bool `json:"is_contract"`
	// The full name of the rate plan.
	PublicName string `json:"public_name"`
	// The scope that this rate plan applies to.
	Scope string `json:"scope"`
	// The list of sets this rate plan applies to.
	Sets []string                                 `json:"sets"`
	JSON userSubscriptionListResponseRatePlanJSON `json:"-"`
}

// userSubscriptionListResponseRatePlanJSON contains the JSON metadata for the
// struct [UserSubscriptionListResponseRatePlan]
type userSubscriptionListResponseRatePlanJSON struct {
	ID                apijson.Field
	Currency          apijson.Field
	ExternallyManaged apijson.Field
	IsContract        apijson.Field
	PublicName        apijson.Field
	Scope             apijson.Field
	Sets              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserSubscriptionListResponseRatePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The state that the subscription is in.
type UserSubscriptionListResponseState string

const (
	UserSubscriptionListResponseStateTrial           UserSubscriptionListResponseState = "Trial"
	UserSubscriptionListResponseStateProvisioned     UserSubscriptionListResponseState = "Provisioned"
	UserSubscriptionListResponseStatePaid            UserSubscriptionListResponseState = "Paid"
	UserSubscriptionListResponseStateAwaitingPayment UserSubscriptionListResponseState = "AwaitingPayment"
	UserSubscriptionListResponseStateCancelled       UserSubscriptionListResponseState = "Cancelled"
	UserSubscriptionListResponseStateFailed          UserSubscriptionListResponseState = "Failed"
	UserSubscriptionListResponseStateExpired         UserSubscriptionListResponseState = "Expired"
)

// A simple zone object. May have null properties if not a zone subscription.
type UserSubscriptionListResponseZone struct {
	// Identifier
	ID string `json:"id"`
	// The domain name
	Name string                               `json:"name"`
	JSON userSubscriptionListResponseZoneJSON `json:"-"`
}

// userSubscriptionListResponseZoneJSON contains the JSON metadata for the struct
// [UserSubscriptionListResponseZone]
type userSubscriptionListResponseZoneJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionListResponseZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionDeleteResponse struct {
	// Subscription identifier tag.
	SubscriptionID string                             `json:"subscription_id"`
	JSON           userSubscriptionDeleteResponseJSON `json:"-"`
}

// userSubscriptionDeleteResponseJSON contains the JSON metadata for the struct
// [UserSubscriptionDeleteResponse]
type userSubscriptionDeleteResponseJSON struct {
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserSubscriptionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [UserSubscriptionReplaceResponseUnknown] or
// [shared.UnionString].
type UserSubscriptionReplaceResponse interface {
	ImplementsUserSubscriptionReplaceResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSubscriptionReplaceResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserSubscriptionListResponseEnvelope struct {
	Errors   []UserSubscriptionListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserSubscriptionListResponseEnvelopeMessages `json:"messages,required"`
	Result   []UserSubscriptionListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserSubscriptionListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserSubscriptionListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userSubscriptionListResponseEnvelopeJSON       `json:"-"`
}

// userSubscriptionListResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserSubscriptionListResponseEnvelope]
type userSubscriptionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    userSubscriptionListResponseEnvelopeErrorsJSON `json:"-"`
}

// userSubscriptionListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [UserSubscriptionListResponseEnvelopeErrors]
type userSubscriptionListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    userSubscriptionListResponseEnvelopeMessagesJSON `json:"-"`
}

// userSubscriptionListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [UserSubscriptionListResponseEnvelopeMessages]
type userSubscriptionListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserSubscriptionListResponseEnvelopeSuccess bool

const (
	UserSubscriptionListResponseEnvelopeSuccessTrue UserSubscriptionListResponseEnvelopeSuccess = true
)

type UserSubscriptionListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       userSubscriptionListResponseEnvelopeResultInfoJSON `json:"-"`
}

// userSubscriptionListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [UserSubscriptionListResponseEnvelopeResultInfo]
type userSubscriptionListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionReplaceParams struct {
	App param.Field[UserSubscriptionReplaceParamsApp] `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues param.Field[[]UserSubscriptionReplaceParamsComponentValue] `json:"component_values"`
	// How often the subscription is renewed automatically.
	Frequency param.Field[UserSubscriptionReplaceParamsFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[UserSubscriptionReplaceParamsRatePlan] `json:"rate_plan"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone param.Field[UserSubscriptionReplaceParamsZone] `json:"zone"`
}

func (r UserSubscriptionReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSubscriptionReplaceParamsApp struct {
	// app install id.
	InstallID param.Field[string] `json:"install_id"`
}

func (r UserSubscriptionReplaceParamsApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A component value for a subscription.
type UserSubscriptionReplaceParamsComponentValue struct {
	// The default amount assigned.
	Default param.Field[float64] `json:"default"`
	// The name of the component value.
	Name param.Field[string] `json:"name"`
	// The unit price for the component value.
	Price param.Field[float64] `json:"price"`
	// The amount of the component value assigned.
	Value param.Field[float64] `json:"value"`
}

func (r UserSubscriptionReplaceParamsComponentValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How often the subscription is renewed automatically.
type UserSubscriptionReplaceParamsFrequency string

const (
	UserSubscriptionReplaceParamsFrequencyWeekly    UserSubscriptionReplaceParamsFrequency = "weekly"
	UserSubscriptionReplaceParamsFrequencyMonthly   UserSubscriptionReplaceParamsFrequency = "monthly"
	UserSubscriptionReplaceParamsFrequencyQuarterly UserSubscriptionReplaceParamsFrequency = "quarterly"
	UserSubscriptionReplaceParamsFrequencyYearly    UserSubscriptionReplaceParamsFrequency = "yearly"
)

// The rate plan applied to the subscription.
type UserSubscriptionReplaceParamsRatePlan struct {
	// The ID of the rate plan.
	ID param.Field[interface{}] `json:"id"`
	// The currency applied to the rate plan subscription.
	Currency param.Field[string] `json:"currency"`
	// Whether this rate plan is managed externally from Cloudflare.
	ExternallyManaged param.Field[bool] `json:"externally_managed"`
	// Whether a rate plan is enterprise-based (or newly adopted term contract).
	IsContract param.Field[bool] `json:"is_contract"`
	// The full name of the rate plan.
	PublicName param.Field[string] `json:"public_name"`
	// The scope that this rate plan applies to.
	Scope param.Field[string] `json:"scope"`
	// The list of sets this rate plan applies to.
	Sets param.Field[[]string] `json:"sets"`
}

func (r UserSubscriptionReplaceParamsRatePlan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A simple zone object. May have null properties if not a zone subscription.
type UserSubscriptionReplaceParamsZone struct {
}

func (r UserSubscriptionReplaceParamsZone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSubscriptionReplaceResponseEnvelope struct {
	Errors   []UserSubscriptionReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserSubscriptionReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   UserSubscriptionReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserSubscriptionReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    userSubscriptionReplaceResponseEnvelopeJSON    `json:"-"`
}

// userSubscriptionReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserSubscriptionReplaceResponseEnvelope]
type userSubscriptionReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionReplaceResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    userSubscriptionReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// userSubscriptionReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [UserSubscriptionReplaceResponseEnvelopeErrors]
type userSubscriptionReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionReplaceResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    userSubscriptionReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// userSubscriptionReplaceResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [UserSubscriptionReplaceResponseEnvelopeMessages]
type userSubscriptionReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserSubscriptionReplaceResponseEnvelopeSuccess bool

const (
	UserSubscriptionReplaceResponseEnvelopeSuccessTrue UserSubscriptionReplaceResponseEnvelopeSuccess = true
)
