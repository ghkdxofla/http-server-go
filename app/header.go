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

	if requestMap["Content-Type"] != nil {
		contentLength, err := strconv.Atoi(*requestMap["Content-Length"])
		if err != nil {
			requestHeader.ContentHeader = NewContentHeader(requestMap["Content-Type"], nil, nil)
		} else {
			requestHeader.ContentHeader = NewContentHeader(requestMap["Content-Type"], &contentLength, nil)
		}
	}

	return &requestHeader
}

type ContentHeader struct {
	ContentType   string
	ContentLength int
}

func NewContentHeader(contentType *string, contentLength *int, data any) *ContentHeader {
	if contentType == nil {
		contentType = new(string)
		*contentType = "text/plain"
	}

	if contentLength == nil {
		contentLength = new(int)
		*contentLength = 0
	}

	if data == nil || data == "" {
		return &ContentHeader{
			ContentType:   *contentType,
			ContentLength: *contentLength,
		}
	} else {
		length, err := Length(data)

		if err != nil {
			return &ContentHeader{
				ContentType:   *contentType,
				ContentLength: *contentLength,
			}
		}
		return &ContentHeader{
			ContentType:   *contentType,
			ContentLength: length,
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
