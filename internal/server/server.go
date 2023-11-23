package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"lamoda_task/internal/config"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func InitHttpRouter() *chi.Mux {
	httpRouter := chi.NewRouter()
	httpRouter.Use(middleware.Logger)
	httpRouter.Use(middleware.Recoverer)
	httpRouter.Use(middleware.Timeout(60 * time.Second))
	httpRouter.Mount("/debug", middleware.Profiler())

	httpRouter.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {})

	return httpRouter
}

func RunHttpServer(ctx context.Context, router *chi.Mux, cfg *config.HttpServer) {
	httpServerAddr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	httpServer := http.Server{ //nolint:exhaustruct
		Addr:    httpServerAddr,
		Handler: router,
	}

	log.Printf("Start listening to http://%s/", httpServerAddr)

	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Can't start server: %v", err)
	}
}
