package persistence

import (
	"github.com/stretchr/testify/assert"
	"gotodo/internal/infra/persistence/record"
	"testing"
	"time"
)

func TestTaskRecord(t *testing.T) {
	now := time.Now().UTC()

	task := record.TaskRecord{
		UserID:      1,
		Title:       "Test Task",
		Description: "This is a test task",
		Completed:   false,
		CompletedAt: now,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Test TableName() function
	assert.Equal(t, "tasks", task.TableName())

	// Test struct field values
	assert.Equal(t, uint(0), task.ID) // should be initialized to 0 by default
	assert.Equal(t, 1, task.UserID)
	assert.Equal(t, "Test Task", task.Title)
	assert.Equal(t, "This is a test task", task.Description)
	assert.False(t, task.Completed)
	assert.Equal(t, now, task.CompletedAt)
	assert.Equal(t, now, task.CreatedAt)
	assert.Equal(t, now, task.UpdatedAt)
}
