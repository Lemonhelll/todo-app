package service

import (
	myTodo_app "myTodo-app"
	"myTodo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user myTodo_app.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userId int, list myTodo_app.TodoList) (int, error)
	GetAll(userId int) ([]myTodo_app.TodoList, error)
	GetById(userId, listId int) (myTodo_app.TodoList, error)
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
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
