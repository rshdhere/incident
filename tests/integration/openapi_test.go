package integration_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	v1 "github.com/rshdhere/incident/packages/api/v1"
)

func TestOpenAPIContract(t *testing.T) {
	handler, err := v1.NewHandler(os.DirFS("../.."))
	if err != nil {
		t.Fatal(err)
	}
	get := func(t *testing.T, path string, status int, payload any) {
		t.Helper()
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
		if rec.Code != status || rec.Header().Get("X-Request-ID") == "" {
			t.Fatalf("GET %s: status %d, headers %v", path, rec.Code, rec.Header())
		}
		if got := strings.Split(rec.Header().Get("Content-Type"), ";")[0]; got != "application/json" {
			t.Fatalf("GET %s: content type %q", path, got)
		}
		if err := json.Unmarshal(rec.Body.Bytes(), payload); err != nil {
			t.Fatalf("GET %s: %v", path, err)
		}
	}
	type response struct {
		Ref     string `json:"$ref"`
		Headers map[string]json.RawMessage
		Content map[string]struct {
			Schema struct {
				Ref string `json:"$ref"`
			}
		}
	}
	var spec struct {
		OpenAPI string
		Paths   map[string]struct {
			Get struct{ Responses map[string]response }
		}
		Components struct {
			Responses map[string]response
			Schemas   map[string]struct{ Required []string }
		}
	}
	get(t, "/api/v1/openapi.json", http.StatusOK, &spec)
	if spec.OpenAPI != "3.0.3" {
		t.Fatalf("unexpected OpenAPI version %q", spec.OpenAPI)
	}
	for _, path := range []string{"/healthz", "/api/v1/scenarios", "/api/v1/scenarios/{id}", "/api/v1/scenarios/{id}/debrief", "/api/v1/openapi.json"} {
		if _, ok := spec.Paths[path].Get.Responses["200"].Headers["X-Request-ID"]; !ok {
			t.Errorf("%s is missing a documented 200 response with a correlation ID", path)
		}
	}

	const detail = "/api/v1/scenarios/{id}"
	for _, tc := range []struct {
		path, id, schema string
		status           int
	}{
		{detail, "sqs-backlog", "Exercise", http.StatusOK},
		{detail + "/debrief", "sqs-backlog", "Debrief", http.StatusOK},
		{detail, "missing", "Error", http.StatusNotFound},
		{detail + "/debrief", "missing", "Error", http.StatusNotFound},
	} {
		path := strings.Replace(tc.path, "{id}", tc.id, 1)
		t.Run(path, func(t *testing.T) {
			documented := spec.Paths[tc.path].Get.Responses[strconv.Itoa(tc.status)]
			if documented.Ref != "" {
				documented = spec.Components.Responses[strings.TrimPrefix(documented.Ref, "#/components/responses/")]
			}
			if got := documented.Content["application/json"].Schema.Ref; got != "#/components/schemas/"+tc.schema {
				t.Fatalf("response schema %q, want %s", got, tc.schema)
			}
			var actual map[string]json.RawMessage
			get(t, path, tc.status, &actual)
			required := slices.Clone(spec.Components.Schemas[tc.schema].Required)
			var keys []string
			for key := range actual {
				keys = append(keys, key)
			}
			slices.Sort(keys)
			slices.Sort(required)
			if !slices.Equal(keys, required) {
				t.Errorf("actual fields %v, contract requires %v", keys, required)
			}
		})
	}
}
