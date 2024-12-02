package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			os.Exit(1)
		}

		requestLine, header, body := ParseRequest(string(buf[:n]))

		fmt.Println("Received line: ", requestLine)
		fmt.Println("Received header: ", header)
		fmt.Println("Received body: ", body)

		router := NewRouter()
		AddEndpoint(router)
		route := router.Find(requestLine)
		response := &Response{}

		// TODO: router 처리 개선
		if route == nil {
			response = HandleRequest(Request{
				RequestLine: requestLine,
				Header:      header,
				RequestBody: body,
				PathParams:  nil,
				QueryParams: nil,
			}, nil)
		} else {
			response = HandleRequest(Request{
				RequestLine: requestLine,
				Header:      header,
				RequestBody: body,
				PathParams:  route.PathParams,
				QueryParams: route.QueryParams,
			}, route.Callback)
		}

		fmt.Println("Response: ", response.ToString())

		n, err = conn.Write([]byte(response.ToString()))
		if err != nil {
			fmt.Println("Error writing:", err.Error())
			os.Exit(1)
		}
	}
}
