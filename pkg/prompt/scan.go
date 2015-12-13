package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tmrts/tmplt/pkg/util/tlog"
)

func Ask(msg string) (value string) {
	fmt.Print(msg)

	_, err := fmt.Scanf("%s", &value)
	if err != nil {
		tlog.Warn(err.Error())
		return ""
	}

	return
}

var (
	booleanValues = map[string]bool{
		"y":   true,
		"yes": true,
		"yup": true,

		"n":    false,
		"no":   false,
		"nope": false,
	}
)

func Confirm(msg string) bool {
	fmt.Print(msg)

	var choice string
	_, err := fmt.Scanf("%s", &choice)
	if err != nil {
		tlog.Warn(err.Error())
		return false
	}

	val, ok := booleanValues[choice]
	if !ok {
		tlog.Warn(fmt.Sprintf("unrecognized choice %#q", choice))
		return false
	}

	return val
}

const (
	PromptFormatMessage = "? Please choose a value for %#q [default: %#q]: "
)

// TODO accept boolean, integer values in addition to string
func New(msg, defval string) func() string {
	return func() string {
		// TODO use colored prompts
		fmt.Printf(PromptFormatMessage, msg, defval)

		input := bufio.NewReader(os.Stdin)
		line, err := input.ReadString('\n')
		if err != nil {
			tlog.Warn(err.Error())
			return line
		}

		if line == "\n" {
			return defval
		}

		return strings.TrimSuffix(line, "\n")
	}
}
