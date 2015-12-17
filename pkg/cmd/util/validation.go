package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tmrts/tmplt/pkg/template"
	"github.com/tmrts/tmplt/pkg/util/osutil"
	"github.com/tmrts/tmplt/pkg/util/validate"
)

var (
	ErrUnexpectedArgs = errors.New("unexpected arguments")
	ErrNotEnoughArgs  = errors.New("not enough arguments")
)

const (
	// Error message format string for filling in the details of an invalid arg.
	InvalidArg = "invalid argument for %q: %q, should be a valid %v"
)

func ValidateArgCount(expectedArgNo, argNo int) error {
	switch {
	case expectedArgNo < argNo:
		return ErrUnexpectedArgs
	case expectedArgNo > argNo:
		return ErrNotEnoughArgs
	case expectedArgNo == argNo:
	}

	return nil
}

func ValidateArgs(args []string, validations []validate.Argument) error {
	if err := ValidateArgCount(len(validations), len(args)); err != nil {
		return err
	}

	for i, arg := range validations {
		if ok := arg.Validate(args[i]); !ok {
			return fmt.Errorf(InvalidArg, arg.Name, args[i], arg.Validate.TypeName())
		}
	}

	return nil
}

func testTemplate(path string) error {
	tmpDir, err := ioutil.TempDir("", "tmplt-validation-test")
	if err != nil {
		return err
	} else {
		defer os.RemoveAll(tmpDir)
	}

	tmpl, err := template.Get(path)
	if err != nil {
		return err
	}

	tmpl.UseDefaultValues()

	return tmpl.Execute(tmpDir)
}

func ValidateTemplate(tmplPath string) (bool, error) {
	if exists, err := osutil.DirExists(tmplPath); !exists {
		if err != nil {
			return false, err
		}

		return false, fmt.Errorf("template directory not found")
	}

	if exists, err := osutil.DirExists(filepath.Join(tmplPath, "template")); !exists {
		if err != nil {
			return false, err
		}

		return false, fmt.Errorf("template should contain %q directory", "template")
	}

	if err := testTemplate(tmplPath); err != nil {
		return false, err
	}

	return true, nil
}
