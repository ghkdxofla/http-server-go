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

	response := callback(request)

	if response.ResponseHeader.ContentEncoding != nil {
		for _, encoding := range response.ResponseHeader.ContentEncoding {
			if encoding == "gzip" {
				data, err := CompressGzip(response.Data)

				if err != nil {
					return NewResponse(request.RequestLine.Version, StatusInternalServerError(), nil, nil, nil, err)
				}

				response.Data = data
				response.ResponseHeader.ContentLength = len(data)
			}
		}
	}

	return response
}
