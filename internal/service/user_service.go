package service

import (
	"github.com/google/uuid"
	"github.com/smallretardedfish/cakes/internal/domain"
	"log"
	"time"
)

type UserRepository interface {
	Create(user *domain.User) error
	Get(id uuid.UUID) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id uuid.UUID) error
	ReadAll() ([]*domain.User, error)
}

type UserService struct {
	userRepo UserRepository
}

func (us *UserService) Get(id uuid.UUID) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (us *UserService) Update(user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (us *UserService) Delete(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (us *UserService) ReadAll() ([]*domain.User, error) {
	return us.userRepo.ReadAll()
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) Create(inp *domain.UserSignUpInput) error {
	id, err := uuid.NewUUID()
	if err != nil {
		log.Println(err)
		return err
	}
	user := &domain.User{
		ID:             id,
		Name:           inp.Name,
		Email:          inp.Email,
		DateOfCreation: time.Now(),
		FavoriteCake:   inp.FavoriteCake,
	}
	return us.userRepo.Create(user)
}
