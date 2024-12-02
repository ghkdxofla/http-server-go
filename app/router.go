package main

import (
	"regexp"
)

type Params map[string]string

type Route struct {
	Path     *regexp.Regexp
	Callback any
}

type RouteWithParams struct {
	Route
	PathParams  *Params
	QueryParams *Params
}

type Router struct {
	routes map[string][]Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Add(method string, path *regexp.Regexp, callback any) {
	if r.routes == nil {
		r.routes = make(map[string][]Route)
	}
	if r.routes[method] == nil {
		r.routes[method] = make([]Route, 0)
	}
	r.routes[method] = append(r.routes[method], Route{
		Path:     path,
		Callback: callback,
	})
}

func (r *Router) Get(method string, path string) *RouteWithParams {
	if r.routes == nil {
		return nil
	}
	if r.routes[method] == nil {
		return nil
	}

	for _, route := range r.routes[method] {
		if route.Path.MatchString(path) {
			match := route.Path.FindStringSubmatch(path)
			names := route.Path.SubexpNames()

			params := make(Params)
			for i, name := range names {
				if i > 0 && name != "" { // 첫 번째는 전체 매칭값이므로 제외
					params[name] = match[i]
				}
			}

			return &RouteWithParams{
				Route:      route,
				PathParams: &params,
				// FIXME: QueryParams를 구현하세요.
				QueryParams: nil,
			}
		}
	}

	return nil
}

func (r *Router) Find(requestLine *RequestLine) *RouteWithParams {
	return r.Get(requestLine.Method, requestLine.Path)
}
