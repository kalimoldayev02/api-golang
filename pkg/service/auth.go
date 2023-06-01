package service

import (
	"github.com/kalimoldayev02/api-golang/models"
	"github.com/kalimoldayev02/api-golang/pkg/repository"
)

type AuthService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	return s.repo.CreateUser(user)
}
