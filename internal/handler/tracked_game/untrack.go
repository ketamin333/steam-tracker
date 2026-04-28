package trackedgamehandler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.rest.api/internal/httputil"
	"go.rest.api/internal/middleware"
)

func (h *Handler) Untrack(w http.ResponseWriter, r *http.Request) error {
	gameID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || gameID == 0 {
		return httputil.Error(w, http.StatusBadRequest, "invalid game id")
	}
	
	userID := r.Context().Value(middleware.UserIDKey).(int)
	if err := h.svc.Untrack(r.Context(), userID, gameID); err != nil {
		return httputil.Error(w, http.StatusInternalServerError, "server error")
	}

	return httputil.JSON(w, http.StatusNoContent, nil)
}
