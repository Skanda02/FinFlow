package handlers

import (
	"net/http"
	"finflow/internal/services"
)

type RegisterRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		WriteJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	
	var req RegisterRequest
	// Limit to 1MB
	if err := ReadJSON(r, &req, (1 << 20)); err != nil {
		WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	if req.Email == "" || req.Password == "" {
		WriteJSONError(w, http.StatusBadRequest, "Email and password required")
		return
	}

	data := services.RegisterData{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
	}
	
	if err := services.RegisterNewUser(&data); err != nil {
		WriteJSONError(w, http.StatusNotImplemented, err.Error())
		return
	}

	WriteJSONData(w, http.StatusOK, map[string]string{
		"email": req.Email,
		"name": req.Name,
	})
}
