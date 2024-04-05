// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package speed

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SpeedService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewSpeedService] method instead.
type SpeedService struct {
	Options        []option.RequestOption
	Tests          *TestService
	Schedule       *ScheduleService
	Availabilities *AvailabilityService
	Pages          *PageService
}

// NewSpeedService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSpeedService(opts ...option.RequestOption) (r *SpeedService) {
	r = &SpeedService{}
	r.Options = opts
	r.Tests = NewTestService(opts...)
	r.Schedule = NewScheduleService(opts...)
	r.Availabilities = NewAvailabilityService(opts...)
	r.Pages = NewPageService(opts...)
	return
}

// Deletes a scheduled test for a page.
func (r *SpeedService) Delete(ctx context.Context, url string, params SpeedDeleteParams, opts ...option.RequestOption) (res *SpeedDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the test schedule for a page in a specific region.
func (r *SpeedService) ScheduleGet(ctx context.Context, url string, params SpeedScheduleGetParams, opts ...option.RequestOption) (res *Schedule, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedScheduleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the core web vital metrics trend over time for a specific page.
func (r *SpeedService) TrendsList(ctx context.Context, url string, params SpeedTrendsListParams, opts ...option.RequestOption) (res *Trend, err error) {
	opts = append(r.Options[:], opts...)
	var env SpeedTrendsListResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/trend", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A test region with a label.
type LabeledRegion struct {
	Label string `json:"label"`
	// A test region.
	Value LabeledRegionValue `json:"value"`
	JSON  labeledRegionJSON  `json:"-"`
}

// labeledRegionJSON contains the JSON metadata for the struct [LabeledRegion]
type labeledRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LabeledRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labeledRegionJSON) RawJSON() string {
	return r.raw
}

// A test region.
type LabeledRegionValue string

const (
	LabeledRegionValueAsiaEast1           LabeledRegionValue = "asia-east1"
	LabeledRegionValueAsiaNortheast1      LabeledRegionValue = "asia-northeast1"
	LabeledRegionValueAsiaNortheast2      LabeledRegionValue = "asia-northeast2"
	LabeledRegionValueAsiaSouth1          LabeledRegionValue = "asia-south1"
	LabeledRegionValueAsiaSoutheast1      LabeledRegionValue = "asia-southeast1"
	LabeledRegionValueAustraliaSoutheast1 LabeledRegionValue = "australia-southeast1"
	LabeledRegionValueEuropeNorth1        LabeledRegionValue = "europe-north1"
	LabeledRegionValueEuropeSouthwest1    LabeledRegionValue = "europe-southwest1"
	LabeledRegionValueEuropeWest1         LabeledRegionValue = "europe-west1"
	LabeledRegionValueEuropeWest2         LabeledRegionValue = "europe-west2"
	LabeledRegionValueEuropeWest3         LabeledRegionValue = "europe-west3"
	LabeledRegionValueEuropeWest4         LabeledRegionValue = "europe-west4"
	LabeledRegionValueEuropeWest8         LabeledRegionValue = "europe-west8"
	LabeledRegionValueEuropeWest9         LabeledRegionValue = "europe-west9"
	LabeledRegionValueMeWest1             LabeledRegionValue = "me-west1"
	LabeledRegionValueSouthamericaEast1   LabeledRegionValue = "southamerica-east1"
	LabeledRegionValueUsCentral1          LabeledRegionValue = "us-central1"
	LabeledRegionValueUsEast1             LabeledRegionValue = "us-east1"
	LabeledRegionValueUsEast4             LabeledRegionValue = "us-east4"
	LabeledRegionValueUsSouth1            LabeledRegionValue = "us-south1"
	LabeledRegionValueUsWest1             LabeledRegionValue = "us-west1"
)

func (r LabeledRegionValue) IsKnown() bool {
	switch r {
	case LabeledRegionValueAsiaEast1, LabeledRegionValueAsiaNortheast1, LabeledRegionValueAsiaNortheast2, LabeledRegionValueAsiaSouth1, LabeledRegionValueAsiaSoutheast1, LabeledRegionValueAustraliaSoutheast1, LabeledRegionValueEuropeNorth1, LabeledRegionValueEuropeSouthwest1, LabeledRegionValueEuropeWest1, LabeledRegionValueEuropeWest2, LabeledRegionValueEuropeWest3, LabeledRegionValueEuropeWest4, LabeledRegionValueEuropeWest8, LabeledRegionValueEuropeWest9, LabeledRegionValueMeWest1, LabeledRegionValueSouthamericaEast1, LabeledRegionValueUsCentral1, LabeledRegionValueUsEast1, LabeledRegionValueUsEast4, LabeledRegionValueUsSouth1, LabeledRegionValueUsWest1:
		return true
	}
	return false
}

// The Lighthouse report.
type LighthouseReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType LighthouseReportDeviceType `json:"deviceType"`
	Error      LighthouseReportError      `json:"error"`
	// First Contentful Paint.
	Fcp float64 `json:"fcp"`
	// The URL to the full Lighthouse JSON report.
	JsonReportURL string `json:"jsonReportUrl"`
	// Largest Contentful Paint.
	Lcp float64 `json:"lcp"`
	// The Lighthouse performance score.
	PerformanceScore float64 `json:"performanceScore"`
	// Speed Index.
	Si float64 `json:"si"`
	// The state of the Lighthouse report.
	State LighthouseReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64              `json:"tti"`
	JSON lighthouseReportJSON `json:"-"`
}

// lighthouseReportJSON contains the JSON metadata for the struct
// [LighthouseReport]
type lighthouseReportJSON struct {
	Cls              apijson.Field
	DeviceType       apijson.Field
	Error            apijson.Field
	Fcp              apijson.Field
	JsonReportURL    apijson.Field
	Lcp              apijson.Field
	PerformanceScore apijson.Field
	Si               apijson.Field
	State            apijson.Field
	Tbt              apijson.Field
	Ttfb             apijson.Field
	Tti              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LighthouseReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lighthouseReportJSON) RawJSON() string {
	return r.raw
}

// The type of device.
type LighthouseReportDeviceType string

const (
	LighthouseReportDeviceTypeDesktop LighthouseReportDeviceType = "DESKTOP"
	LighthouseReportDeviceTypeMobile  LighthouseReportDeviceType = "MOBILE"
)

func (r LighthouseReportDeviceType) IsKnown() bool {
	switch r {
	case LighthouseReportDeviceTypeDesktop, LighthouseReportDeviceTypeMobile:
		return true
	}
	return false
}

type LighthouseReportError struct {
	// The error code of the Lighthouse result.
	Code LighthouseReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                    `json:"finalDisplayedUrl"`
	JSON              lighthouseReportErrorJSON `json:"-"`
}

// lighthouseReportErrorJSON contains the JSON metadata for the struct
// [LighthouseReportError]
type lighthouseReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LighthouseReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lighthouseReportErrorJSON) RawJSON() string {
	return r.raw
}

