package main

import (
	"strconv"
)

type Response struct {
	Version string
	Status  Status
	Data    interface{}
}

func NewResponse(version string, status Status, data any) *Response {
	return &Response{
		Version: version,
		Status:  status,
		Data:    data,
	}
}

func (r *Response) ToString() string {
	contentType := "text/plain"
	contentLength := 0
	dataString := ""

	if r.Data != nil {
		contentLength = len(r.Data.(string))
		dataString = r.Data.(string)
	}

	return r.Version + " " + r.Status.ToString() + "\r\n" + "Content-Type: " + contentType + "\r\n" + "Content-Length: " + strconv.Itoa(contentLength) + "\r\n\r\n" + dataString
}
