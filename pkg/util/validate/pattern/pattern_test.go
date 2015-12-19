package pattern_test

import (
	"testing"

	"github.com/tmrts/tmplt/pkg/util/validate/pattern"
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
