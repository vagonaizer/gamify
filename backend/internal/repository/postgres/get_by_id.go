package postgres

import (
	"context"
	"fmt"
	"gamify/internal/models"

	"github.com/jackc/pgx/v5"
)

// GetByID retrieves a user by their ID from the database.
func (r *UserRepository) GetByID(ctx context.Context, id uint64) (*models.User, error) {
	query := `
		SELECT id, nickname, email, password_hash, created_at, updated_at, deleted_at
		FROM users
		WHERE id = $1 AND deleted_at IS NULL;
	`

	// Execute the query and scan the returned values into the variables
	row := r.pool.QueryRow(ctx, query, id)

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
			// Return a custom error indicating that the user was not found
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		// Handle other errors
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return &user, nil
}
