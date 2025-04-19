package handler

import (
	"encoding/json"
	"errors"
	"gamify/internal/service/user"
	"gamify/pkg/dto"
	"net/http"
)

// Login handler
// dto structures defined in pkg/dto/user_dto.go
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse the request body
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate the request
	if err := h.validator.Struct(req); err != nil {
		h.ResponseWithError(w, http.StatusBadRequest, "validation failed: "+err.Error())
		return
	}

	// Call business logic
	userLogin, err := h.userService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, user.ErrInvalidCredentials) {
			h.ResponseWithError(w, http.StatusUnauthorized, "invalid email or password")
			return
		}
		h.ResponseWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	resp := dto.LoginResponse{
		ID:       userLogin.ID,
		Email:    userLogin.Email,
		Nickname: userLogin.Nickname,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
