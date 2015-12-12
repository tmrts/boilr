package tmplt

import "errors"

var (
	ErrTemplateAlreadyExists = errors.New("tmplt: project template already exists")
	ErrUninitializedTmpltDir = errors.New("tmplt: .tmplt directory is not initialized")
)
