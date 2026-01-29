package services

import (
	"context"
	"errors"
	"time"

	"finflow/internal/db"
)

// Plaid/Bank integration errors
var (
	ErrBankConnectionFailed = errors.New("failed to connect to bank")
	ErrBankSyncFailed       = errors.New("failed to sync bank transactions")
)

// PlaidClient interface for bank integration (implement with actual Plaid SDK)
type PlaidClient interface {
	ExchangePublicToken(publicToken string) (accessToken string, itemID string, err error)
	GetTransactions(accessToken string, startDate, endDate time.Time) ([]BankTransaction, error)
	GetInstitutionName(itemID string) (string, error)
}

type BankTransaction struct {
	TransactionID string
	Amount        float64
	Description   string
	Date          time.Time
	IsDebit       bool // true for expenses, false for income
}

type BankService struct {
	plaidClient PlaidClient
}

func NewBankService(plaidClient PlaidClient) *BankService {
	return &BankService{
		plaidClient: plaidClient,
	}
}

// LinkBankAccount exchanges public token and stores bank connection
func (s *BankService) LinkBankAccount(ctx context.Context, userID int, publicToken string) error {
	// Exchange public token for access token
	accessToken, itemID, err := s.plaidClient.ExchangePublicToken(publicToken)
	if err != nil {
		return ErrBankConnectionFailed
	}

	// Get institution name
	bankName, err := s.plaidClient.GetInstitutionName(itemID)
	if err != nil {
		bankName = "Unknown Bank"
	}

	// Store bank connection
	_, err = db.CreateBankConnection(ctx, userID, bankName, accessToken, itemID)
	if err != nil {
		return ErrInternal
	}

	return nil
}

// SyncBankTransactions fetches and stores transactions from all connected banks
func (s *BankService) SyncBankTransactions(ctx context.Context, userID int) error {
	// Get all active bank connections
	connections, err := db.GetBankConnectionsByUserID(ctx, userID)
	if err != nil {
		return ErrInternal
	}

	// Sync transactions from last 30 days
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -30)

	for _, conn := range connections {
		// Fetch transactions from bank
		bankTxns, err := s.plaidClient.GetTransactions(conn.AccessToken, startDate, endDate)
		if err != nil {
			// Log error but continue with other connections
			continue
		}

		// Store transactions
		for _, txn := range bankTxns {
			// Check if transaction already exists
			exists, err := db.BankTransactionExists(ctx, userID, txn.TransactionID)
			if err != nil || exists {
				continue
			}

			// Determine transaction type
			txnType := "income"
			amount := txn.Amount
			if txn.IsDebit {
				txnType = "expense"
			} else {
				// Income amounts should be positive
				amount = -amount
			}

			// Create transaction
			_, err = db.CreateBankTransaction(ctx, userID, amount, txn.Description, txnType, txn.TransactionID, txn.Date)
			if err != nil {
				// Log error but continue
				continue
			}
		}

		// Update last synced time
		db.UpdateBankConnectionSyncTime(ctx, conn.ID)
	}

	return nil
}

// UnlinkBankAccount deactivates a bank connection
func (s *BankService) UnlinkBankAccount(ctx context.Context, userID, connectionID int) error {
	return db.DeactivateBankConnection(ctx, connectionID, userID)
}

// GetBankConnections retrieves all active bank connections for a user
func (s *BankService) GetBankConnections(ctx context.Context, userID int) ([]*db.BankConnection, error) {
	return db.GetBankConnectionsByUserID(ctx, userID)
}
