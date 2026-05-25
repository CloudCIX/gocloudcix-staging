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

func TestComputeBackupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Compute.Backups.New(context.TODO(), gocloudcix.ComputeBackupNewParams{
		InstanceID: 456,
		ProjectID:  1,
		Name:       gocloudcix.String("critical-windows-backup"),
		Type:       gocloudcix.String("hyperv"),
	})
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputeBackupGet(t *testing.T) {
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
	_, err := client.Compute.Backups.Get(context.TODO(), 0)
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputeBackupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Compute.Backups.Update(
		context.TODO(),
		0,
		gocloudcix.ComputeBackupUpdateParams{
			ComputeBackupUpdate: gocloudcix.ComputeBackupUpdateParam{
				Name:  gocloudcix.String("old-backup-to-remove"),
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

func TestComputeBackupListWithOptionalParams(t *testing.T) {
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
	_, err := client.Compute.Backups.List(context.TODO(), gocloudcix.ComputeBackupListParams{
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

func TestComputeBackupDelete(t *testing.T) {
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
	err := client.Compute.Backups.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
