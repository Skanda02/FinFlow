package handlers

import (
	"net/http"

	"finflow/internal/http_helpers"
)

func Health(w http.ResponseWriter, r *http.Request) {
	http_helpers.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Okay",
	})
}
