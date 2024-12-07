package main

import (
	"errors"
	"strings"
)

type RequestHeader struct {
	Host      string
	UserAgent string
	Accept    string
}

func NewRequestHeader(requests ...string) *RequestHeader {
	requestMap := make(map[string]string)
	for _, request := range requests {
		key, value, err := separateKeyValue(request)
		if err != nil {
			continue
		}
		requestMap[key] = value
	}

	return &RequestHeader{
		Host:      requestMap["Host"],
		UserAgent: requestMap["User-Agent"],
		Accept:    requestMap["Accept"],
	}
}

func (h *RequestHeader) SetHeader() map[string]string {
	return map[string]string{
		"Host":       h.Host,
		"User-Agent": h.UserAgent,
		"Accept":     h.Accept,
	}
}

type ResponseHeader struct {
	ContentType   string
	ContentLength int
}

func NewResponseHeader(contentType *string, data any) *ResponseHeader {
	if contentType == nil {
		contentType = new(string)
		*contentType = "text/plain"
	}

	if data == nil || data == "" {
		return &ResponseHeader{
			ContentType:   *contentType,
			ContentLength: 0,
		}
	} else {
		length, err := Length(data)
		if err != nil {
			return &ResponseHeader{
				ContentType:   *contentType,
				ContentLength: 0,
			}
		}
		return &ResponseHeader{
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
