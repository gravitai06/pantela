package taskServise

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task   string `json:"task"`    // Название задачи
	IsDone bool   `json:"is_done"` // Статус выполнения задачи
}
