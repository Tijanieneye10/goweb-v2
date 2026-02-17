package models

import (
	"errors"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail    = errors.New("email already exists")
	ErrInvalidCredential = errors.New("invalid email or password")
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
}

type UserStore struct {
	mu    sync.RWMutex
	users map[string]*User
	nextID int
}

func NewUserStore() *UserStore {
	return &UserStore{
		users:  make(map[string]*User),
		nextID: 1,
	}
}

func (s *UserStore) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[email]; exists {
		return ErrDuplicateEmail
	}

	s.users[email] = &User{
		ID:             s.nextID,
		Name:           name,
		Email:          email,
		HashedPassword: hashedPassword,
	}
	s.nextID++
	return nil
}

func (s *UserStore) Authenticate(email, password string) (int, error) {
	s.mu.RLock()
	user, exists := s.users[email]
	s.mu.RUnlock()

	if !exists {
		return 0, ErrInvalidCredential
	}

	err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
	if err != nil {
		return 0, ErrInvalidCredential
	}

	return user.ID, nil
}
