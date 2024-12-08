package main

import (
	"errors"
	"strconv"
	"strings"
)

type RequestHeader struct {
	Host          *string
	UserAgent     *string
	Accept        *string
	ContentHeader *ContentHeader
}

func NewRequestHeader(requests ...string) *RequestHeader {
	requestMap := make(map[string]*string)
	for _, request := range requests {
		key, value, err := separateKeyValue(request)
		if err != nil {
			continue
		}
		requestMap[key] = &value
	}

	requestHeader := RequestHeader{
		Host:          requestMap["Host"],
		UserAgent:     requestMap["User-Agent"],
		Accept:        requestMap["Accept"],
		ContentHeader: nil,
	}
	var err error
	var contentLength int
	if requestMap["Content-Length"] == nil {
		contentLength = 0
	} else {
		contentLength, err = strconv.Atoi(*requestMap["Content-Length"])
		CheckError(err)
	}

	var acceptEncoding []string
	if requestMap["Accept-Encoding"] == nil {
		acceptEncoding = nil
	} else {
		acceptEncoding = strings.Split(*requestMap["Accept-Encoding"], ", ")
	}

	requestHeader.ContentHeader = NewContentHeader(requestMap["Content-Type"], &contentLength, acceptEncoding, nil)

	return &requestHeader
}

type ContentHeader struct {
	ContentType     string
	ContentLength   int
	ContentEncoding []string
}

func NewContentHeader(contentType *string, contentLength *int, acceptEncoding []string, data any) *ContentHeader {
	if contentType == nil {
		contentType = new(string)
		*contentType = "text/plain"
	}

	if contentLength == nil {
		contentLength = new(int)
		*contentLength = 0
	}

	var contentEncoding []string

	if acceptEncoding != nil {
		for _, encoding := range acceptEncoding {
			if encoding == "gzip" {
				contentEncoding = append(contentEncoding, encoding)
			}
		}

		if len(contentEncoding) == 0 {
			contentEncoding = nil
		}
	}

	if data == nil || data == "" {
		return &ContentHeader{
			ContentType:     *contentType,
			ContentLength:   *contentLength,
			ContentEncoding: contentEncoding,
		}
	} else {
		length, err := Length(data)

		if err != nil {
			return &ContentHeader{
				ContentType:     *contentType,
				ContentLength:   *contentLength,
				ContentEncoding: contentEncoding,
			}
		}
		return &ContentHeader{
			ContentType:     *contentType,
			ContentLength:   length,
			ContentEncoding: contentEncoding,
		}
	}
}

func separateKeyValue(header string) (string, string, error) {
	splits := strings.Split(header, ": ")
	if len(splits) != 2 {
		return "", "", errors.New("invalid header")
	}

	return splits[0], splits[1], nil
}
