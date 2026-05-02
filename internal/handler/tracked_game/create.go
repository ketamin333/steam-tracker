package trackedgamehandler

import (
	"encoding/json"
	"errors"
	"net/http"

	"steam-tracker/internal/apperr"
	"steam-tracker/internal/httputil"
	"steam-tracker/internal/middleware"
	"steam-tracker/internal/model"
)

type createRequest struct {
	GameID int      `json:"game_id"`
	Price  *float64 `json:"price"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) error {
	var req createRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return httputil.Error(w, http.StatusBadRequest, "invalid request body")
	}

	if req.GameID == 0 {
		return httputil.Error(w, http.StatusBadRequest, "game_id is required")
	}

	user := r.Context().Value(middleware.UserKey).(*model.User)

	g, err := h.svc.Create(r.Context(), user, req.GameID, req.Price)
	if errors.Is(err, apperr.ErrAlreadyExists) {
		return httputil.Error(w, http.StatusConflict, "game already tracked")
	}
	if err != nil {
		return httputil.Error(w, http.StatusInternalServerError, "server error")
	}

	return httputil.JSON(w, http.StatusCreated, g)
}
