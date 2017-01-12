package prompt_test

import (
	"reflect"
	"testing"

	"github.com/tmrts/boilr/pkg/prompt"
)

func TestNewStringPromptFunc(t *testing.T) {
	defval, name := "defaultValue", "fieldName"

	sp := prompt.Func(defval)

	msg := sp.PromptMessage(name)

	expectedPromptMsg := "Please choose a value for \"fieldName\""
	if msg != expectedPromptMsg {
		t.Errorf("strPrompt(%q).PromptMessage(%q) expected %q got %q", defval, name, expectedPromptMsg, msg)
	}

	choiceCases := []struct {
		choice      string
		expectedVal string
	}{
		{"", defval},
		{"chosenValue", "chosenValue"},
	}

	for _, c := range choiceCases {
		val, err := sp.EvaluateChoice(c.choice)
		if err != nil {
			t.Errorf("strPrompt(%q).EvaluateChoice(%q) got error %q", defval, c.choice, err)
			continue
		}

		if !reflect.DeepEqual(val, c.expectedVal) {
			t.Errorf("strPrompt(%q).EvaluateChoice(%q) expected %q got %q", defval, c.choice, c.expectedVal, val)
		}
	}
}

func TestNewBooleanPromptFunc(t *testing.T) {
	defval, name := true, "fieldName"

	boolPrompt := prompt.Func(defval)

	msg := boolPrompt.PromptMessage(name)

	expectedPromptMsg := "Please choose a value for \"fieldName\""
	if msg != expectedPromptMsg {
		t.Errorf("boolPrompt(%q).PromptMessage(%q) expected %q got %q", defval, name, expectedPromptMsg, msg)
	}

	choiceCases := []struct {
		choice      string
		expectedVal bool
	}{
		{"", true},

		{"n", false},
		{"no", false},
		{"nope", false},
		{"false", false},

		{"", true},
		{"y", true},
		{"yes", true},
		{"yup", true},
		{"true", true},
	}

	for _, c := range choiceCases {
		val, err := boolPrompt.EvaluateChoice(c.choice)
		if err != nil {
			t.Errorf("boolPrompt(%q).EvaluateChoice(%q) got error %q", defval, c.choice, err)
			continue
		}

		if !reflect.DeepEqual(val, c.expectedVal) {
			t.Errorf("boolPrompt(%#v).EvaluateChoice(%#v) expected %#v got %#v", defval, c.choice, c.expectedVal, val)
		}
	}
}

func TestMultipleChoicePromptFunc(t *testing.T) {
	defval, name := []interface{}{"choice1", "choice2"}, "fieldName"

	slicePrompt := prompt.Func(defval)

	msg := slicePrompt.PromptMessage(name)

	expectedPromptMsg := "Please choose an option for \"fieldName\""
	if msg != expectedPromptMsg {
		t.Errorf("slicePrompt(%q).PromptMessage(%q) expected %q got %q", defval, name, expectedPromptMsg, msg)
	}

	choiceCases := []struct {
		choice      string
		expectedVal string
	}{
		{"", "choice1"},

		{"1", "choice1"},
		{"2", "choice2"},

		{"0", "choice1"},
		{"3", "choice1"},
		{"4", "choice1"},

		{"aklsjdflska", "choice1"},
		{"-1", "choice1"},
	}

	for _, c := range choiceCases {
		val, err := slicePrompt.EvaluateChoice(c.choice)
		if err != nil {
			t.Errorf("slicePrompt(%q).EvaluateChoice(%q) got error %q", defval, c.choice, err)
			continue
		}

		if !reflect.DeepEqual(val, c.expectedVal) {
			t.Errorf("slicePrompt(%q).EvaluateChoice(%q) expected %q got %q", defval, c.choice, c.expectedVal, val)
		}
	}
}
