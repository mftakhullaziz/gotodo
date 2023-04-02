package persistence

import (
	"github.com/stretchr/testify/assert"
	"gotodo/internal/persistence/record"
	"reflect"
	"testing"
	"time"
)

func TestTaskRecord(t *testing.T) {
	t.Run("Test Task Record TableName", func(t *testing.T) {
		var task record.TaskRecord
		assert.NotNil(t, task)
		assert.Equal(t, "tasks", task.TableName())
	})

	t.Run("Test Task Record Struct Field Values", func(t *testing.T) {
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

		// Test struct field values
		assert.Equal(t, uint(0), task.ID) // should be initialized to 0 by default
		assert.Equal(t, 1, task.UserID)
		assert.Equal(t, "Test Task", task.Title)
		assert.Equal(t, "This is a test task", task.Description)
		assert.False(t, task.Completed)
		assert.Equal(t, now, task.CompletedAt)
		assert.Equal(t, now, task.CreatedAt)
		assert.Equal(t, now, task.UpdatedAt)
	})

	t.Run("Test Task Record DataType", func(t *testing.T) {
		var task record.TaskRecord

		idType := reflect.TypeOf(task.ID).Kind().String()
		assert.Equal(t, "uint", idType)

		userIdType := reflect.TypeOf(task.UserID).Kind().String()
		assert.Equal(t, "int", userIdType)

		titleType := reflect.TypeOf(task.Title).Kind().String()
		assert.Equal(t, "string", titleType)

		descriptionType := reflect.TypeOf(task.Description).Kind().String()
		assert.Equal(t, "string", descriptionType)

		completedType := reflect.TypeOf(task.Completed).Kind().String()
		assert.Equal(t, "bool", completedType)

		completedAt := reflect.TypeOf(task.CompletedAt).String()
		assert.Equal(t, "time.Time", completedAt)

		createAt := reflect.TypeOf(task.CreatedAt).String()
		assert.Equal(t, "time.Time", createAt)

		updateAt := reflect.TypeOf(task.UpdatedAt).String()
		assert.Equal(t, "time.Time", updateAt)
	})
}
