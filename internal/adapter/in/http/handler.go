package http

import "net/http"

type Route interface {
	http.Handler
	Method() string
	Path() string
}

type Handler interface {
	Routes() []Route
}
