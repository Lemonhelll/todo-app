package repository

import (
	"github.com/jmoiron/sqlx"
	myTodo_app "myTodo-app"
)

type Authorization interface {
	CreateUser(user myTodo_app.User) (int, error)
	GetUser(username, password string) (myTodo_app.User, error)
}

type TodoList interface {
	Create(userId int, list myTodo_app.TodoList) (int, error)
	GetAll(userId int) ([]myTodo_app.TodoList, error)
	GetById(userId, listId int) (myTodo_app.TodoList, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
