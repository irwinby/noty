package middleware

import (
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type ResponseWriter struct {
	http.ResponseWriter

	status int
}

func (l *Logger) WrapResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
	}
}

func (w *ResponseWriter) Status() int {
	return w.status
}

func (w *ResponseWriter) WriteHeader(code int) {
	w.status = code

	w.ResponseWriter.WriteHeader(code)
}

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Log() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()

			_w := l.WrapResponseWriter(w)

			next.ServeHTTP(_w, r)

			body, _ := io.ReadAll(r.Body)

			zap.L().Info(
				r.Method,
				zap.String("body", string(body)),
				zap.String("host", r.Host),
				zap.String("path", r.URL.EscapedPath()),
				zap.String("query", r.URL.RawQuery),
				zap.Int("status", _w.Status()),
				zap.Duration("duration", time.Since(now)),
			)
		})
	}
}
