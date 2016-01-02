package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/google/go-github/github"
	cli "github.com/spf13/cobra"
	"github.com/tmrts/boilr/pkg/boilr"
	"github.com/tmrts/boilr/pkg/util/exit"
	"github.com/tmrts/boilr/pkg/util/validate"
)

type transport struct {
	Username string
	Password string
}

func (t transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(t.Username, t.Password)
	return http.DefaultTransport.RoundTrip(req)
}

func (t *transport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func readPassword() (transport, error) {
	var name string
	fmt.Printf("Username for github: ")
	fmt.Scanf("%s", &name)

	fmt.Printf("Password for %s: ", name)
	pass, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return transport{}, err
	}
	fmt.Println()

	return transport{
		Username: name,
		Password: string(pass),
	}, nil
}

func getIssue() (*github.IssueRequest, error) {
	dir, err := ioutil.TempDir("", "boilr-report")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)

	fname := filepath.Join(dir, "issue.markdown")

	f, err := os.Create(fname)
	if err != nil {
		return nil, err
	}

	_, err = io.WriteString(f, "<!-- Replace with Issue Title -->\n\n<!-- Replace with Issue Body -->")
	if err != nil {
		return nil, err
	}
	f.Close()

	cmd := exec.Command("vi", fname)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Start(); err != nil {
		return nil, err
	}

	if err = cmd.Wait(); err != nil {
		return nil, err
	}

	issueFile, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer issueFile.Close()

	buf, err := ioutil.ReadAll(issueFile)
	if err != nil {
		return nil, err
	}

	slices := strings.SplitAfterN(string(buf), "\n", 2)

	// TODO Abort if any field is empty
	title, body := slices[0], slices[1]

	title = strings.TrimPrefix(title, "\n")
	title = strings.TrimSuffix(title, "\n")

	body = strings.TrimPrefix(body, "\n")
	body = strings.TrimSuffix(body, "\n")

	if title == "" || body == "" {
		return nil, fmt.Errorf("issue field is empty, report aborting")
	}

	return &github.IssueRequest{
		Title: &title,
		Body:  &body,
	}, nil
}

func createIssue() (string, error) {
	req, err := getIssue()
	if err != nil {
		return "", err
	}

	t, err := readPassword()
	if err != nil {
		return "", err
	}

	client := github.NewClient(t.Client())
	issue, _, err := client.Issues.Create(boilr.GithubOwner, boilr.GithubRepo, req)
	if err != nil {
		return "", err
	}

	return *issue.HTMLURL, nil
}

// Report contains the cli-command for creating github issues.
var Report = &cli.Command{
	Use:   "report",
	Short: "Create an issue in the github repository",
	Run: func(c *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{})

		url, err := createIssue()
		if err != nil {
			exit.Error(fmt.Errorf("Failed to create an issue: %v", err))
		}

		exit.OK("Successfully created an issue %v", url)
	},
}
