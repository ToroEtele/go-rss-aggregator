package main

import (
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/ToroEtele/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAd: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondWithJson(w, 200, databaseUserToUser(user))
}