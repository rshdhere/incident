// Package events represents timestamped evidence in a scenario timeline.
package events

type Evidence struct {
	Time        string `json:"time"`
	Source      string `json:"source"`
	Observation string `json:"observation"`
}
