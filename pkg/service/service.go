package service

import "restful-api/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
