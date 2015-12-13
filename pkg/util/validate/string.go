package validate

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/tmrts/tmplt/pkg/util/validate/pattern"
)

type String func(string) bool

// TypeName returns the type expected to be validated by the validate.String function.
func (s String) TypeName() string {
	fullPath := runtime.FuncForPC(reflect.ValueOf(s).Pointer()).Name()

	return strings.ToLower(fullPath[strings.LastIndex(fullPath, ".")+1:])
}

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
