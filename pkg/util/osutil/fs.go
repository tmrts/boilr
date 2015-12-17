package osutil

import (
	"fmt"
	"os"
)

func FileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	if info.IsDir() {
		return false, fmt.Errorf("%v: is a directory, expected file")
	}

	return true, nil
}

func DirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	if !info.IsDir() {
		return false, fmt.Errorf("%v: is a file, expected directory")
	}

	return true, nil
}

func CreateDirs(dirPaths ...string) error {
	for _, path := range dirPaths {
		if err := os.MkdirAll(path, 0744); err != nil {
			return err
		}
	}

	return nil
}
