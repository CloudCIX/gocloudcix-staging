// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocloudcix

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/CloudCIX/gocloudcix/auth"
	"github.com/CloudCIX/gocloudcix/config"
	"github.com/CloudCIX/gocloudcix/internal/requestconfig"
	"github.com/CloudCIX/gocloudcix/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the gocloudcix API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options      []option.RequestOption
	Compute      ComputeService
	Network      NetworkService
	// Management of Cloud Projects
	//
	// This module provides API endpoints for managing cloud projects in the CloudCIX
	// Compute platform. Projects are logical containers that organise and group your
	// cloud infrastructure resources such as virtual machines, routers, firewalls, and
	// storage. Each project belongs to a specific region and has its own isolated
	// network environment.
	//
	// Available operations:
	//
	// - List and filter projects across your organization
	// - Create new projects in available cloud regions
	// - Retrieve detailed project information including region and manager
	// - Update project details such as name and notes
	//
	// Each project includes its associated address, region, manager, and creation
	// metadata.
	Project      ProjectService
	Storage      StorageService
	tokenManager *auth.TokenManager
	settings     *config.Settings
}

// DefaultClientOptions read from the environment (GOCLOUDCIX_API_KEY,
// GOCLOUDCIX_BASE_URL). This should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithHTTPClient(defaultHTTPClient()), option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("GOCLOUDCIX_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("GOCLOUDCIX_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("GOCLOUDCIX_CUSTOM_HEADERS"); ok {
		for _, line := range strings.Split(o, "\n") {
			colon := strings.Index(line, ":")
			if colon >= 0 {
				defaults = append(defaults, option.WithHeader(strings.TrimSpace(line[:colon]), strings.TrimSpace(line[colon+1:])))
			}
		}
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (GOCLOUDCIX_API_KEY, GOCLOUDCIX_BASE_URL). The option passed in as
// arguments are applied after these default arguments, and all option will be
// passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{Options: opts}

	r.Compute = NewComputeService(opts...)
	r.Network = NewNetworkService(opts...)
	r.Project = NewProjectService(opts...)
	r.Storage = NewStorageService(opts...)

	return
}

// NewClientWithAuth creates a client with automatic authentication from environment variables.
// Returns an error if credentials are not configured.
func NewClientWithAuth(opts ...option.RequestOption) (*Client, error) {
	settings, err := config.LoadSettings()
	if err != nil {
		return nil, fmt.Errorf("failed to load settings: %w", err)
	}

	if err := settings.Validate(); err != nil {
		return nil, fmt.Errorf("authentication credentials not configured: %w", err)
	}

	client := NewClientWithSettings(settings, opts...)
	return &client, nil
}

// NewClientFromFile creates a client loading credentials from a settings file.
func NewClientFromFile(settingsFile string, opts ...option.RequestOption) (*Client, error) {
	settings, err := config.LoadSettings(settingsFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load settings from file: %w", err)
	}

	if err := settings.Validate(); err != nil {
		return nil, fmt.Errorf("invalid settings: %w", err)
	}

	client := NewClientWithSettings(settings, opts...)
	return &client, nil
}

// NewClientWithSettings creates a client with the provided settings and automatic authentication.
func NewClientWithSettings(settings *config.Settings, opts ...option.RequestOption) Client {
	tokenManager := auth.NewTokenManager(settings)

	// Add authentication middleware
	authOpts := []option.RequestOption{
		option.WithBaseURL(settings.ComputeURL()),
		option.WithMiddleware(auth.AuthRetryMiddleware(tokenManager)),
		auth.WithAutoAuth(tokenManager),
	}

	opts = append(authOpts, opts...)
	opts = append(DefaultClientOptions(), opts...)

	client := Client{
		Options:      opts,
		tokenManager: tokenManager,
		settings:     settings,
	}

	client.Compute = NewComputeService(opts...)
	client.Network = NewNetworkService(opts...)
	client.Project = NewProjectService(opts...)
	client.Storage = NewStorageService(opts...)

	return client
}

// RefreshToken manually refreshes the authentication token
func (r *Client) RefreshToken(ctx context.Context) error {
	if r.tokenManager == nil {
		return fmt.Errorf("client not configured with authentication")
	}

	r.tokenManager.InvalidateToken()
	_, err := r.tokenManager.GetValidToken(ctx)
	return err
}

// GetSettings returns the client's settings if configured with authentication
func (r *Client) GetSettings() *config.Settings {
	return r.settings
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params any, res any, opts ...option.RequestOption) error {
	opts = slices.Concat(r.Options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
