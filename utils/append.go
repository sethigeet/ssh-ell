package utils

import "os"

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
