// Package tlog implements logging utilities for tmplt
package tlog

import "github.com/Sirupsen/logrus"

func Debug(msg string) {
	logrus.Debug(msg)
}

func Info(msg string) {
	logrus.Info(msg)
}

func Warn(msg string) {
	logrus.Warning(msg)
}

func Error(msg string) {
	logrus.Error(msg)
}

func Fatal(msg string) {
	logrus.Fatal(msg)
}

func Panic(msg string) {
	logrus.Panic(msg)
}
