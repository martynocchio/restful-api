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

func (s *StructTaskService) GetAll(userId, listId int) ([]restful_api.StructTask, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *StructTaskService) GetById(userId, taskId int) (restful_api.StructTask, error) {
	return s.repo.GetById(userId, taskId)
}

func (s *StructTaskService) Delete(userId, taskId int) error {
	return s.repo.Delete(userId, taskId)
}

func (s *StructTaskService) Update(userId, taskId int, input restful_api.UpdateTaskInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, taskId, input)
}
