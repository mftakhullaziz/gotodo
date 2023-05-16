package config

import (
	errs "gotodo/internal/utils/errors"
	"os"
	"path/filepath"
)

func LoadEnv(name string) string {
	dir, err := os.Getwd()
	errs.PanicIfError(err)
	// rootDir := filepath.Dir(filepath.Dir(filepath.Dir(dir)))
	// Construct the full path to the .env.test file
	envPath := filepath.Join(dir, name)
	return envPath
}
