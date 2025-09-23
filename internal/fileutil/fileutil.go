package fileutil

import (
	"strings"
	"os"
)

func CheckExtension(path string, ext string) bool {
	if ext == "" {
		return false
	}

	return strings.HasSuffix(path, ext)
}

func IsAvailable(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
