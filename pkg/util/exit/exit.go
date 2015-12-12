package exit

import (
	"fmt"
	"os"

	"github.com/tmrts/tmplt/pkg/util/tlog"
)

const (
	CodeOK    = 0
	CodeError = 2
)

func Error(err error) {
	tlog.Error(fmt.Sprint(err))

	os.Exit(CodeError)
}

func OK(msg string) {
	os.Exit(CodeOK)
}
