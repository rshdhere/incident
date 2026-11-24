package bench_test

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	v1 "github.com/rshdhere/incident/packages/api/v1"
)

func BenchmarkScenarioDetail(b *testing.B) {
	logger := slog.New(slog.NewJSONHandler(io.Discard, nil))
	handler, err := v1.NewHandlerWithLogger(os.DirFS("../.."), logger)
	if err != nil {
		b.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "/api/v1/scenarios/sqs-backlog", nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, request)
		if response.Code != http.StatusOK {
			b.Fatalf("scenario detail returned HTTP %d", response.Code)
		}
	}
}
