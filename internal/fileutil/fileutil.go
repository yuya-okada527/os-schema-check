package fileutil

import "strings"

func CheckExtension(path string, ext string) bool {
	if ext == "" {
		return false
	}

	return strings.HasSuffix(path, ext)
}
