package willy

import "net/http/httptest"

type Response struct {
	*httptest.ResponseRecorder
}

func (r *Response) Location() string {
	return r.Header().Get("Location")
}
