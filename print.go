// Copyright 2021 AreSZerA. All rights reserved.
// This file provides functions for displaying information, warnings, and errors.

package crimson

import (
	"fmt"
	"time"
)

const (
	logLevelInfo    = "Info"
	logLevelWarning = "Warning"
	logLevelError   = "Error"
)

var logLevelColor = map[string]string{
	logLevelInfo:    "\u001B[32m[", // green
	logLevelWarning: "\u001B[33m[", // yellow
	logLevelError:   "\u001B[31m[", // red
}

func printSomething(logLevel string, contents ...interface{}) {
	fmt.Print(logLevelColor[logLevel] + time.Now().Format(time.UnixDate) + "] " + logLevel + ": ")
	for _, content := range contents {
		fmt.Print(content)
		fmt.Print(" ")
	}
	fmt.Println("\033[0m")
}

// PrintInfo displays green information in console.
func PrintInfo(contents ...interface{}) {
	printSomething(logLevelInfo, contents...)
}

// PrintWarning displays yellow warnings in console.
func PrintWarning(contents ...interface{}) {
	printSomething(logLevelWarning, contents...)
}

// PrintError displays red error in console.
func PrintError(contents ...interface{}) {
	printSomething(logLevelError, contents...)
}
