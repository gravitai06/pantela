package taskServise

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Migrate() error {
	return r.db.AutoMigrate(&Task{})
}

func (r *Repository) CreateTask(task *Task) error {
	return r.db.Create(task).Error
}

func (r *Repository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *Repository) GetTasksForUser(userID uint) ([]Task, error) {
	var tasks []Task
	result := r.db.Where("user_id = ?", userID).Find(&tasks)
	return tasks, result.Error
}

func (r *Repository) GetTaskByID(id uint) (*Task, error) {
	var task Task
	err := r.db.First(&task, id).Error
	return &task, err
}

func (r *Repository) UpdateTask(id uint, updateData map[string]interface{}) (*Task, error) {
	var task Task
	err := r.db.Model(&task).Where("id = ?", id).Updates(updateData).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *Repository) DeleteTask(id uint) error {
	return r.db.Delete(&Task{}, id).Error
}
