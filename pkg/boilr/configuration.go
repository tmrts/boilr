package boilr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tmrts/boilr/pkg/util/exit"
	"github.com/tmrts/boilr/pkg/util/osutil"
	"github.com/tmrts/boilr/pkg/util/tlog"
)

const (
	// Name of the application
	AppName = "boilr"

	// Version of the application
	Version = "0.1.0"

	// Configuration Directory of the application
	ConfigDirPath = ".config/boilr"

	// Configuration File Name of the application
	ConfigFileName = "config.json"

	// Directory that contains the template registry
	TemplateDir = "templates"

	// Name of the file that contains the context values for the template
	ContextFileName = "project.json"

	// Name of the directory that contains the template files in a boilr template
	TemplateDirName = "template"

	// Name of the file that contains the metadata about the template saved in registry
	TemplateMetadataName = "__metadata.json"

	// Owner of the github repository
	GithubOwner = "tmrts"

	// Name of the github repository
	GithubRepo = "boilr"
)

// Configuration contains the values for needed for boilr to operate.
// These values can be overridden by the inclusion of a boilr.json
// file in the configuration directory.
var Configuration = struct {
	FilePath        string
	TemplateDirPath string
}{}

// TemplatePath returns the absolute path of a template given the name of the template.
func TemplatePath(name string) (string, error) {
	return filepath.Join(Configuration.TemplateDirPath, name), nil
}

func init() {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		// FIXME is this really necessary?
		exit.Error(fmt.Errorf("environment variable ${HOME} should be set"))
	}

	Configuration.FilePath = filepath.Join(homeDir, ConfigDirPath, ConfigFileName)
	Configuration.TemplateDirPath = filepath.Join(homeDir, ConfigDirPath, TemplateDir)

	IsTemplateDirInitialized, err := osutil.DirExists(Configuration.TemplateDirPath)
	if err != nil {
		exit.Error(err)
	}

	// TODO perform this in related commands only with ValidateInitialization
	if !IsTemplateDirInitialized {
		tlog.Warn("Template registry is not initialized. Please run `init` command to create it.")
		return
	}

	// Read .config/boilr/config.json if exists
	// TODO use defaults if config.json doesn't exist
	hasConfig, err := osutil.FileExists(Configuration.FilePath)
	if err != nil {
		exit.Error(err)
	}

	if !hasConfig {
		// TODO report the absence of config.json
		//tlog.Debug("Couldn't find %s user configuration", ConfigFileName)
		return
	}

	buf, err := ioutil.ReadFile(Configuration.FilePath)
	if err != nil {
		exit.Error(err)
	}

	if err := json.Unmarshal(buf, &Configuration); err != nil {
		exit.Error(err)
	}
}
