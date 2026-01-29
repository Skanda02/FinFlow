package handlers

import (
	"net/http"

	"finflow/internal/http_helpers"
	"finflow/internal/middleware"
	"finflow/internal/services"
)

type LinkBankRequest struct {
	PublicToken string `json:"public_token"`
}

type UnlinkBankRequest struct {
	ConnectionID int `json:"connection_id"`
}

var bankService *services.BankService

// SetBankService sets the bank service instance
func SetBankService(bs *services.BankService) {
	bankService = bs
}

// LinkBank handles bank account linking
func LinkBank(w http.ResponseWriter, r *http.Request) {
	if bankService == nil {
		http_helpers.WriteJSONError(w, http.StatusServiceUnavailable, "Bank integration is not configured. Please contact administrator.")
		return
	}

	var req LinkBankRequest

	if !http_helpers.GetRequest(w, r, &req) {
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http_helpers.WriteJSONError(w, http.StatusUnauthorized, "user not authenticated")
		return
	}

	err := bankService.LinkBankAccount(r.Context(), userID, req.PublicToken)
	if err != nil {
		http_helpers.HandleServiceError(w, err)
		return
	}

	http_helpers.WriteJSONData(w, http.StatusCreated, map[string]interface{}{
		"message": "Bank account linked successfully",
	})
}

// UnlinkBank handles bank account unlinking
func UnlinkBank(w http.ResponseWriter, r *http.Request) {
	if bankService == nil {
		http_helpers.WriteJSONError(w, http.StatusServiceUnavailable, "Bank integration is not configured. Please contact administrator.")
		return
	}

	var req UnlinkBankRequest

	if !http_helpers.GetRequest(w, r, &req) {
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http_helpers.WriteJSONError(w, http.StatusUnauthorized, "user not authenticated")
		return
	}

	err := bankService.UnlinkBankAccount(r.Context(), userID, req.ConnectionID)
	if err != nil {
		http_helpers.HandleServiceError(w, err)
		return
	}

	http_helpers.WriteJSONData(w, http.StatusOK, map[string]interface{}{
		"message": "Bank account unlinked successfully",
	})
}

// SyncBankTransactions triggers a sync of bank transactions
func SyncBankTransactions(w http.ResponseWriter, r *http.Request) {
	if bankService == nil {
		http_helpers.WriteJSONError(w, http.StatusServiceUnavailable, "Bank integration is not configured. Please contact administrator.")
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http_helpers.WriteJSONError(w, http.StatusUnauthorized, "user not authenticated")
		return
	}

	err := bankService.SyncBankTransactions(r.Context(), userID)
	if err != nil {
		http_helpers.HandleServiceError(w, err)
		return
	}

	http_helpers.WriteJSONData(w, http.StatusOK, map[string]interface{}{
		"message": "Bank transactions synced successfully",
	})
}

// GetBankConnections retrieves all connected banks
func GetBankConnections(w http.ResponseWriter, r *http.Request) {
	if bankService == nil {
		http_helpers.WriteJSONError(w, http.StatusServiceUnavailable, "Bank integration is not configured. Please contact administrator.")
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		http_helpers.WriteJSONError(w, http.StatusUnauthorized, "user not authenticated")
		return
	}

	connections, err := bankService.GetBankConnections(r.Context(), userID)
	if err != nil {
		http_helpers.HandleServiceError(w, err)
		return
	}

	// Format response (don't expose access tokens)
	var response []map[string]interface{}
	for _, conn := range connections {
		response = append(response, map[string]interface{}{
			"id":             conn.ID,
			"bank_name":      conn.BankName,
			"last_synced_at": conn.LastSyncedAt,
			"is_active":      conn.IsActive,
		})
	}

	http_helpers.WriteJSONData(w, http.StatusOK, response)
}

