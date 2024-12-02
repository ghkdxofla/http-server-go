package main

import "regexp"

func AddEndpoint(router *Router) {
	router.Add("GET", regexp.MustCompile(`^/echo/(?P<param>[a-zA-Z0-9]+)$`), Echo)
	router.Add("GET", regexp.MustCompile(`^/$`), Root)
}
