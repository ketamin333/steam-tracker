package trackedgamehandler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.rest.api/internal/apperr"
	"go.rest.api/internal/httputil"
	"go.rest.api/internal/middleware"
	"go.rest.api/internal/model"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) error {
	gameID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || gameID == 0 {
		return httputil.Error(w, http.StatusBadRequest, "invalid game id")
	}

	user := r.Context().Value(middleware.UserKey).(*model.User)
	err = h.svc.Delete(r.Context(), user, gameID)

	if errors.Is(err, apperr.ErrNotFound) {
		return httputil.Error(w, http.StatusNotFound, "game not found")
	}

	if err != nil {
		return httputil.Error(w, http.StatusInternalServerError, "server error")
	}

	return httputil.JSON(w, http.StatusNoContent, nil)
}
