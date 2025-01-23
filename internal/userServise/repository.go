package userService

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Migrate() error {
	return r.db.AutoMigrate(&User{})
}

func (r *Repository) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *Repository) GetUserByID(id uint) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *Repository) UpdateUser(id uint, updateData map[string]interface{}) (*User, error) {
	var user User
	err := r.db.Model(&user).Where("id = ?", id).Updates(updateData).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) DeleteUser(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
