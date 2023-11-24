package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Goel-Bhavye8302/RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

func(apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, req *http.Request, user database.User){
	type parameters struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON : %v", err))
		return 
	}

	feed, feedErr := apiCfg.DB.CreateFeed(req.Context(), database.CreateFeedParams{
		ID : uuid.New(),
		CreatedAt : time.Now().UTC(),
		UpdatedAt : time.Now().UTC(),
		Name : params.Name,
		Url : params.Url,
		UserID : user.ID,
	})

	if feedErr != nil{
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed : %v", feedErr))
		return 
	}

	respondWithJSON(w, 201, feed)
}

func(apiCfg *apiConfig) handlerGetAllFeed(w http.ResponseWriter, req *http.Request){

	feeds, feedErr := apiCfg.DB.GetFeeds(req.Context());

	if feedErr != nil{
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed follows : %v", feedErr))
		return 
	}

	respondWithJSON(w, 201, feeds)
}