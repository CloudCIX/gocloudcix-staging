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

func TestProjectNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Project.New(context.TODO(), gocloudcix.ProjectNewParams{
		Name:     "production-environment",
		RegionID: 0,
		Note:     gocloudcix.String("Main production infrastructure for web services"),
	})
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectGet(t *testing.T) {
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
	_, err := client.Project.Get(context.TODO(), 0)
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Project.Update(
		context.TODO(),
		0,
		gocloudcix.ProjectUpdateParams{
			ProjectUpdate: gocloudcix.ProjectUpdateParam{
				Name: gocloudcix.String("name"),
				Note: gocloudcix.String("note"),
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

func TestProjectListWithOptionalParams(t *testing.T) {
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
	_, err := client.Project.List(context.TODO(), gocloudcix.ProjectListParams{
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

func TestProjectDelete(t *testing.T) {
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
	err := client.Project.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *gocloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
