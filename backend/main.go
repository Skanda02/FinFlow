package main

import (
	"log"
	"os"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	var port string = os.Getenv("PORT")
	var mux *http.ServeMux = http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("running!"))
	})

	log.Println("Running server at :" + port)
	http.ListenAndServe(":"+port, mux)
}
