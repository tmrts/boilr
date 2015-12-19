package template_test

import (
	"encoding/json"
	"testing"

	"github.com/tmrts/boilr/pkg/template"
)

func TestMarshalsTime(t *testing.T) {
	jsonT := template.NewTime()

	b, err := jsonT.MarshalJSON()
	if err != nil {
		t.Error(err)
	}

	var unmarshaledT template.JSONTime
	if err := json.Unmarshal(b, &unmarshaledT); err != nil {
		t.Error(err)
	}

	expected, got := jsonT.String(), unmarshaledT.String()
	if expected != got {
		t.Errorf("marshaled and unmarshaled time should've been equal expected %q, got %q", expected, got)
	}
}
