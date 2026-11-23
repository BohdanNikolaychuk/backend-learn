package repository

import (
	"context"

	"github.com/BohdanNikolaychuk/backend-learn/internal/features/user/dto"
	"github.com/BohdanNikolaychuk/backend-learn/internal/features/user/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, req dto.CreateUserRequest) (*model.User, error) {
	query := `INSERT INTO learn_backend.users (email,password,full_name)
			  VALUES ($1,$2,$3)
			  RETURNING  id, version, email, password, full_name, created_at, updated_at`
	user := &model.User{}

	err := r.db.QueryRow(ctx, query, req.Email, req.Password, req.FullName).Scan(&user.ID,
		&user.Version,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.CreatedAt,
		&user.UpdatedAt)

	return user, err
}
