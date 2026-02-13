package util

import (
	"errors"
	"fmt"
	"os"
)

func DoesFileExist(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, fmt.Errorf("failed to check if file %s exists: %w", path, err)
}
