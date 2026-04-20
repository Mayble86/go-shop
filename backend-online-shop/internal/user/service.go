package user

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func (s *Service) Login(email, password string) (*User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func NewService(repo *Repository) *Service {
	if repo == nil {
		log.Fatal("repo is nil!")
	}
	return &Service{repo: repo}
}

func (s *Service) Register(email, password string, roleID int) (*User, error) {
	existing, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("user already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Email:    email,
		Password: string(hashed),
		RoleID:   roleID,
	}

	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetAll() ([]User, error) {
	return s.repo.GetAll()
}
