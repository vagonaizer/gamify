package user

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrAlreadyExists      = errors.New("user already exists")
)
