package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// LoggingMiddleware logs the details of each HTTP request and response.
func LoggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Create a response writer to capture the status code
			rec := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			// Call the next handler
			next.ServeHTTP(rec, r)

			// Log the request details
			logger.Info("request",
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()),
				zap.Int("status", rec.statusCode),
				zap.Duration("duration", time.Since(start)),
			)
		})
	}
}

// Logging middleware with Zap
func LoggingMiddleware2(logger *zap.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a response writer to capture the status code
		rec := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Log the incoming request
		logger.Info("Incoming request",
			zap.String("method", r.Method),
			zap.String("url", r.URL.String()),
			zap.Int("status", rec.statusCode),
			zap.Duration("duration", time.Since(start)),
		)

		next.ServeHTTP(rec, r) // Call the next handler
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
