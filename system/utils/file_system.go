package utils

import (
	"os"
)

// FileIsExist is a function to check FileIsExist
func FileIsExist(file string) (bool, error) {
	_, err := os.Stat(file)
	if err != nil {
		return false, err
	}
	return true, nil
}