// The error code of the Lighthouse result.
type LighthouseReportErrorCode string

const (
	LighthouseReportErrorCodeNotReachable      LighthouseReportErrorCode = "NOT_REACHABLE"
	LighthouseReportErrorCodeDNSFailure        LighthouseReportErrorCode = "DNS_FAILURE"
	LighthouseReportErrorCodeNotHTML           LighthouseReportErrorCode = "NOT_HTML"
	LighthouseReportErrorCodeLighthouseTimeout LighthouseReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	LighthouseReportErrorCodeUnknown           LighthouseReportErrorCode = "UNKNOWN"
)

func (r LighthouseReportErrorCode) IsKnown() bool {
	switch r {
	case LighthouseReportErrorCodeNotReachable, LighthouseReportErrorCodeDNSFailure, LighthouseReportErrorCodeNotHTML, LighthouseReportErrorCodeLighthouseTimeout, LighthouseReportErrorCodeUnknown:
		return true
	}
	return false
}

// The state of the Lighthouse report.
type LighthouseReportState string

const (
	LighthouseReportStateRunning  LighthouseReportState = "RUNNING"
	LighthouseReportStateComplete LighthouseReportState = "COMPLETE"
	LighthouseReportStateFailed   LighthouseReportState = "FAILED"
)

