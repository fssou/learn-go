package handler

import (
	"encoding/json"
	http2 "github.com/fssou/learn-go/internal/adapter/in/http"
	"net/http"
)

type HelloResponse struct {
	Message string  `json:"message"`
}

type Hello struct {
}
func NewHello() *Hello {
	handler := &Hello{}
	return handler
}
func (h *Hello) Routes() []http2.Route {
	return []http2.Route{
		&helloHandlerGet{},
	}
}

type helloHandlerGet struct {}
func (h *helloHandlerGet) Method() string { return http.MethodGet }
func (h *helloHandlerGet) Path() string { return "/hello" }
func (h *helloHandlerGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&HelloResponse{"Hello World!"})
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
