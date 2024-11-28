package main

type Router struct {
	routes map[string]map[string]any
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Add(method string, path string, callback any) {
	if r.routes == nil {
		r.routes = make(map[string]map[string]any)
	}
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]any)
	}
	r.routes[method][path] = callback
}

func (r *Router) Get(method string, path string) any {
	return r.routes[method][path]
}

func (r *Router) Find(requestLine *RequestLine) any {
	return r.Get(requestLine.Method, requestLine.Path)
}
