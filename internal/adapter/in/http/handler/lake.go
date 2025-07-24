package handler

import (
	http2 "github.com/fssou/learn-go/internal/adapter/in/http"
	"github.com/fssou/learn-go/internal/application/lake"
	"net/http"
)

type LakeResponse struct {
	Name string `json:"name"`
	Price int    `json:"price"`
}

type Lake struct {
	service lake.Service
}
func NewLake(service lake.Service) *Lake {
	return &Lake{ service }
}
func (lake *Lake) Routes() []http2.Route {
	return []http2.Route{
		&lakeHandlerGet{lake.service},
	}
}

type lakeHandlerGet struct {
	service lake.Service
}
func (handler *lakeHandlerGet) Method() string { return http.MethodGet }
func (handler *lakeHandlerGet) Path() string { return "/lakes" }
func (handler *lakeHandlerGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"message": "This is the lake handler"}`))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
