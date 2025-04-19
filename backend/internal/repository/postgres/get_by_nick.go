package postgres

import (
	"context"
	"fmt"
	"gamify/internal/models"

	"github.com/jackc/pgx/v5"
)

// GetByNickname retrieves a user by their nickname.
func (r *UserRepository) GetByNickname(ctx context.Context, nickname string) (*models.User, error) {
	query := `
		SELECT id, email, nickname, password_hash, created_at, updated_at, deleted_at
		FROM users
		WHERE nickname = $1 AND deleted_at IS NULL;
	`

	row := r.pool.QueryRow(ctx, query, nickname)

	var user models.User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Nickname,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("user fetch by nickname error: %w", err)
	}

	return &user, nil
}
