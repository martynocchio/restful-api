package service

import (
	restful_api "restful-api"
	"restful-api/pkg/repository"
)

type StructListService struct {
	repo repository.StructList
}

func NewStructListService(repo repository.StructList) *StructListService {
	return &StructListService{repo: repo}
}

func (s *StructListService) Create(userId int, list restful_api.StructList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *StructListService) GetAll(userId int) ([]restful_api.StructList, error) {
	return s.repo.GetAll(userId)
}
