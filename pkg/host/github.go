package host

import (
	"regexp"
	"strings"
)

const githubStorageURL = "https://codeload.github.com"

// ZipURL returns the URL of the zip archive given a github repository URL.
func ZipURL(repo string) string {
	var version = "master"

	repo = strings.TrimSuffix(strings.TrimPrefix(repo, "/"), "/")

	zipRegex, _ := regexp.Compile(`zip/(\S+)$`)
	if zipRegex.MatchString(repo) {
		return repo
	}

	// So this could identify a port number, but since we only support github
	// I don't believe using it as a version modifier is a problem. Though
	// perhaps we should use something else instead?
	if strings.Contains(repo, ":") {
		parts := strings.SplitAfter(repo, ":")

		repo = parts[0]
		version = parts[1]

		repo = strings.TrimSuffix(repo, ":")
	}

	urlTokens := []string{githubStorageURL, repo, "zip", version}

	return strings.Join(urlTokens, "/")
}
