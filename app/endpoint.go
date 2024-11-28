package main

func AddEndpoint(router *Router) {
	router.Add("GET", "/", Echo)
	router.Add("GET", "/echo", Echo)
}
