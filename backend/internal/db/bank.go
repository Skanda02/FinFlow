package db

import (
	"context"
	"time"
)

type BankConnection struct {
	ID           int
	UserID       int
	BankName     string
	AccessToken  string
	ItemID       string
	LastSyncedAt *time.Time
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// CreateBankConnection creates a new bank connection
func CreateBankConnection(ctx context.Context, userID int, bankName, accessToken, itemID string) (*BankConnection, error) {
	query := `
		INSERT INTO bank_connections (user_id, bank_name, access_token, item_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id, user_id, bank_name, access_token, item_id, last_synced_at, is_active, created_at, updated_at
	`

	var conn BankConnection
	err := DB.QueryRowContext(ctx, query, userID, bankName, accessToken, itemID).Scan(
		&conn.ID,
		&conn.UserID,
		&conn.BankName,
		&conn.AccessToken,
		&conn.ItemID,
		&conn.LastSyncedAt,
		&conn.IsActive,
		&conn.CreatedAt,
		&conn.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &conn, nil
}

// GetBankConnectionsByUserID retrieves all bank connections for a user
func GetBankConnectionsByUserID(ctx context.Context, userID int) ([]*BankConnection, error) {
	query := `
		SELECT id, user_id, bank_name, access_token, item_id, last_synced_at, is_active, created_at, updated_at
		FROM bank_connections
		WHERE user_id = $1 AND is_active = true
		ORDER BY created_at DESC
	`

	rows, err := DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var connections []*BankConnection
	for rows.Next() {
		var conn BankConnection
		err := rows.Scan(
			&conn.ID,
			&conn.UserID,
			&conn.BankName,
			&conn.AccessToken,
			&conn.ItemID,
			&conn.LastSyncedAt,
			&conn.IsActive,
			&conn.CreatedAt,
			&conn.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		connections = append(connections, &conn)
	}

	return connections, nil
}

// UpdateBankConnectionSyncTime updates the last synced timestamp
func UpdateBankConnectionSyncTime(ctx context.Context, id int) error {
	query := `
		UPDATE bank_connections
		SET last_synced_at = NOW(), updated_at = NOW()
		WHERE id = $1
	`

	_, err := DB.ExecContext(ctx, query, id)
	return err
}

// DeactivateBankConnection marks a bank connection as inactive
func DeactivateBankConnection(ctx context.Context, id int, userID int) error {
	query := `
		UPDATE bank_connections
		SET is_active = false, updated_at = NOW()
		WHERE id = $1 AND user_id = $2
	`

	_, err := DB.ExecContext(ctx, query, id, userID)
	return err
}
