package main

import (
	"errors"
	"strings"
)

type Header struct {
	Host      string
	UserAgent string
	Accept    string
}

func NewHeader(requests ...string) *Header {
	requestMap := make(map[string]string)
	for _, request := range requests {
		key, value, err := separateKeyValue(request)
		if err != nil {
			continue
		}
		requestMap[key] = value
	}

	return &Header{
		Host:      requestMap["Host"],
		UserAgent: requestMap["User-Agent"],
		Accept:    requestMap["Accept"],
	}
}

func (h *Header) SetHeader() map[string]string {
	return map[string]string{
		"Host":       h.Host,
		"User-Agent": h.UserAgent,
		"Accept":     h.Accept,
	}
}

func separateKeyValue(header string) (string, string, error) {
	splits := strings.Split(header, ": ")
	if len(splits) != 2 {
		return "", "", errors.New("invalid header")
	}

	return splits[0], splits[1], nil
}
