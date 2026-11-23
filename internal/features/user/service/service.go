package service

import (
	"context"

	"github.com/BohdanNikolaychuk/backend-learn/internal/features/user/dto"
	"github.com/BohdanNikolaychuk/backend-learn/internal/features/user/model"
	"github.com/BohdanNikolaychuk/backend-learn/internal/features/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, dto dto.CreateUserRequest) (*model.User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	dto.Password = string(hashedPass)

	return s.repo.Create(ctx, dto)
}
