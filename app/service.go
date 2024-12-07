package main

import (
	"flag"
)

func Root(request Request) Response {
	return NewResponse(request.RequestLine.Version, StatusOK(), nil, "")
}

func Echo(request Request) Response {
	if request.PathParams == nil {
		return NewResponse(request.RequestLine.Version, StatusOK(), nil, "")
	}

	return NewResponse(request.RequestLine.Version, StatusOK(), nil, (*request.PathParams)["param"])
}

func File(request Request) Response {
	if request.PathParams == nil {
		return NewResponse(request.RequestLine.Version, StatusNotFound(), nil, nil)
	}

	root := flag.Lookup("directory").Value.(flag.Getter).Get().(string)

	file, err := ReadFile(root, (*request.PathParams)["param"])

	if err != nil {
		return NewResponse(request.RequestLine.Version, StatusNotFound(), nil, err)
	}

	contentType := "application/octet-stream"
	return NewResponse(request.RequestLine.Version, StatusOK(), &contentType, file)
}

func UserAgent(request Request) Response {
	if request.Header == nil {
		return NewResponse(request.RequestLine.Version, StatusOK(), nil, "")
	}

	return NewResponse(request.RequestLine.Version, StatusOK(), nil, request.Header.UserAgent)
}
