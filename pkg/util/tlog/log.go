// Package tlog implements logging utilities for boilr
package tlog

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// TODO default to ASCII if Unicode is not supported
const (
	// DebugMark character indicates debug message.
	DebugMark = "☹"

	// CheckMark character indicates success message.
	CheckMark = "✔"

	// InfoMark character indicates information message.
	InfoMark = "i"

	// WarnMark character indicates warning message.
	WarnMark = "!"

	// ErrorMark character indicates error message.
	ErrorMark = "✘"

	// QuestionMark character indicates prompt message.
	QuestionMark = "?"
)

const (
    LevelDebug = 32
    LevelFatal = 16
    LevelError = 8
    LevelWarn = 4
    LevelInfo = 2
    LevelSuccess = 1
)

var LogLevel uint16

func SetLogLevel(LogLevelString string) {
    switch strings.ToLower(LogLevelString) {
    case "debug":
		LogLevel |= (LevelSuccess | LevelError | LevelFatal | LevelWarn | LevelInfo | LevelDebug)
	case "info":
        LogLevel |= (LevelSuccess | LevelError | LevelFatal | LevelWarn | LevelInfo)
	case "warn":
        LogLevel |= (LevelSuccess | LevelError | LevelFatal | LevelWarn)
	default:
        LogLevel |= (LevelSuccess | LevelError | LevelFatal)
    }
}

// TODO add log levels
func coloredPrintMsg(icon string, msg string, iC color.Attribute, mC color.Attribute) {
	fmt.Println(
		color.New(mC).SprintFunc()("["+icon+"]"),
		color.New(color.Bold, iC).SprintFunc()(msg))
}

// Debug logs the given message as a debug message.
func Debug(msg string) {
	if 0 == LogLevel & LevelDebug {
		return
	}

	coloredPrintMsg(DebugMark, msg, color.FgYellow, color.FgYellow)
}

// Success logs the given message as a success message.
func Success(msg string) {
	coloredPrintMsg(CheckMark, msg, color.FgWhite, color.FgGreen)
}

// Info logs the given message as a info message.
func Info(msg string) {
	if 0 == LogLevel & LevelInfo {
		return
	}

	coloredPrintMsg(InfoMark, msg, color.FgWhite, color.FgBlue)
}

// Warn logs the given message as a warn message.
func Warn(msg string) {
	if 0 == LogLevel & LevelWarn {
		return
	}

	coloredPrintMsg(WarnMark, msg, color.FgMagenta, color.FgMagenta)
}

// Error logs the given message as a error message.
func Error(msg string) {
	coloredPrintMsg(ErrorMark, msg, color.FgRed, color.FgRed)
}

// Fatal logs the given message as a fatal message.
func Fatal(msg string) {
	Error(msg)
}

// Prompt outputs the given message as a question along with a default value.
func Prompt(msg string, defval interface{}) {
	fmt.Print(strings.Join([]string{
		color.New(color.FgBlue).SprintFunc()("[" + QuestionMark + "]"),
		color.New(color.Bold, color.FgWhite).SprintFunc()(msg),
		color.New(color.FgBlue).SprintFunc()(fmt.Sprintf("[default: %#v]: ", defval)),
	}, " "))
}

// TODO use dependency injection wrapper for fmt.Print usage in the code base
