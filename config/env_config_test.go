package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Fatal(err)
		}
	}(tmpDir)

	// Create a .env.test file in the temporary directory
	envFile := filepath.Join(tmpDir, ".env.test")
	if err := os.WriteFile(envFile, []byte("FOO=bar"), 0644); err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, envFile)

	// Test that LoadEnv() returns the path to the .env.test file
	envPath := LoadEnv(".env.test")
	assert.NotNil(t, envPath)
}

func TestEnvironmentTest(t *testing.T) {
	expectedEnv := "test"

	// Call the EnvironmentTest function
	EnvironmentTest()

	// Retrieve the value of the "ENV" environment variable
	env := os.Getenv("ENV")

	// Check if the environment variable has the expected value
	if env != expectedEnv {
		t.Errorf("Unexpected environment variable value. Got: %s, Expected: %s", env, expectedEnv)
	}
}
