package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	cli "github.com/spf13/cobra"

	"github.com/tmrts/boilr/pkg/boilr"
	"github.com/tmrts/boilr/pkg/template"
	"github.com/tmrts/boilr/pkg/util/exit"
	"github.com/tmrts/boilr/pkg/util/validate"
)

// TemplateInRegistry checks whether the given name exists in the template registry.
func TemplateInRegistry(name string) (bool, error) {
	names, err := ListTemplates()
	if err != nil {
		return false, err
	}

	_, ok := names[name]
	return ok, nil
}

// TODO add --use-cache flag to execute a template from previous answers to prompts
// Use contains the cli-command for using templates located in the local template registry.
var Use = &cli.Command{
	Use:   "use <template-tag> <target-dir>",
	Short: "Execute a project template in the given directory",
	Run: func(cmd *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{
			{"template-tag", validate.UnixPath},
			{"target-dir", validate.UnixPath},
		})

		MustValidateTemplateDir()

		tmplName := args[0]
		targetDir, err := filepath.Abs(args[1])
		if err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		templateFound, err := TemplateInRegistry(tmplName)
		if err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		if !templateFound {
			exit.Fatal(fmt.Errorf("Template %q couldn't be found in the template registry", tmplName))
		}

		tmplPath, err := boilr.TemplatePath(tmplName)
		if err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		tmpl, err := template.Get(tmplPath)
		if err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		if shouldUseDefaults := GetBoolFlag(cmd, "use-defaults"); shouldUseDefaults {
			tmpl.UseDefaultValues()
		}

		tmpDir, err := ioutil.TempDir("", "boilr-use-template")
		if err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		}
		defer os.RemoveAll(tmpDir)

		if err := os.Mkdir(targetDir, 0744); err != nil {
			if os.IsNotExist(err) {
				exit.Fatal(fmt.Errorf("use: directory %q doesn't exist", filepath.Dir(targetDir)))
			}

			if !os.IsExist(err) {
				exit.Fatal(fmt.Errorf("use: %s", err))
			}
		}

		if err := tmpl.Execute(tmpDir); err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		// Complete the template execution transaction by copying the temporary dir to
		// the target directory.
		if err := filepath.Walk(tmpDir, func(fname string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			relPath, err := filepath.Rel(tmpDir, fname)
			if err != nil {
				return err
			}

			mirrorPath := filepath.Join(targetDir, relPath)

			if info.IsDir() {
				if err := os.Mkdir(mirrorPath, 0744); err != nil {
					if !os.IsExist(err) {
						return err
					}
				}
			} else {
				fi, err := os.Lstat(fname)
				if err != nil {
					return err
				}

				tmpf, err := os.Open(fname)
				if err != nil {
					return err
				}
				defer tmpf.Close()

				f, err := os.OpenFile(mirrorPath, os.O_CREATE|os.O_WRONLY, fi.Mode())
				if err != nil {
					return err
				}
				defer f.Close()

				if _, err := io.Copy(f, tmpf); err != nil {
					return err
				}
			}

			return nil
		}); err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		exit.OK("Successfully executed the project template %v in %v", tmplName, targetDir)
	},
}
