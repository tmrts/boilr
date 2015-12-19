package stringutil_test

import (
	"testing"

	"github.com/tmrts/boilr/pkg/util/stringutil"
)

func TestReadsFromString(t *testing.T) {
	s := stringutil.NewString("Test")

	buf := make([]byte, 4)

	n, err := s.Read(buf)
	if n != 4 || err != nil {
		t.Fatalf("s.Read(buf) should have read the contents to the buffer")
	}

	if string(buf) != "Test" {
		t.Errorf("s.Read(buf) should have read the contents to the buffer -> expected %q got %q", "Test", string(buf))
	}
}

func TestWritesToString(t *testing.T) {
	s := stringutil.NewString("")

	buf := []byte("Test")

	n, err := s.Write(buf)
	if n != 4 || err != nil {
		t.Fatalf("s.Read(buf) should have read the contents to the buffer")
	}

	if s.String() != "Test" {
		t.Errorf("s.Read(buf) should have read the contents to the buffer -> expected %q got %q", "Test", s.String())
	}
}
