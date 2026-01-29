package services

import (
	"context"
	"errors"

	"finflow/internal/db"
)

// Custom error types for transaction operations
var (
	ErrInvalidAmount = errors.New("invalid amount")
)

type TransactionData struct {
	UserID      int
	Amount      float64
	Description string
}

type TransactionResponse struct {
	ID              int
	UserID          int
	Amount          float64
	Description     string
	TransactionType string
}

// AddIncome creates a new income transaction
func AddIncome(ctx context.Context, data *TransactionData) (*TransactionResponse, error) {
	// Validate amount
	if data.Amount <= 0 {
		return nil, ErrInvalidAmount
	}

	// Create income transaction
	transaction, err := db.CreateTransaction(ctx, data.UserID, data.Amount, data.Description, "income")
	if err != nil {
		return nil, ErrInternal
	}

	return &TransactionResponse{
		ID:              transaction.ID,
		UserID:          transaction.UserID,
		Amount:          transaction.Amount,
		Description:     transaction.Description,
		TransactionType: transaction.TransactionType,
	}, nil
}

// AddExpense creates a new expense transaction
func AddExpense(ctx context.Context, data *TransactionData) (*TransactionResponse, error) {
	// Validate amount
	if data.Amount <= 0 {
		return nil, ErrInvalidAmount
	}

	// Create expense transaction
	transaction, err := db.CreateTransaction(ctx, data.UserID, data.Amount, data.Description, "expense")
	if err != nil {
		return nil, ErrInternal
	}

	return &TransactionResponse{
		ID:              transaction.ID,
		UserID:          transaction.UserID,
		Amount:          transaction.Amount,
		Description:     transaction.Description,
		TransactionType: transaction.TransactionType,
	}, nil
}
