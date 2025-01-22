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

//package taskServise
//
//import (
//	"gorm.io/gorm"
//)
//
//type Repository struct {
//	db *gorm.DB
//}
//
//func NewRepository(db *gorm.DB) *Repository {
//	return &Repository{db: db}
//}
//
//func (r *Repository) GetAll() ([]Task, error) {
//	var tasks []Task
//	if err := r.db.Find(&tasks).Error; err != nil {
//		return nil, err
//	}
//	return tasks, nil
//}
//
//func (r *Repository) GetByID(id uint) (*Task, error) {
//	var task Task
//	if err := r.db.First(&task, id).Error; err != nil {
//		return nil, err
//	}
//	return &task, nil
//}
//
//func (r *Repository) Create(task *Task) error {
//	return r.db.Create(task).Error
//}
//
//func (r *Repository) Update(task *Task, updateData map[string]interface{}) error {
//	return r.db.Model(task).Updates(updateData).Error
//}
//
//func (r *Repository) Delete(task *Task) error {
//	return r.db.Delete(task).Error
//}
