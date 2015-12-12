package tmplt

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/tmrts/tmplt/pkg/util/exit"
	"github.com/tmrts/tmplt/pkg/util/inpututil"
	"github.com/tmrts/tmplt/pkg/util/osutil"
)

const (
	Version = "0.0.1"
)

var (
	Identifier    = "tmplt"
	DotIdentifier = "." + Identifier

	AppDataDirPath = filepath.Join("/var/lib", Identifier)

	DBPath = filepath.Join(AppDataDirPath, "config.db")

	ConfigPath      = DotIdentifier + "rc"
	ConfigDirPath   = DotIdentifier
	TemplateDirPath = filepath.Join(DotIdentifier, "templates")
)

func TemplatePath(name string) (string, error) {
	return filepath.Join(TemplateDirPath, name), nil
}

// Use init function to read .rcconfig file
func init() {
	homeDir, err := osutil.GetUserHomeDir()
	if err != nil {
		exit.Error(err)
	}

	ConfigPath = filepath.Join(homeDir, ConfigPath)
	ConfigDirPath = filepath.Join(homeDir, ConfigDirPath)
	TemplateDirPath = filepath.Join(homeDir, TemplateDirPath)

	if err := osutil.DirExists(TemplateDirPath); os.IsNotExist(err) {
		shouldInitialize, err := inpututil.ScanYesOrNo("Template directory doesn't exist. Initialize?", true)
		if err != nil {
			exit.Error(err)
		}

		if shouldInitialize {
			if err := Initialize(); err != nil {
				exit.Error(err)
			}
		} else {
			exit.Error(ErrUninitializedTmpltDir)
		}
	} else if err != nil {
		exit.Error(err)
	}
}

func Initialize() error {
	dirs := []string{
		DBPath,
		TemplateDirPath,
	}

	for _, path := range dirs {
		if _, err := exec.Command("/usr/bin/mkdir", "-p", path).Output(); err != nil {
			return err
		}
	}

	return initializeDB()
}
