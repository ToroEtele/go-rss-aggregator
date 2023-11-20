package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/ToroEtele/rss-aggregator/internal/database"
)


type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAd time.Time `json:"created_ad"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAd: dbUser.CreatedAd,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
	}
}