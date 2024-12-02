package main

func Root(request Request) (string, error) {
	return "", nil
}

func Echo(request Request) (string, error) {
	if request.PathParams == nil {
		return "", nil
	}

	return (*request.PathParams)["param"], nil
}

func UserAgent(request Request) (string, error) {
	if request.Header == nil {
		return "", nil
	}

	return request.Header.UserAgent, nil
}
