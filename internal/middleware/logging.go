package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type wrappedResonseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedResonseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func logging(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapped := &wrappedResonseWriter{
			ResponseWriter: w,
		}
		wrapped.statusCode = http.StatusOK

		start := time.Now()
		next.ServeHTTP(wrapped, r)
		duration := time.Since(start)

		logger.Info(
			"handled request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int64("duration_ns", duration.Nanoseconds()),
			slog.Int("status", wrapped.statusCode),
			slog.String("remote_addr", r.RemoteAddr),
			slog.String("x-forwarded-for", r.Header.Get("X-Forwareded-For")))
	})
}

// Logging middleware is used to write log information out to the console
// on each request/response.
func Logging(logger *slog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return logging(logger, next)
	}
}
