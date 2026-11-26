package repository

import (
	"context"
	"errors"

	"github.com/BohdanNikolaychuk/backend-learn/internal/features/user/dto"
	"github.com/BohdanNikolaychuk/backend-learn/internal/features/user/model"
	"github.com/jackc/pgx/v5"
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

func (r *UserRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	query := `SELECT id, version, email, password, full_name, created_at, updated_at 
			  FROM learn_backend.users
				WHERE id = $1 AND deleted_at IS NULL
				`
	user := &model.User{}

	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID,
		&user.Version,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.CreatedAt,
		&user.UpdatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("user not found")
	}

	return user, err
}

func (r *UserRepository) DeleteByID(ctx context.Context, id string) (*model.User, error) {
	query := `UPDATE learn_backend.users
			  SET deleted_at = NOW()
			  WHERE id = $1 AND deleted_at IS NULL
			  RETURNING id, version, email, password, full_name, created_at, updated_at, deleted_at`
	user := &model.User{}

	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID,
		&user.Version,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("user not found")
	}

	return user, err
}

func (r *UserRepository) Update(ctx context.Context, id string, req dto.UpdateUserRequest) (*model.User, error) {
	query := `UPDATE learn_backend.users
              SET full_name = COALESCE(NULLIF($1, ''), full_name),
                  updated_at = NOW()
              WHERE id = $2 AND deleted_at IS NULL
              RETURNING id, version, email, password, full_name, created_at, updated_at`
	user := &model.User{}

	err := r.db.QueryRow(ctx, query, req.FullName, id).Scan(
		&user.ID,
		&user.Version,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("user not found")
	}

	return user, err
}

func (r *UserRepository) ChangePassword(ctx context.Context, id string, newPasswordHash string) error {
	query := `UPDATE learn_backend.users
              SET password = $1, updated_at = NOW()
              WHERE id = $2 AND deleted_at IS NULL`

	tag, err := r.db.Exec(ctx, query, newPasswordHash, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT id, version, email, password, full_name, created_at, updated_at 
			  FROM learn_backend.users
				WHERE email = $1 AND deleted_at IS NULL
				`

	user := &model.User{}

	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Version,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("user not found")
	}

	return user, err
}
