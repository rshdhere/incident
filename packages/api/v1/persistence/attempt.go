// Package persistence owns the v1 learner-attempt schema.
package persistence

import (
	"time"

	"gorm.io/gorm"
)

// Attempt records one learner's reasoning against a particular exercise revision.
type Attempt struct {
	ID               uint      `gorm:"primaryKey"`
	LearnerID        string    `gorm:"not null;index"`
	ScenarioID       string    `gorm:"not null"`
	ScenarioRevision string    `gorm:"not null"`
	Reasoning        string    `gorm:"not null"`
	CreatedAt        time.Time `gorm:"not null"`
}

// Migrate prepares the initial attempt schema in the caller's database.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Attempt{})
}
