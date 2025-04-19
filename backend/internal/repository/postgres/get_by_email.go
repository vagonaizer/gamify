package postgres

import (
	"context"
	"fmt"
	"gamify/internal/models"

	"github.com/jackc/pgx/v5"
)

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, nickname, email, password_hash, created_at, updated_at, deleted_at
		FROM users
		WHERE email = $1 AND deleted_at IS NULL;
	`

	row := r.pool.QueryRow(ctx, query, email)

	var user models.User
	err := row.Scan(
		&user.ID,
		&user.Nickname,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows { // No rows found
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		// Handle other errors
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return &user, nil
}
