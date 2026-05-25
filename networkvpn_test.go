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

func TestNetworkVpnNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Network.Vpns.New(context.TODO(), gocloudcix.NetworkVpnNewParams{
		ProjectID: 1,
		Metadata: gocloudcix.NetworkVpnNewParamsMetadata{
			ChildSas: []gocloudcix.NetworkVpnNewParamsMetadataChildSa{{
				LocalTs:  gocloudcix.String("10.0.0.0/24"),
				RemoteTs: gocloudcix.String("172.16.10.0/27"),
			}},
			IkeAuthentication:   gocloudcix.String("SHA256"),
			IkeDhGroups:         gocloudcix.String("Group 24"),
			IkeEncryption:       gocloudcix.String("256 bit AES-CBC"),
			IkeGatewayType:      gocloudcix.String("hostname"),
			IkeGatewayValue:     gocloudcix.String("hostname.example.com"),
			IkeLifetime:         gocloudcix.Int(0),
			IkePreSharedKey:     gocloudcix.String("R4nd0mKey!"),
			IkeVersion:          gocloudcix.String("v2-only"),
			IpsecAuthentication: gocloudcix.String("SHA256"),
			IpsecEncryption:     gocloudcix.String("AES 256 GCM"),
			IpsecEstablishTime:  gocloudcix.String("On Traffic"),
			IpsecLifetime:       gocloudcix.Int(3000),
			IpsecPfsGroups:      gocloudcix.String("Group 24"),
			TrafficSelector:     gocloudcix.Bool(true),
		},
		Name: gocloudcix.String("Cork Office to Dublin Office"),
		Type: gocloudcix.String("site-to-site"),
	})
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkVpnGet(t *testing.T) {
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
	_, err := client.Network.Vpns.Get(context.TODO(), 0)
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkVpnUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Network.Vpns.Update(
		context.TODO(),
		0,
		gocloudcix.NetworkVpnUpdateParams{
			NetworkVpnUpdate: gocloudcix.NetworkVpnUpdateParam{
				Metadata: gocloudcix.NetworkVpnUpdateMetadataParam{
					ChildSas: []gocloudcix.NetworkVpnUpdateMetadataChildSaParam{{
						LocalTs:  gocloudcix.String("local_ts"),
						RemoteTs: gocloudcix.String("remote_ts"),
					}},
					IkeAuthentication:   gocloudcix.String("ike_authentication"),
					IkeDhGroups:         gocloudcix.String("ike_dh_groups"),
					IkeEncryption:       gocloudcix.String("ike_encryption"),
					IkeGatewayType:      gocloudcix.String("ike_gateway_type"),
					IkeGatewayValue:     gocloudcix.String("ike_gateway_value"),
					IkeLifetime:         gocloudcix.Int(0),
					IkePreSharedKey:     gocloudcix.String("ike_pre_shared_key"),
					IkeVersion:          gocloudcix.String("ike_version"),
					IpsecAuthentication: gocloudcix.String("ipsec_authentication"),
					IpsecEncryption:     gocloudcix.String("ipsec_encryption"),
					IpsecEstablishTime:  gocloudcix.String("ipsec_establish_time"),
					IpsecLifetime:       gocloudcix.Int(0),
					IpsecPfsGroups:      gocloudcix.String("ipsec_pfs_groups"),
					TrafficSelector:     gocloudcix.Bool(true),
				},
				Name:  gocloudcix.String("Cork Office to Dublin Office"),
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

func TestNetworkVpnListWithOptionalParams(t *testing.T) {
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
	_, err := client.Network.Vpns.List(context.TODO(), gocloudcix.NetworkVpnListParams{
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

func TestNetworkVpnDelete(t *testing.T) {
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
	err := client.Network.Vpns.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
