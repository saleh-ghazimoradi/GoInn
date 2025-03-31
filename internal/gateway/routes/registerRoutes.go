package routes

import (
	"github.com/saleh-ghazimoradi/GoInn/internal/gateway/handlers"
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	healthHandler := handlers.NewHealthHandler()

	HealthRoute(mux, healthHandler)

	return mux
}
