package userService

import (
	"crypto/rand"
	"encoding/hex"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func generateID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

func (s *Service) CreateUser(user *User) error {
	user.ID = generateID()
	return s.repo.CreateUser(user)
}

func (s *Service) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}

func (s *Service) UpdateUser(id string, updateData map[string]interface{}) (*User, error) {
	return s.repo.UpdateUser(id, updateData)
}
