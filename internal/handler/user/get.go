package userhandler

import (
	"net/http"
	"steam-tracker/internal/httputil"
	"steam-tracker/internal/middleware"
	"steam-tracker/internal/model"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) error {
	user := r.Context().Value(middleware.UserKey).(*model.User)

	return httputil.JSON(w, http.StatusOK, user)
}
