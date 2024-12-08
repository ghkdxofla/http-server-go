package main

import "strconv"

// enum for status code
const (
	StatusCodeOK                  = 200
	StatusCodeCreated             = 201
	StatusCodeBadRequest          = 400
	StatusCodeNotFound            = 404
	StatusCodeInternalServerError = 500
)

// enum for status message
const (
	StatusMessageOK                  = "OK"
	StatusMessageCreated             = "Created"
	StatusMessageBadRequest          = "Bad Request"
	StatusMessageNotFound            = "Not Found"
	StatusMessageInternalServerError = "Internal Server Error"
)

// StatusField struct
type StatusField struct {
	Code    int
	Message string
}

// Status interface
type Status struct {
	StatusField
}

func (s *Status) IsEqual(status *Status) bool {
	return s.Code == status.Code && s.Message == status.Message
}

func (s *Status) ToString() string {
	return strconv.Itoa(s.Code) + " " + s.Message
}

// StatusOK returns status OK
func StatusOK() Status {
	return Status{
		StatusField: StatusField{
			Code:    StatusCodeOK,
			Message: StatusMessageOK,
		},
	}
}

// StatusCreated returns status Created
func StatusCreated() Status {
	return Status{
		StatusField: StatusField{
			Code:    StatusCodeCreated,
			Message: StatusMessageCreated,
		},
	}
}

// StatusBadRequest returns status Bad Request
func StatusBadRequest() Status {
	return Status{
		StatusField: StatusField{
			Code:    StatusCodeBadRequest,
			Message: StatusMessageBadRequest,
		},
	}
}

// StatusNotFound returns status Not Found
func StatusNotFound() Status {
	return Status{
		StatusField: StatusField{
			Code:    StatusCodeNotFound,
			Message: StatusMessageNotFound,
		},
	}
}

// StatusInternalServerError returns status Internal Server Error
func StatusInternalServerError() Status {
	return Status{
		StatusField: StatusField{
			Code:    StatusCodeInternalServerError,
			Message: StatusMessageInternalServerError,
		},
	}
}
