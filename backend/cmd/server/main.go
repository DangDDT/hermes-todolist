// @title           Hermes TodoList API
// @version         1.0
// @description     REST API for Hermes TodoList — a production-grade task management application.
// @contact.name    DangDDT
// @host            localhost:8080
// @BasePath        /api/v1
// @schemes         http https
// @securityDefinitions.apikey BearerAuth
// @in              header
// @name            Authorization
// @description     JWT token. Use "Bearer <token>" or cookie "access_token".
package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httprate "github.com/go-chi/httprate"

	"github.com/DangDDT/hermes-todolist/backend/internal/config"
	"github.com/DangDDT/hermes-todolist/backend/internal/feature/auth_login"
	"github.com/DangDDT/hermes-todolist/backend/internal/feature/auth_register"
	"github.com/DangDDT/hermes-todolist/backend/internal/feature/tag_list"
	"github.com/DangDDT/hermes-todolist/backend/internal/feature/task_comment"
	"github.com/DangDDT/hermes-todolist/backend/internal/feature/task_create"
	"github.com/DangDDT/hermes-todolist/backend/internal/feature/task_delete"
	"github.com/DangDDT/hermes-todolist/backend/internal/feature/task_get"
	"github.com/DangDDT/hermes-todolist/backend/internal/feature/task_list"
	"github.com/DangDDT/hermes-todolist/backend/internal/feature/task_update"
	mw "github.com/DangDDT/hermes-todolist/backend/internal/infra/middleware"
	"github.com/DangDDT/hermes-todolist/backend/internal/infra/postgres"
	"github.com/DangDDT/hermes-todolist/backend/internal/infra/repository"
	"github.com/DangDDT/hermes-todolist/backend/internal/shared/response"
)

func main() {
	// Load configuration.
	cfg := config.MustLoad()

	// Setup structured logger.
	logLevel := slog.LevelInfo
	if cfg.Log.Level == "debug" {
		logLevel = slog.LevelDebug
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	slog.SetDefault(logger)

	// Create database connection pool.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := postgres.NewPool(ctx, cfg.Database.URL)
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer pool.Close()
	logger.Info("connected to database")

	// Create repositories.
	taskRepo := repository.NewTaskRepo(pool)
	userRepo := repository.NewUserRepo(pool)
	commentRepo := repository.NewCommentRepo(pool)

	// Create usecases.
	registerUC := auth_register.NewUsecase(userRepo)
	loginUC := auth_login.NewUsecase(userRepo, cfg.JWT.Secret, cfg.JWT.ExpiryHours)
	taskCreateUC := task_create.NewUsecase(taskRepo)
	taskListUC := task_list.NewUsecase(taskRepo)
	taskGetUC := task_get.NewUsecase(taskRepo)
	taskUpdateUC := task_update.NewUsecase(taskRepo)
	taskDeleteUC := task_delete.NewUsecase(taskRepo)
	taskCommentUC := task_comment.NewUsecase(taskRepo, commentRepo)

	// Create router.
	r := chi.NewRouter()

	// Global middleware.
	r.Use(middleware.RequestID)
	r.Use(mw.Tracing)
	r.Use(mw.Logger(logger))
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{cfg.CORS.Origin},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Use(mw.JWTAuth(cfg.JWT.Secret))
	r.Use(httprate.LimitByIP(100, 1*time.Minute))

	// Health check.
	r.Get("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, map[string]string{
			"status":  "ok",
			"version": "1.0.0",
		})
	})

	// API v1 routes.
	r.Route("/api/v1", func(r chi.Router){
		r.Route("/auth", func(r chi.Router){
			r.Mount("/register", auth_register.Routes(registerUC))
			r.Mount("/login", auth_login.Routes(loginUC))
		})

		r.Route("/tasks", func(r chi.Router){
			taskCreateH := task_create.NewHandler(taskCreateUC)
			taskListH := task_list.NewHandler(taskListUC)
			taskGetH := task_get.NewHandler(taskGetUC)
			taskUpdateH := task_update.NewHandler(taskUpdateUC)
			taskDeleteH := task_delete.NewHandler(taskDeleteUC)

			r.Post("/", taskCreateH.Create)
			r.Get("/", taskListH.List)
			r.Get("/{id}", taskGetH.Get)
			r.Put("/{id}", taskUpdateH.Update)
			r.Delete("/{id}", taskDeleteH.Delete)
			r.Mount("/{id}/comments", task_comment.Routes(taskCommentUC))
		})

		r.Mount("/tags", tag_list.Routes())
	})

	// Swagger UI.
	r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	})
	r.Get("/swagger/*", func(w http.ResponseWriter, r *http.Request) {
		// Serve index.html for /swagger/ root, otherwise serve static file
		path := strings.TrimPrefix(r.URL.Path, "/swagger/")
		if path == "" || path == "/" {
			http.ServeFile(w, r, "./docs/swagger/index.html")
			return
		}
		http.ServeFile(w, r, "./docs/swagger/"+path)
	})

	// Start server.
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		logger.Info("shutting down...")
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer shutdownCancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			logger.Error("forced shutdown", "error", err)
		}
	}()

	logger.Info("starting server", "addr", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("server failed", "error", err)
		os.Exit(1)
	}
	logger.Info("server stopped")
}
