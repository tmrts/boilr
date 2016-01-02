package cmd

import (
	"fmt"

	"github.com/tmrts/boilr/pkg/boilr"
	"github.com/tmrts/boilr/pkg/cmd/util"
	"github.com/tmrts/boilr/pkg/util/exit"
	"github.com/tmrts/boilr/pkg/util/validate"
)

// MustValidateVarArgs validates given variadic arguments with the supplied validation function.
// If there are any errors it exits the execution.
func MustValidateVarArgs(args []string, v validate.Argument) {
	if err := util.ValidateVarArgs(args, v); err != nil {
		exit.Error(err)
	}
}

// MustValidateArgs validates given arguments with the supplied validation functions.
// If there are any errors it exits the execution.
func MustValidateArgs(args []string, validations []validate.Argument) {
	if err := util.ValidateArgs(args, validations); err != nil {
		exit.Error(err)
	}
}

// MustValidateTemplate validates a template given it's absolut path.
// If there are any errors it exits the execution.
func MustValidateTemplate(path string) {
	isValid, err := util.ValidateTemplate(path)
	if err != nil {
		exit.Fatal(fmt.Errorf("validate: %s", err))
	}

	// FIXME redundant
	if !isValid {
		exit.Fatal(fmt.Errorf("validate: %s", ErrTemplateInvalid))
	}
}

// MustValidateTemplateDir ensures that template directory is initialized.
func MustValidateTemplateDir() {
	isInitialized, err := boilr.IsTemplateDirInitialized()
	if err != nil {
		exit.Error(err)
	}

	if !isInitialized {
		exit.Error(fmt.Errorf("Template registry is not initialized. Please run `init` command to initialize it."))
	}
}
