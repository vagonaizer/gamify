package handler

import (
	"encoding/json"
	"gamify/internal/service/user"
	"gamify/pkg/dto"
	"net/http"

	"github.com/go-playground/validator"
)

type UserHandler struct {
	userService user.UserService
	validator   *validator.Validate
}

func NewUserHandler(userService *user.UserService) *UserHandler {
	return &UserHandler{
		userService: *userService,
		validator:   validator.New(),
	}
}

// some methods to make life easier

func (h *UserHandler) ResponseWithError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(dto.ErrorResponse{Error: message})
}
