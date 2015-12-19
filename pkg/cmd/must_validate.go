package cmd

import (
	"fmt"

	"github.com/tmrts/boilr/pkg/cmd/util"
	"github.com/tmrts/boilr/pkg/util/exit"
	"github.com/tmrts/boilr/pkg/util/validate"
)

func MustValidateVarArgs(args []string, v validate.Argument) {
	if err := util.ValidateVarArgs(args, v); err != nil {
		exit.Error(err)
	}
}

func MustValidateArgs(args []string, validations []validate.Argument) {
	if err := util.ValidateArgs(args, validations); err != nil {
		exit.Error(err)
	}
}

// TODO use defaults option while executing for validation
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
