package tasks

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gotodo/config"
	"gotodo/config/database"
	"gotodo/internal/persistence/record"
	"gotodo/internal/utils"
	"time"
)

// Mock function environment testing
func mockEnvTest() {
	// The param path is dynamic following testing path file
	path := config.LoadEnvFromFile("../../../..")
	_ = godotenv.Load(path)
}

// Mock function insert data account user
func mockInsertUserData(db *gorm.DB) {
	// Perform mock data insertion for account table
	mockAccount := &record.AccountRecord{
		// Set the desired mock account data
		AccountID: 1,
		UserID:    1,
		Username:  "@johndoe_test",
		Password:  utils.HashPasswordAndSalt([]byte("password")),
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_ = db.Create(mockAccount)

	// Perform mock data insertion for user detail table
	mockUserDetail := &record.UserDetailRecord{
		// Set the desired mock user detail data
		UserID:      uint(mockAccount.UserID),
		Username:    mockAccount.Username,
		Password:    mockAccount.Password,
		Email:       "johndoe@mail.com",
		Name:        "John Doe",
		MobilePhone: 6282299812,
		Address:     "Jakarta",
		Status:      "active",
		CreatedAt:   mockAccount.CreatedAt,
		UpdatedAt:   mockAccount.UpdatedAt,
	}
	_ = db.Create(mockUserDetail)

	mockTask := &record.TaskRecord{
		TaskID:      1,
		UserID:      mockAccount.UserID,
		Title:       "Create first tasks to create fixing bugs",
		Description: "Completed tasks to finish fixing bugs hotfix in production",
		Completed:   false,
		TaskStatus:  "active",
		CompletedAt: time.Time{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
	_ = db.Create(mockTask)
}

// Mock function database sqlite
func mockDBTest() *gorm.DB {
	mockEnvTest()

	// Create an in-memory SQLite database for testing
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	// Run database migrations or any other initialization steps
	_ = database.MigrateDatabase(
		db,
		&record.TaskRecord{},
		&record.AccountRecord{},
		&record.UserDetailRecord{},
		&record.AccountLoginHistoriesRecord{},
	)

	mockInsertUserData(db)

	return db
}
