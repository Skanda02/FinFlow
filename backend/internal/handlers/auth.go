package handlers

import (
	"net/http"

	"finflow/internal/http_helpers"
	"finflow/internal/services"
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

	if !http_helpers.GetRequest(w, r, &req) {
		return
	}

	data := services.RegisterData{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := services.RegisterNewUser(r.Context(), &data)
	if err != nil {
		http_helpers.HandleServiceError(w, err)
		return
	}

	http_helpers.WriteJSONData(w, http.StatusCreated, map[string]interface{}{
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

	if !http_helpers.GetRequest(w, r, &req) {
		return
	}

	data := services.LoginData{
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := services.LoginUser(r.Context(), &data)
	if err != nil {
		http_helpers.HandleServiceError(w, err)
		return
	}

	http_helpers.WriteJSONData(w, http.StatusOK, map[string]interface{}{
		"user": map[string]interface{}{
			"id":    user.UserID,
			"email": user.Email,
			"name":  user.Name,
		},
		"token": user.Token,
	})
}
