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
