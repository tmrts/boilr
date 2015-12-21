package osutil

import (
	"fmt"
	"os"
)

// FileExists checks whether the given path exists and belongs to a file.
func FileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	if info.IsDir() {
		return false, fmt.Errorf("%v: is a directory, expected file", path)
	}

	return true, nil
}

// DirExists checks whether the given path exists and belongs to a directory.
func DirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	if !info.IsDir() {
		return false, fmt.Errorf("%v: is a file, expected directory", path)
	}

	return true, nil
}

// CreateDirs creates directories from the given directory path arguments.
func CreateDirs(dirPaths ...string) error {
	for _, path := range dirPaths {
		if err := os.MkdirAll(path, 0744); err != nil {
			return err
		}
	}

	return nil
}
