package osutil

import (
	"os"
	"os/exec"
)

func FileExists(filename string) error {
	_, err := os.Stat(filename)
	if err != nil {
		return err
	}

	return nil
}

func DirExists(dirname string) error {
	_, err := os.Stat(dirname)
	if err != nil {
		return err
	}

	return nil
}

func CreateDirs(dirPaths ...string) error {
	for _, path := range dirPaths {
		if _, err := exec.Command("/usr/bin/mkdir", "-p", path).Output(); err != nil {
			return err
		}
	}

	return nil
}
