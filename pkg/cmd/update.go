package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/boilr/pkg/boilr"
	"github.com/tmrts/boilr/pkg/util/exit"
	"github.com/tmrts/boilr/pkg/util/git"
	"github.com/tmrts/boilr/pkg/util/osutil"
	"github.com/tmrts/boilr/pkg/util/validate"
	gogit "gopkg.in/src-d/go-git.v4"
)

// CLI command for updating a template
var Update = &cli.Command{
	Use:   "update <template-tag>",
	Short: "Update a project template from a github repository",
	Run: func(c *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{
			{"template-tag", validate.AlphanumericExt},
		})

		MustValidateTemplateDir()

		templateName := args[0]

		targetDir, err := boilr.TemplatePath(templateName)
		if err != nil {
			exit.Error(fmt.Errorf("download: %s", err))
		}

		switch exists, err := osutil.DirExists(targetDir); {
		case err != nil:
			exit.Error(fmt.Errorf("updatee: %s", err))
		case !exists:
      exit.OK("Couldn't find '%s' template to update. Have you downloaded it?", templateName)
		}

    // We don't mind if this fails, as if the file doesn't exist; we can update the repo.
    // Further work to be carried out with regards to metadata
    // See: https://github.com/tmrts/boilr/pull/43#issuecomment-289391044<Paste>
    metadataPath := filepath.Join(targetDir, boilr.TemplateMetadataName)
    os.Remove(metadataPath)

    repository, err := git.Open(targetDir)
    if err != nil {
      exit.Error(fmt.Errorf("Failed to update repository '%s': %s", targetDir, err))
    }

    if err := repository.Pull(&gogit.PullOptions{RemoteName: "origin"}); err != nil {
      exit.Error(fmt.Errorf("Failed to update repository '%s': %s", targetDir, err))
    }

		exit.OK("Successfully updated the template %#v", templateName)
	},
}
