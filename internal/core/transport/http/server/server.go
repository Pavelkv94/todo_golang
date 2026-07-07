package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/Pavelkv94/todo_golang/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux    *http.ServeMux
	config HTTPServerConfig
	logger *core_logger.Logger
}

func NewServer(config HTTPServerConfig, logger *core_logger.Logger) *HTTPServer {
	return &HTTPServer{mux: http.NewServeMux(), config: config, logger: logger}
}

func (s *HTTPServer) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:    s.config.Addr,
		Handler: s.mux,
	}
	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		s.logger.Info("starting HTTP server", zap.String("addr", s.config.Addr))

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		return fmt.Errorf("HTTP server error: %w", err)
	case <-ctx.Done():
		s.logger.Warn("HTTP server shutdown", zap.String("addr", s.config.Addr))
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), s.config.ShutdownTimeout)
		defer shutdownCancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("HTTP server shutdown error: %w", err)
		}

		s.logger.Warn("HTTP server shutdown completed", zap.String("addr", s.config.Addr))
	}

	return nil
}
