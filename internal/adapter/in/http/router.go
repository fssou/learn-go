package http

import (
	"fmt"
	"net/http"
)

type Router struct {
	handlers []Handler
}

func NewRouter(
	handlers []Handler,
) *Router {
	router := &Router{
		handlers,
	}
	return router
}

func (r *Router) Register() *http.ServeMux {
	serveMux := http.NewServeMux()
	for _, handler := range r.handlers {
		for _, route := range handler.Routes() {
			pattern := fmt.Sprintf("%s %s", route.Method(), route.Path())
			serveMux.Handle(pattern, route)
		}
	}
	return serveMux
}
