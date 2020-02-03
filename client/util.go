package client

import (
	"encoding/json"
	"net/http"
)

// IsErrorDefault is a default function for client to judge the request is successful or not by the response.
func IsErrorDefault(resp *http.Response) bool {
	return resp.StatusCode >= 400
}

// ParseRespDefault is the default function for client to process the succeeded request's response.
func ParseRespDefault(resp *http.Response, output interface{}) error {
	if output == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(output)
}

// ParseErrorRespDefault is the default function for client to process the failed request's response.
func ParseErrorRespDefault(resp *http.Response) error {
	a := &struct {
		Errors *Error
	}{}
	if err := json.NewDecoder(resp.Body).Decode(a); err != nil {
		return err
	}
	return a.Errors
}
