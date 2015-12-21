package boilr

import "errors"

var (
	// Indicates that a template is already present in the local registry.
	ErrTemplateAlreadyExists = errors.New("boilr: project template already exists")
)
