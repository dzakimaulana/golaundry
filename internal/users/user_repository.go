package users

import (
	"context"

	"github.com/dzakimaulana/golaundry/pkg/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	db.AutoMigrate(&models.User{})
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, "username = ?", username).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

func (r *repository) GetAllUser(ctx context.Context) (*[]models.User, error) {
	var user []models.User
	if err := r.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	if err := r.DB.Model(&user).Preload("Transactions").First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) ResetPassword(ctx context.Context, id string, password string) error {
	var user models.User
	if err := r.DB.Model(&user).Where("id = ?", id).Update("password", password).Error; err != nil {
		return err
	}
	return nil
}
