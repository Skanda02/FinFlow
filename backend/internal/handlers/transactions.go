package handlers

import (
	"net/http"

	"finflow/internal/http_helpers"
	"finflow/internal/middleware"
	"finflow/internal/services"
)

type AddIncomeRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

type AddExpenseRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

func AddIncome(w http.ResponseWriter, r *http.Request) {
	var req AddIncomeRequest

	if !http_helpers.GetRequest(w, r, &req) {
		return
	}

	// Extract user ID from context (set by auth middleware)
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http_helpers.WriteJSONError(w, http.StatusUnauthorized, "user not authenticated")
		return
	}

	data := services.TransactionData{
		UserID:      userID,
		Amount:      req.Amount,
		Description: req.Description,
	}

	transaction, err := services.AddIncome(r.Context(), &data)
	if err != nil {
		http_helpers.HandleServiceError(w, err)
		return
	}

	http_helpers.WriteJSONData(w, http.StatusCreated, map[string]interface{}{
		"id":          transaction.ID,
		"amount":      transaction.Amount,
		"description": transaction.Description,
		"type":        transaction.TransactionType,
	})
}

func AddExpense(w http.ResponseWriter, r *http.Request) {
	var req AddExpenseRequest

	if !http_helpers.GetRequest(w, r, &req) {
		return
	}

	// Extract user ID from context (set by auth middleware)
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http_helpers.WriteJSONError(w, http.StatusUnauthorized, "user not authenticated")
		return
	}

	data := services.TransactionData{
		UserID:      userID,
		Amount:      req.Amount,
		Description: req.Description,
	}

	transaction, err := services.AddExpense(r.Context(), &data)
	if err != nil {
		http_helpers.HandleServiceError(w, err)
		return
	}

	http_helpers.WriteJSONData(w, http.StatusCreated, map[string]interface{}{
		"id":          transaction.ID,
		"amount":      transaction.Amount,
		"description": transaction.Description,
		"type":        transaction.TransactionType,
	})
}
