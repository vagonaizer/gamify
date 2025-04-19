package handler

import (
	"encoding/json"
	"errors"
	"gamify/internal/service/user"
	"gamify/pkg/dto"
	"log"
	"net/http"
)

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Initializing req object of type dto.RegisterRequest
	w.Header().Set("Content-Type", "application/json")

	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	userRegister, err := h.userService.Register(r.Context(), req.Email, req.Nickname, req.Password)
	if err != nil {
		if errors.Is(err, user.ErrAlreadyExists) {
			http.Error(w, "user already exists", http.StatusConflict)
			return
		}
		log.Printf("Register error: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	resp := dto.RegisterResponse{
		ID:       userRegister.ID,
		Email:    userRegister.Email,
		Nickname: userRegister.Nickname,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Failed to encode register response: %v", err)
	}
}
