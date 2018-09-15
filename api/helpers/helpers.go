package helpers

import (
	"net/http"
)

// GetBody ...
func GetBody(r *http.Request) []byte {
	// read the body content
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	return body
}
