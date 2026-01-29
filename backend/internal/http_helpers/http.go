package http_helpers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
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

func ReadJSON(r *http.Request, dst interface{}, maxBytes int64) error {
	r.Body = http.MaxBytesReader(nil, r.Body, maxBytes)
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		return err
	}

	if dec.More() {
		return errors.New("extra data in request body")
	}

	return nil
}

func GetRequest(w http.ResponseWriter, r *http.Request, req interface{}) bool {
	if r.Method != http.MethodPost {
		WriteJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return false
	}

	// Limit to 1MB
	if err := ReadJSON(r, &req, (1 << 20)); err != nil {
		WriteJSONError(w, http.StatusBadRequest, err.Error())
		return false
	}

	return true
}