func (r LighthouseReportState) IsKnown() bool {
	switch r {
	case LighthouseReportStateRunning, LighthouseReportStateComplete, LighthouseReportStateFailed:
		return true
	}
	return false
}

// The test schedule.
type Schedule struct {
	// The frequency of the test.
	Frequency ScheduleFrequency `json:"frequency"`
	// A test region.
	Region ScheduleRegion `json:"region"`
	// A URL.
	URL  string       `json:"url"`
	JSON scheduleJSON `json:"-"`
}

// scheduleJSON contains the JSON metadata for the struct [Schedule]
type scheduleJSON struct {
	Frequency   apijson.Field
	Region      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Schedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleJSON) RawJSON() string {
	return r.raw
}

// The frequency of the test.
type ScheduleFrequency string

const (
	ScheduleFrequencyDaily  ScheduleFrequency = "DAILY"
	ScheduleFrequencyWeekly ScheduleFrequency = "WEEKLY"
)

func (r ScheduleFrequency) IsKnown() bool {
	switch r {
	case ScheduleFrequencyDaily, ScheduleFrequencyWeekly:
		return true
	}
	return false
}

// A test region.
type ScheduleRegion string

const (
	ScheduleRegionAsiaEast1           ScheduleRegion = "asia-east1"
	ScheduleRegionAsiaNortheast1      ScheduleRegion = "asia-northeast1"
	ScheduleRegionAsiaNortheast2      ScheduleRegion = "asia-northeast2"
	ScheduleRegionAsiaSouth1          ScheduleRegion = "asia-south1"
	ScheduleRegionAsiaSoutheast1      ScheduleRegion = "asia-southeast1"
	ScheduleRegionAustraliaSoutheast1 ScheduleRegion = "australia-southeast1"
	ScheduleRegionEuropeNorth1        ScheduleRegion = "europe-north1"
	ScheduleRegionEuropeSouthwest1    ScheduleRegion = "europe-southwest1"
	ScheduleRegionEuropeWest1         ScheduleRegion = "europe-west1"
	ScheduleRegionEuropeWest2         ScheduleRegion = "europe-west2"
	ScheduleRegionEuropeWest3         ScheduleRegion = "europe-west3"
	ScheduleRegionEuropeWest4         ScheduleRegion = "europe-west4"
	ScheduleRegionEuropeWest8         ScheduleRegion = "europe-west8"
	ScheduleRegionEuropeWest9         ScheduleRegion = "europe-west9"
	ScheduleRegionMeWest1             ScheduleRegion = "me-west1"
	ScheduleRegionSouthamericaEast1   ScheduleRegion = "southamerica-east1"
	ScheduleRegionUsCentral1          ScheduleRegion = "us-central1"
	ScheduleRegionUsEast1             ScheduleRegion = "us-east1"
	ScheduleRegionUsEast4             ScheduleRegion = "us-east4"
	ScheduleRegionUsSouth1            ScheduleRegion = "us-south1"
	ScheduleRegionUsWest1             ScheduleRegion = "us-west1"
)

func (r ScheduleRegion) IsKnown() bool {
	switch r {
	case ScheduleRegionAsiaEast1, ScheduleRegionAsiaNortheast1, ScheduleRegionAsiaNortheast2, ScheduleRegionAsiaSouth1, ScheduleRegionAsiaSoutheast1, ScheduleRegionAustraliaSoutheast1, ScheduleRegionEuropeNorth1, ScheduleRegionEuropeSouthwest1, ScheduleRegionEuropeWest1, ScheduleRegionEuropeWest2, ScheduleRegionEuropeWest3, ScheduleRegionEuropeWest4, ScheduleRegionEuropeWest8, ScheduleRegionEuropeWest9, ScheduleRegionMeWest1, ScheduleRegionSouthamericaEast1, ScheduleRegionUsCentral1, ScheduleRegionUsEast1, ScheduleRegionUsEast4, ScheduleRegionUsSouth1, ScheduleRegionUsWest1:
		return true
	}
	return false
}

// The test schedule.
type ScheduleParam struct {
	// The frequency of the test.
	Frequency param.Field[ScheduleFrequency] `json:"frequency"`
	// A test region.
	Region param.Field[ScheduleRegion] `json:"region"`
	// A URL.
	URL param.Field[string] `json:"url"`
}

