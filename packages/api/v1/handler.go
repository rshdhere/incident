// Package v1 serves the first version of the incident catalog API.
package v1

import (
	_ "embed"
	"encoding/json"
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/rshdhere/incident/packages/api/v1/services"
)

//go:embed openapi.json
var openAPI []byte

// NewHandler loads the registered AWS exercises and returns the catalog HTTP handler.
func NewHandler(root fs.FS) (http.Handler, error) {
	return NewHandlerWithLogger(root, slog.Default())
}

// NewHandlerWithLogger sends request logs to the supplied logger.
func NewHandlerWithLogger(root fs.FS, logger *slog.Logger) (http.Handler, error) {
	catalog, err := services.Load(root)
	if err != nil {
		return nil, err
	}
	if logger == nil {
		logger = slog.Default()
	}
	writeJSON := func(w http.ResponseWriter, status int, data any) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			logger.Error("write JSON response", "request_id", w.Header().Get(requestIDHeader), "error", err)
		}
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(openAPI); err != nil {
			logger.Error("write OpenAPI response", "request_id", w.Header().Get(requestIDHeader), "error", err)
		}
	})
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /api/v1/scenarios", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, catalog.Summaries)
	})
	mux.HandleFunc("GET /api/v1/scenarios/{id}", func(w http.ResponseWriter, r *http.Request) {
		scenario, ok := catalog.Find(r.PathValue("id"))
		if !ok {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "unknown scenario"})
			return
		}
		writeJSON(w, http.StatusOK, scenario.Exercise)
	})
	mux.HandleFunc("GET /api/v1/scenarios/{id}/debrief", func(w http.ResponseWriter, r *http.Request) {
		scenario, ok := catalog.Find(r.PathValue("id"))
		if !ok {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "unknown scenario"})
			return
		}
		writeJSON(w, http.StatusOK, scenario.WorkedResponse)
	})
	return withRequestLog(mux, logger), nil
}
