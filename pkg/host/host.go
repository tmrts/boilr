package host

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/howeyc/gopass"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"

	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

var (
	scpLikeUrlRegExp   = regexp.MustCompile("^(ssh://)?(?:(?P<user>[^@]+)(?:@))?(?P<host>[^:|/]+)(?::|/)?(?P<path>.+)$")
	isHttpSchemeRegExp = regexp.MustCompile("^(http|https)://")
)

func IsRepoURL(url string) bool {
	return isHttpSchemeRegExp.MatchString(url) || scpLikeUrlRegExp.MatchString(url)
}

func IsSSH(url string) bool {
	if !isHttpSchemeRegExp.MatchString(url) && scpLikeUrlRegExp.MatchString(url) {
		return true
	}

	return false
}

func AuthMethodForURL(url string) transport.AuthMethod {

	if IsSSH(url) {
		// 1:scheme, 2:user, 3:host, 4:path
		m := scpLikeUrlRegExp.FindStringSubmatch(url)
		a, _ := ssh.NewSSHAgentAuth(m[2])
		return a
	}

	ba := http.NewBasicAuth("", "")
	ba.CredentialsProvider = promptForCredentials

	return ba
}

func promptForCredentials() (string, string) {
	var u = ""
	var p = ""

	uname, err := promptForUsername()

	if err == nil {
		u = uname
	}

	pbytes, err := gopass.GetPasswdPrompt("password: ", true, os.Stdin, os.Stdout)

	if err == nil {
		p = string(pbytes)
	}

	return u, p
}

func promptForUsername() (string, error) {
	consolereader := bufio.NewReader(os.Stdin)

	fmt.Print("username: ")
	response, err := consolereader.ReadString('\n')

	if err != nil {
		return "", err
	}

	if len(response) < 1 {
		return promptForUsername()
	}

	return strings.TrimSuffix(response, "\n"), nil
}
