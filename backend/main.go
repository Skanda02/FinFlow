package main

import (
	"log"
	"os"
	"net/http"

	"github.com/joho/godotenv"

	"finflow/internal/db"
	"finflow/internal/router"
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

	var port string = os.Getenv("PORT")
	var mux *http.ServeMux = http.NewServeMux()

	router.Register(mux)

	log.Println("Running server at :" + port)
	http.ListenAndServe(":"+port, mux)
}
