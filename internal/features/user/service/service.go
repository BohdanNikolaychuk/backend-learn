package service

import (
	"context"
	"errors"

	"github.com/BohdanNikolaychuk/backend-learn/internal/features/user/dto"
	"github.com/BohdanNikolaychuk/backend-learn/internal/features/user/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(ctx context.Context, req dto.CreateUserRequest) (*model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
	DeleteByID(ctx context.Context, id string) (*model.User, error)
	Update(ctx context.Context, id string, dto dto.UpdateUserRequest) (*model.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
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

func (s *UserService) GetByID(ctx context.Context, id string) (*model.User, error) {
	if id == "" {
		return nil, errors.New("invalid id")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) DeleteByID(ctx context.Context, id string) (*model.User, error) {
	if id == "" {
		return nil, errors.New("invalid id")
	}
	return s.repo.DeleteByID(ctx, id)
}

func (s *UserService) Update(ctx context.Context, id string, dto dto.UpdateUserRequest) (*model.User, error) {
	if id == "" {
		return nil, errors.New("invalid id")
	}

	if dto.FullName == "" {
		return nil, errors.New("nothing to update")
	}

	return s.repo.Update(ctx, id, dto)
}
