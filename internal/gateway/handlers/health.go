package handlers

import (
	"github.com/saleh-ghazimoradi/GoInn/config"
	"github.com/saleh-ghazimoradi/GoInn/internal/helper"
	"net/http"
)

type HealthHandler struct {
	config *config.Config
	error  *helper.Error
}

func (h *HealthHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	env := helper.Envelope{
		"status": "available",
		"system_info": map[string]any{
			"environment": h.config.Application.Env,
			"version":     h.config.Application.Version,
		},
	}
	if err := helper.WriteJSON(w, http.StatusOK, env, nil); err != nil {
		h.error.ServerErrorResponse(w, r, err)
	}
}

func NewHealthHandler(config *config.Config, error *helper.Error) *HealthHandler {
	return &HealthHandler{
		config: config,
		error:  error,
	}
}
