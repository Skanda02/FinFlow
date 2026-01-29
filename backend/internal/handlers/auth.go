package handlers

import (
	"finflow/internal/services"
	"net/http"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if !GetRequest(w, r, &req) {
		return
	}

	if req.Email == "" || req.Password == "" {
		WriteJSONError(w, http.StatusBadRequest, "email and password required")
		return
	}

	data := services.RegisterData{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := services.RegisterNewUser(&data); err != nil {
		WriteJSONError(w, http.StatusNotImplemented, err.Error())
		return
	}

	WriteJSONData(w, http.StatusOK, map[string]string{
		"email": req.Email,
		"name":  req.Name,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if !GetRequest(w, r, &req) {
		return
	}

	if req.Email == "" || req.Password == "" {
		WriteJSONError(w, http.StatusBadRequest, "email and password required")
		return
	}

	data := services.LoginData{
		Email:    req.Email,
		Password: req.Password,
	}

	if err := services.LoginUser(&data); err != nil {
		WriteJSONError(w, http.StatusNotImplemented, err.Error())
		return
	}

	WriteJSONData(w, http.StatusOK, map[string]string{
		"email": req.Email,
	})
}
