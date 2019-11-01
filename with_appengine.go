package firepower

import (
	"context"
	"net/http"
)

func withContext(ctx context.Context, req *http.Request) *http.Request {
	// No-op because App Engine adds context to a request differently.
	return req
}
