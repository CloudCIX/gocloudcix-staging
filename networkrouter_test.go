// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocloudcix_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/CloudCIX/gocloudcix"
	"github.com/CloudCIX/gocloudcix/internal/testutil"
	"github.com/CloudCIX/gocloudcix/option"
)

func TestNetworkRouterNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gocloudcix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Network.Routers.New(context.TODO(), gocloudcix.NetworkRouterNewParams{
		ProjectID: 1,
		Metadata: gocloudcix.NetworkRouterNewParamsMetadata{
			Destination: gocloudcix.String("destination"),
			Nat:         gocloudcix.Bool(true),
			Nexthop:     gocloudcix.String("nexthop"),
		},
		Networks: []gocloudcix.NetworkRouterNewParamsNetwork{{
			Ipv4: gocloudcix.String("10.0.1.0/24"),
			Ipv6: gocloudcix.String("ipv6"),
			Name: gocloudcix.String("web-tier"),
			Vlan: gocloudcix.Int(0),
		}, {
			Ipv4: gocloudcix.String("10.0.2.0/24"),
			Ipv6: gocloudcix.String("ipv6"),
			Name: gocloudcix.String("app-tier"),
			Vlan: gocloudcix.Int(0),
		}, {
			Ipv4: gocloudcix.String("10.0.3.0/24"),
			Ipv6: gocloudcix.String("ipv6"),
			Name: gocloudcix.String("db-tier"),
			Vlan: gocloudcix.Int(0),
		}},
		Type: gocloudcix.String("router"),
	})
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkRouterGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gocloudcix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Network.Routers.Get(context.TODO(), 0)
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkRouterUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gocloudcix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Network.Routers.Update(
		context.TODO(),
		0,
		gocloudcix.NetworkRouterUpdateParams{
			NetworkRouterUpdate: gocloudcix.NetworkRouterUpdateParam{
				State: "state",
				Metadata: gocloudcix.NetworkRouterUpdateMetadataParam{
					Destination: gocloudcix.String("destination"),
					Nat:         gocloudcix.Bool(true),
					Nexthop:     gocloudcix.String("nexthop"),
				},
				Networks: []gocloudcix.NetworkRouterUpdateNetworkParam{{
					Ipv4: gocloudcix.String("10.0.1.0/24"),
					Ipv6: gocloudcix.String("fd00::/64"),
					Name: gocloudcix.String("existing-network"),
					Vlan: gocloudcix.Int(100),
				}, {
					Ipv4: gocloudcix.String("10.0.50.0/24"),
					Ipv6: gocloudcix.String("ipv6"),
					Name: gocloudcix.String("new-database-network"),
					Vlan: gocloudcix.Int(0),
				}},
			},
		},
	)
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkRouterListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gocloudcix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Network.Routers.List(context.TODO(), gocloudcix.NetworkRouterListParams{
		Exclude: map[string]any{
			"foo": "bar",
		},
		Limit: gocloudcix.Int(0),
		Order: gocloudcix.String("order"),
		Page:  gocloudcix.Int(0),
		Search: map[string]any{
			"foo": "bar",
		},
	})
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkRouterDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gocloudcix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Network.Routers.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
