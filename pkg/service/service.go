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

type StructList interface {
	Create(userId int, list restful_api.StructList) (int, error)
	GetAll(userId int) ([]restful_api.StructList, error)
	GetById(userId, listId int) (restful_api.StructList, error)
	Update(userId, listId int, input restful_api.UpdateListInput) error
	Delete(userId, listId int) error
}

type StructTask interface {
	Create(userId, listId int, task restful_api.StructTask) (int, error)
}

type StructSubtask interface {
}

type Service struct {
	Authorization
	StructList
	StructTask
	StructSubtask
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		StructList:    NewStructListService(repos.StructList),
	}
}
