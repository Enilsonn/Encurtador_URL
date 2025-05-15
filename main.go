package main

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/Enilsonn/Encurtador_URL.git/api"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code",
			"error", err,
		)
		return
	}
	slog.Info("All systems offline")
}

func run() error {
	db := make(map[string]string)

	handler := api.NewHandler(db)

	s := http.Server{
		Addr:         "localhost:8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	if err := s.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	return nil
}
