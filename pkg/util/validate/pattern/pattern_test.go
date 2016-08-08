package pattern_test

import (
	"testing"

	"github.com/tmrts/boilr/pkg/util/validate/pattern"
)

func TestUnixPathPattern(t *testing.T) {
	tests := []struct {
		String string
		Valid  bool
	}{
		{"", false},
		{"/", true},
		{"/root", true},
		{"/tmp-dir", true},
		{"/tmp-dir/new_dir", true},
		{"/TMP/dir", true},
		{"rel/dir", true},
	}

	for _, test := range tests {
		if ok := pattern.UnixPath.MatchString(test.String); ok != test.Valid {
			t.Errorf("pattern.UnixPath.MatchString(%q) expected to be %v", test.String, test.Valid)
		}
	}
}

func TestAlphanumericPattern(t *testing.T) {
	tests := []struct {
		String string
		Valid  bool
	}{
		{" ", false},
		{"/", false},
		{"root", true},
		{"tmp-dir", false},
		{"TMPDIR", true},
		{"L33T", true},
		{"L@@T", false},
	}

	for _, test := range tests {
		if ok := pattern.Alphanumeric.MatchString(test.String); ok != test.Valid {
			t.Errorf("pattern.Alphanumeric.MatchString(%q) expected to be %v", test.String, test.Valid)
		}
	}
}

func TestAlphanumericExtPattern(t *testing.T) {
	tests := []struct {
		String string
		Valid  bool
	}{
		{" ", false},
		{"/", false},
		{"root", true},
		{"tmp-dir", true},
		{"tmp-dir_now", true},
		{"TMPDIR", true},
		{"L33T", true},
		{"L@@T", false},
	}

	for _, test := range tests {
		if ok := pattern.AlphanumericExt.MatchString(test.String); ok != test.Valid {
			t.Errorf("pattern.AlphanumericExt.MatchString(%q) expected to be %v", test.String, test.Valid)
		}
	}
}

func TestIntegerPattern(t *testing.T) {
	tests := []struct {
		String string
		Valid  bool
	}{
		{"", false},
		{" ", false},
		{"/", false},
		{"root", false},
		{"L33T", false},
		{"", false},
	}

	for _, test := range tests {
		if ok := pattern.Numeric.MatchString(test.String); ok != test.Valid {
			t.Errorf("pattern.Numeric.MatchString(%q) expected to be %v", test.String, test.Valid)
		}
	}
}

func TestURLPattern(t *testing.T) {
	tests := []struct {
		String string
		Valid  bool
	}{
		{"", false},
		{" ", false},
		{"/", false},
		{"http://", false},
		{"http://github.com/tmrts/boilr", true},
		{"https://github.com/tmrts/boilr", true},
		{"github.com/tmrts/boilr", true},
		{"rawcontent.github.com/tmrts/boilr", true},
		{"github.com:80/tmrts/boilr", true},
	}

	for _, test := range tests {
		if ok := pattern.URL.MatchString(test.String); ok != test.Valid {
			t.Errorf("pattern.URL.MatchString(%q) expected to be %v", test.String, test.Valid)
		}
	}
}
