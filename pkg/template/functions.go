package template

import (
	"os"
	"strings"
	"text/template"
	"time"
)

var (
	FuncMap = template.FuncMap{
		// TODO confirmation prompt
		// TODO value prompt
		// TODO encoding utilities (e.g. toBinary)
		// TODO GET, POST utilities
		// TODO Hostname(Also accesible through $HOSTNAME), IP addr, etc.
		// TODO add validate for custom regex and expose validate package
		"env":      os.Getenv,
		"time":     CurrentTimeInFmt,
		"hostname": func() string { return os.Getenv("HOSTNAME") },

		// String utilities
		"toLower": strings.ToLower,
		"toUpper": strings.ToUpper,
		"toTitle": strings.ToTitle,

		"trimSpace": strings.TrimSpace,

		"repeat": strings.Repeat,
	}

	Options = []string{
		// TODO ignore a field if no value is found instead of writing <no value>
		"missingkey=invalid",
	}
)

func CurrentTimeInFmt(fmt string) string {
	t := time.Now()

	return t.Format(fmt)
}
