package userService

import (
	"pantela/internal/taskServise"
)

type Service struct {
	repo        *Repository
	taskService *taskServise.Service
}

func NewService(repo *Repository, taskService *taskServise.Service) *Service {
	return &Service{
		repo:        repo,
		taskService: taskService,
	}
}

func (s *Service) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *Service) CreateUser(user *User) error {
	return s.repo.CreateUser(user)
}

func (s *Service) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}

func (s *Service) UpdateUser(id uint, updateData map[string]interface{}) (*User, error) {
	return s.repo.UpdateUser(id, updateData)
}

func (s *Service) GetTasksForUser(userID uint) ([]taskServise.Task, error) {
	return s.taskService.GetTasksForUser(userID)
}
