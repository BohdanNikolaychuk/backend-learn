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
	ChangePassword(ctx context.Context, id string, newPasswordHash string) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, req dto.CreateUserRequest) (*model.User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	req.Password = string(hashedPass)

	return s.repo.Create(ctx, req)
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

func (s *UserService) Update(ctx context.Context, id string, req dto.UpdateUserRequest) (*model.User, error) {
	if id == "" {
		return nil, errors.New("invalid id")
	}

	if req.FullName == "" {
		return nil, errors.New("nothing to update")
	}

	return s.repo.Update(ctx, id, req)
}

func (s *UserService) ChangePassword(ctx context.Context, id string, req dto.ChangePasswordUserRequest) error {
	if id == "" {
		return errors.New("invalid id")
	}

	user, err := s.repo.GetByID(ctx, id)

	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return errors.New("invalid old password")
	}

	newPassHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	return s.repo.ChangePassword(ctx, id, string(newPassHash))
}
