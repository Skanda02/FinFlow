package router

import (
	"net/http"

	"finflow/internal/handlers"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/health", handlers.Health)

	mux.HandleFunc("/register", handlers.Register)
	mux.HandleFunc("/login", handlers.Login)

	mux.HandleFunc("/add_income", handlers.AddIncome)
	mux.HandleFunc("/add_expense", handlers.AddExpense)
}
