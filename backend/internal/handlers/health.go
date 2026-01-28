package handlers

import (
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	WriteData(w, http.StatusOK, map[string]interface{}{
		"message": "Okay",
	})
}
