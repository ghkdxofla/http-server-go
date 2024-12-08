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

func NewResponse(version string, status Status, contentType *string, contentLength *int, contentEncoding []string, data any) Response {
	responseHeader := NewContentHeader(contentType, contentLength, contentEncoding, data)

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

	result := r.Version + " " +
		r.Status.ToString() + "\r\n" +
		"Content-Type: " + r.ResponseHeader.ContentType + "\r\n"

	if r.ResponseHeader.ContentEncoding != nil {
		contentEncodingString := ""
		for _, encoding := range r.ResponseHeader.ContentEncoding {
			contentEncodingString += encoding + ", "
		}
		contentEncodingString = contentEncodingString[:len(contentEncodingString)-2]

		result += "Content-Encoding: " + contentEncodingString + "\r\n"
	}

	result += "Content-Length: " + strconv.Itoa(r.ResponseHeader.ContentLength) + "\r\n\r\n" + dataString

	return result
}
