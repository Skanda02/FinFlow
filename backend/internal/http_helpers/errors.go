package http_helpers

import (
	"errors"
	"net/http"

	"finflow/internal/services"
)

// HandleServiceError maps service errors to appropriate HTTP responses
func HandleServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, services.ErrValidation):
		WriteJSONError(w, http.StatusBadRequest, "email and password are required")
	case errors.Is(err, services.ErrUserExists):
		WriteJSONError(w, http.StatusConflict, "user with this email already exists")
	case errors.Is(err, services.ErrInvalidCredentials):
		WriteJSONError(w, http.StatusUnauthorized, "invalid email or password")
	case errors.Is(err, services.ErrInvalidAmount):
		WriteJSONError(w, http.StatusBadRequest, "amount must be greater than zero")
	default:
		WriteJSONError(w, http.StatusInternalServerError, "internal server error")
	}
}
