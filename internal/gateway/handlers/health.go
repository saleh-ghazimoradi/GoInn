package handlers

import (
	"github.com/saleh-ghazimoradi/GoInn/config"
	"github.com/saleh-ghazimoradi/GoInn/internal/helper"
	"net/http"
)

type HealthHandler struct {
	config config.Config
	error  helper.Error
}

func (h *HealthHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	env := helper.Envelope{
		"status": "available",
		"system_info": map[string]any{
			"environment": h.config.Server.Env,
			"version":     h.config.Server.Version,
		},
	}
	if err := helper.WriteJSON(w, http.StatusOK, env, nil); err != nil {
		h.error.ServerErrorResponse(w, r, err)
	}
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}
