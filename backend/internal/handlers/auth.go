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

	data := services.RegisterData{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := services.RegisterNewUser(r.Context(), &data)
	if err != nil {
		HandleServiceError(w, err)
		return
	}

	WriteJSONData(w, http.StatusCreated, map[string]interface{}{
		"user": map[string]interface{}{
			"id":    user.UserID,
			"email": user.Email,
			"name":  user.Name,
		},
		"token": user.Token,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if !GetRequest(w, r, &req) {
		return
	}

	data := services.LoginData{
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := services.LoginUser(r.Context(), &data)
	if err != nil {
		HandleServiceError(w, err)
		return
	}

	WriteJSONData(w, http.StatusOK, map[string]interface{}{
		"user": map[string]interface{}{
			"id":    user.UserID,
			"email": user.Email,
			"name":  user.Name,
		},
		"token": user.Token,
	})
}
