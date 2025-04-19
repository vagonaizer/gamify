package app

import (
	"context"
	"fmt"
	"gamify/internal/config"
	"gamify/internal/logging"
	"gamify/internal/repository/postgres"
	"gamify/internal/server/handler"
	"gamify/internal/server/routes"
	"gamify/internal/service/user"
	"net/http"

	"go.uber.org/zap"
)

type App struct {
	Config *config.Config
	Router http.Handler
}

// NewApp инициализирует все компоненты приложения.
func NewApp(cfg *config.Config) (*App, error) {
	// Инициализация логгера
	if err := logging.Init(cfg.Logger.Level); err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}
	logging.Log.Info("Logger initialized", zap.String("level", cfg.Logger.Level))

	// Инициализация контекста
	ctx := context.Background()

	// Инициализация репозитория
	userRepo, err := postgres.New(ctx, cfg)
	if err != nil {
		logging.Log.Fatal("Failed to initialize Postgres", zap.Error(err))
	}

	// Инициализация сервиса пользователей
	userService := user.New(userRepo)

	// Инициализация хендлеров
	h := handler.NewUserHandler(userService)

	// Инициализация маршрутов
	router := routes.InitRoutes(h)
	logging.Log.Info("Router initialized: chi router")

	return &App{
		Config: cfg,
		Router: router,
	}, nil
}

// Run запускает HTTP сервер
func (a *App) Run() error {
	addr := fmt.Sprintf("%s:%s", a.Config.Server.BaseIp, a.Config.Server.Port)
	logging.Log.Info("Starting server", zap.String("address", addr))

	return http.ListenAndServe(addr, a.Router)
}
