package main

import (
	"fmt"
	"net/http"

	"github.com/Goel-Bhavye8302/RSS-Aggregator/internal/auth"
	"github.com/Goel-Bhavye8302/RSS-Aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request){
		apiKey, err := auth.GetApiKey(req.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth Error : %v", err))
			return 
		}

		user, usrErr := apiCfg.DB.GetUserByAPIKey(req.Context(), apiKey)
		if usrErr != nil {
			respondWithError(w, 403, fmt.Sprintf("Couldn't get user : %v", usrErr))
			return 
		}

		handler(w, req, user)
	}
}

