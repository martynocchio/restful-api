package service

import (
	restful_api "restful-api"
	"restful-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user restful_api.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type List interface {
}

type Task interface {
}

type Subtask interface {
}

type Service struct {
	Authorization
	List
	Task
	Subtask
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
