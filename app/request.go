package main

type Request struct {
	RequestLine *RequestLine
	Header      *RequestHeader
	RequestBody *RequestBody
	PathParams  *Params
	QueryParams *Params
}

type ServiceFunc func(Request) Response

func HandleRequest(request Request, callback ServiceFunc) Response {
	if callback == nil {
		return NewResponse(request.RequestLine.Version, StatusNotFound(), nil, nil, nil)
	}

	response := callback(request)

	return response
}
