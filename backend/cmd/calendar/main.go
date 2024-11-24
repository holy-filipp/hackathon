package main

import (
	"calendar-backend/internal/config"
	"calendar-backend/internal/http-server/handlers/events/bymonth"
	"calendar-backend/internal/http-server/handlers/parser/parse"
	mwLogger "calendar-backend/internal/http-server/middleware/logger"
	"calendar-backend/internal/lib/logger/handlers/slogpretty"
	"calendar-backend/internal/lib/logger/sl"
	"calendar-backend/internal/storage/postgresql"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() { // Точка входа
	cfg := config.MustLoad() // Загружаем конфиг

	log := setupLogger(cfg.Env) // Инициализируем логгер

	log.Info("starting calendar api")

	storage, err := postgresql.New(cfg.Postgresql.Host, cfg.Postgresql.Port, cfg.Postgresql.Username, cfg.Postgresql.Database, cfg.Postgresql.Password, cfg.Postgresql.SSLMode) // Инициализация храналища
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	router := chi.NewRouter() // Создаём роутер

	router.Use(middleware.RequestID) // Подключаем мидлвари
	router.Use(middleware.RealIP)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/parser/parse", parse.New(log, *cfg, storage))
	router.Post("/events/bymonth", bymonth.New(log, storage))

	log.Info("starting server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stopped")
}

func setupLogger(env string) *slog.Logger { // Инициализация логгера
	var log *slog.Logger

	switch env { // Логгер будет поддерживать три режима окружающей среды
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

func setupPrettySlog() *slog.Logger { // Логгер для локальной разработки
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
