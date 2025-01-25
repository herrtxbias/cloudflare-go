// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/herrtxbias/cloudflare-go/v4"
	"github.com/herrtxbias/cloudflare-go/v4/internal/testutil"
	"github.com/herrtxbias/cloudflare-go/v4/option"
	"github.com/herrtxbias/cloudflare-go/v4/radar"
)

func TestAttackLayer7TimeseriesWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Radar.Attacks.Layer7.Timeseries(context.TODO(), radar.AttackLayer7TimeseriesParams{
		AggInterval:       cloudflare.F(radar.AttackLayer7TimeseriesParamsAggInterval15m),
		ASN:               cloudflare.F([]string{"string"}),
		Attack:            cloudflare.F([]radar.AttackLayer7TimeseriesParamsAttack{radar.AttackLayer7TimeseriesParamsAttackDDoS}),
		Continent:         cloudflare.F([]string{"string"}),
		DateEnd:           cloudflare.F([]time.Time{time.Now()}),
		DateRange:         cloudflare.F([]string{"7d"}),
		DateStart:         cloudflare.F([]time.Time{time.Now()}),
		Format:            cloudflare.F(radar.AttackLayer7TimeseriesParamsFormatJson),
		HTTPMethod:        cloudflare.F([]radar.AttackLayer7TimeseriesParamsHTTPMethod{radar.AttackLayer7TimeseriesParamsHTTPMethodGet}),
		HTTPVersion:       cloudflare.F([]radar.AttackLayer7TimeseriesParamsHTTPVersion{radar.AttackLayer7TimeseriesParamsHTTPVersionHttPv1}),
		IPVersion:         cloudflare.F([]radar.AttackLayer7TimeseriesParamsIPVersion{radar.AttackLayer7TimeseriesParamsIPVersionIPv4}),
		Location:          cloudflare.F([]string{"string"}),
		MitigationProduct: cloudflare.F([]radar.AttackLayer7TimeseriesParamsMitigationProduct{radar.AttackLayer7TimeseriesParamsMitigationProductDDoS}),
		Name:              cloudflare.F([]string{"string"}),
		Normalization:     cloudflare.F(radar.AttackLayer7TimeseriesParamsNormalizationPercentageChange),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
