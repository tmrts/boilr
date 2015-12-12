package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	cli "github.com/spf13/cobra"
	"github.com/tmrts/cookie/pkg/cookie"
	"github.com/tmrts/cookie/pkg/template"
	"github.com/tmrts/cookie/pkg/util/osutil"
)

var Use = &cli.Command{
	Use:   "use",
	Short: "Executes a project template",
	Run: func(_ *cli.Command, args []string) {
		tmplPath, err := cookie.TemplatePath(args[0])
		if err != nil {
			panic(err)
		}

		tmpl, err := template.Get(tmplPath)
		if err != nil {
			panic(err)
		}

		metadata := template.Metadata{
			Name:    "test-project-1",
			Author:  "test-author",
			Email:   "test@mail.com",
			Date:    time.Now().Format("Mon Jan 2 2006 15:04:05"),
			Version: "0.0.1",
		}

		err = tmpl.Execute(args[1], metadata)
		if err != nil {
			panic(err)
		}
	},
}

var Save = &cli.Command{
	Use:   "save",
	Short: "Saves a project template to template registry",
	Run: func(_ *cli.Command, args []string) {
		templateName, sourceDir := args[0], args[1]

		targetDir := filepath.Join(cookie.TemplateDirPath, templateName)

		switch err := osutil.FileExists(targetDir); {
		case !os.IsNotExist(err):
			// Template Already Exists Ask If Should be Replaced
			panic(err)
		}

		if _, err := exec.Command("/usr/bin/cp", "-r", sourceDir, targetDir).Output(); err != nil {
			fmt.Println(sourceDir, targetDir)
			panic(err)
		}
	},
}
