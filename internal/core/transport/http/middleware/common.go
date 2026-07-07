package core_http_middleware

import (
	"context"
	"net/http"
	"time"

	core_logger "github.com/Pavelkv94/todo_golang/internal/core/logger"
	core_http_response "github.com/Pavelkv94/todo_golang/internal/core/transport/http/response"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	requestIDHeader = "X-Request-ID"
)

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.NewString()
			}
			r.Header.Set(requestIDHeader, requestID)
			w.Header().Set(requestIDHeader, requestID)
			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			log.With(
				zap.String("request_id", requestID),
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()),
			)

			ctx := context.WithValue(r.Context(), "logger", log) // передаем логгер в контекст

			next.ServeHTTP(w, r.WithContext(ctx)) // передаем контекст в следующий обработчик
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

			defer func() {
				if err := recover(); err != nil {
					responseHandler.PanicResponse(err, "panic occurred")
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := core_logger.FromContext(ctx)
			responseWriter := core_http_response.NewHTTPResponseWriter(w)

			before := time.Now().UTC()
			logger.Debug(">>> incoming request", zap.Time("time", before))

			next.ServeHTTP(responseWriter, r)

			logger.Debug(
				">>> outgoing response",
				zap.Duration("latency", time.Since(before)),
				zap.Int("status_code", responseWriter.GetStatusCode()),
			)
		})
	}
}
