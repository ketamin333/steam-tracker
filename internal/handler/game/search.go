package gamehandler

import (
	"net/http"

	"go.rest.api/internal/httputil"
)

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query().Get("search")

	if query == "" {
		return httputil.Error(w, http.StatusBadRequest, "search is empty")
	}

	g, err := h.svc.Search(r.Context(), query)

	if err != nil {
		return httputil.Error(w, http.StatusInternalServerError, "server error")
	}

	return httputil.JSON(w, http.StatusOK, g)
}
