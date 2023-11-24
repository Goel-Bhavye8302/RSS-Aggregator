package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Goel-Bhavye8302/RSS-Aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	envErr := godotenv.Load()

	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	Port := os.Getenv("PORT")
	if Port == "" {
		log.Fatal("Port not found in ENV file")
	}

	DbUrl := os.Getenv("DB_URL")
	if DbUrl == "" {
		log.Fatal("Database Connection URL not found in ENV file")
	}

	conn, dbErr := sql.Open("postgres", DbUrl)

	if dbErr != nil {
		log.Fatal("Could'nt connect to the database")
	}

	db := database.New(conn)
	apiCnfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	server := &http.Server{
		Handler: router,
		Addr:    ":" + Port,
	}

	router.Get("/ready", handlerReadiness)

	router.Get("/err", func(w http.ResponseWriter, r *http.Request) {
		respondWithError(w, 400, "Something Went Wrong")
	})

	router.Post("/users", apiCnfg.handlerCreateUser)

	router.Get("/users", apiCnfg.middlewareAuth(apiCnfg.handlerGetUserByAPIKey))

	router.Post("/feeds", apiCnfg.middlewareAuth(apiCnfg.handlerCreateFeed))

	router.Get("/feeds", apiCnfg.handlerGetAllFeed)

	router.Get("/posts", apiCnfg.middlewareAuth(apiCnfg.handlerGetPostsForUser))

	router.Post("/feed_follows", apiCnfg.middlewareAuth(apiCnfg.handlerCreateFeedFollow))

	router.Get("/feed_follows", apiCnfg.middlewareAuth(apiCnfg.handlerGetAllFeedFollows))

	router.Delete("/feed_follows/{feedFollowID}", apiCnfg.middlewareAuth(apiCnfg.handlerDeleteFeedFollows))


	srvErr := server.ListenAndServe()

	if srvErr != nil {
		log.Fatal("Server Error : ", srvErr.Error())
	}
}
