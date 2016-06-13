package osutil

import "os/user"

// GetUserHomeDir returns the home directory of the user.
func GetUserHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return usr.HomeDir, nil
}
