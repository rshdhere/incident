package integration_test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	v1 "github.com/rshdhere/incident/packages/api/v1"
)

func TestRequestCorrelation(t *testing.T) {
	var logs bytes.Buffer
	handler, err := v1.NewHandlerWithLogger(os.DirFS("../.."), slog.New(slog.NewJSONHandler(&logs, nil)))
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		method, path, route string
		status              int
		headers             []string
		wantID              string
	}{
		{"GET", "/healthz", "/healthz", 200, []string{"mobile_9.2-01"}, "mobile_9.2-01"},
		{"GET", "/api/v1/scenarios/private-value", "/api/v1/scenarios/{id}", 404, nil, ""},
		{"POST", "/healthz", "unmatched", 405, []string{"id with spaces"}, ""},
		{"GET", "/private-value", "unmatched", 404, []string{strings.Repeat("x", 65)}, ""},
		{"GET", "/healthz", "/healthz", 200, []string{"one", "two"}, ""},
	}
	generated := make(map[string]bool)
	for _, c := range cases {
		logs.Reset()
		req := httptest.NewRequest(c.method, c.path+"?token=private-value", nil)
		req.Header.Set("Authorization", "Bearer private-value")
		for _, value := range c.headers {
			req.Header.Add("X-Request-ID", value)
		}
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, req)
		id := response.Header().Get("X-Request-ID")
		if c.wantID != "" {
			if id != c.wantID {
				t.Fatalf("expected supplied request ID %q, got %q", c.wantID, id)
			}
		} else {
			b, err := hex.DecodeString(id)
			if err != nil || len(b) != 16 || generated[id] {
				t.Fatalf("expected unique generated request ID, got %q", id)
			}
			generated[id] = true
		}
		if strings.Contains(logs.String(), "private-value") {
			t.Fatal("access log contains private request data")
		}
		var entry struct {
			RequestID string   `json:"request_id"`
			Method    string   `json:"method"`
			Route     string   `json:"route"`
			Status    int      `json:"status"`
			Duration  *float64 `json:"duration_ms"`
		}
		if err := json.Unmarshal(logs.Bytes(), &entry); err != nil {
			t.Fatal(err)
		}
		if entry.RequestID != id || entry.Route != c.route || entry.Method != c.method ||
			entry.Status != c.status || response.Code != c.status || entry.Duration == nil || *entry.Duration < 0 {
			t.Fatalf("%s %s returned %d with log %s", c.method, c.path, response.Code, logs.String())
		}
		if c.status == http.StatusMethodNotAllowed && !strings.Contains(response.Header().Get("Allow"), "GET") {
			t.Fatal("missing Allow header on method rejection")
		}
	}
}
