package core

import (
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const (
	// Chars used when generating image file names in the database
	filenameChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var prng *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// fileExtensionLUT maps browser file type strings to the corresponding file
// extension.
var fileExtensionLUT = map[string] string {
	"image/jpeg": ".jpg",
	"image/png": ".png",
}

// ParentDir returns the 'level'th parent of the filepath.
// For example, ParentDir("/a/b/c.txt", 2) returns "/a" since the directory
// "/a" is the second-level parent of file "/a/b/c.txt".
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

// GetFilesInDir returns a slice of all the files in the given directory.
func GetFilesInDir(directory string) ([]string, error) {
	var files []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return []string{}, err
	} else {
		return files, err
	}
}

// DeleteAllFilesInDirectory deletes all of the files in the given directory.
func DeleteAllFilesInDirectory(directory string) error {
	files, err := GetFilesInDir(directory)
	for _, file := range files {
		DeleteFile(file)
	}
	return err
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

// RandomAlphanumericString returns a random string of the given length.
// containing only alphanumeric characters.
func RandomAlphanumericString(length int) string {
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = filenameChars[prng.Intn(len(filenameChars))]
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
