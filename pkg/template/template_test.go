package template_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/tmrts/cookie/pkg/template"
)

func TestRetrievesProjectTemplate(t *testing.T) {
	tmpl, err := template.Get("test-data/test-template")
	if err != nil {
		t.Fatalf("template.Get(%q) got error -> %v", "test-template", err)
	}

	tmpDirPath, err := ioutil.TempDir("", "template-test")
	if err != nil {
		t.Fatalf("ioutil.TempDir() got error -> %v", err)
	} else {
		defer os.RemoveAll(tmpDirPath)
	}

	metadata := template.Metadata{
		Name:    "test-project-1",
		Author:  "test-author",
		Email:   "test@mail.com",
		Date:    time.Now().Format("Mon Jan 2 2006 15:04:05"),
		Version: "0.0.1",
	}

	err = tmpl.Execute(tmpDirPath, metadata)
	if err != nil {
		t.Fatalf("template.Execute(metadata) got error -> %v", err)
	}

	_, err = os.Open(filepath.Join(tmpDirPath, "test-template/"+metadata.Name))
	if err != nil {
		t.Errorf("template.Execute(metadata) directory %q should exist", metadata.Name)
	}
}
