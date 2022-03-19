package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID             uuid.UUID
	Name           string
	Email          string
	DateOfCreation time.Time
	FavoriteCake   string
}
