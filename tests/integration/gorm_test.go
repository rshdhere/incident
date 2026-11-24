package integration_test

import (
	"path/filepath"
	"testing"

	"github.com/rshdhere/incident/packages/api/v1/persistence"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestAttemptMigration(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(filepath.Join(t.TempDir(), "practice.db")), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := sqlDB.Close(); err != nil {
			t.Errorf("close test database: %v", err)
		}
	})

	if err := persistence.Migrate(db); err != nil {
		t.Fatal(err)
	}
	attempt := persistence.Attempt{
		LearnerID: "learner-1", ScenarioID: "sqs-backlog", ScenarioRevision: "fixture-v1",
		Reasoning: "Check the consumer before increasing queue capacity.",
	}
	if err := db.Create(&attempt).Error; err != nil {
		t.Fatal(err)
	}
	if err := persistence.Migrate(db); err != nil {
		t.Fatal(err)
	}
	var saved persistence.Attempt
	if err := db.First(&saved, attempt.ID).Error; err != nil {
		t.Fatal(err)
	}
	if saved.LearnerID != attempt.LearnerID || saved.ScenarioID != attempt.ScenarioID ||
		saved.ScenarioRevision != attempt.ScenarioRevision || saved.Reasoning != attempt.Reasoning ||
		!saved.CreatedAt.Equal(attempt.CreatedAt) {
		t.Fatalf("attempt changed after migration: %+v", saved)
	}
}
