// Package types defines the JSON contract for incident exercises.
package types

import "github.com/rshdhere/incident/packages/api/v1/events"

type Summary struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Platform string `json:"platform"`
}

type Exercise struct {
	Summary
	Brief    string            `json:"brief"`
	Question string            `json:"question"`
	Evidence []events.Evidence `json:"evidence"`
}

type Debrief struct {
	Diagnosis  string `json:"diagnosis"`
	Mitigation string `json:"mitigation"`
	Recovery   string `json:"recovery"`
}

type Scenario struct {
	Exercise
	WorkedResponse Debrief `json:"workedResponse"`
}
