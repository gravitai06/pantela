package taskServise

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task   string `json:"task"`    // Название задачи
	IsDone bool   `json:"is_done"` // Статус выполнения задачи
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Task{})
}

func CreateTask(db *gorm.DB, task *Task) error {
	return db.Create(task).Error
}

func GetAllTasks(db *gorm.DB) ([]Task, error) {
	var tasks []Task
	err := db.Find(&tasks).Error
	return tasks, err
}

func GetTaskByID(db *gorm.DB, id uint) (*Task, error) {
	var task Task
	err := db.First(&task, id).Error
	return &task, err
}

func UpdateTask(db *gorm.DB, id uint, updateData map[string]interface{}) (*Task, error) {
	var task Task
	err := db.Model(&task).Where("id = ?", id).Updates(updateData).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func DeleteTask(db *gorm.DB, id uint) error {
	return db.Delete(&Task{}, id).Error
}
