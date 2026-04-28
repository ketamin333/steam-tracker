package trackedgamehandler

import (
	"encoding/json"
	"net/http"

	"go.rest.api/internal/httputil"
	"go.rest.api/internal/middleware"
)

type trackRequest struct {
	GameID int `json:"game_id"`
}

func (h *Handler) Track(w http.ResponseWriter, r *http.Request) error {
	var req trackRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return httputil.Error(w, http.StatusBadRequest, "invalid request body")
	}

	if req.GameID == 0 {
		return httputil.Error(w, http.StatusBadRequest, "game_id is required")
	}

	userID := r.Context().Value(middleware.UserIDKey).(int)
	if err := h.svc.Track(r.Context(), userID, req.GameID); err != nil {
		return httputil.Error(w, http.StatusInternalServerError, "server error")
	}

	return httputil.JSON(w, http.StatusCreated, nil)
}
