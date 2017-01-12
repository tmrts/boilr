package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tmrts/boilr/pkg/util/tlog"
)

type templateFunc func() interface{}

type Interface interface {
	// PromptMessage returns a proper prompt message for the given field with the given default value.
	PromptMessage(string) string
	EvaluateChoice(string) (interface{}, error)
}

type Chain struct {
	Prompts []Interface
}

type strPrompt string

func (p strPrompt) PromptMessage(name string) string {
	return fmt.Sprintf("Please choose a value for %q", name)
}

func (p strPrompt) EvaluateChoice(c string) (interface{}, error) {
	if c != "" {
		return c, nil
	}

	return string(p), nil
}

type boolPrompt bool

func (p boolPrompt) PromptMessage(name string) string {
	return fmt.Sprintf("Please choose a value for %q", name)
}

var (
	booleanValues = map[string]bool{
		"y":    true,
		"yes":  true,
		"yup":  true,
		"true": true,

		"n":     false,
		"no":    false,
		"nope":  false,
		"false": false,
	}
)

func (p boolPrompt) EvaluateChoice(c string) (interface{}, error) {
	if val, ok := booleanValues[c]; ok {
		return val, nil
	}

	return bool(p), nil
}

// TODO: add proper format messages for multiple choices
type multipleChoicePrompt []string

func (p multipleChoicePrompt) PromptMessage(name string) string {
	return fmt.Sprintf("Please choose an option for %q", name)
}

func (p multipleChoicePrompt) EvaluateChoice(c string) (interface{}, error) {
	if c != "" {
		index, err := strconv.Atoi(c)
		if err != nil || index < 1 || index > len(p) {
			tlog.Warn(fmt.Sprintf("Unrecognized choice %v, using the default choice", index))

			return p[0], nil
		}

		return p[index-1], nil
	}

	return p[0], nil
}

// TODO add deep pretty printer
// TODO handle TOML
func Func(defval interface{}) Interface {
	switch defval := defval.(type) {
	case bool:
		return boolPrompt(defval)
	case []interface{}:
		if len(defval) == 0 {
			tlog.Warn(fmt.Sprintf("empty list of choices"))
			return nil
		}

		var s []string
		for _, v := range defval {
			s = append(s, fmt.Sprint(v))
		}

		return multipleChoicePrompt(s)
	}

	return strPrompt(fmt.Sprint(defval))
}

func scanLine() (string, error) {
	input := bufio.NewReader(os.Stdin)
	line, err := input.ReadString('\n')
	if err != nil {
		return line, err
	}

	return strings.TrimSuffix(line, "\n"), nil
}

// New returns a prompt closure when executed asks for
// user input once and caches it for further invocations
// and has a default value that returns result.
func New(fieldName string, defval interface{}) func() interface{} {
	prompt := Func(defval)

	var cachedValue interface{}
	return func() interface{} {
		if cachedValue == nil {
			msg := prompt.PromptMessage(fieldName)

			tlog.Prompt(msg, defval)

			choice, err := scanLine()
			if err != nil {
				tlog.Warn(err.Error())
			}

			cachedValue, err = prompt.EvaluateChoice(choice)
			if err != nil {
				tlog.Warn(err.Error())
			}
		}

		return cachedValue
	}
}
