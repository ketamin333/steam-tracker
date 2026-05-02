package userhandler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"steam-tracker/internal/apperr"
	"steam-tracker/internal/httputil"
	"steam-tracker/internal/middleware"
	"steam-tracker/internal/model"
)

type updateRequest struct {
	Email *string `json:"email"`
	Lang  *string `json:"lang"`
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) error {
	var req updateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return httputil.Error(w, http.StatusBadRequest, apperr.ErrInvalidBody.Error())
	}

	user := r.Context().Value(middleware.UserKey).(*model.User)
	u, err := h.svc.Update(r.Context(), user, req.Email, req.Lang)

	if errors.Is(err, sql.ErrNoRows) {
		return httputil.Error(w, http.StatusNotFound, apperr.ErrNotFound.Error())
	}

	if err != nil {
		return httputil.Error(w, http.StatusInternalServerError, apperr.ErrInternal.Error())
	}

	return httputil.JSON(w, http.StatusOK, u)
}
