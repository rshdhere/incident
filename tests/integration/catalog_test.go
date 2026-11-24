package integration_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	v1 "github.com/rshdhere/incident/packages/api/v1"
)

func TestExerciseJourney(t *testing.T) {
	handler, err := v1.NewHandler(os.DirFS("../.."))
	if err != nil {
		t.Fatal(err)
	}
	request := func(method, path string, wantStatus int, value any) {
		t.Helper()
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, httptest.NewRequest(method, path, nil))
		if rec.Code != wantStatus {
			t.Fatalf("%s %s returned %d: %s", method, path, rec.Code, rec.Body.String())
		}
		if value != nil {
			if err := json.Unmarshal(rec.Body.Bytes(), value); err != nil {
				t.Fatal(err)
			}
		}
	}

	var health map[string]string
	request("GET", "/healthz", http.StatusOK, &health)
	if health["status"] != "ok" {
		t.Fatalf("unexpected health response: %v", health)
	}

	var summaries []map[string]string
	request("GET", "/api/v1/scenarios", http.StatusOK, &summaries)
	var previousID string
	var foundSQS bool
	for _, summary := range summaries {
		id := summary["id"]
		if len(summary) != 3 || summary["platform"] != "aws" || summary["title"] == "" || id <= previousID {
			t.Fatalf("unexpected summary or listing order: %v", summaries)
		}
		previousID = id
		if id == "sqs-backlog" {
			foundSQS = true
		}
	}
	if !foundSQS {
		t.Fatal("SQS backlog exercise missing from catalog")
	}

	const detail = "/api/v1/scenarios/sqs-backlog"
	var exercise struct {
		ID             string
		Question       string
		Evidence       []struct{ Observation string }
		WorkedResponse json.RawMessage
	}
	request("GET", detail, http.StatusOK, &exercise)
	if exercise.WorkedResponse != nil {
		t.Fatal("exercise exposed the worked response before debrief")
	}
	if exercise.ID != "sqs-backlog" || exercise.Question == "" || len(exercise.Evidence) == 0 || exercise.Evidence[0].Observation == "" {
		t.Fatalf("unexpected exercise: %+v", exercise)
	}

	var debrief map[string]string
	request("GET", detail+"/debrief", http.StatusOK, &debrief)
	if debrief["diagnosis"] == "" || debrief["mitigation"] == "" || debrief["recovery"] == "" {
		t.Fatalf("incomplete debrief: %v", debrief)
	}
	request("GET", "/api/v1/scenarios/unknown", http.StatusNotFound, nil)
	request("GET", "/api/v1/scenarios/unknown/debrief", http.StatusNotFound, nil)
	request("POST", "/api/v1/scenarios", http.StatusMethodNotAllowed, nil)
	request("GET", "/api/scenarios", http.StatusNotFound, nil)
	request("GET", "/api/v2/scenarios", http.StatusNotFound, nil)
}
