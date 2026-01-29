package handlers

import (
	"net/http"
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

	if !GetRequest(w, r, &req) {
		return
	}

	if req.Amount < 0.0 {
		WriteJSONError(w, http.StatusBadRequest, "negetive amount")
		return
	}

	WriteJSONError(w, http.StatusNotImplemented, "Yet to implement")
}

func AddExpense(w http.ResponseWriter, r *http.Request) {
	var req AddExpenseRequest

	if !GetRequest(w, r, &req) {
		return
	}

	if req.Amount < 0.0 {
		WriteJSONError(w, http.StatusBadRequest, "negetive amount")
		return
	}

	WriteJSONError(w, http.StatusNotImplemented, "Yet to implement")
}
