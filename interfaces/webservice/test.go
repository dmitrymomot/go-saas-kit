package webservice

import "net/http"

// Test struct
type Test struct {
}

// GetData handler
func (ws *Test) GetData() (Response, error) {
	return Response{Status: http.StatusOK, Data: "test data"}, nil
}
