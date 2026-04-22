package game_handler

import (
	"net/http"

	"go.rest.api/internal/handlers"
	"go.rest.api/internal/models"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) error {
	var game models.Game

	if err := handlers.DecodeJSON(r, &game); err != nil {
		return handlers.Error(w, http.StatusBadRequest, "Invalid JSON payload")
	}

	if err := handlers.Validate(game); err != nil {
		return handlers.Error(w, http.StatusUnprocessableEntity, err.Error())
	}

	if err := h.repo.Create(&game); err != nil {
		return handlers.Error(w, http.StatusInternalServerError, err.Error())
	}

	return handlers.JSON(w, http.StatusCreated, game)
}
