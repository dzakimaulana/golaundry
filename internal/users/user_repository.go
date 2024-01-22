package users

import (
	"context"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	if err := r.DB.First(&user, "username = ?", username).Error; err != nil {
		return &user, err
	}
	return &user, nil
}
