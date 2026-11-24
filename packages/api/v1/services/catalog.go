// Package services loads the authored exercises into a read-only catalog.
package services

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path"
	"sort"
	"strings"

	"github.com/rshdhere/incident/packages/api/v1/types"
	"github.com/rshdhere/incident/packages/api/v1/validators"
)

// Register only authored content. Provider directories can also hold private CLI state.
var sources = []string{
	".aws/sqs-backlog.json",
}

type Catalog struct {
	byID      map[string]types.Scenario
	Summaries []types.Summary
}

func Load(root fs.FS) (*Catalog, error) {
	catalog := &Catalog{byID: make(map[string]types.Scenario)}
	for _, file := range sources {
		raw, err := fs.ReadFile(root, file)
		if err != nil {
			return nil, err
		}
		var scenario types.Scenario
		if err := json.Unmarshal(raw, &scenario); err != nil {
			return nil, fmt.Errorf("%s: %w", file, err)
		}
		if err := validators.Scenario(scenario); err != nil {
			return nil, fmt.Errorf("%s: %w", file, err)
		}
		if scenario.ID != strings.TrimSuffix(path.Base(file), ".json") {
			return nil, fmt.Errorf("%s: id must match filename", file)
		}
		if _, exists := catalog.byID[scenario.ID]; exists {
			return nil, fmt.Errorf("%s: duplicate id %q", file, scenario.ID)
		}
		scenario.Platform = strings.TrimPrefix(path.Dir(file), ".")
		catalog.byID[scenario.ID] = scenario
		catalog.Summaries = append(catalog.Summaries, scenario.Summary)
	}
	sort.Slice(catalog.Summaries, func(i, j int) bool {
		return catalog.Summaries[i].ID < catalog.Summaries[j].ID
	})
	return catalog, nil
}

func (c *Catalog) Find(id string) (types.Scenario, bool) {
	s, ok := c.byID[id]
	return s, ok
}
