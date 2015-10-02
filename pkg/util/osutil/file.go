package osutil

import "os"

func FileExists(filename string) error {
	_, err := os.Stat(filename)
	if err != nil {
		return err
	}

	return nil
}
