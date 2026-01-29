package db

import (
	"context"
	"time"
)

type Transaction struct {
	ID              int
	UserID          int
	Amount          float64
	Description     string
	TransactionType string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// CreateTransaction inserts a new transaction into the database
func CreateTransaction(ctx context.Context, userID int, amount float64, description, transactionType string) (*Transaction, error) {
	query := `
		INSERT INTO transactions (user_id, amount, description, transaction_type, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id, user_id, amount, description, transaction_type, created_at, updated_at
	`

	var transaction Transaction
	err := DB.QueryRowContext(ctx, query, userID, amount, description, transactionType).Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.Amount,
		&transaction.Description,
		&transaction.TransactionType,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
