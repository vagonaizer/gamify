package routes

import (
	"gamify/internal/server/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRoutes(h *handler.UserHandler) *chi.Mux {
	r := chi.NewRouter()

	// connecting middleware features
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/register", h.Register)
		r.Post("/login", h.Login)
	})

	return r
}
