package cmd_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/tmrts/cookie/pkg/cmd"
)

func TestBuildExecutesProjectTemplate(t *testing.T) {
	tmpDirPath, err := ioutil.TempDir("", "template-test")
	if err != nil {
		t.Fatalf("ioutil.TempDir() got error -> %v", err)
	} else {
		//defer os.RemoveAll(tmpDirPath)
	}

	if err := os.MkdirAll(filepath.Join(tmpDirPath, "input", "{{Project.Name}}", "{{Project.Date}}"), 0744); err != nil {
		t.Fatalf("os.MkdirAll should have created template directories -> got error %v", err)
	}

	inputDir, outputDir := filepath.Join(tmpDirPath, "input", "{{Project.Name}}"), filepath.Join(tmpDirPath, "output")

	args := []string{inputDir, outputDir}

	cmd.Build.Run(cmd.Build, args)
}
