package core

import (
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

var prng *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

var fileExtensionLUT = map[string] string {
	"image/jpeg": ".jpg",
}

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

// DeleteFile deletes the given file or empty directory if it exists.
func DeleteFile(filepath string) error {
	return os.Remove(filepath)
}

// WriteFile writes the given file to filepath.
func WriteFile(filepath string, file io.Reader) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	return err
}

// RandomAlphanumericString returns a random string of the given length
// containing only alphanumeric characters.
func RandomAlphanumericString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = chars[prng.Intn(len(chars))]
	}
	return string(bytes)
}

// FileExtensionFromFormat returns the corresponding file extension for the
// given format. If the format does not have a known file extension, the
// return value is an empty string. For example, if the input format is
// "image/jpeg", the return value is ".jpg".
func FileExtensionFromFormat(format string) string {
	extension, ok := fileExtensionLUT[format]
	if !ok {
		return ""
	}
	return extension
}
