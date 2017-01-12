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

var (
	logLevel Level
)

// Level is a 16-bit set holding the enabled log levels.
type Level uint16

const (
	LevelDebug   = 1 << 5
	LevelFatal   = 1 << 4
	LevelWarn    = 1 << 3
	LevelError   = 1 << 2
	LevelInfo    = 1 << 1
	LevelSuccess = 1 << 0
)

// Set enables the levels upto and including the given log level.
func (lvl *Level) Set(newLvl Level) {
	*lvl = (newLvl << 1) - 1
}

// Permits queries whether the given log level is enabled or not.
func (lvl Level) Permits(queryLvl Level) bool {
	return lvl&queryLvl > 0
}

// SetLogLevel sets the global logging level.
func SetLogLevel(LogLevelString string) {
	levels := map[string]Level{
		"debug":   LevelDebug,
		"fatal":   LevelFatal,
		"warn":    LevelWarn,
		"error":   LevelError,
		"info":    LevelInfo,
		"success": LevelSuccess,
	}

	newLevel, ok := levels[strings.ToLower(LogLevelString)]
	if !ok {
		Error(fmt.Sprintf("unknown log level %s", LogLevelString))
		return
	}

	logLevel.Set(newLevel)
}

// TODO add log levels
func coloredPrintMsg(icon string, msg string, iC color.Attribute, mC color.Attribute) {
	fmt.Println(
		color.New(mC).SprintFunc()("["+icon+"]"),
		color.New(color.Bold, iC).SprintFunc()(msg))
}

// Debug logs the given message as a debug message.
func Debug(msg string) {
	if !logLevel.Permits(LevelDebug) {
		return
	}

	coloredPrintMsg(DebugMark, msg, color.FgYellow, color.FgYellow)
}

// Success logs the given message as a success message.
func Success(msg string) {
	if !logLevel.Permits(LevelSuccess) {
		return
	}

	coloredPrintMsg(CheckMark, msg, color.FgWhite, color.FgGreen)
}

// Info logs the given message as a info message.
func Info(msg string) {
	if !logLevel.Permits(LevelInfo) {
		return
	}

	coloredPrintMsg(InfoMark, msg, color.FgWhite, color.FgBlue)
}

// Warn logs the given message as a warn message.
func Warn(msg string) {
	if !logLevel.Permits(LevelWarn) {
		return
	}

	coloredPrintMsg(WarnMark, msg, color.FgMagenta, color.FgMagenta)
}

// Error logs the given message as a error message.
func Error(msg string) {
	if !logLevel.Permits(LevelError) {
		return
	}

	coloredPrintMsg(ErrorMark, msg, color.FgRed, color.FgRed)
}

// Fatal logs the given message as a fatal message.
func Fatal(msg string) {
	// Fatal level is being deprecated
	Error(msg)
}

// Prompt outputs the given message as a question along with a default value.
func Prompt(msg string, defval interface{}) {
	tokens := []string{
		color.New(color.FgBlue).SprintFunc()("[" + QuestionMark + "]"),
		color.New(color.Bold, color.FgWhite).SprintFunc()(msg),
	}

	// TODO refactor & eliminate duplication
	switch val := defval.(type) {
	case []interface{}:
		tokens = append(tokens, "\n")
		for i, v := range val {
			tokens = append(tokens, color.New(color.Bold, color.FgWhite).SprintFunc()(fmt.Sprintf("   %v - %#v\n", i+1, v)))
		}

		tokens = append(tokens, color.New(color.Bold, color.FgWhite).SprintFunc()(fmt.Sprintf("   Choose from %v..%v", 1, len(val))))

		tokens = append(tokens, color.New(color.Bold, color.FgBlue).SprintFunc()(fmt.Sprintf("[default: %v]: ", 1)))
	default:
		tokens = append(tokens, color.New(color.Bold, color.FgBlue).SprintFunc()(fmt.Sprintf("[default: %#v]: ", defval)))
	}

	fmt.Print(strings.Join(tokens, " "))
}

// TODO use dependency injection wrapper for fmt.Print usage in the code base
func init() {
	logLevel.Set(LevelError)
}
