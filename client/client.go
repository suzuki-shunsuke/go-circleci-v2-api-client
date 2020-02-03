package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path/filepath"
)

type (
	// Client is a CircleCI v2 API client.
	// Client should be created by the function NewClient .
	Client struct {
		endpoint       string
		token          string
		httpClient     *http.Client
		isError        IsError
		parseResp      ParseResp
		parseErrorResp ParseErrorResp
	}

	// ParseResp parses a succeeded API response.
	ParseResp func(resp *http.Response, output interface{}) error
	// ParseErrorResp parses an API error response.
	ParseErrorResp func(resp *http.Response) error
	// IsError decides whether the request successes or not.
	IsError func(resp *http.Response) bool
)

var (
	// DefaultEndpoint is the default CircleCI v2 API endpoint.
	DefaultEndpoint = "https://circleci.com/api/v2"
)

// NewClient returns a new client.
func NewClient(token string) *Client {
	return &Client{
		token:          token,
		endpoint:       DefaultEndpoint,
		httpClient:     http.DefaultClient,
		isError:        IsErrorDefault,
		parseResp:      ParseRespDefault,
		parseErrorResp: ParseErrorRespDefault,
	}
}

func (c *Client) getResp(
	ctx context.Context, method, path string, body interface{}, query url.Values,
) (*http.Response, error) {
	endpoint, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	if _, ok := query["circle-token"]; !ok {
		query.Add("circle-token", c.token)
	}

	endpoint.Path = filepath.Join(endpoint.Path, path)
	endpoint.RawQuery = query.Encode()
	var req *http.Request
	if body == nil {
		req, err = http.NewRequest(method, endpoint.String(), nil)
	} else {
		reqBody := &bytes.Buffer{}
		if err := json.NewEncoder(reqBody).Encode(body); err != nil {
			return nil, err
		}
		req, err = http.NewRequestWithContext(ctx, method, endpoint.String(), reqBody)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return c.httpClient.Do(req)
}

func (c *Client) parseResponse(
	resp *http.Response, output interface{},
) error {
	if c.isError(resp) {
		return c.parseErrorResp(resp)
	}
	return c.parseResp(resp, output)
}
