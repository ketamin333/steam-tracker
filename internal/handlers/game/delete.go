package game_handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.rest.api/internal/handlers"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		return handlers.Error(w, http.StatusBadRequest, "Invalid ID")
	}

	if err := h.repo.Delete(id); err != nil {
		return handlers.Error(w, http.StatusInternalServerError, err.Error())
	}

	return handlers.NoContent(w, http.StatusNoContent)
}
