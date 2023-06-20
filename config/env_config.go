package config

import (
	"gotodo/internal/utils"
	"os"
	"path/filepath"
)

func LoadEnv(name string) string {
	dir, err := os.Getwd()
	utils.PanicIfError(err)
	// rootDir := filepath.Dir(filepath.Dir(filepath.Dir(dir)))
	// Construct the full path to the .env.test file
	envPath := filepath.Join(dir, name)
	return envPath
}

func EnvironmentTest() {
	_ = os.Setenv("ENV", "test")
}
