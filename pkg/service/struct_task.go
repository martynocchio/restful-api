package service

import (
	restful_api "restful-api"
	"restful-api/pkg/repository"
)

type StructTaskService struct {
	repo     repository.StructTask
	listRepo repository.StructList
}

func NewStructTaskService(repo repository.StructTask, listRepo repository.StructList) *StructTaskService {
	return &StructTaskService{
		repo:     repo,
		listRepo: listRepo,
	}
}

func (s *StructTaskService) Create(userId, listId int, task restful_api.StructTask) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, task)
}