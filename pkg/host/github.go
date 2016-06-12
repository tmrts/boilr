package host

import "strings"

// ZipURL returns the URL of the zip archive given a github repository URL.
func ZipURL(url string) string {
	var version = "master"

	url = strings.TrimSuffix(url, "/")
	url = strings.TrimPrefix(url, "/")

	if strings.HasSuffix(url, "zip/master") {
		return url
	}

	// So this could identify a port number, but since we only support github
	// I don't believe using it as a version modifier is a problem. Though
	// perhaps we should use something else instead?
	if strings.Contains(url, ":") {
		parts := strings.SplitAfter(url, ":")

		url = parts[0]
		version = parts[1]

		url = strings.TrimSuffix(url, ":")
	}

	// BUG filepath.Join trims slashes use url.Join
	return "https://codeload.github.com/" + url + "/zip/" + version
}
