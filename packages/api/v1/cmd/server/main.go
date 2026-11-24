package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	v1 "github.com/rshdhere/incident/packages/api/v1"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	if err := run(logger); err != nil {
		logger.Error("server failed", "error", err)
		os.Exit(1)
	}
}

func run(logger *slog.Logger) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	handler, err := v1.NewHandlerWithLogger(os.DirFS(cfg.root), logger)
	if err != nil {
		return fmt.Errorf("load catalog: %w", err)
	}
	server := &http.Server{
		Handler:           handler,
		ErrorLog:          slog.NewLogLogger(logger.Handler(), slog.LevelError),
		ReadHeaderTimeout: min(5*time.Second, cfg.readTimeout),
		ReadTimeout:       cfg.readTimeout,
		WriteTimeout:      cfg.writeTimeout,
		IdleTimeout:       cfg.idleTimeout,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.port))
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}
	logger.Info("serving catalog", "address", listener.Addr().String())
	served := make(chan error, 1)
	go func() { served <- server.Serve(listener) }()

	select {
	case err := <-served:
		return err
	case <-ctx.Done():
		stop()
	}

	logger.Info("draining HTTP requests", "timeout", cfg.shutdownTimeout.String())
	drain, cancel := context.WithTimeout(context.Background(), cfg.shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(drain); err != nil {
		if closeErr := server.Close(); closeErr != nil {
			return errors.Join(fmt.Errorf("drain HTTP requests: %w", err), fmt.Errorf("force close: %w", closeErr))
		}
		return fmt.Errorf("drain HTTP requests: %w", err)
	}
	if err := <-served; !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	logger.Info("server stopped")
	return nil
}
