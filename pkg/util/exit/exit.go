package exit

import (
	"fmt"
	"os"

	"github.com/tmrts/boilr/pkg/util/tlog"
)

const (
	// CodeOK indicates successful execution.
	CodeOK = 0

	// CodeError indicates erroneous execution.
	CodeError = 1

	// CodeFatal indicates erroneous use by user.
	CodeFatal = 2
)

// Fatal terminates execution using fatal exit code.
func Fatal(err error) {
	tlog.Fatal(fmt.Sprint(err))

	os.Exit(CodeFatal)
}

// Error terminates execution using unsuccessful execution exit code.
func Error(err error) {
	tlog.Error(err.Error())

	os.Exit(CodeError)
}

// GoodEnough terminates execution successfully, but indicates that something is missing.
func GoodEnough(fmtStr string, s ...interface{}) {
	tlog.Info(fmt.Sprintf(fmtStr, s...))

	os.Exit(CodeOK)
}

// OK terminates execution successfully.
func OK(fmtStr string, s ...interface{}) {
	tlog.Success(fmt.Sprintf(fmtStr, s...))

	os.Exit(CodeOK)
}
