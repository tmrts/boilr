package inpututil

import (
	"fmt"
	"strings"
)

var (
	booleanValues = map[string]bool{
		"y":   true,
		"yes": true,
		"yup": true,

		"n":    false,
		"no":   false,
		"nope": false,
	}
)

func ScanYesOrNo(msg string, def bool) (bool, error) {
	var choices string
	if def {
		choices = "Y/n"
	} else {
		choices = "y/N"
	}

	fmt.Print(msg, " ", choices, " ")

	var choice string
	_, err := fmt.Scanf("%s", &choice)
	if err != nil {
		return def, err
	}

	switch val, ok := booleanValues[strings.ToLower(choice)]; {
	case ok:
		return val, nil
	case choice == "":
		return def, nil
	}

	return def, fmt.Errorf("unrecognized choice %#q", choice)
}

func ScanValue(msg string) (string, error) {
	fmt.Print(msg)

	var value string
	_, err := fmt.Scanf("%s", &value)
	if err != nil {
		return "", err
	}

	return value, nil
}
