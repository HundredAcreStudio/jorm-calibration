package store

import (
	"fmt"
	"sync"
	"time"

	"github.com/HundredAcreStudio/jorm-calibration/internal/model"
)

// Store defines the interface for user persistence.
type Store interface {
	List() ([]model.User, error)
	Get(id string) (*model.User, error)
	Create(user model.User) (*model.User, error)
	Delete(id string) error
}

// MemoryStore is an in-memory implementation of Store.
type MemoryStore struct {
	mu    sync.RWMutex
	users map[string]model.User
	seq   int
}

// NewMemoryStore creates a new in-memory store.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users: make(map[string]model.User),
	}
}

func (s *MemoryStore) List() ([]model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]model.User, 0, len(s.users))
	for _, u := range s.users {
		users = append(users, u)
	}
	return users, nil
}

func (s *MemoryStore) Get(id string) (*model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	u, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user %q not found", id)
	}
	return &u, nil
}

func (s *MemoryStore) Create(user model.User) (*model.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.seq++
	user.ID = fmt.Sprintf("%d", s.seq)
	user.CreatedAt = time.Now()
	s.users[user.ID] = user
	return &user, nil
}

func (s *MemoryStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[id]; !ok {
		return fmt.Errorf("user %q not found", id)
	}
	delete(s.users, id)
	return nil
}
