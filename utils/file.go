package utils

import (
	"os"
)

// IsExist checks whether a file or directory exists or not.
// @name: filename or pathname
// It returns true when the file or directory already existed, return false on the contrary
func IsExist(name string) bool {
	_, err := os.Stat(name)
	return err == nil || os.IsExist(err)
}
