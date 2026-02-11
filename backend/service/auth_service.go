package service

import (
	"errors"
	"time"

	"automobile-parts-backend/config"
	"automobile-parts-backend/model"
	"automobile-parts-backend/utils"
)

type AuthService struct {
	cfg config.Config
}

func NewAuthService(cfg config.Config) *AuthService {
	return &AuthService{cfg: cfg}
}

func (s *AuthService) Login(req model.LoginRequest) (model.AuthResponse, error) {
	if req.Username == "" || req.Password == "" {
		return model.AuthResponse{}, errors.New("invalid credentials")
	}
	userID := req.Username
	role := "user"
	token, err := utils.GenerateToken(userID, role, s.cfg.JWTSecret, s.cfg.JWTExpireHours)
	if err != nil {
		return model.AuthResponse{}, err
	}
	return model.AuthResponse{
		Token:  token,
		UserID: userID,
		Role:   role,
	}, nil
}

func (s *AuthService) Register(req model.RegisterRequest) (model.AuthResponse, error) {
	if req.Username == "" || req.Password == "" || req.Role == "" || req.Org == "" {
		return model.AuthResponse{}, errors.New("invalid registration data")
	}
	userID := req.Username
	role := req.Role
	token, err := utils.GenerateToken(userID, role, s.cfg.JWTSecret, s.cfg.JWTExpireHours)
	if err != nil {
		return model.AuthResponse{}, err
	}
	return model.AuthResponse{
		Token:  token,
		UserID: userID,
		Role:   role,
	}, nil
}

func (s *AuthService) Now() string {
	return time.Now().Format(time.RFC3339)
}
