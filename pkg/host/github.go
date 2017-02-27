package host

import (
	"regexp"
	"strings"
)

const (
	githubURL        = "https://github.com"
	githubStorageURL = "https://codeload.github.com"
)

// ZipURL returns the URL of the zip archive given a github repository URL.
func ZipURL(repo string) string {
	var version = "master"

	repo = strings.TrimSuffix(strings.TrimPrefix(repo, "/"), "/")

	zipRegex := regexp.MustCompile(`zip/(\S+)$`)
	if zipRegex.MatchString(repo) {
		return repo
	}

	// FIXME(tmrts): this check could also identify a port number, but since
	// we only support github I don't believe using it as a version modifier
	// is a problem. Perhaps we should reconsider?
	if strings.Contains(repo, ":") {
		parts := strings.SplitAfter(repo, ":")

		repo = parts[0]
		version = parts[1]

		repo = strings.TrimSuffix(repo, ":")
	}

	urlTokens := []string{githubStorageURL, repo, "zip", version}

	return strings.Join(urlTokens, "/")
}

// URL returns the normalized URL of a GitHub repository.
func URL(repo string) string {
	githubRegex := regexp.MustCompile(githubURL + `/(\S+)$`)
	if githubRegex.MatchString(repo) {
		return repo
	}

	return strings.Join([]string{githubURL, repo}, "/")
}
