package service

import (
	"github.com/cakes/internal/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	Add(user *domain.User) error
	Get(id uuid.UUID) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id uuid.UUID) error
	ReadAll() ([]*domain.User, error)
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}
