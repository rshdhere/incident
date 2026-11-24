// Package validators checks authored incident content before it is served.
package validators

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rshdhere/incident/packages/api/v1/types"
)

func Scenario(s types.Scenario) error {
	if missing(s.ID, s.Title, s.Brief, s.Question) {
		return errors.New("expected an id, title, brief and question")
	}
	if len(s.Evidence) == 0 {
		return errors.New("expected at least one evidence item")
	}
	for i, item := range s.Evidence {
		if missing(item.Time, item.Source, item.Observation) {
			return fmt.Errorf("evidence item %d needs a time, source and observation", i+1)
		}
	}
	if missing(s.WorkedResponse.Diagnosis, s.WorkedResponse.Mitigation, s.WorkedResponse.Recovery) {
		return errors.New("workedResponse needs a diagnosis, mitigation and recovery")
	}
	return nil
}

func missing(values ...string) bool {
	for _, v := range values {
		if strings.TrimSpace(v) == "" {
			return true
		}
	}
	return false
}
