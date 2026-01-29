package handlers

import (
	"net/http"
	"encoding/json"
	"errors"
)

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

