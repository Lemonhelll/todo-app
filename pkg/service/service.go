package service

import (
	myTodo_app "myTodo-app"
	"myTodo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user myTodo_app.User) (int, error)
	GenerateToken(username string, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
