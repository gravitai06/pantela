package userService

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *Service) CreateUser(user *User) error {
	return s.repo.CreateUser(user)
}

func (s *Service) DeleteUser(id uint) error {
	return s.repo.DeleteUser(int(id))
}

func (s *Service) UpdateUser(id uint, updateData map[string]interface{}) (*User, error) {
	return s.repo.UpdateUser(int(id), updateData)
}
