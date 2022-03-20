package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	DateOfCreation time.Time `json:"date-of-creation"`
	FavoriteCake   string    `json:"favorite-cake"`
}

type UserSignUpInput struct {
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	DateOfCreation time.Time `json:"date-of-creation"`
	FavoriteCake   string    `json:"favorite-cake"`
}
