package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tmrts/boilr/pkg/util/tlog"
)

// TODO align brackets used in the prompt message
const (
	// PromptFormatMessage is a format message for value prompts.
	PromptFormatMessage = "[?] Please choose a value for %#q [default: %#v]: "

	// PromptChoiceFormatMessage is a format message for choice prompts.
	PromptChoiceFormatMessage = "[?] Please choose an option for %#q\n%v    Select from %v..%v [default: %#v]: "
)

func scanLine() (string, error) {
	input := bufio.NewReader(os.Stdin)
	line, err := input.ReadString('\n')
	if err != nil {
		return line, err
	}

	return strings.TrimSuffix(line, "\n"), nil
}

// TODO add GetLine method using a channel
// TODO use interfaces to eliminate code duplication
func newString(name string, defval interface{}) func() interface{} {
	var cache interface{}
	return func() interface{} {
		if cache == nil {
			cache = func() interface{} {
				// TODO use colored prompts
				fmt.Printf(PromptFormatMessage, name, defval)

				line, err := scanLine()
				if err != nil {
					tlog.Warn(err.Error())
					return line
				}

				if line == "" {
					return defval
				}

				return line
			}()
		}

		return cache
	}
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

func newBool(name string, defval bool) func() interface{} {
	var cache interface{}
	return func() interface{} {
		if cache == nil {
			cache = func() interface{} {
				fmt.Printf(PromptFormatMessage, name, defval)

				choice, err := scanLine()
				if err != nil {
					tlog.Warn(err.Error())
					return choice
				}

				if choice == "" {
					return defval
				}

				val, ok := booleanValues[strings.ToLower(choice)]
				if !ok {
					tlog.Warn(fmt.Sprintf("Unrecognized choice %q, using the default", choice))

					return defval
				}

				return val
			}()
		}

		return cache
	}
}

// Choice contains the values for a choice
type Choice struct {
	// Default choice
	Default int

	// List of choices
	Choices []string
}

func formattedChoices(cs []string) (s string) {
	for i, c := range cs {
		s += fmt.Sprintf("    %v -  %q\n", i+1, c)
	}

	return
}

func newSlice(name string, choices []string) func() interface{} {
	var cache interface{}
	return func() interface{} {
		if cache == nil {
			defindex := 0
			defval := choices[defindex]
			cache = func() interface{} {
				s := formattedChoices(choices)
				fmt.Printf(PromptChoiceFormatMessage, name, s, 1, len(choices), defindex+1)

				choice, err := scanLine()
				if err != nil {
					tlog.Warn(err.Error())
					return choice
				}

				if choice == "" {
					return defval
				}

				index, err := strconv.Atoi(choice)
				if err != nil {
					return err
				}

				if index > len(choices)+1 || index < 1 {
					tlog.Warn(fmt.Sprintf("Unrecognized choice %v, using the default", index))

					return defval
				}

				return choices[index-1]
			}()
		}

		return cache
	}
}

// New returns a prompt closure when executed asks for
// user input and has a default value that returns result.
func New(name string, defval interface{}) func() interface{} {
	// TODO use reflect package
	// TODO add a prompt as such "How many Items will you enter", "Enter each" use in "{{range Items}}"
	switch defval := defval.(type) {
	case bool:
		return newBool(name, defval)
	case []interface{}:
		if len(defval) == 0 {
			tlog.Warn(fmt.Sprintf("empty list for %q choices", name))
			return nil
		}

		var s []string
		for _, v := range defval {
			s = append(s, v.(string))
		}

		return newSlice(name, s)
	}

	return newString(name, defval)
}
