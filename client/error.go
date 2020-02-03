package client

type (
	// Error is CircleCI v2 API's error response body.
	Error struct {
		Message string `json:"message"`
	}
)

// Error returns an error's description.
func (e *Error) Error() string {
	return ""
}
