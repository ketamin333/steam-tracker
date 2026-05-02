package trackedgamehandler

import (
	"errors"
	"net/http"
	"strconv"

	"steam-tracker/internal/apperr"
	"steam-tracker/internal/httputil"
	"steam-tracker/internal/middleware"
	"steam-tracker/internal/model"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) error {
	gameID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || gameID == 0 {
		return httputil.Error(w, http.StatusBadRequest, apperr.ErrInvalidParams.Error())
	}

	user := r.Context().Value(middleware.UserKey).(*model.User)
	err = h.svc.Delete(r.Context(), user, gameID)

	if errors.Is(err, apperr.ErrNotFound) {
		return httputil.Error(w, http.StatusNotFound, "game not found")
	}

	if err != nil {
		return httputil.Error(w, http.StatusInternalServerError, apperr.ErrInternal.Error())
	}

	return httputil.JSON(w, http.StatusNoContent, nil)
}
