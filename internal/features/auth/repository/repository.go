package repository

import (
	"context"
	"errors"
	"time"

	"github.com/BohdanNikolaychuk/backend-learn/internal/features/auth/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Create(ctx context.Context, userID string, token string, expiresAt time.Time) error {
	query := `INSERT INTO learn_backend.refresh_tokens(user_id,token,expires_at)
			VALUES($1,$2,$3)`

	_, err := r.db.Exec(ctx, query, userID, token, expiresAt)
	return err
}

func (r *AuthRepository) GetByToken(ctx context.Context, token string) (*model.RefreshToken, error) {
	query := `SELECT id, user_id, token, expires_at, created_at 
	FROM learn_backend.refresh_tokens 
	WHERE token = $1`

	refreshToken := &model.RefreshToken{}

	err := r.db.QueryRow(ctx, query, token).Scan(
		&refreshToken.ID,
		&refreshToken.UserID,
		&refreshToken.Token,
		&refreshToken.ExpiresAt,
		&refreshToken.CreatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("token not found")
	}

	return refreshToken, err

}

func (r *AuthRepository) DeleteByToken(ctx context.Context, token string) error {
	query := `DELETE FROM learn_backend.refresh_tokens 
	WHERE token = $1`

	tag, err := r.db.Exec(ctx, query, token)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return errors.New("token not found")
	}
	return nil
}
