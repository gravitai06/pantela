package taskServise

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *Service) GetTasksForUser(userID uint) ([]Task, error) {
	return s.repo.GetTasksForUser(userID)
}

func (s *Service) CreateTask(task *Task) error {
	return s.repo.CreateTask(task)
}

func (s *Service) DeleteTask(id uint) error {
	return s.repo.DeleteTask(id)
}

func (s *Service) UpdateTask(id uint, updateData map[string]interface{}) (*Task, error) {
	return s.repo.UpdateTask(id, updateData)
}
