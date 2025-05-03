package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aryansharma2k4/rss-aggregator/internal/auth"
	"github.com/aryansharma2k4/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request){
	
	type parameters struct{
		Name string `json:name`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w,400,fmt.Sprint("Error Parsing JSON", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: 		uuid.New(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now().UTC(),
		Name:    	params.Name,
	})

	if err != nil{
		respondWithError(w,400,fmt.Sprintf("Couldn't create user %s", err))
	}
	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request){
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w,400,fmt.Sprintf("Couldn't create user: %v",err))
		return
	}

	user, err := apiKey.DB.GetUserByAPIKey(r.Context(), apiKey)

	if err != nil {
		respondWithError(w,400,fmt.Sprintf("Auth Error: %v",err))
	}

	respondWithJSON(w,200,databaseUserToUser(user))


}