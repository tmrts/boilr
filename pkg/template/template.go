package template

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/tmrts/tmplt/pkg/util/stringutil"
)

// TODO Use JSON Dictionary
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
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	md, err := func(fname string) (map[string]interface{}, error) {
		f, err := os.Open(fname)
		if err != nil {
			if os.IsNotExist(err) {
				return nil, nil
			}

			return nil, err
		} else {
			defer f.Close()
		}

		buf, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}

		var metadata map[string]interface{}
		if err := json.Unmarshal(buf, &metadata); err != nil {
			return nil, err
		}

		return metadata, nil
	}(filepath.Join(absPath, "metadata.json"))

	return &dirTemplate{Path: absPath, Metadata: md}, err
}

type dirTemplate struct {
	Path     string
	Metadata map[string]interface{}
}

// Execute fills the template with the given metadata.
func (d *dirTemplate) Execute(dirPrefix string, md Metadata) error {
	FuncMap["Project"] = func() map[string]interface{} {
		return d.Metadata
	}

	// TODO(tmrts): create io.ReadWriter from string
	// TODO(tmrts): refactor command execution
	// TODO(tmrts): refactor name manipulation
	return filepath.Walk(d.Path, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Path relative to the root of the template directory
		oldName, err := filepath.Rel(filepath.Dir(d.Path), filename)
		if err != nil {
			return err
		}

		buf := stringutil.NewString("")

		fnameTmpl := template.Must(template.
			New("filename").
			Option(Options...).
			Funcs(FuncMap).
			Parse(oldName))

		if err := fnameTmpl.Execute(buf, nil); err != nil {
			return err
		}

		newName := buf.String()

		target := filepath.Join(dirPrefix, newName)

		if info.IsDir() {
			if _, err := exec.Command("/bin/mkdir", target).Output(); err != nil {
				fmt.Println(target)
				return err
			}
		} else {
			fi, err := os.Lstat(filename)
			if err != nil {
				return err
			}

			f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, fi.Mode())
			if err != nil {
				return err
			} else {
				defer f.Close()
			}

			contentsTmpl := template.Must(template.
				New("filecontents").
				Option(Options...).
				Funcs(FuncMap).
				ParseFiles(filename))

			fileTemplateName := filepath.Base(filename)

			if err := contentsTmpl.ExecuteTemplate(f, fileTemplateName, nil); err != nil {
				return err
			}
		}

		return nil
	})
}
