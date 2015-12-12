package cookie

import (
	"path/filepath"

	"github.com/tmrts/cookie/pkg/util/osutil"
)

var (
	ConfigPath = ".cookierc"

	ConfigDirPath   = ".cookie/"
	TemplateDirPath = ".cookie/templates/"
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
