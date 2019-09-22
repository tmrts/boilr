package cmd

import (
	"fmt"
	"os"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/boilr/pkg/boilr"
	"github.com/tmrts/boilr/pkg/host"
	"github.com/tmrts/boilr/pkg/util/exit"
	"github.com/tmrts/boilr/pkg/util/git"
	"github.com/tmrts/boilr/pkg/util/osutil"
	"github.com/tmrts/boilr/pkg/util/validate"
)

// Download contains the cli-command for downloading templates from github.
var Download = &cli.Command{
	Use:   "download <template-repo> <template-tag>",
	Short: "Download a project template from a github repository to template registry",
	// FIXME Half-Updates leave messy templates
	Run: func(c *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{
			{"template-repo", validate.UnixPath},
			{"template-tag", validate.AlphanumericExt},
		})

		MustValidateTemplateDir()

		templateURL, templateName, templateSubFolder := args[0], args[1], GetStringFlag(c, "sub-path")

		targetDir, err := boilr.TemplatePath(templateName)
		if err != nil {
			exit.Error(fmt.Errorf("download: %s", err))
		}

		targetTmpDir := targetDir

		switch exists, err := osutil.DirExists(targetDir); {
		case err != nil:
			exit.Error(fmt.Errorf("download: %s", err))
		case exists:
			if shouldOverwrite := GetBoolFlag(c, "force"); !shouldOverwrite {
				exit.OK("Template %v already exists use -f to overwrite the template", templateName)
			}

			// TODO(tmrts): extract `template delete` helper and use that one
			if err := os.RemoveAll(targetDir); err != nil {
				exit.Error(fmt.Errorf("download: %s", err))
			}
		}
		// In case if we are copying template from repository sub-folder, clone repo to temp folder
		if templateSubFolder != "" {
			targetTmpDir, err = boilr.TemplateTempPath(templateName)
			if err != nil {
				exit.Error(fmt.Errorf("download: %s", err))
			}
			exists, err := osutil.DirExists(targetTmpDir)
			if exists || (!exists && err != nil) {
				if err := os.RemoveAll(targetTmpDir); err != nil {
					exit.Error(fmt.Errorf("download: %s", err))
				}
			}
		}

		// TODO(tmrts): allow fetching other branches than 'master'
		if err := git.Clone(targetTmpDir, git.CloneOptions{
			URL: host.URL(templateURL),
		}); err != nil {
			exit.Error(fmt.Errorf("download: %s", err))
		}

		// Copy content from sub-folder to target folder
		if templateSubFolder != "" {
			// Ensure sub-folder exists
			templateTmpDir := osutil.JoinPaths(targetTmpDir, templateSubFolder)
			exists, err := osutil.DirExists(templateTmpDir)
			if err != nil {
				exit.Error(fmt.Errorf("download: %s", err))
			}
			if !exists {
				exit.Error(fmt.Errorf("download: sub-folder doesn't exist"))
			}
			// Check target folder exists, and copy contents
			if exists, err = osutil.DirExists(targetDir); err != nil {
				exit.Error(fmt.Errorf("download: %s", err))
			}
			if !exists {
				if err = osutil.CreateDirs(targetDir); err != nil {
					exit.Error(fmt.Errorf("download: %s", err))

				}
			}
			if err = osutil.CopyRecursively(templateTmpDir, targetDir); err != nil {
				exit.Error(fmt.Errorf("download: Error copying files from temp %s", err))
			}
			// Delete all temp files
			if err := os.RemoveAll(targetTmpDir); err != nil {
				exit.Error(fmt.Errorf("download: Error deleting temp files %s", err))
			}
		}
		// TODO(tmrts): use git-notes as metadata storage or boltdb
		if err := serializeMetadata(templateName, templateURL, targetDir); err != nil {
			exit.Error(fmt.Errorf("download: %s", err))
		}

		exit.OK("Successfully downloaded the template %#v", templateName)
	},
}
