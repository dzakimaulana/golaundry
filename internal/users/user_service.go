package users

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dzakimaulana/golaundry/pkg/utils"
	"github.com/google/uuid"
)

type service struct {
	UserRepository
	timeout time.Duration
}

func NewService(repository UserRepository) UserService {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Password: hashedPassword,
	}

	r, err := s.UserRepository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:       uuid.New(),
		Username: r.Username,
	}

	return res, nil
}

func (s *service) Login(ctx context.Context, req *LoginReq) (*LoginRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	u, err := s.UserRepository.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return &LoginRes{}, err
	}

	err = utils.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &LoginRes{}, err
	}

	// Generate JWT
	claims := jwt.MapClaims{}
	claims["username"] = u.Username
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	token, err := utils.GenerateToken(&claims)
	if err != nil {
		return &LoginRes{}, err
	}

	return &LoginRes{
		accessToken: token,
		ID:          u.ID,
		Username:    u.Username,
	}, nil
}
