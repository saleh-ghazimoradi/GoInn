package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/saleh-ghazimoradi/GoInn/internal/gateway/handlers"
	"net/http"
)

func HealthRoute(router httprouter.Router, handler *handlers.HealthHandler) {
	router.HandlerFunc(http.MethodGet, "/v1/health", handler.HealthCheckHandler)
}
