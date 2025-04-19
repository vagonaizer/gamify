package postgres

import (
	"context"
	"fmt"
	"gamify/internal/models"
)

// Update modifies an existing user information.
func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	query := `
		UPDATE users
		SET email = $1, nickname = $2, password_hash = $3, updated_at = NOW()
		WHERE id = $4 AND deleted_at IS NULL;
	`

	// Выполнение запроса на обновление информации о пользователе
	_, err := r.pool.Exec(ctx, query, user.Email, user.Nickname, user.PasswordHash, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
