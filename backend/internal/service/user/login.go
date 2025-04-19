package user

import (
	"context"
	"gamify/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) Login(ctx context.Context, email, password string) (*models.User, error) {
	// Geting user by email
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidCredentials
	}

	// comparing hash password with the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
