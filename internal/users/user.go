package users

import (
	"context"

	"github.com/dzakimaulana/golaundry/pkg/models"
)

type CreateUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResetPasswordReq struct {
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetAllUser(ctx context.Context) (*[]models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	ResetPassword(ctx context.Context, id string, password string) error
}

type UserService interface {
	CreateUser(ctx context.Context, req *CreateUserReq) (*models.UserRes, error)
	Login(ctx context.Context, req *LoginReq) (*models.LoginRes, error)
	GetAllUser(ctx context.Context) (*[]models.UserRes, error)
	GetUserByID(ctx context.Context, id string) (*models.UserResByID, error)
	ResetPassword(ctx context.Context, id string, req *ResetPasswordReq) error
}
