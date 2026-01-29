package plaid

import (
	"errors"
	"time"

	"finflow/internal/services"
)

// MockPlaidClient implements PlaidClient interface for testing
// Replace with actual Plaid SDK implementation in production
type MockPlaidClient struct{}

func NewMockPlaidClient() *MockPlaidClient {
	return &MockPlaidClient{}
}

func (m *MockPlaidClient) ExchangePublicToken(publicToken string) (string, string, error) {
	// In production, use actual Plaid SDK:
	// client, _ := plaid.NewClient(plaid.ClientOptions{...})
	// response, err := client.ExchangePublicToken(publicToken)

	if publicToken == "" {
		return "", "", errors.New("invalid public token")
	}

	// Mock response
	return "access-token-" + publicToken, "item-id-12345", nil
}

func (m *MockPlaidClient) GetTransactions(accessToken string, startDate, endDate time.Time) ([]services.BankTransaction, error) {
	// In production, use actual Plaid SDK:
	// client, _ := plaid.NewClient(plaid.ClientOptions{...})
	// response, err := client.GetTransactions(accessToken, startDate, endDate)

	// Mock transactions
	return []services.BankTransaction{
		{
			TransactionID: "txn-001",
			Amount:        50.00,
			Description:   "Grocery Store",
			Date:          time.Now().AddDate(0, 0, -1),
			IsDebit:       true,
		},
		{
			TransactionID: "txn-002",
			Amount:        1500.00,
			Description:   "Salary Deposit",
			Date:          time.Now().AddDate(0, 0, -5),
			IsDebit:       false,
		},
	}, nil
}

func (m *MockPlaidClient) GetInstitutionName(itemID string) (string, error) {
	// In production, use actual Plaid SDK to get institution details
	return "Mock Bank", nil
}
