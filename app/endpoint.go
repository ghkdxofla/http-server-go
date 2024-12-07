package main

import "regexp"

func AddEndpoint(router *Router) {
	router.Add("GET", regexp.MustCompile(`^/echo/(?P<param>[a-zA-Z0-9_-]+)$`), Echo)
	router.Add("GET", regexp.MustCompile(`^/user-agent$`), UserAgent)
	router.Add("GET", regexp.MustCompile(`^/files/(?P<param>[a-zA-Z0-9_-]+)$`), File)
	router.Add("GET", regexp.MustCompile(`^/$`), Root)
}
