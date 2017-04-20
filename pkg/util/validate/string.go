package validate

import (
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"github.com/tmrts/boilr/pkg/util/validate/pattern"
)

var (
	repoUrlRegExp = regexp.MustCompile("^(ssh://|https://)?(?:(?P<user>[^@]+)(?:@))?(?P<host>[^:|/]+)(?::|/)?(?P<path>.+)$")
)

// String is the validation function used for checking whether a string is valid or not.
type String func(string) bool

// TypeName returns the type expected to be validated by the validate.String function.
func (s String) TypeName() string {
	fullPath := runtime.FuncForPC(reflect.ValueOf(s).Pointer()).Name()

	return strings.ToLower(fullPath[strings.LastIndex(fullPath, ".")+1:])
}

// Integer validates whether a string is an integer string.
func Integer(n string) bool {
	return pattern.Integer.MatchString(n)
}

// URL validates whether a string is an URL string.
func URL(url string) bool {
	return pattern.URL.MatchString(url)
}

// UnixPath validates whether a string is an unix path string.
func UnixPath(path string) bool {
	return pattern.UnixPath.MatchString(path)
}

// Alphanumeric validates whether a string is an alphanumeric string.
func Alphanumeric(s string) bool {
	return pattern.Alphanumeric.MatchString(s)
}

// AlphanumericExt validates whether a string is an alphanumeric string, but a
// small set of extra characters allowed
func AlphanumericExt(s string) bool {
	return pattern.AlphanumericExt.MatchString(s)
}

// RepoURL validates whether a string is a valid repo url.
func RepoURL(s string) bool {
	isurl := false
	if repoUrlRegExp.MatchString(s) {
		isurl = true
		m := repoUrlRegExp.FindStringSubmatch(s)

		//bad scheme
		if m[1] == "http://" {
			isurl = false
		}

		//blank host
		if len(strings.TrimSpace(m[3])) < 1 {
			isurl = false
		}

		//blank path
		if len(strings.TrimSpace(m[4])) < 1 {
			isurl = false
		}
	}

	return isurl
}
