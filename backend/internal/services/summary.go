package services

import (
	"context"
	"time"

	"finflow/internal/db"
)

type TransactionSummary struct {
	TotalIncome  float64                  `json:"total_income"`
	TotalExpense float64                  `json:"total_expense"`
	NetAmount    float64                  `json:"net_amount"`
	Transactions []TransactionDetail      `json:"transactions"`
	BySource     map[string]SourceSummary `json:"by_source"`
}

type TransactionDetail struct {
	ID              int       `json:"id"`
	Amount          float64   `json:"amount"`
	Description     string    `json:"description"`
	Type            string    `json:"type"`
	Source          string    `json:"source"`
	TransactionDate time.Time `json:"transaction_date"`
}

type SourceSummary struct {
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
}

// GetWeeklySummary returns transaction summary for the current week
func GetWeeklySummary(ctx context.Context, userID int) (*TransactionSummary, error) {
	now := time.Now()
	// Start of week (Monday)
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	startOfWeek := now.AddDate(0, 0, -(weekday - 1)).Truncate(24 * time.Hour)
	endOfWeek := startOfWeek.AddDate(0, 0, 7)

	return getSummaryForPeriod(ctx, userID, startOfWeek, endOfWeek)
}

// GetMonthlySummary returns transaction summary for the current month
func GetMonthlySummary(ctx context.Context, userID int) (*TransactionSummary, error) {
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0)

	return getSummaryForPeriod(ctx, userID, startOfMonth, endOfMonth)
}

// GetCustomSummary returns transaction summary for a custom date range
func GetCustomSummary(ctx context.Context, userID int, startDate, endDate time.Time) (*TransactionSummary, error) {
	return getSummaryForPeriod(ctx, userID, startDate, endDate)
}

func getSummaryForPeriod(ctx context.Context, userID int, startDate, endDate time.Time) (*TransactionSummary, error) {
	transactions, err := db.GetTransactionsByDateRange(ctx, userID, startDate, endDate)
	if err != nil {
		return nil, ErrInternal
	}

	summary := &TransactionSummary{
		Transactions: make([]TransactionDetail, 0),
		BySource: map[string]SourceSummary{
			"manual": {Income: 0, Expense: 0},
			"bank":   {Income: 0, Expense: 0},
		},
	}

	for _, txn := range transactions {
		// Add to totals
		if txn.TransactionType == "income" {
			summary.TotalIncome += txn.Amount
			sourceSummary := summary.BySource[txn.Source]
			sourceSummary.Income += txn.Amount
			summary.BySource[txn.Source] = sourceSummary
		} else {
			summary.TotalExpense += txn.Amount
			sourceSummary := summary.BySource[txn.Source]
			sourceSummary.Expense += txn.Amount
			summary.BySource[txn.Source] = sourceSummary
		}

		// Add to transaction details
		summary.Transactions = append(summary.Transactions, TransactionDetail{
			ID:              txn.ID,
			Amount:          txn.Amount,
			Description:     txn.Description,
			Type:            txn.TransactionType,
			Source:          txn.Source,
			TransactionDate: txn.TransactionDate,
		})
	}

	summary.NetAmount = summary.TotalIncome - summary.TotalExpense

	return summary, nil
}
