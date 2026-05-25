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

func TestComputeInstanceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Compute.Instances.New(context.TODO(), gocloudcix.ComputeInstanceNewParams{
		Metadata: gocloudcix.ComputeInstanceNewParamsMetadata{
			DNS:          gocloudcix.String("dns"),
			InstanceType: gocloudcix.String("instance_type"),
			Userdata:     gocloudcix.String("userdata"),
		},
		ProjectID: 1,
		Specs: []gocloudcix.ComputeInstanceNewParamsSpec{{
			Quantity: gocloudcix.Int(8),
			SKUName:  gocloudcix.String("RAM_001"),
		}, {
			Quantity: gocloudcix.Int(4),
			SKUName:  gocloudcix.String("vCPU_002"),
		}, {
			Quantity: gocloudcix.Int(100),
			SKUName:  gocloudcix.String("SSD_001"),
		}, {
			Quantity: gocloudcix.Int(1),
			SKUName:  gocloudcix.String("MSServer2022"),
		}},
		GracePeriod: gocloudcix.Int(72),
		Interfaces: []gocloudcix.ComputeInstanceNewParamsInterface{{
			Gateway: gocloudcix.Bool(true),
			Ipv4Addresses: []gocloudcix.ComputeInstanceNewParamsInterfaceIpv4Address{{
				Address: gocloudcix.String("10.0.0.10"),
				Nat:     gocloudcix.Bool(true),
			}},
			Ipv6Addresses: []gocloudcix.ComputeInstanceNewParamsInterfaceIpv6Address{{
				Address: gocloudcix.String("2a02:2078:9:1234::20"),
			}},
		}},
		Name: gocloudcix.String("windows-server"),
		Type: gocloudcix.String("hyperv"),
	})
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputeInstanceGet(t *testing.T) {
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
	_, err := client.Compute.Instances.Get(context.TODO(), 0)
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputeInstanceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Compute.Instances.Update(
		context.TODO(),
		0,
		gocloudcix.ComputeInstanceUpdateParams{
			ComputeInstanceUpdate: gocloudcix.ComputeInstanceUpdateParam{
				Metadata: gocloudcix.ComputeInstanceUpdateMetadataParam{
					DNS:          gocloudcix.String("dns"),
					InstanceType: gocloudcix.String("instance_type"),
					Userdata:     gocloudcix.String("userdata"),
				},
				Specs: []gocloudcix.ComputeInstanceUpdateSpecParam{{
					Quantity: gocloudcix.Int(0),
					SKUName:  gocloudcix.String("sku_name"),
				}},
				GracePeriod: gocloudcix.Int(0),
				Interfaces: []gocloudcix.ComputeInstanceUpdateInterfaceParam{{
					Gateway: gocloudcix.Bool(true),
					Ipv4Addresses: []gocloudcix.ComputeInstanceUpdateInterfaceIpv4AddressParam{{
						Address: gocloudcix.String("address"),
						Nat:     gocloudcix.Bool(true),
					}},
					Ipv6Addresses: []gocloudcix.ComputeInstanceUpdateInterfaceIpv6AddressParam{{
						Address: gocloudcix.String("address"),
					}},
				}},
				Name:  gocloudcix.String("name"),
				State: gocloudcix.String("delete"),
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

func TestComputeInstanceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Compute.Instances.List(context.TODO(), gocloudcix.ComputeInstanceListParams{
		Exclude: map[string]any{},
		Limit:   gocloudcix.Int(0),
		Order:   gocloudcix.String("order"),
		Page:    gocloudcix.Int(0),
		Search:  map[string]any{},
	})
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputeInstanceDelete(t *testing.T) {
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
	err := client.Compute.Instances.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
