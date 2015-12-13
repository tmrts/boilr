package host

import "strings"

func ZipURL(url string) string {
	url = strings.TrimSuffix(url, "/")
	url = strings.TrimPrefix(url, "/")

	if strings.HasSuffix(url, "zip/master") {
		return url
	}

	return "https://codeload.github.com/" + url + "/zip/master"
}
