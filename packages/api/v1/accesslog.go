package v1

import (
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const requestIDHeader = "X-Request-ID"

var requestIDFormat = regexp.MustCompile(`^[A-Za-z0-9._-]{1,64}$`)

func withRequestLog(mux *http.ServeMux, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		id, err := requestID(r.Header.Values(requestIDHeader))
		if err != nil {
			logger.ErrorContext(r.Context(), "generate request ID", "error", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set(requestIDHeader, id)
		response := &statusWriter{ResponseWriter: w}
		mux.ServeHTTP(response, r)

		if logger.Enabled(r.Context(), slog.LevelInfo) {
			// ServeMux.Handler exposes registered patterns on Go 1.22. Only log
			// method-qualified patterns, since redirects can return a URL path.
			_, pattern := mux.Handler(r)
			_, route, matched := strings.Cut(pattern, " ")
			if !matched {
				route = "unmatched"
			}
			status := response.status
			if status == 0 {
				status = http.StatusOK
			}
			logger.InfoContext(r.Context(), "http request",
				"request_id", id, "method", r.Method, "route", route,
				"status", status, "duration_ms", float64(time.Since(start))/float64(time.Millisecond))
		}
	})
}

func requestID(values []string) (string, error) {
	if len(values) == 1 && len(values[0]) <= 64 && requestIDFormat.MatchString(values[0]) {
		return values[0], nil
	}
	var id [16]byte
	if _, err := rand.Read(id[:]); err != nil {
		return "", err
	}
	return hex.EncodeToString(id[:]), nil
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) Unwrap() http.ResponseWriter { return w.ResponseWriter }

func (w *statusWriter) WriteHeader(code int) {
	if w.status == 0 && (code >= 200 || code == http.StatusSwitchingProtocols) {
		w.status = code
	}
	w.ResponseWriter.WriteHeader(code)
}

func (w *statusWriter) Write(data []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	return w.ResponseWriter.Write(data)
}
