package cmd

import (
	"fmt"

	"github.com/tmrts/tmplt/pkg/cmd/util"
	"github.com/tmrts/tmplt/pkg/util/exit"
	"github.com/tmrts/tmplt/pkg/util/validate"
)

func MustValidateArgs(args []string, validations []validate.Argument) {
	if err := util.ValidateArgs(args, validations); err != nil {
		exit.Error(err)
	}
}

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
