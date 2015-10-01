package template

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

// ?
type Metadata struct {
	Name    string
	Author  string
	Email   string
	Date    string
	Version string
}

type Interface interface {
	Execute(string, Metadata) error
}

func Get(path string) (Interface, error) {
	return &dirTemplate{Path: path}, nil
}

type dirTemplate struct {
	Path string
}

// Execute fills the template with the given metadata.
func (d *dirTemplate) Execute(dirPrefix string, md Metadata) error {
	fill := func(filepath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		oldName := info.Name()

		// BUG:
		buf := bytes.NewBuffer([]byte{})

		tmpl := template.Must(template.New("x").Parse(oldName))

		if err := tmpl.Execute(buf, md); err != nil {
			return err
		}
		newName := buf.String()

		fmt.Print(oldName, "\n")
		fmt.Print(newName, "\n")
		if oldName != newName {
			if err := os.Rename(filepath, newName); err != nil {
				return err
			}
		}

		if !info.IsDir() && false {
			f, err := os.Open(filepath)
			if err != nil {
				return err
			} else {
				defer f.Close()
			}
		}

		return nil
	}

	// TODO: {{Project.Name}} vs app/etc.
	if _, err := exec.Command("/bin/cp", "-r", d.Path, filepath.Join(dirPrefix, "test")).Output(); err != nil {
		return err
	}

	return filepath.Walk(filepath.Join(dirPrefix, "test"), fill)
}
