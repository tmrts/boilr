package template

import "text/template"

var (
	FuncMap = template.FuncMap{}
	Options = []string{
		"missingkey=default",
	}

	BaseTemplate = template.New("base").Option(Options...).Funcs(FuncMap)
)
