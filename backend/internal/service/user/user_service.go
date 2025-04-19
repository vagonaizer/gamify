package user

import repository "gamify/internal/repository/interfaces"

type UserService struct {
	repo repository.UserRepository
}

func New(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
