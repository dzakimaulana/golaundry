package users

import (
	"context"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `jsin:"username"`
	Password string    `json:"password"`
}

type CreateUserReq struct {
	Username string `jsin:"username"`
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
	accessToken string
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
}

type UserService interface {
	CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error)
	Login(ctx context.Context, req *LoginReq) (*LoginRes, error)
}
