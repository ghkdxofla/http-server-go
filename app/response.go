package main

import (
	"fmt"
	"strconv"
)

type Response struct {
	Version        string
	Status         Status
	ResponseHeader *ContentHeader
	Data           any
}

func NewResponse(version string, status Status, contentType *string, contentLength *int, data any) Response {
	responseHeader := NewContentHeader(contentType, contentLength, data)

	return Response{
		Version:        version,
		Status:         status,
		ResponseHeader: responseHeader,
		Data:           data,
	}
}

func (r *Response) ToString() string {
	dataString := ""

	if r.Data != nil {
		dataString = fmt.Sprintf("%s", r.Data)
	}

	return r.Version + " " + r.Status.ToString() + "\r\n" + "Content-Type: " + r.ResponseHeader.ContentType + "\r\n" + "Content-Length: " + strconv.Itoa(r.ResponseHeader.ContentLength) + "\r\n\r\n" + dataString
}
