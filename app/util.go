package main

import (
	"errors"
	"os"
	"reflect"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile(root string, path string) ([]byte, error) {
	var absolutePath string
	if root[len(root)-1] == '/' {
		absolutePath = root + path
	} else {
		absolutePath = root + "/" + path
	}

	dat, err := os.ReadFile(absolutePath)

	return dat, err
}

func Length(v any) (int, error) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Array, reflect.Slice, reflect.String, reflect.Map, reflect.Chan:
		return rv.Len(), nil
	default:
		return 0, errors.New("해당 타입에 대해서 length를 구할 수 없습니다.")
	}
}
