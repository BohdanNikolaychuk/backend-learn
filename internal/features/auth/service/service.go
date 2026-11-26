package service

import (
	"context"
	"errors"
	"time"

	"github.com/BohdanNikolaychuk/backend-learn/internal/features/auth/dto"
	authModel "github.com/BohdanNikolaychuk/backend-learn/internal/features/auth/model"
	userModel "github.com/BohdanNikolaychuk/backend-learn/internal/features/user/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*userModel.User, error)
}

type AuthRepository interface {
	Create(ctx context.Context, userID string, token string, expiresAt time.Time) error
	GetByToken(ctx context.Context, token string) (*authModel.RefreshToken, error)
	DeleteByToken(ctx context.Context, token string) error
}

type AuthService struct {
	userRepo  UserRepository
	authRepo  AuthRepository
	jwtSecret string
}

func NewAuthService(userRepo UserRepository, authRepo AuthRepository, jwtSecret string) *AuthService {
	return &AuthService{userRepo: userRepo, authRepo: authRepo, jwtSecret: jwtSecret}
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, err
	}

	refreshToken := uuid.New().String()
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	if err := s.authRepo.Create(ctx, user.ID, refreshToken, expiresAt); err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}

func (s *AuthService) Refresh(ctx context.Context, req dto.RefreshRequest) (*dto.LoginResponse, error) {
	tokenData, err := s.authRepo.GetByToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	if time.Now().After(tokenData.ExpiresAt) {
		return nil, errors.New("refresh token expired")
	}
	claims := jwt.MapClaims{
		"user_id": tokenData.UserID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := jwtToken.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: req.RefreshToken,
	}, nil

}

func (s *AuthService) Logout(ctx context.Context, req dto.RefreshRequest) error {
	err := s.authRepo.DeleteByToken(ctx, req.RefreshToken)
	return err
}
