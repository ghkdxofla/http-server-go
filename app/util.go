package main

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
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

func CompressGzip(data any) (string, error) {
	var buf bytes.Buffer
	var err error
	handler := gzip.NewWriter(&buf)

	rv := reflect.ValueOf(data)
	switch rv.Kind() {
	case reflect.Array, reflect.Slice:
		_, err = handler.Write(rv.Bytes())
	case reflect.String:
		_, err = handler.Write([]byte(rv.String()))
	default:
		_, err = handler.Write([]byte(rv.String()))
	}

	if err != nil {
		return "", errors.New("gzip 파일을 쓸 수 없습니다")
	}

	err = handler.Close()
	if err != nil {
		return "", errors.New("gzip 파일을 닫을 수 없습니다")
	}

	return buf.String(), nil
}

func DecompressGzip(data any) (string, error) {
	var handler *gzip.Reader
	var err error

	rv := reflect.ValueOf(data)
	switch rv.Kind() {
	case reflect.Array, reflect.Slice:
		handler, err = gzip.NewReader(
			bytes.NewReader(rv.Bytes()),
		)
	case reflect.String:
		if rv.String() == "" {
			return "", nil
		}

		handler, err = gzip.NewReader(
			bytes.NewReader([]byte(rv.String())),
		)
	default:
		handler, err = gzip.NewReader(
			bytes.NewReader(rv.Bytes()),
		)
	}

	if err != nil {
		return "", errors.New("gzip 파일을 열 수 없습니다")
	}
	defer func(handler *gzip.Reader) {
		err := handler.Close()
		CheckError(err)
	}(handler)

	result, err := io.ReadAll(handler)

	if err != nil {
		return "", errors.New("gzip 파일을 읽을 수 없습니다")
	}
	return string(result), nil
}

func getAbsolutePath(root string, path string) string {
	if root[len(root)-1] == '/' {
		return root + path
	} else {
		return root + "/" + path
	}
}
