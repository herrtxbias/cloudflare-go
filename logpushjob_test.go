// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestLogpushJobNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Logpush.Jobs.New(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.LogpushJobNewParams{
			DestinationConf: cloudflare.F("s3://mybucket/logs?region=us-west-2"),
			Dataset:         cloudflare.F("http_requests"),
			Enabled:         cloudflare.F(false),
			Frequency:       cloudflare.F(cloudflare.LogpushJobNewParamsFrequencyHigh),
			LogpullOptions:  cloudflare.F("fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339"),
			Name:            cloudflare.F("example.com"),
			OutputOptions: cloudflare.F(cloudflare.LogpushJobNewParamsOutputOptions{
				Cve2021_4428:    cloudflare.F(true),
				BatchPrefix:     cloudflare.F("string"),
				BatchSuffix:     cloudflare.F("string"),
				FieldDelimiter:  cloudflare.F("string"),
				FieldNames:      cloudflare.F([]string{"ClientIP", "EdgeStartTimestamp", "RayID"}),
				OutputType:      cloudflare.F(cloudflare.LogpushJobNewParamsOutputOptionsOutputTypeNdjson),
				RecordDelimiter: cloudflare.F("string"),
				RecordPrefix:    cloudflare.F("string"),
				RecordSuffix:    cloudflare.F("string"),
				RecordTemplate:  cloudflare.F("string"),
				SampleRate:      cloudflare.F(0.000000),
				TimestampFormat: cloudflare.F(cloudflare.LogpushJobNewParamsOutputOptionsTimestampFormatUnixnano),
			}),
			OwnershipChallenge: cloudflare.F("00000000000000000000"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogpushJobList(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Logpush.Jobs.List(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogpushJobDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Logpush.Jobs.Delete(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		int64(1),
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogpushJobGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Logpush.Jobs.Get(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		int64(1),
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogpushJobReplaceWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Logpush.Jobs.Replace(
		context.TODO(),
		"string",
		"023e105f4ecef8ad9ca31a8372d0c353",
		int64(1),
		cloudflare.LogpushJobReplaceParams{
			DestinationConf: cloudflare.F("s3://mybucket/logs?region=us-west-2"),
			Enabled:         cloudflare.F(false),
			Frequency:       cloudflare.F(cloudflare.LogpushJobReplaceParamsFrequencyHigh),
			LogpullOptions:  cloudflare.F("fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339"),
			OutputOptions: cloudflare.F(cloudflare.LogpushJobReplaceParamsOutputOptions{
				Cve2021_4428:    cloudflare.F(true),
				BatchPrefix:     cloudflare.F("string"),
				BatchSuffix:     cloudflare.F("string"),
				FieldDelimiter:  cloudflare.F("string"),
				FieldNames:      cloudflare.F([]string{"ClientIP", "EdgeStartTimestamp", "RayID"}),
				OutputType:      cloudflare.F(cloudflare.LogpushJobReplaceParamsOutputOptionsOutputTypeNdjson),
				RecordDelimiter: cloudflare.F("string"),
				RecordPrefix:    cloudflare.F("string"),
				RecordSuffix:    cloudflare.F("string"),
				RecordTemplate:  cloudflare.F("string"),
				SampleRate:      cloudflare.F(0.000000),
				TimestampFormat: cloudflare.F(cloudflare.LogpushJobReplaceParamsOutputOptionsTimestampFormatUnixnano),
			}),
			OwnershipChallenge: cloudflare.F("00000000000000000000"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
