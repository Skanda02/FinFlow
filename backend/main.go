package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"finflow/internal/db"
	"finflow/internal/handlers"
	"finflow/internal/plaid"
	"finflow/internal/router"
	"finflow/internal/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Initialize database connection
	if err := db.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.CloseDB()

	// Initialize bank service with Plaid client
	plaidClient, err := plaid.NewRealPlaidClient()
	if err != nil {
		log.Printf("Warning: Plaid client initialization failed: %v", err)
		log.Println("Bank integration features will not be available")
		log.Println("Please set PLAID_CLIENT_ID and PLAID_SECRET environment variables")
		// Continue without bank service - will return errors if bank endpoints are called
	} else {
		bankService := services.NewBankService(plaidClient)
		handlers.SetBankService(bankService)
		log.Println("Plaid integration initialized successfully")
	}

	var port string = os.Getenv("PORT")
	var mux *http.ServeMux = http.NewServeMux()

	router.Register(mux)

	log.Println("Running server at :" + port)
	http.ListenAndServe(":"+port, mux)
}
