package main

type RequestBody struct {
	data string
}

func NewRequestBody(data string) *RequestBody {
	return &RequestBody{
		data: data,
	}
}
