package user

import (
	"context"
	"fmt"

	"gamify/internal/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) Login(ctx context.Context, email, password string) (*models.User, error) {
	// Geting user by email
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		log.Printf("Error getting user by email: %v", err) // Логируем ошибку
		return nil, err                                    // Возвращаем ошибку, если она возникла при получении пользователя
	}
	if user == nil {
		log.Printf("User with email %s not found", email) // Логируем, если пользователь не найден
		return nil, ErrInvalidCredentials                 // Возвращаем ошибку, если пользователь не найден
	}

	log.Printf("Stored hash: %s\n", user.PasswordHash)

	log.Printf("Plain password: %s\n", password)

	// comparing hash password with the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		fmt.Println("Password compare failed:", err)
		return nil, ErrInvalidCredentials
	}
	fmt.Println("Login successful!")
	return user, nil
}