func (r ScheduleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Trend struct {
	// Cumulative Layout Shift trend.
	Cls []float64 `json:"cls"`
	// First Contentful Paint trend.
	Fcp []float64 `json:"fcp"`
	// Largest Contentful Paint trend.
	Lcp []float64 `json:"lcp"`
	// The Lighthouse score trend.
	PerformanceScore []float64 `json:"performanceScore"`
	// Speed Index trend.
	Si []float64 `json:"si"`
	// Total Blocking Time trend.
	Tbt []float64 `json:"tbt"`
	// Time To First Byte trend.
	Ttfb []float64 `json:"ttfb"`
	// Time To Interactive trend.
	Tti  []float64 `json:"tti"`
	JSON trendJSON `json:"-"`
}

// trendJSON contains the JSON metadata for the struct [Trend]
type trendJSON struct {
	Cls              apijson.Field
	Fcp              apijson.Field
	Lcp              apijson.Field
	PerformanceScore apijson.Field
	Si               apijson.Field
	Tbt              apijson.Field
	Ttfb             apijson.Field
	Tti              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Trend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trendJSON) RawJSON() string {
	return r.raw
}

type SpeedDeleteResponse struct {
	// Number of items affected.
	Count float64                 `json:"count"`
	JSON  speedDeleteResponseJSON `json:"-"`
}

// speedDeleteResponseJSON contains the JSON metadata for the struct
// [SpeedDeleteResponse]
type speedDeleteResponseJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SpeedDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A test region.
	Region param.Field[SpeedDeleteParamsRegion] `query:"region"`
}

// URLQuery serializes [SpeedDeleteParams]'s query parameters as `url.Values`.
func (r SpeedDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type SpeedDeleteParamsRegion string

const (
	SpeedDeleteParamsRegionAsiaEast1           SpeedDeleteParamsRegion = "asia-east1"
	SpeedDeleteParamsRegionAsiaNortheast1      SpeedDeleteParamsRegion = "asia-northeast1"
	SpeedDeleteParamsRegionAsiaNortheast2      SpeedDeleteParamsRegion = "asia-northeast2"
	SpeedDeleteParamsRegionAsiaSouth1          SpeedDeleteParamsRegion = "asia-south1"
	SpeedDeleteParamsRegionAsiaSoutheast1      SpeedDeleteParamsRegion = "asia-southeast1"
	SpeedDeleteParamsRegionAustraliaSoutheast1 SpeedDeleteParamsRegion = "australia-southeast1"
	SpeedDeleteParamsRegionEuropeNorth1        SpeedDeleteParamsRegion = "europe-north1"
	SpeedDeleteParamsRegionEuropeSouthwest1    SpeedDeleteParamsRegion = "europe-southwest1"
	SpeedDeleteParamsRegionEuropeWest1         SpeedDeleteParamsRegion = "europe-west1"
	SpeedDeleteParamsRegionEuropeWest2         SpeedDeleteParamsRegion = "europe-west2"
	SpeedDeleteParamsRegionEuropeWest3         SpeedDeleteParamsRegion = "europe-west3"
	SpeedDeleteParamsRegionEuropeWest4         SpeedDeleteParamsRegion = "europe-west4"
	SpeedDeleteParamsRegionEuropeWest8         SpeedDeleteParamsRegion = "europe-west8"
	SpeedDeleteParamsRegionEuropeWest9         SpeedDeleteParamsRegion = "europe-west9"
	SpeedDeleteParamsRegionMeWest1             SpeedDeleteParamsRegion = "me-west1"
	SpeedDeleteParamsRegionSouthamericaEast1   SpeedDeleteParamsRegion = "southamerica-east1"
	SpeedDeleteParamsRegionUsCentral1          SpeedDeleteParamsRegion = "us-central1"
	SpeedDeleteParamsRegionUsEast1             SpeedDeleteParamsRegion = "us-east1"
	SpeedDeleteParamsRegionUsEast4             SpeedDeleteParamsRegion = "us-east4"
	SpeedDeleteParamsRegionUsSouth1            SpeedDeleteParamsRegion = "us-south1"
	SpeedDeleteParamsRegionUsWest1             SpeedDeleteParamsRegion = "us-west1"
)

func (r SpeedDeleteParamsRegion) IsKnown() bool {
	switch r {
	case SpeedDeleteParamsRegionAsiaEast1, SpeedDeleteParamsRegionAsiaNortheast1, SpeedDeleteParamsRegionAsiaNortheast2, SpeedDeleteParamsRegionAsiaSouth1, SpeedDeleteParamsRegionAsiaSoutheast1, SpeedDeleteParamsRegionAustraliaSoutheast1, SpeedDeleteParamsRegionEuropeNorth1, SpeedDeleteParamsRegionEuropeSouthwest1, SpeedDeleteParamsRegionEuropeWest1, SpeedDeleteParamsRegionEuropeWest2, SpeedDeleteParamsRegionEuropeWest3, SpeedDeleteParamsRegionEuropeWest4, SpeedDeleteParamsRegionEuropeWest8, SpeedDeleteParamsRegionEuropeWest9, SpeedDeleteParamsRegionMeWest1, SpeedDeleteParamsRegionSouthamericaEast1, SpeedDeleteParamsRegionUsCentral1, SpeedDeleteParamsRegionUsEast1, SpeedDeleteParamsRegionUsEast4, SpeedDeleteParamsRegionUsSouth1, SpeedDeleteParamsRegionUsWest1:
		return true
	}
	return false
}

type SpeedDeleteResponseEnvelope struct {
	Errors   interface{}                     `json:"errors,required"`
	Messages interface{}                     `json:"messages,required"`
	Success  interface{}                     `json:"success,required"`
	Result   SpeedDeleteResponse             `json:"result"`
	JSON     speedDeleteResponseEnvelopeJSON `json:"-"`
}

// speedDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedDeleteResponseEnvelope]
type speedDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SpeedScheduleGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A test region.
	Region param.Field[SpeedScheduleGetParamsRegion] `query:"region"`
}

