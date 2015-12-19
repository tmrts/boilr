package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/tmrts/tmplt/pkg/template"
	"github.com/tmrts/tmplt/pkg/tmplt"
)

func serializeMetadata(tag string, repo string, targetDir string) error {
	fname := filepath.Join(targetDir, tmplt.TemplateMetadataName)

	f, err := os.Create(fname)
	if err != nil {
		return err
	} else {
		defer f.Close()
	}

	enc := json.NewEncoder(f)

	t := template.Metadata{tag, repo, template.NewTime()}
	if err := enc.Encode(&t); err != nil {
		return err
	}

	return nil
}
