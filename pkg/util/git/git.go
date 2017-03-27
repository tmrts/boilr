// Package git is a facade for git methods used by boilr
package git

import git "gopkg.in/src-d/go-git.v4"

// CloneOptions are used when cloning a git repository
type CloneOptions git.CloneOptions

// Clone clones a git repository with the given options
func Clone(dir string, opts CloneOptions) error {
	o := git.CloneOptions(opts)

	_, err := git.PlainClone(dir, false, &o)
	return err
}