// URLQuery serializes [SpeedScheduleGetParams]'s query parameters as `url.Values`.
func (r SpeedScheduleGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type SpeedScheduleGetParamsRegion string

const (
	SpeedScheduleGetParamsRegionAsiaEast1           SpeedScheduleGetParamsRegion = "asia-east1"
	SpeedScheduleGetParamsRegionAsiaNortheast1      SpeedScheduleGetParamsRegion = "asia-northeast1"
	SpeedScheduleGetParamsRegionAsiaNortheast2      SpeedScheduleGetParamsRegion = "asia-northeast2"
	SpeedScheduleGetParamsRegionAsiaSouth1          SpeedScheduleGetParamsRegion = "asia-south1"
	SpeedScheduleGetParamsRegionAsiaSoutheast1      SpeedScheduleGetParamsRegion = "asia-southeast1"
	SpeedScheduleGetParamsRegionAustraliaSoutheast1 SpeedScheduleGetParamsRegion = "australia-southeast1"
	SpeedScheduleGetParamsRegionEuropeNorth1        SpeedScheduleGetParamsRegion = "europe-north1"
	SpeedScheduleGetParamsRegionEuropeSouthwest1    SpeedScheduleGetParamsRegion = "europe-southwest1"
	SpeedScheduleGetParamsRegionEuropeWest1         SpeedScheduleGetParamsRegion = "europe-west1"
	SpeedScheduleGetParamsRegionEuropeWest2         SpeedScheduleGetParamsRegion = "europe-west2"
	SpeedScheduleGetParamsRegionEuropeWest3         SpeedScheduleGetParamsRegion = "europe-west3"
	SpeedScheduleGetParamsRegionEuropeWest4         SpeedScheduleGetParamsRegion = "europe-west4"
	SpeedScheduleGetParamsRegionEuropeWest8         SpeedScheduleGetParamsRegion = "europe-west8"
	SpeedScheduleGetParamsRegionEuropeWest9         SpeedScheduleGetParamsRegion = "europe-west9"
	SpeedScheduleGetParamsRegionMeWest1             SpeedScheduleGetParamsRegion = "me-west1"
	SpeedScheduleGetParamsRegionSouthamericaEast1   SpeedScheduleGetParamsRegion = "southamerica-east1"
	SpeedScheduleGetParamsRegionUsCentral1          SpeedScheduleGetParamsRegion = "us-central1"
	SpeedScheduleGetParamsRegionUsEast1             SpeedScheduleGetParamsRegion = "us-east1"
	SpeedScheduleGetParamsRegionUsEast4             SpeedScheduleGetParamsRegion = "us-east4"
	SpeedScheduleGetParamsRegionUsSouth1            SpeedScheduleGetParamsRegion = "us-south1"
	SpeedScheduleGetParamsRegionUsWest1             SpeedScheduleGetParamsRegion = "us-west1"
)

func (r SpeedScheduleGetParamsRegion) IsKnown() bool {
	switch r {
	case SpeedScheduleGetParamsRegionAsiaEast1, SpeedScheduleGetParamsRegionAsiaNortheast1, SpeedScheduleGetParamsRegionAsiaNortheast2, SpeedScheduleGetParamsRegionAsiaSouth1, SpeedScheduleGetParamsRegionAsiaSoutheast1, SpeedScheduleGetParamsRegionAustraliaSoutheast1, SpeedScheduleGetParamsRegionEuropeNorth1, SpeedScheduleGetParamsRegionEuropeSouthwest1, SpeedScheduleGetParamsRegionEuropeWest1, SpeedScheduleGetParamsRegionEuropeWest2, SpeedScheduleGetParamsRegionEuropeWest3, SpeedScheduleGetParamsRegionEuropeWest4, SpeedScheduleGetParamsRegionEuropeWest8, SpeedScheduleGetParamsRegionEuropeWest9, SpeedScheduleGetParamsRegionMeWest1, SpeedScheduleGetParamsRegionSouthamericaEast1, SpeedScheduleGetParamsRegionUsCentral1, SpeedScheduleGetParamsRegionUsEast1, SpeedScheduleGetParamsRegionUsEast4, SpeedScheduleGetParamsRegionUsSouth1, SpeedScheduleGetParamsRegionUsWest1:
		return true
	}
	return false
}

type SpeedScheduleGetResponseEnvelope struct {
	Errors   interface{} `json:"errors,required"`
	Messages interface{} `json:"messages,required"`
	Success  interface{} `json:"success,required"`
	// The test schedule.
	Result Schedule                             `json:"result"`
	JSON   speedScheduleGetResponseEnvelopeJSON `json:"-"`
}

// speedScheduleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedScheduleGetResponseEnvelope]
type speedScheduleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedScheduleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedScheduleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SpeedTrendsListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The type of device.
	DeviceType param.Field[SpeedTrendsListParamsDeviceType] `query:"deviceType,required"`
	// A comma-separated list of metrics to include in the results.
	Metrics param.Field[string] `query:"metrics,required"`
	// A test region.
	Region param.Field[SpeedTrendsListParamsRegion] `query:"region,required"`
	Start  param.Field[time.Time]                   `query:"start,required" format:"date-time"`
	// The timezone of the start and end timestamps.
	Tz  param.Field[string]    `query:"tz,required"`
	End param.Field[time.Time] `query:"end" format:"date-time"`
}

