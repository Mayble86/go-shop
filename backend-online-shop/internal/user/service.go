package user

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	if repo == nil {
		log.Fatal("repo is nil!")
	}
	return &Service{repo: repo}
}

func (s *Service) Register(email, password string, roleID int) (*User, error) {
	existing, _ := s.repo.GetByEmail(email)
	if existing != nil {
		return nil, errors.New("user already exists")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &User{
		Email:    email,
		Password: string(hashed),
		RoleID:   roleID,
	}

	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
