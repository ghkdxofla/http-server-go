package main

func Echo(_ *RequestLine, _ *Header, requestBody *RequestBody) (string, error) {
	return requestBody.data, nil
}
