package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()

	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	Port := os.Getenv("PORT")

	fmt.Println("Port =", Port)

	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + Port,
	}

	router.Get("/ready", handlerReadiness)

	router.Get("/err", func(w http.ResponseWriter, r *http.Request) {
		respondWithError(w, 400, "Something Went Wrong")
	})

	srvErr := server.ListenAndServe()

	if srvErr != nil {
		log.Fatal("Server Error : ", srvErr.Error())
	}
}
