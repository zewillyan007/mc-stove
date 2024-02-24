package util

import "os"

func DirectoryExists(path string) bool {

	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	}
	return false
}
