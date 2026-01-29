package handlers

import (
	"net/http"
	"time"

	"finflow/internal/http_helpers"
	"finflow/internal/middleware"
	"finflow/internal/services"
)

type CustomSummaryRequest struct {
	StartDate string `json:"start_date"` // Format: "2006-01-02"
	EndDate   string `json:"end_date"`   // Format: "2006-01-02"
}

// GetWeeklySummary returns transaction summary for current week
func GetWeeklySummary(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http_helpers.WriteJSONError(w, http.StatusUnauthorized, "user not authenticated")
		return
	}

	summary, err := services.GetWeeklySummary(r.Context(), userID)
	if err != nil {
		http_helpers.HandleServiceError(w, err)
		return
	}

	http_helpers.WriteJSONData(w, http.StatusOK, summary)
}

// GetMonthlySummary returns transaction summary for current month
func GetMonthlySummary(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http_helpers.WriteJSONError(w, http.StatusUnauthorized, "user not authenticated")
		return
	}

	summary, err := services.GetMonthlySummary(r.Context(), userID)
	if err != nil {
		http_helpers.HandleServiceError(w, err)
		return
	}

	http_helpers.WriteJSONData(w, http.StatusOK, summary)
}

// GetCustomSummary returns transaction summary for custom date range
func GetCustomSummary(w http.ResponseWriter, r *http.Request) {
	var req CustomSummaryRequest

	if !http_helpers.GetRequest(w, r, &req) {
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http_helpers.WriteJSONError(w, http.StatusUnauthorized, "user not authenticated")
		return
	}

	// Parse dates
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		http_helpers.WriteJSONError(w, http.StatusBadRequest, "invalid start_date format")
		return
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		http_helpers.WriteJSONError(w, http.StatusBadRequest, "invalid end_date format")
		return
	}

	summary, err := services.GetCustomSummary(r.Context(), userID, startDate, endDate)
	if err != nil {
		http_helpers.HandleServiceError(w, err)
		return
	}

	http_helpers.WriteJSONData(w, http.StatusOK, summary)
}