// URLQuery serializes [SpeedTrendsListParams]'s query parameters as `url.Values`.
func (r SpeedTrendsListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of device.
type SpeedTrendsListParamsDeviceType string

const (
	SpeedTrendsListParamsDeviceTypeDesktop SpeedTrendsListParamsDeviceType = "DESKTOP"
	SpeedTrendsListParamsDeviceTypeMobile  SpeedTrendsListParamsDeviceType = "MOBILE"
)

func (r SpeedTrendsListParamsDeviceType) IsKnown() bool {
	switch r {
	case SpeedTrendsListParamsDeviceTypeDesktop, SpeedTrendsListParamsDeviceTypeMobile:
		return true
	}
	return false
}

// A test region.
type SpeedTrendsListParamsRegion string

const (
	SpeedTrendsListParamsRegionAsiaEast1           SpeedTrendsListParamsRegion = "asia-east1"
	SpeedTrendsListParamsRegionAsiaNortheast1      SpeedTrendsListParamsRegion = "asia-northeast1"
	SpeedTrendsListParamsRegionAsiaNortheast2      SpeedTrendsListParamsRegion = "asia-northeast2"
	SpeedTrendsListParamsRegionAsiaSouth1          SpeedTrendsListParamsRegion = "asia-south1"
	SpeedTrendsListParamsRegionAsiaSoutheast1      SpeedTrendsListParamsRegion = "asia-southeast1"
	SpeedTrendsListParamsRegionAustraliaSoutheast1 SpeedTrendsListParamsRegion = "australia-southeast1"
	SpeedTrendsListParamsRegionEuropeNorth1        SpeedTrendsListParamsRegion = "europe-north1"
	SpeedTrendsListParamsRegionEuropeSouthwest1    SpeedTrendsListParamsRegion = "europe-southwest1"
	SpeedTrendsListParamsRegionEuropeWest1         SpeedTrendsListParamsRegion = "europe-west1"
	SpeedTrendsListParamsRegionEuropeWest2         SpeedTrendsListParamsRegion = "europe-west2"
	SpeedTrendsListParamsRegionEuropeWest3         SpeedTrendsListParamsRegion = "europe-west3"
	SpeedTrendsListParamsRegionEuropeWest4         SpeedTrendsListParamsRegion = "europe-west4"
	SpeedTrendsListParamsRegionEuropeWest8         SpeedTrendsListParamsRegion = "europe-west8"
	SpeedTrendsListParamsRegionEuropeWest9         SpeedTrendsListParamsRegion = "europe-west9"
	SpeedTrendsListParamsRegionMeWest1             SpeedTrendsListParamsRegion = "me-west1"
	SpeedTrendsListParamsRegionSouthamericaEast1   SpeedTrendsListParamsRegion = "southamerica-east1"
	SpeedTrendsListParamsRegionUsCentral1          SpeedTrendsListParamsRegion = "us-central1"
	SpeedTrendsListParamsRegionUsEast1             SpeedTrendsListParamsRegion = "us-east1"
	SpeedTrendsListParamsRegionUsEast4             SpeedTrendsListParamsRegion = "us-east4"
	SpeedTrendsListParamsRegionUsSouth1            SpeedTrendsListParamsRegion = "us-south1"
	SpeedTrendsListParamsRegionUsWest1             SpeedTrendsListParamsRegion = "us-west1"
)

func (r SpeedTrendsListParamsRegion) IsKnown() bool {
	switch r {
	case SpeedTrendsListParamsRegionAsiaEast1, SpeedTrendsListParamsRegionAsiaNortheast1, SpeedTrendsListParamsRegionAsiaNortheast2, SpeedTrendsListParamsRegionAsiaSouth1, SpeedTrendsListParamsRegionAsiaSoutheast1, SpeedTrendsListParamsRegionAustraliaSoutheast1, SpeedTrendsListParamsRegionEuropeNorth1, SpeedTrendsListParamsRegionEuropeSouthwest1, SpeedTrendsListParamsRegionEuropeWest1, SpeedTrendsListParamsRegionEuropeWest2, SpeedTrendsListParamsRegionEuropeWest3, SpeedTrendsListParamsRegionEuropeWest4, SpeedTrendsListParamsRegionEuropeWest8, SpeedTrendsListParamsRegionEuropeWest9, SpeedTrendsListParamsRegionMeWest1, SpeedTrendsListParamsRegionSouthamericaEast1, SpeedTrendsListParamsRegionUsCentral1, SpeedTrendsListParamsRegionUsEast1, SpeedTrendsListParamsRegionUsEast4, SpeedTrendsListParamsRegionUsSouth1, SpeedTrendsListParamsRegionUsWest1:
		return true
	}
	return false
}

type SpeedTrendsListResponseEnvelope struct {
	Errors   interface{}                         `json:"errors,required"`
	Messages interface{}                         `json:"messages,required"`
	Success  interface{}                         `json:"success,required"`
	Result   Trend                               `json:"result"`
	JSON     speedTrendsListResponseEnvelopeJSON `json:"-"`
}

// speedTrendsListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpeedTrendsListResponseEnvelope]
type speedTrendsListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTrendsListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedTrendsListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
