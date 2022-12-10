package http

// HttpResponse is a response struct for http handler
type HttpResponse struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}
