// Package tlog implements logging utilities for tmplt
package tlog

import (
	"fmt"

	"github.com/fatih/color"
)

func Debug(msg string) {
	fmt.Println(
		color.New(color.BgYellow).SprintFunc()("DEBUG"),
		color.New(color.FgYellow).SprintFunc()(msg))
}

func Success(msg string) {
	fmt.Println(
		color.New(color.BgGreen).SprintFunc()("SUCCESS"),
		color.New(color.FgGreen).SprintFunc()(msg))
}

func Info(msg string) {
	fmt.Println(
		color.New(color.BgBlue).SprintFunc()("INFO"),
		color.New(color.FgBlue).SprintFunc()(msg))
}

func Warn(msg string) {
	fmt.Println(
		color.New(color.BgMagenta).SprintFunc()("WARN"),
		color.New(color.FgMagenta).SprintFunc()(msg))
}

func Error(msg string) {
	fmt.Println(
		color.New(color.BgRed).SprintFunc()("ERROR"),
		color.New(color.Bold, color.FgRed).SprintFunc()(msg))
}

func Fatal(msg string) {
	Error(msg)
}
