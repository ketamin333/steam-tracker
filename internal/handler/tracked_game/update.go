package trackedgamehandler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"steam-tracker/internal/apperr"
	"steam-tracker/internal/httputil"
	"steam-tracker/internal/middleware"
	"steam-tracker/internal/model"

	"github.com/go-chi/chi/v5"
)

type updateRequest struct {
	Price *float64 `json:"price"`
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) error {
	var req updateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return httputil.Error(w, http.StatusBadRequest, "invalid request body")
	}

	gameID, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil || gameID == 0 {
		return httputil.Error(w, http.StatusBadRequest, "invalid game id")
	}

	user := r.Context().Value(middleware.UserKey).(*model.User)
	g, err := h.svc.Update(r.Context(), user, gameID, req.Price)

	if errors.Is(err, apperr.ErrNotFound) {
		return httputil.Error(w, http.StatusNotFound, "game not found")
	}

	if err != nil {
		return httputil.Error(w, http.StatusInternalServerError, "server error")
	}

	return httputil.JSON(w, http.StatusOK, g)
}
