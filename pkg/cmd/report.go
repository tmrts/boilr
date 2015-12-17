package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/google/go-github/github"
	cli "github.com/spf13/cobra"
	"github.com/tmrts/tmplt/pkg/tmplt"
	"github.com/tmrts/tmplt/pkg/util/exit"
	"github.com/tmrts/tmplt/pkg/util/validate"
)

type Transport struct {
	Username string
	Password string
}

func (t Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(t.Username, t.Password)
	return http.DefaultTransport.RoundTrip(req)
}

func (t *Transport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func readPassword() (Transport, error) {
	var name string
	fmt.Printf("Username for github: ")
	fmt.Scanf("%s", &name)

	fmt.Printf("Password for %s: ", name)
	pass, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return Transport{}, err
	}
	fmt.Println()

	return Transport{
		Username: name,
		Password: string(pass),
	}, nil
}

func getIssue() (*github.IssueRequest, error) {
	dir, err := ioutil.TempDir("", "tmplt-report")
	if err != nil {
		return nil, err
	} else {
		defer os.RemoveAll(dir)
	}

	f, err := os.Create(dir + "/issue.md")
	if err != nil {
		return nil, err
	}

	_, err = io.WriteString(f, "<!-- Replace with Issue Title -->\n\n<!-- Replace with Issue Body -->")
	if err != nil {
		return nil, err
	}
	f.Close()

	// TODO allow gedit, vi, emacs
	cmd := exec.Command("vim", dir+"/issue.md")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Start(); err != nil {
		return nil, err
	}

	if err = cmd.Wait(); err != nil {
		return nil, err
	}

	issueFile, err := os.Open(dir + "/issue.md")
	if err != nil {
		return nil, err
	} else {
		defer issueFile.Close()
	}

	buf, err := ioutil.ReadAll(issueFile)
	if err != nil {
		return nil, err
	}

	// TODO handle empty files
	slices := strings.SplitAfterN(string(buf), "\n", 3)

	title, body := slices[0], strings.Join(slices[1:], "\n")

	return &github.IssueRequest{
		Title: &title,
		Body:  &body,
	}, nil
}

func CreateIssue() (string, error) {
	req, err := getIssue()
	if err != nil {
		return "", err
	}

	t, err := readPassword()
	if err != nil {
		return "", err
	}

	client := github.NewClient(t.Client())
	issue, _, err := client.Issues.Create(tmplt.GithubOwner, tmplt.GithubRepo, req)
	if err != nil {
		return "", err
	}

	return *issue.HTMLURL, nil
}

var Report = &cli.Command{
	Use:   "report",
	Short: "Creates an issue in the github repository",
	Run: func(c *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{})

		url, err := CreateIssue()
		if err != nil {
			exit.Error(fmt.Errorf("Failed to create an issue, %v", err))
		}

		exit.OK("Successfully created an issue %v", url)
	},
}
