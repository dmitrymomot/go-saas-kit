package http

import "net/http"

// Predefined response header contants
const (
	contentTypeHeader = "Content-Type"
	jsonContentType   = "application/json; charset=UTF-8"
)

// API response interface
type response interface {
	GetCode() int
	GetData() interface{}
	GetMeta() map[string]interface{}
}

// API endpoint function
type endpointFunc func(params map[string]interface{}) (response, error)

// API endpoint wrapper
func apiHandler(e endpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentTypeHeader, jsonContentType)

		resp, err := e(map[string]interface{}{})
		if err != nil {
			w.WriteHeader(resp.GetCode())
		} else {
			w.WriteHeader(resp.GetCode())
		}
	})
}
