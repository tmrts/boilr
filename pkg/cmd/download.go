package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/tmplt/pkg/cmd/util"
	"github.com/tmrts/tmplt/pkg/host"
	"github.com/tmrts/tmplt/pkg/tmplt"
	"github.com/tmrts/tmplt/pkg/util/exit"
	"github.com/tmrts/tmplt/pkg/util/osutil"
	"github.com/tmrts/tmplt/pkg/util/validate"
)

func downloadZip(URL, targetDir string) error {
	f, err := ioutil.TempFile("", "tmplt-download")
	if err != nil {
		return err
	} else {
		defer f.Close()
		defer os.RemoveAll(f.Name())
	}

	resp, err := http.Get(URL)
	if err != nil {
		return err
	} else {
		defer resp.Body.Close()
	}

	if _, err := io.Copy(f, resp.Body); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	r, err := zip.OpenReader(f.Name())
	if err != nil {
		return err
	} else {
		defer r.Close()
	}

	extractFile := func(f *zip.File, dest string) error {
		rc, err := f.Open()
		if err != nil {
			return err
		} else {
			defer rc.Close()
		}

		// split first token of f.Name since it's zip file name
		path := filepath.Join(dest, strings.SplitAfterN(f.Name, "/", 2)[1])

		if f.FileInfo().IsDir() {
			err := os.MkdirAll(path, f.Mode())
			if err != nil {
				return err
			}
		} else {
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			} else {
				defer f.Close()
			}

			if _, err := io.Copy(f, rc); err != nil {
				return err
			}
		}

		return nil
	}

	for _, f := range r.File {
		if err := extractFile(f, targetDir); err != nil {
			return err
		}
	}

	_, err := util.ValidateTemplate(path)
	if err != nil {
		return err
	}

	return nil
}

// FIXME Half-Updates leave messy templates
var Download = &cli.Command{
	Use:   "download <template-repo> <template-name>",
	Short: "Download a project template from a github repository to template registry",
	Run: func(c *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{
			{"template-repo", validate.UnixPath},
			{"template-name", validate.Alphanumeric},
		})

		templateURL, templateName := args[0], args[1]

		targetDir, err := tmplt.TemplatePath(templateName)
		if err != nil {
			exit.Error(fmt.Errorf("download: %s", err))
		}

		switch exists, err := osutil.DirExists(targetDir); {
		case err != nil:
			exit.Error(fmt.Errorf("download: %s", err))
		case exists:
			if shouldOverwrite := GetBoolFlag(c, "force"); !shouldOverwrite {
				exit.OK("Template %v already exists use -u to update the template", templateName)
			}
		case !exists:
			if err := os.MkdirAll(targetDir, 0755); err != nil {
				exit.Error(fmt.Errorf("download: %s", err))
			}
		}

		/*
		 *if !strings.Contains(templateURL, "github.com") {
		 *    exit.Error(fmt.Errorf("download only supports project templates hosted on github at the moment"))
		 *}
		 */

		zipURL := host.ZipURL(templateURL)

		// TODO validate template as well
		if err := downloadZip(zipURL, targetDir); err != nil {
			// Delete if download transaction fails
			defer os.RemoveAll(targetDir)

			exit.Error(fmt.Errorf("download: %s", err))
		}

		exit.OK("Successfully downloaded the template %v", templateName)
	},
}
