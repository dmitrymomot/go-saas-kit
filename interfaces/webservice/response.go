package webservice

// Response structure
type Response struct {
	Status int
	Data   interface{}
	Meta   map[string]interface{}
}

// GetCode returns http status code of the response
func (r *Response) GetCode() int {
	return r.Status
}

// GetData returns the response data
func (r *Response) GetData() interface{} {
	return r.Data
}

// GetMeta returns a map with response meta data
func (r *Response) GetMeta() map[string]interface{} {
	return r.Meta
}
