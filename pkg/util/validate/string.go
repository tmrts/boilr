package validate

import "github.com/tmrts/tmplt/pkg/util/validate/pattern"

type String func(string) bool

func Integer(n string) bool {
	return pattern.Integer.MatchString(n)
}

func URL(url string) bool {
	return pattern.URL.MatchString(url)
}

func UnixPath(path string) bool {
	return pattern.UnixPath.MatchString(path)
}

func Alphanumeric(s string) bool {
	return pattern.Alphanumeric.MatchString(s)
}
