package taskServise

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID     uint   `json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID string `json:"user_id"`
}

//type Task struct {
//	gorm.Model
//	Task   string `json:"task"`
//	IsDone bool   `json:"is_done"`
//	UserID uint   `json:"user_id"`
//}
