package taskServise

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll() ([]Task, error) {
	var tasks []Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) GetByID(id uint) (*Task, error) {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *Repository) Create(task *Task) error {
	return r.db.Create(task).Error
}

func (r *Repository) Update(task *Task, updateData map[string]interface{}) error {
	return r.db.Model(task).Updates(updateData).Error
}

func (r *Repository) Delete(task *Task) error {
	return r.db.Delete(task).Error
}
