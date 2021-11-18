package utils

import (
	"testing"
)

func TestIsExist(t *testing.T) {
	fileStatusMap := map[string]bool{
		"file.go":       true,
		"file_debug.go": false,
	}
	for filename, expected := range fileStatusMap {
		existed := IsExist(filename)
		if existed != expected {
			t.Errorf("failed to test IsExist, expected %t, but got %t", expected, existed)
		}
	}
}
