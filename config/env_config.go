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

func LoadEnvFromFile(path string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		return ""
	}

	gotodoDir := filepath.Join(currentDir, path, "gotodo")
	envFilePath := filepath.Join(gotodoDir, ".env.test")
	return envFilePath
}
