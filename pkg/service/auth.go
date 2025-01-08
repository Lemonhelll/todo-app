package service

import (
	"crypto/sha1"
	"fmt"
	myTodo_app "myTodo-app"
	"myTodo-app/pkg/repository"
)

const salt = "dgfjhdhsgar1312fas"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user myTodo_app.User) (int, error) {
	user.Password = generatePassHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePassHash(password string) string {

	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
