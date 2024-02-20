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

func TestUserLoadBalancerPoolNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.LoadBalancers.Pools.New(context.TODO(), cloudflare.UserLoadBalancerPoolNewParams{
		Name: cloudflare.F("primary-dc-1"),
		Origins: cloudflare.F([]cloudflare.UserLoadBalancerPoolNewParamsOrigin{{
			Address: cloudflare.F("0.0.0.0"),
			Enabled: cloudflare.F(true),
			Header: cloudflare.F(cloudflare.UserLoadBalancerPoolNewParamsOriginsHeader{
				Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
			}),
			Name:             cloudflare.F("app-server-1"),
			VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
			Weight:           cloudflare.F(0.600000),
		}, {
			Address: cloudflare.F("0.0.0.0"),
			Enabled: cloudflare.F(true),
			Header: cloudflare.F(cloudflare.UserLoadBalancerPoolNewParamsOriginsHeader{
				Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
			}),
			Name:             cloudflare.F("app-server-1"),
			VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
			Weight:           cloudflare.F(0.600000),
		}, {
			Address: cloudflare.F("0.0.0.0"),
			Enabled: cloudflare.F(true),
			Header: cloudflare.F(cloudflare.UserLoadBalancerPoolNewParamsOriginsHeader{
				Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
			}),
			Name:             cloudflare.F("app-server-1"),
			VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
			Weight:           cloudflare.F(0.600000),
		}}),
		CheckRegions: cloudflare.F([]cloudflare.UserLoadBalancerPoolNewParamsCheckRegion{cloudflare.UserLoadBalancerPoolNewParamsCheckRegionWeu, cloudflare.UserLoadBalancerPoolNewParamsCheckRegionEnam}),
		Description:  cloudflare.F("Primary data center - Provider XYZ"),
		Enabled:      cloudflare.F(false),
		Latitude:     cloudflare.F(0.000000),
		LoadShedding: cloudflare.F(cloudflare.UserLoadBalancerPoolNewParamsLoadShedding{
			DefaultPercent: cloudflare.F(0.000000),
			DefaultPolicy:  cloudflare.F(cloudflare.UserLoadBalancerPoolNewParamsLoadSheddingDefaultPolicyRandom),
			SessionPercent: cloudflare.F(0.000000),
			SessionPolicy:  cloudflare.F(cloudflare.UserLoadBalancerPoolNewParamsLoadSheddingSessionPolicyHash),
		}),
		Longitude:         cloudflare.F(0.000000),
		MinimumOrigins:    cloudflare.F(int64(0)),
		Monitor:           cloudflare.F[any](map[string]interface{}{}),
		NotificationEmail: cloudflare.F("someone@example.com,sometwo@example.com"),
		NotificationFilter: cloudflare.F(cloudflare.UserLoadBalancerPoolNewParamsNotificationFilter{
			Origin: cloudflare.F(cloudflare.UserLoadBalancerPoolNewParamsNotificationFilterOrigin{
				Disable: cloudflare.F(true),
				Healthy: cloudflare.F(true),
			}),
			Pool: cloudflare.F(cloudflare.UserLoadBalancerPoolNewParamsNotificationFilterPool{
				Disable: cloudflare.F(true),
				Healthy: cloudflare.F(false),
			}),
		}),
		OriginSteering: cloudflare.F(cloudflare.UserLoadBalancerPoolNewParamsOriginSteering{
			Policy: cloudflare.F(cloudflare.UserLoadBalancerPoolNewParamsOriginSteeringPolicyRandom),
		}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoadBalancerPoolUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.LoadBalancers.Pools.Update(
		context.TODO(),
		"17b5962d775c646f3f9725cbc7a53df4",
		cloudflare.UserLoadBalancerPoolUpdateParams{
			CheckRegions: cloudflare.F([]cloudflare.UserLoadBalancerPoolUpdateParamsCheckRegion{cloudflare.UserLoadBalancerPoolUpdateParamsCheckRegionWeu, cloudflare.UserLoadBalancerPoolUpdateParamsCheckRegionEnam}),
			Description:  cloudflare.F("Primary data center - Provider XYZ"),
			Enabled:      cloudflare.F(false),
			Latitude:     cloudflare.F(0.000000),
			LoadShedding: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsLoadShedding{
				DefaultPercent: cloudflare.F(0.000000),
				DefaultPolicy:  cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicyRandom),
				SessionPercent: cloudflare.F(0.000000),
				SessionPolicy:  cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsLoadSheddingSessionPolicyHash),
			}),
			Longitude:         cloudflare.F(0.000000),
			MinimumOrigins:    cloudflare.F(int64(0)),
			Monitor:           cloudflare.F[any](map[string]interface{}{}),
			Name:              cloudflare.F("primary-dc-1"),
			NotificationEmail: cloudflare.F("someone@example.com,sometwo@example.com"),
			NotificationFilter: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsNotificationFilter{
				Origin: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsNotificationFilterOrigin{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(true),
				}),
				Pool: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsNotificationFilterPool{
					Disable: cloudflare.F(true),
					Healthy: cloudflare.F(false),
				}),
			}),
			OriginSteering: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsOriginSteering{
				Policy: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsOriginSteeringPolicyRandom),
			}),
			Origins: cloudflare.F([]cloudflare.UserLoadBalancerPoolUpdateParamsOrigin{{
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}, {
				Address: cloudflare.F("0.0.0.0"),
				Enabled: cloudflare.F(true),
				Header: cloudflare.F(cloudflare.UserLoadBalancerPoolUpdateParamsOriginsHeader{
					Host: cloudflare.F([]string{"example.com", "example.com", "example.com"}),
				}),
				Name:             cloudflare.F("app-server-1"),
				VirtualNetworkID: cloudflare.F("a5624d4e-044a-4ff0-b3e1-e2465353d4b4"),
				Weight:           cloudflare.F(0.600000),
			}}),
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

func TestUserLoadBalancerPoolListWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.LoadBalancers.Pools.List(context.TODO(), cloudflare.UserLoadBalancerPoolListParams{
		Monitor: cloudflare.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoadBalancerPoolDelete(t *testing.T) {
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
	_, err := client.Users.LoadBalancers.Pools.Delete(context.TODO(), "17b5962d775c646f3f9725cbc7a53df4")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoadBalancerPoolGet(t *testing.T) {
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
	_, err := client.Users.LoadBalancers.Pools.Get(context.TODO(), "17b5962d775c646f3f9725cbc7a53df4")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
