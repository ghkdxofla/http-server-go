package main

import "strings"

type RequestLine struct {
	Method  string
	Path    string
	Version string
}

func NewRequestLine(request string) *RequestLine {
	method, path, version := separateRequestLine(request)

	return &RequestLine{
		Method:  method,
		Path:    path,
		Version: version,
	}
}

func separateRequestLine(requestLine string) (string, string, string) {
	splits := strings.Split(requestLine, " ")

	return splits[0], splits[1], splits[2]
}
