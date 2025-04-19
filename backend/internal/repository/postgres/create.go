package postgres

import (
	"context"
	"fmt"
	"gamify/internal/models"
	"time"
)

// Create inserts a new user into the database and returns the user's ID.
func (r *UserRepository) Create(ctx context.Context, user *models.User) (uint64, error) {
	// Implementation for creating a user in the database
	query := `
		INSERT INTO users (nickname, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at;
	`

	var id uint64
	var createdAt, updatedAt time.Time

	// Execute the query and scan the returned values into the variables
	// Note: TODO making sure to handle the error properly
	err := r.pool.QueryRow(ctx, query, user.Nickname, user.Email, user.PasswordHash).Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	user.ID = id
	return id, nil
}
