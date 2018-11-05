package middleware

import "net/http"

// Predefined response header contants
const (
	headerContentType   = "Content-Type"
	mimeApplicationJSON = "application/json; charset=UTF-8"
)

// APIResponseHeaders middleware
// Adds application/json content type to each API response
func APIResponseHeaders(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
		w.Header().Set(headerContentType, mimeApplicationJSON)
	}
}
