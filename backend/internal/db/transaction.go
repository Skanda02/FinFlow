package db

import (
	"context"
	"time"
)

type Transaction struct {
	ID                int
	UserID            int
	Amount            float64
	Description       string
	TransactionType   string
	Source            string
	BankTransactionID *string
	TransactionDate   time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// CreateTransaction inserts a new transaction into the database
func CreateTransaction(ctx context.Context, userID int, amount float64, description, transactionType string) (*Transaction, error) {
	query := `
		INSERT INTO transactions (user_id, amount, description, transaction_type, source, transaction_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, 'manual', NOW(), NOW(), NOW())
		RETURNING id, user_id, amount, description, transaction_type, source, bank_transaction_id, transaction_date, created_at, updated_at
	`

	var transaction Transaction
	err := DB.QueryRowContext(ctx, query, userID, amount, description, transactionType).Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.Amount,
		&transaction.Description,
		&transaction.TransactionType,
		&transaction.Source,
		&transaction.BankTransactionID,
		&transaction.TransactionDate,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

// CreateBankTransaction inserts a bank-sourced transaction
func CreateBankTransaction(ctx context.Context, userID int, amount float64, description, transactionType, bankTransactionID string, transactionDate time.Time) (*Transaction, error) {
	query := `
		INSERT INTO transactions (user_id, amount, description, transaction_type, source, bank_transaction_id, transaction_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, 'bank', $5, $6, NOW(), NOW())
		RETURNING id, user_id, amount, description, transaction_type, source, bank_transaction_id, transaction_date, created_at, updated_at
	`

	var transaction Transaction
	err := DB.QueryRowContext(ctx, query, userID, amount, description, transactionType, bankTransactionID, transactionDate).Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.Amount,
		&transaction.Description,
		&transaction.TransactionType,
		&transaction.Source,
		&transaction.BankTransactionID,
		&transaction.TransactionDate,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

// GetTransactionsByDateRange retrieves transactions within a date range
func GetTransactionsByDateRange(ctx context.Context, userID int, startDate, endDate time.Time) ([]*Transaction, error) {
	query := `
		SELECT id, user_id, amount, description, transaction_type, source, bank_transaction_id, transaction_date, created_at, updated_at
		FROM transactions
		WHERE user_id = $1 AND transaction_date >= $2 AND transaction_date < $3
		ORDER BY transaction_date DESC
	`

	rows, err := DB.QueryContext(ctx, query, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*Transaction
	for rows.Next() {
		var t Transaction
		err := rows.Scan(
			&t.ID,
			&t.UserID,
			&t.Amount,
			&t.Description,
			&t.TransactionType,
			&t.Source,
			&t.BankTransactionID,
			&t.TransactionDate,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &t)
	}

	return transactions, nil
}

// BankTransactionExists checks if a bank transaction ID already exists
func BankTransactionExists(ctx context.Context, userID int, bankTransactionID string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM transactions WHERE user_id = $1 AND bank_transaction_id = $2)`

	var exists bool
	err := DB.QueryRowContext(ctx, query, userID, bankTransactionID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
