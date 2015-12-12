package tmplt

import (
	"path/filepath"

	"github.com/tmrts/tmplt/pkg/util/osutil"
)

var (
	ConfigPath = ".tmpltrc"

	ConfigDirPath   = ".tmplt/"
	TemplateDirPath = ".tmplt/templates/"
)

func TemplatePath(name string) (string, error) {
	return filepath.Join(TemplateDirPath, name), nil
}

// Use init function to read .rcconfig file
func init() {
	homeDir, err := osutil.GetUserHomeDir()
	if err != nil {
		panic(err)
	}

	ConfigPath = filepath.Join(homeDir, ConfigPath)
	ConfigDirPath = filepath.Join(homeDir, ConfigDirPath)
	TemplateDirPath = filepath.Join(homeDir, TemplateDirPath)
}
