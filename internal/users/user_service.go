package users

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dzakimaulana/golaundry/pkg/models"
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

func (s *service) CreateUser(ctx context.Context, req *CreateUserReq) (*models.UserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &models.User{
		ID:       uuid.New(),
		Username: req.Username,
		Password: hashedPassword,
		Role:     "employee",
	}

	r, err := s.UserRepository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &models.UserRes{
		ID:       r.ID,
		Username: r.Username,
	}

	return res, nil
}

func (s *service) Login(ctx context.Context, req *LoginReq) (*models.LoginRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	u, err := s.UserRepository.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return &models.LoginRes{}, err
	}

	err = utils.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &models.LoginRes{}, err
	}

	// Generate JWT
	claims := jwt.MapClaims{}
	claims["sub"] = u.ID
	claims["role"] = u.Role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	token, err := utils.GenerateToken(&claims)
	if err != nil {
		return &models.LoginRes{}, err
	}

	return &models.LoginRes{
		AccessToken: token,
		ID:          u.ID,
		Username:    u.Username,
	}, nil
}

func (s *service) GetAllUser(ctx context.Context) (*[]models.UserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	r, err := s.UserRepository.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}

	var resUser []models.UserRes
	for _, dataUser := range *r {
		resUser = append(resUser, models.UserRes{
			ID:       dataUser.ID,
			Username: dataUser.Username,
		})
	}
	return &resUser, nil
}

func (s *service) GetUserByID(ctx context.Context, id string) (*models.UserResByID, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	r, err := s.UserRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	var resTs []models.TransactionsResUs
	for _, dataTs := range *r.Transactions {
		resTs = append(resTs, models.TransactionsResUs{
			ID:         dataTs.ID,
			CustomerID: dataTs.CustomerID,
			TimeIn:     dataTs.TimeIn,
			TimeOut:    dataTs.TimeOut,
			Total:      dataTs.Total,
		})
	}

	res := &models.UserResByID{
		ID:           r.ID,
		Username:     r.Username,
		Transactions: resTs,
	}
	return res, nil
}

func (s *service) ResetPassword(ctx context.Context, id string, req *ResetPasswordReq) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	r, err := s.UserRepository.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	err = utils.CheckPassword(req.Password, r.Password)
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	err = s.UserRepository.ResetPassword(ctx, r.ID.String(), hashedPassword)
	if err != nil {
		return err
	}
	return nil
}
