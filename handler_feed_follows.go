package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Goel-Bhavye8302/RSS-Aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func(apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, req *http.Request, user database.User){
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON : %v", err))
		return 
	}

	feedFollow, feedErr := apiCfg.DB.CreateFeedFollow(req.Context(), database.CreateFeedFollowParams{
		ID : uuid.New(),
		CreatedAt : time.Now().UTC(),
		UpdatedAt : time.Now().UTC(),
		UserID : user.ID,
		FeedID: params.FeedId,
	})

	if feedErr != nil{
		respondWithError(w, 400, fmt.Sprintf("Couldn't follow feed : %v", feedErr))
		return 
	}

	respondWithJSON(w, 201, feedFollow)
}

func(apiCfg *apiConfig) handlerGetAllFeedFollows(w http.ResponseWriter, req *http.Request, user database.User){

	feed_follows, feedErr := apiCfg.DB.GetFeedFollows(req.Context(), user.ID);

	if feedErr != nil{
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed follows : %v", feedErr))
		return 
	}

	respondWithJSON(w, 201, feed_follows)
}

func(apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, req *http.Request, user database.User){

	feedIdStr := chi.URLParam(req, "feedFollowID")

	feed_id, err := uuid.Parse(feedIdStr)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Feed id provided is wrong : %v", err))
		return 
	}

	feedErr := apiCfg.DB.DeleteFeedFollow(req.Context(), database.DeleteFeedFollowParams{
		FeedID: feed_id,
		UserID: user.ID,
	})

	if feedErr != nil{
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follows : %v", feedErr))
		return 
	}

	respondWithJSON(w, 201, struct{}{})
}