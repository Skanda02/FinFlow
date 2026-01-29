package router

import (
	"net/http"

	"finflow/internal/handlers"
	"finflow/internal/middleware"
)

// applyMiddleware chains multiple middleware functions
func applyMiddleware(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// Register sets up all routes with appropriate middleware
func Register(mux *http.ServeMux) {
	// Health check - no middleware (for load balancers, monitoring)
	mux.HandleFunc("/health", handlers.Health)

	// Common middleware for all other routes
	commonMiddleware := []func(http.HandlerFunc) http.HandlerFunc{
		middleware.LoggingMiddleware,
		middleware.CORSMiddleware,
	}

	// Public routes - no authentication required
	publicMiddleware := append(commonMiddleware, middleware.RateLimitMiddleware)

	mux.HandleFunc("/register", applyMiddleware(handlers.Register, publicMiddleware...))
	mux.HandleFunc("/login", applyMiddleware(handlers.Login, publicMiddleware...))

	// Protected routes - authentication required
	protectedMiddleware := append(commonMiddleware, middleware.AuthMiddleware)

	mux.HandleFunc("/add_income", applyMiddleware(handlers.AddIncome, protectedMiddleware...))
	mux.HandleFunc("/add_expense", applyMiddleware(handlers.AddExpense, protectedMiddleware...))

	// Bank connection endpoints
	mux.HandleFunc("/bank/link", applyMiddleware(handlers.LinkBank, protectedMiddleware...))
	mux.HandleFunc("/bank/unlink", applyMiddleware(handlers.UnlinkBank, protectedMiddleware...))
	mux.HandleFunc("/bank/sync", applyMiddleware(handlers.SyncBankTransactions, protectedMiddleware...))
	mux.HandleFunc("/bank/connections", applyMiddleware(handlers.GetBankConnections, protectedMiddleware...))

	// Summary endpoints
	mux.HandleFunc("/summary/weekly", applyMiddleware(handlers.GetWeeklySummary, protectedMiddleware...))
	mux.HandleFunc("/summary/monthly", applyMiddleware(handlers.GetMonthlySummary, protectedMiddleware...))
	mux.HandleFunc("/summary/custom", applyMiddleware(handlers.GetCustomSummary, protectedMiddleware...))
}
