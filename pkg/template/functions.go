package template

import (
	"os"
	"text/template"
	"time"

	"github.com/tmrts/tmplt/pkg/prompt"
)

var (
	FuncMap = template.FuncMap{
		"env":     os.Getenv,
		"now":     CurrentTime,
		"ask":     prompt.Ask,
		"confirm": prompt.Confirm,
	}

	Options = []string{
		"missingkey=default",
	}

	//BaseTemplate = template.New("base").Option(Options...).Funcs(FuncMap)
)

func CurrentTime(fmt string) string {
	t := time.Now()

	return t.Format(fmt)
}
