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
	AppName       = "boilr"
	Version       = "0.0.1"
	ConfigDirPath = ".config/boilr"

	ConfigFileName = "config.json"
	TemplateDir    = "templates"

	ContextFileName      = "project.json"
	TemplateDirName      = "template"
	TemplateMetadataName = "__metadata.json"

	GithubOwner = "tmrts"
	GithubRepo  = "boilr"
)

var Configuration = struct {
	FilePath        string
	TemplateDirPath string
}{}

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
