package host

import (
	"path/filepath"
	"strings"
)

// ZipURL returns the URL of the zip archive given a github repository URL.
func ZipURL(url string) string {
	url = strings.TrimSuffix(strings.TrimPrefix(url, "/"), "/")

	if strings.HasSuffix(url, "zip/master") {
		return url
	}

	return "https://codeload.github.com/" + filepath.Join(url, "/zip/master")
}
