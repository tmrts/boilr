// Package tlog implements logging utilities for boilr
package tlog

import (
	"fmt"

	"github.com/fatih/color"
)

// TODO default to ASCII if Unicode is not supported
const (
	DebugMark = "☹"
	CheckMark = "✔"
	InfoMark  = "i"
	WarnMark  = "!"
	ErrorMark = "✘"

	// TODO use for prompts
	QuestionMark = "?"
)

func coloredPrintMsg(icon string, msg string, iC color.Attribute, mC color.Attribute) {
	fmt.Println(
		color.New(mC).SprintFunc()("["+icon+"]"),
		color.New(color.Bold, iC).SprintFunc()(msg))
}

// TODO add log levels
func Debug(msg string) {
	coloredPrintMsg(DebugMark, msg, color.FgYellow, color.FgYellow)
}

func Success(msg string) {
	coloredPrintMsg(CheckMark, msg, color.FgWhite, color.FgGreen)
}

func Info(msg string) {
	coloredPrintMsg(InfoMark, msg, color.FgBlue, color.FgBlue)
}

func Warn(msg string) {
	coloredPrintMsg(WarnMark, msg, color.FgMagenta, color.FgMagenta)
}

func Error(msg string) {
	coloredPrintMsg(ErrorMark, msg, color.FgRed, color.FgRed)
}

func Fatal(msg string) {
	Error(msg)
}

// TODO use dependency injection wrapper for fmt.Print usage in the code base
