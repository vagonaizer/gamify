package repository

import (
	"context"
	"gamify/internal/models"
)

type UserRepository interface {
	// Create a new user.
	Create(ctx context.Context, user *models.User) (uint64, error)
	// GetByID retrieves a user by their ID.
	GetByID(ctx context.Context, id uint64) (*models.User, error)
	// GetByEmail retrieves a user by their email.
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	// GetByNickname retrieves a user by their nickname.
	GetByNickname(ctx context.Context, nickname string) (*models.User, error)
	// List retrieves a list of users with pagination.
	List(ctx context.Context, limit, offset int) ([]*models.User, error)
	// Delete removes a user by their ID.
	Delete(ctx context.Context, id uint64) error
	// Update modifies an existing user information.
	Update(ctx context.Context, user *models.User) error
}
