// Package git is a facade for git methods used by boilr
package git

import git "gopkg.in/src-d/go-git.v4"

// Options are used when cloning or pulling from a git repository
type CloneOptions git.CloneOptions
type PullOptions git.PullOptions

// Clone clones a git repository with the given options
func Clone(dir string, opts CloneOptions) error {
	o := git.CloneOptions(opts)

	_, err := git.PlainClone(dir, false, &o)
	return err
}

func Open(path string) (*git.Repository, error) {
  return git.PlainOpen(path)
}

