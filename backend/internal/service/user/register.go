package user

import (
	"context"
	"fmt"
	"gamify/internal/models"

	"golang.org/x/crypto/bcrypt"
)

// Register func is responsible for registering a new user.
func (s *UserService) Register(ctx context.Context, email, nickname, password string) (*models.User, error) {
	// checking if email/nickname already exists
	if existingUser, _ := s.repo.GetByEmail(ctx, email); existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", email)
	}
	if existingUser, _ := s.repo.GetByNickname(ctx, nickname); existingUser != nil {
		return nil, fmt.Errorf("user with nickname %s already exists", nickname)
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create a new user
	user := &models.User{
		Email:        email,
		Nickname:     nickname,
		PasswordHash: string(hashedPassword),
	}

	// Saving up
	id, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	user.ID = id

	return user, nil
}
