package main

import (
	"strings"
)

func ParseRequest(rawRequest string) (*RequestLine, *Header, *RequestBody) {
	rawRequestLine, rawHeader, rawRequestBody := separateRequest(rawRequest)

	requestLine := NewRequestLine(rawRequestLine)
	header := NewHeader(rawHeader...)
	requestBody := NewRequestBody(rawRequestBody)

	return requestLine, header, requestBody
}

func separateRequest(request string) (string, []string, string) {
	firstSeparated := strings.Split(request, "\r\n\r\n")
	rawRequestBody := firstSeparated[len(firstSeparated)-1]

	secondSeparated := strings.Split(firstSeparated[0], "\r\n")
	rawRequestLine := secondSeparated[0]
	rawHeader := secondSeparated[1:]

	return rawRequestLine, rawHeader, rawRequestBody
}
