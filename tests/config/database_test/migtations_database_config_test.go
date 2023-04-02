package database

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gotodo/config"
	"gotodo/config/database"
	"gotodo/internal/persistence/record"
	"testing"
)

func TestMigrationsDatabaseConfig(t *testing.T) {
	t.Run("Test Migration Record", func(t *testing.T) {
		nameEnv := config.LoadEnv(".env.test")
		assert.NotNil(t, nameEnv)

		gdb, err := database.NewDatabaseConnection(context.Background(), nameEnv)
		assert.NoError(t, err)

		// Init record
		task := &record.TaskRecord{}
		account := &record.AccountRecord{}
		user := &record.UserDetailRecord{}

		err = database.MigrateDatabase(gdb, task, account, user)
		assert.NoError(t, err)
	})

	t.Run("Test Migration Record With In Validate", func(t *testing.T) {
		// Init record
		task := &record.TaskRecord{}
		account := &record.AccountRecord{}
		user := &record.UserDetailRecord{}

		nameEnv := config.LoadEnv(".env.test")
		assert.NotNil(t, nameEnv)

		gdb, err := database.NewDatabaseConnection(context.Background(), nameEnv)
		assert.NoError(t, err)

		migrate := gdb.AutoMigrate(task, account, user)
		assert.NoError(t, migrate)

		if !gdb.Migrator().HasTable(task) {
			t.Errorf("expected table 'tasks' to be created")
		}

		if !gdb.Migrator().HasColumn(task, "id") ||
			!gdb.Migrator().HasColumn(task, "user_id") ||
			!gdb.Migrator().HasColumn(task, "title") ||
			!gdb.Migrator().HasColumn(task, "description") ||
			!gdb.Migrator().HasColumn(task, "completed") ||
			!gdb.Migrator().HasColumn(task, "completed_at") ||
			!gdb.Migrator().HasColumn(task, "created_at") ||
			!gdb.Migrator().HasColumn(task, "updated_at") {
			t.Errorf("expected columns 'id','user_id','title','description'," +
				"'completed','completed_at','create_at','update_at' to be created in table 'tasks'")
		}

		if !gdb.Migrator().HasTable(account) {
			t.Errorf("expected table 'accounts' to be created")
		}

		if !gdb.Migrator().HasColumn(account, "account_id") ||
			!gdb.Migrator().HasColumn(account, "user_id") ||
			!gdb.Migrator().HasColumn(account, "username") ||
			!gdb.Migrator().HasColumn(account, "password") ||
			!gdb.Migrator().HasColumn(account, "status") ||
			!gdb.Migrator().HasColumn(account, "created_at") ||
			!gdb.Migrator().HasColumn(account, "updated_at") {
			t.Errorf("expected columns 'account_id','user_id','username'," +
				"'password','status','created_at','updated_at' to be created in table 'accounts'")
		}

		if !gdb.Migrator().HasTable(user) {
			t.Errorf("expected table 'user_details' to be created")
		}

		if !gdb.Migrator().HasColumn(user, "user_id") ||
			!gdb.Migrator().HasColumn(user, "username") ||
			!gdb.Migrator().HasColumn(user, "password") ||
			!gdb.Migrator().HasColumn(user, "email") ||
			!gdb.Migrator().HasColumn(user, "name") ||
			!gdb.Migrator().HasColumn(user, "mobile_phone") ||
			!gdb.Migrator().HasColumn(user, "address") ||
			!gdb.Migrator().HasColumn(user, "status") ||
			!gdb.Migrator().HasColumn(user, "created_at") ||
			!gdb.Migrator().HasColumn(user, "updated_at") {
			t.Errorf("expected columns 'user_id','username','password'," +
				"'email','name','mobile_phone','address','status','created_at'," +
				"'updated_at' to be created in table 'accounts'")
		}

		// Get the current database name and compare it to the expected value
		currentDB := gdb.Migrator().CurrentDatabase()
		expectedDB := "todo_app"
		if currentDB != expectedDB {
			t.Errorf("expected database name '%s', but got '%s'", expectedDB, currentDB)
		}

	})
}
