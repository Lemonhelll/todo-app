package service

import (
	myTodo_app "myTodo-app"
	"myTodo-app/pkg/repository"
)

type AuthService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user myTodo_app.User, err error) {
	return s.repo.CreateUser(user)
}
