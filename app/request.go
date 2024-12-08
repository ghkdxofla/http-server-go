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
		return NewResponse(request.RequestLine.Version, StatusNotFound(), nil, nil, nil, nil)
	}

	if request.Header != nil && request.Header.ContentHeader != nil {
		if request.Header.ContentHeader.ContentEncoding == nil {
			// do nothing
		} else if *request.Header.ContentHeader.ContentEncoding == "gzip" {
			data, err := DecompressGzip(request.RequestBody.data)
			if err != nil {
				return NewResponse(request.RequestLine.Version, StatusInternalServerError(), nil, nil, nil, err)
			}
			request.RequestBody.data = data
		}
	}

	response := callback(request)

	return response
}
