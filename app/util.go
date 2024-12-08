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
	absolutePath := getAbsolutePath(root, path)

	dat, err := os.ReadFile(absolutePath)

	return dat, err
}

func WriteFile(root string, data []byte, filename string) error {
	absolutePath := getAbsolutePath(root, filename)

	err := os.WriteFile(absolutePath, data, 0644)

	return err
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

func getAbsolutePath(root string, path string) string {
	if root[len(root)-1] == '/' {
		return root + path
	} else {
		return root + "/" + path
	}
}
