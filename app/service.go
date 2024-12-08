package main

import (
	"flag"
)

func Root(request Request) Response {
	return NewResponse(request.RequestLine.Version, StatusOK(), nil, nil, request.Header.ContentHeader.ContentEncoding, "")
}

func Echo(request Request) Response {
	if request.PathParams == nil {
		return NewResponse(request.RequestLine.Version, StatusOK(), nil, nil, request.Header.ContentHeader.ContentEncoding, "")
	}

	return NewResponse(request.RequestLine.Version, StatusOK(), nil, nil, request.Header.ContentHeader.ContentEncoding, (*request.PathParams)["param"])
}

func GetFile(request Request) Response {
	if request.PathParams == nil {
		return NewResponse(request.RequestLine.Version, StatusNotFound(), nil, nil, request.Header.ContentHeader.ContentEncoding, nil)
	}

	root := flag.Lookup("directory").Value.(flag.Getter).Get().(string)

	file, err := ReadFile(root, (*request.PathParams)["param"])

	if err != nil {
		return NewResponse(request.RequestLine.Version, StatusNotFound(), nil, nil, request.Header.ContentHeader.ContentEncoding, err)
	}

	contentType := "application/octet-stream"
	return NewResponse(request.RequestLine.Version, StatusOK(), &contentType, nil, request.Header.ContentHeader.ContentEncoding, file)
}

func CreateFile(request Request) Response {
	if request.RequestBody == nil {
		return NewResponse(request.RequestLine.Version, StatusBadRequest(), nil, nil, request.Header.ContentHeader.ContentEncoding, nil)
	}

	root := flag.Lookup("directory").Value.(flag.Getter).Get().(string)

	err := WriteFile(root, []byte(request.RequestBody.data), (*request.PathParams)["param"])

	if err != nil {
		return NewResponse(request.RequestLine.Version, StatusInternalServerError(), nil, nil, request.Header.ContentHeader.ContentEncoding, err)
	}

	return NewResponse(request.RequestLine.Version, StatusCreated(), nil, nil, request.Header.ContentHeader.ContentEncoding, nil)
}

func UserAgent(request Request) Response {
	if request.Header == nil {
		return NewResponse(request.RequestLine.Version, StatusOK(), nil, nil, request.Header.ContentHeader.ContentEncoding, "")
	}
	return NewResponse(request.RequestLine.Version, StatusOK(), nil, nil, request.Header.ContentHeader.ContentEncoding, *request.Header.UserAgent)
}
