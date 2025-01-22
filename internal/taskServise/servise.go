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

func (s *Service) CreateTask(task *Task) error {
	return s.repo.CreateTask(task)
}

func (s *Service) DeleteTask(id uint) error {
	return s.repo.DeleteTask(id)
}

func (s *Service) UpdateTask(id uint, updateData map[string]interface{}) (*Task, error) {
	return s.repo.UpdateTask(id, updateData)
}

//package taskServise
//
//type Service struct {
//	repo *Repository
//}
//
//func NewService(repo *Repository) *Service {
//	return &Service{repo: repo}
//}
//
//func (s *Service) GetAllTasks() ([]Task, error) {
//	return s.repo.GetAll()
//}
//
//func (s *Service) GetTaskByID(id uint) (*Task, error) {
//	return s.repo.GetByID(id)
//}
//
//func (s *Service) CreateTask(task *Task) error {
//	return s.repo.Create(task)
//}
//
//func (s *Service) UpdateTask(id uint, updateData map[string]interface{}) (*Task, error) {
//	task, err := s.repo.GetByID(id)
//	if err != nil {
//		return nil, err
//	}
//
//	if err := s.repo.Update(task, updateData); err != nil {
//		return nil, err
//	}
//
//	return task, nil
//}
//
//func (s *Service) DeleteTask(id uint) error {
//	task, err := s.repo.GetByID(id)
//	if err != nil {
//		return err
//	}
//
//	return s.repo.Delete(task)
//}
