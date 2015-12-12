package util

import (
	"errors"
	"fmt"

	"github.com/tmrts/tmplt/pkg/util/validate"
)

var (
	ErrUnexpectedArgs = errors.New("unexpected arguments")
	ErrNotEnoughArgs  = errors.New("not enough arguments")
	ErrInvalidArgType = errors.New("invalid argument type")
	ErrInvalidArg     = errors.New("invalid argument")
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

func ValidateArgs(args []string, validations []validate.String) []error {
	if err := ValidateArgCount(len(validations), len(args)); err != nil {
		return []error{err}
	}

	var errors []error
	for i, validate := range validations {
		if ok := validate(args[i]); !ok {
			errors = append(errors, fmt.Errorf("%v %q", ErrInvalidArg, args[i]))
		}
	}

	return errors
}
