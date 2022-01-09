package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// AppendToFile appends the provided string to the given filename
func AppendToFile(filename string, toWrite []byte, perm os.FileMode) error {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm)
	if err != nil {
		return err
	}
	if _, err := f.Write(toWrite); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

// GetAbsPath turns the regular path provided into an absolute path for other
// functions to use
func GetAbsPath(path string) (string, error) {
	var absPath string
	if strings.HasPrefix(path, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		absPath = strings.Replace(path, "~", homeDir, 1)
	}
	absPath, err := filepath.Abs(absPath)
	if err != nil {
		return "", err
	}

	return absPath, nil
}
