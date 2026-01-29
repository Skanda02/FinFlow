package plaid

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/plaid/plaid-go/v26/plaid"

	"finflow/internal/services"
)

// RealPlaidClient implements PlaidClient interface using Plaid SDK
type RealPlaidClient struct {
	client *plaid.APIClient
	ctx    context.Context
}

// NewRealPlaidClient creates a new Plaid client
func NewRealPlaidClient() (*RealPlaidClient, error) {
	clientID := os.Getenv("PLAID_CLIENT_ID")
	secret := os.Getenv("PLAID_SECRET")
	environment := os.Getenv("PLAID_ENVIRONMENT") // sandbox, development, or production

	if clientID == "" || secret == "" {
		return nil, errors.New("PLAID_CLIENT_ID and PLAID_SECRET environment variables must be set")
	}

	if environment == "" {
		environment = "sandbox"
	}

	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", clientID)
	configuration.AddDefaultHeader("PLAID-SECRET", secret)

	// Set environment
	switch environment {
	case "sandbox":
		configuration.UseEnvironment(plaid.Sandbox)
	case "development":
		configuration.UseEnvironment(plaid.Development)
	case "production":
		configuration.UseEnvironment(plaid.Production)
	default:
		configuration.UseEnvironment(plaid.Sandbox)
	}

	client := plaid.NewAPIClient(configuration)

	return &RealPlaidClient{
		client: client,
		ctx:    context.Background(),
	}, nil
}

// ExchangePublicToken exchanges a public token for an access token
func (c *RealPlaidClient) ExchangePublicToken(publicToken string) (string, string, error) {
	request := c.client.PlaidApi.ItemPublicTokenExchange(c.ctx)
	request = request.ItemPublicTokenExchangeRequest(*plaid.NewItemPublicTokenExchangeRequest(publicToken))

	response, _, err := request.Execute()
	if err != nil {
		return "", "", err
	}

	return response.GetAccessToken(), response.GetItemId(), nil
}

// GetTransactions retrieves transactions from Plaid
func (c *RealPlaidClient) GetTransactions(accessToken string, startDate, endDate time.Time) ([]services.BankTransaction, error) {
	request := c.client.PlaidApi.TransactionsGet(c.ctx)
	request = request.TransactionsGetRequest(*plaid.NewTransactionsGetRequest(
		accessToken,
		startDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
	))

	response, _, err := request.Execute()
	if err != nil {
		return nil, err
	}

	var transactions []services.BankTransaction
	for _, txn := range response.GetTransactions() {
		// Parse transaction date
		txnDate, err := time.Parse("2006-01-02", txn.GetDate())
		if err != nil {
			txnDate = time.Now()
		}

		// Plaid returns positive amounts for debits (expenses)
		// and negative amounts for credits (income)
		amount := txn.GetAmount()
		isDebit := amount > 0

		// Convert to absolute value
		if amount < 0 {
			amount = -amount
		}

		transactions = append(transactions, services.BankTransaction{
			TransactionID: txn.GetTransactionId(),
			Amount:        amount,
			Description:   txn.GetName(),
			Date:          txnDate,
			IsDebit:       isDebit,
		})
	}

	return transactions, nil
}

// GetInstitutionName retrieves the name of the financial institution
func (c *RealPlaidClient) GetInstitutionName(itemID string) (string, error) {
	// First get the item to retrieve the institution ID
	itemRequest := c.client.PlaidApi.ItemGet(c.ctx)
	itemRequest = itemRequest.ItemGetRequest(*plaid.NewItemGetRequest(itemID))

	itemResponse, _, err := itemRequest.Execute()
	if err != nil {
		return "", err
	}

	item := itemResponse.GetItem()
	institutionID := item.GetInstitutionId()
	if institutionID == "" {
		return "Unknown Bank", nil
	}

	// Get institution details
	instRequest := c.client.PlaidApi.InstitutionsGetById(c.ctx)
	instRequest = instRequest.InstitutionsGetByIdRequest(*plaid.NewInstitutionsGetByIdRequest(
		institutionID,
		[]plaid.CountryCode{plaid.COUNTRYCODE_US},
	))

	instResponse, _, err := instRequest.Execute()
	if err != nil {
		return "Unknown Bank", nil
	}

	institution := instResponse.GetInstitution()
	name := institution.GetName()
	if name == "" {
		return "Unknown Bank", nil
	}
	
	return name, nil
}
