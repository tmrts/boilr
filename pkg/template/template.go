package template

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/tmrts/tmplt/pkg/prompt"
	"github.com/tmrts/tmplt/pkg/tmplt"
	"github.com/tmrts/tmplt/pkg/util/stringutil"
)

type Interface interface {
	Execute(string) error
	UseDefaultValues()
}

func Get(path string) (Interface, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	// TODO make context optional
	ctxt, err := func(fname string) (map[string]interface{}, error) {
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
	}(filepath.Join(absPath, tmplt.ContextFileName))

	return &dirTemplate{
		Context: ctxt,
		FuncMap: FuncMap,
		Path:    filepath.Join(absPath, tmplt.TemplateDirName),
	}, err
}

type dirTemplate struct {
	Path    string
	Context map[string]interface{}
	FuncMap template.FuncMap

	ShouldUseDefaults bool
}

func (t *dirTemplate) UseDefaultValues() {
	t.ShouldUseDefaults = true
}

func (t *dirTemplate) BindPrompts() {
	if t.ShouldUseDefaults {
		for s, v := range t.Context {
			t.FuncMap[s] = v
		}
	} else {
		for s, v := range t.Context {
			t.FuncMap[s] = prompt.New(s, v)
		}
	}
}

// Execute fills the template with the project metadata.
func (t *dirTemplate) Execute(dirPrefix string) error {
	t.BindPrompts()

	// TODO create io.ReadWriter from string
	// TODO refactor name manipulation
	return filepath.Walk(t.Path, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Path relative to the root of the template directory
		oldName, err := filepath.Rel(t.Path, filename)
		if err != nil {
			return err
		}

		buf := stringutil.NewString("")

		// TODO translate errors into meaningful ones
		fnameTmpl := template.Must(template.
			New("file name template").
			Option(Options...).
			Funcs(FuncMap).
			Parse(oldName))

		if err := fnameTmpl.Execute(buf, nil); err != nil {
			return err
		}

		newName := buf.String()

		target := filepath.Join(dirPrefix, newName)

		if info.IsDir() {
			if err := os.MkdirAll(target, 0744); err != nil {
				return err
			}
		} else {
			fi, err := os.Lstat(filename)
			if err != nil {
				return err
			}

			// Delete target file if it exists
			if err := os.Remove(target); err != nil && !os.IsNotExist(err) {
				return err
			}

			f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, fi.Mode())
			if err != nil {
				return err
			} else {
				defer f.Close()
			}

			contentsTmpl := template.Must(template.
				New("file contents template").
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
