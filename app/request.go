package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Request struct {
	RequestLine RequestLine
	Header      Header
	RequestBody RequestBody
}

// CallFunc 공통 호출 함수
func CallFunc(fn any, args ...any) (any, error) {
	fnValue := reflect.ValueOf(fn)
	fnType := fnValue.Type()

	// 함수 타입 검증
	if fnType.Kind() != reflect.Func {
		return nil, errors.New("provided argument is not a function")
	}

	// 인자 검증
	if len(args) != fnType.NumIn() {
		return nil, fmt.Errorf("expected %d arguments, got %d", fnType.NumIn(), len(args))
	}

	// 인자 변환
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		if reflect.TypeOf(arg) != fnType.In(i) {
			return nil, fmt.Errorf("argument %d must be of type %s", i, fnType.In(i))
		}
		in[i] = reflect.ValueOf(arg)
	}

	// 함수 호출
	out := fnValue.Call(in)

	// 결과 반환
	if len(out) == 2 && out[1].Interface() != nil {
		return out[0].Interface(), out[1].Interface().(error)
	}
	return out[0].Interface(), nil
}

func HandleRequest(requestLine *RequestLine, header *Header, requestBody *RequestBody, callback any) *Response {
	if callback == nil {
		return NewResponse(requestLine.Version, StatusNotFound(), nil)
	}

	result, err := CallFunc(callback, requestLine, header, requestBody)
	if err != nil {
		return NewResponse(requestLine.Version, StatusInternalServerError(), err.Error())
	}

	return NewResponse(requestLine.Version, StatusOK(), result)
}
