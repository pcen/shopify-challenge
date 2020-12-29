package core

import (
	"os"
	"path/filepath"
)

// Filepath Utilities

// ParentDir returns the 'level'th parent of the filepath.
func ParentDir(path string, level int) string {
	for i := 0; i < level; i++ {
		path = filepath.Dir(path)
	}
	parent, _ := filepath.Abs(path)
	return parent
}

// EnsureDirExists creates the specified directory if it does not exist.
func EnsureDirExists(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0700)
	}
}
