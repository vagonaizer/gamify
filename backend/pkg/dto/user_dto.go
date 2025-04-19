package dto

/* login */
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`    // User's email address
	Password string `json:"password" validate:"required,min=8"` // User's password
}

type LoginResponse struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

/* register */
type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`           // User's email address
	Nickname string `json:"nickname" validate:"required,min=3,max=32"` // User's nickname
	Password string `json:"password" validate:"required,min=8,max=64"` // User's password
}

type RegisterResponse struct {
	ID       uint64 `json:"id"`       // Unique identifier for the user
	Email    string `json:"email"`    // User's email address
	Nickname string `json:"nickname"` // User's nickname
}
