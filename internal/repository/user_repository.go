package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/smallretardedfish/cakes/internal/domain"
	"github.com/smallretardedfish/cakes/pkg/errors"
	"sync"
)

type MemoryStorage struct {
	lock    sync.RWMutex
	storage map[string]*domain.User
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		lock:    sync.RWMutex{},
		storage: make(map[string]*domain.User),
	}
}

func (m *MemoryStorage) Create(user *domain.User) error {
	m.lock.Lock()
	if _, ok := m.storage[user.Email]; ok {
		fmt.Println(m.storage)
		return errors.ErrUserAlreadyExist
	}
	m.storage[user.Email] = user
	m.lock.Unlock()
	fmt.Println(len(m.storage))
	return nil

}

func (m MemoryStorage) Get(id uuid.UUID) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m MemoryStorage) Update(user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (m MemoryStorage) Delete(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (m MemoryStorage) ReadAll() ([]*domain.User, error) {
	users := make([]*domain.User, 0, len(m.storage))
	for _, val := range m.storage {
		users = append(users, val)
	}
	return users, nil
}
