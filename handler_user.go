package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	// "github.com/Goel-Bhavye8302/RSS-Aggregator/internal/auth"
	"github.com/Goel-Bhavye8302/RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

func(apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, req *http.Request){
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		// log.Fatal("Error parsing JSON :", err)
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON : %v", err))
		return 
	}

	user, usrErr := apiCfg.DB.CreateUser(req.Context(), database.CreateUserParams{
		ID : uuid.New(),
		CreatedAt : time.Now().UTC(),
		UpdatedAt : time.Now().UTC(),
		Name : params.Name,
	})

	if usrErr != nil{
		// log.Fatal("Couldn't create user :", err)
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user : %v", err))
		return 
	}

	respondWithJSON(w, 201, user)
}

func(apiCfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, req *http.Request, user database.User){
	respondWithJSON(w, 201, user)	
}

func(apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, req *http.Request, user database.User){
	posts, err := apiCfg.DB.GetPostsForUser(req.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get posts : %v", err))
		return
	}

	respondWithJSON(w, 200, posts)
}