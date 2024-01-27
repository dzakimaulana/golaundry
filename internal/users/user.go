package users

import (
	"context"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"id"`
	Username string    `jsin:"username"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
}

type CreateUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserRes struct {
	ID       uuid.UUID `json:"id"`
	Username string    `jsin:"username"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRes struct {
	ID          uuid.UUID `json:"id"`
	AccessToken string    `json:"access_token"`
	Username    string    `json:"username"`
}

type UserRes struct {
	ID       uuid.UUID `json:"id"`
	Username string    `jsin:"username"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
}

type UserService interface {
	CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error)
	Login(ctx context.Context, req *LoginReq) (*LoginRes, error)
}
