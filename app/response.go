package main

type Response struct {
	Version string
	Status  Status
	Data    interface{}
}

func NewResponse(version string, status Status, data interface{}) *Response {
	return &Response{
		Version: version,
		Status:  status,
		Data:    data,
	}
}

func (r *Response) ToString() string {
	dataString := ""
	if r.Data != nil && r.Data.(string) != "" {
		dataString = " " + r.Data.(string)
	}

	return r.Version + " " + r.Status.ToString() + dataString + "\r\n\r\n"
}
