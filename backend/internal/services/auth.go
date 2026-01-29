package services

import (
	"context"
	"errors"
	"strings"

	"finflow/internal/db"
	"finflow/internal/utils"
)

// Custom error types for better error handling
var (
	ErrValidation         = errors.New("validation error")
	ErrUserExists         = errors.New("user with this email already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrInternal           = errors.New("internal server error")
)

type RegisterData struct {
	Name     string
	Email    string
	Password string
}

type LoginData struct {
	Email    string
	Password string
}

type AuthResponse struct {
	UserID int
	Name   string
	Email  string
	Token  string
}

func RegisterNewUser(ctx context.Context, data *RegisterData) (*AuthResponse, error) {
	// Validate input
	if data.Email == "" || data.Password == "" {
		return nil, ErrValidation
	}

	// Normalize email
	data.Email = strings.ToLower(strings.TrimSpace(data.Email))

	// Check if user already exists
	exists, err := db.UserExists(ctx, data.Email)
	if err != nil {
		return nil, ErrInternal
	}
	if exists {
		return nil, ErrUserExists
	}

	// Hash the password
	passwordHash, err := utils.HashPassword(data.Password)
	if err != nil {
		return nil, ErrInternal
	}

	// Create user in database
	user, err := db.CreateUser(ctx, data.Name, data.Email, passwordHash)
	if err != nil {
		return nil, ErrInternal
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return nil, ErrInternal
	}

	return &AuthResponse{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Token:  token,
	}, nil
}

func LoginUser(ctx context.Context, data *LoginData) (*AuthResponse, error) {
	// Validate input
	if data.Email == "" || data.Password == "" {
		return nil, ErrValidation
	}

	// Normalize email
	data.Email = strings.ToLower(strings.TrimSpace(data.Email))

	// Get user from database
	user, err := db.GetUserByEmail(ctx, data.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Check password
	if !utils.CheckPassword(data.Password, user.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return nil, ErrInternal
	}

	return &AuthResponse{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Token:  token,
	}, nil
}
