package boilr

import "errors"

var (
	// ErrTemplateAlreadyExists indicates that a template is already present in the local registry.
	ErrTemplateAlreadyExists = errors.New("boilr: project template already exists")
)
