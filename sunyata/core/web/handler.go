package web

import "net/http"

// HTTPHandler Protocol represents an HTTP Handler
type HTTPHandler interface {
	HandleRequest(r *http.Request)
}
