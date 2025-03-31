package gateway

import (
	"github.com/saleh-ghazimoradi/GoInn/internal/gateway/routes"
	"net/http"
)

func Server() error {
	mux := routes.RegisterRoutes()
	if err := http.ListenAndServe(":8080", mux); err != nil {
		return err
	}

	return nil
}
