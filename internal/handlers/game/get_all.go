package game_handler

import (
	"net/http"

	"go.rest.api/internal/handlers"
)

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) error {
	games, err := h.repo.GetAll()

	if err != nil {
		return handlers.Error(w, http.StatusInternalServerError, err.Error())
	}

	return handlers.JSON(w, http.StatusOK, games)
}
