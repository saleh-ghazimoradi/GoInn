package routes

import (
	"github.com/saleh-ghazimoradi/GoInn/internal/gateway/handlers"
	"net/http"
)

func HealthRoute(mux *http.ServeMux, handler *handlers.HealthHandler) {
	mux.HandleFunc("GET /v1/health", handler.HealthCheckHandler)
}
