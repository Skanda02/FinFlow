package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"finflow/internal/services"
)

func WriteJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Println("JSON encoding failed!")
	}
}

func WriteJSONData(w http.ResponseWriter, status int, data interface{}) {
	WriteJSON(w, status, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

func WriteJSONError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]interface{}{
		"success": false,
		"message": message,
	})
}

// HandleServiceError maps service errors to appropriate HTTP responses
func HandleServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, services.ErrValidation):
		WriteJSONError(w, http.StatusBadRequest, "email and password are required")
	case errors.Is(err, services.ErrUserExists):
		WriteJSONError(w, http.StatusConflict, "user with this email already exists")
	case errors.Is(err, services.ErrInvalidCredentials):
		WriteJSONError(w, http.StatusUnauthorized, "invalid email or password")
	default:
		WriteJSONError(w, http.StatusInternalServerError, "internal server error")
	}
}
