package handler

import (
	"encoding/json"
	myhttp "github.com/fssou/learn-go/internal/adapter/in/http"
	"net/http"
)

type HealthCheckResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}


type Health struct {
}

func NewHealth() *Health {
	return &Health{
	}
}
func (handler *Health) Routes() []myhttp.Route {
	return []myhttp.Route{
		&healthRouteGet{},
	}
}

type healthRouteGet struct {

}
func (h *healthRouteGet) Method() string { return http.MethodGet }
func (h *healthRouteGet) Path() string { return "/health" }
func (h *healthRouteGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&HealthCheckResponse{
		Status:  "healthy",
		Message: "The service is running smoothly.",
	})
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
