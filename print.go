// Copyright 2021 AreSZerA. All rights reserved.
// This file provides functions for displaying information, warnings, and errors.

package crimson

import (
	"fmt"
	"time"
)

var logFontColor = map[string]string{
	logLevelInfo:    fontGreen,
	logLevelWarning: fontYellow,
	logLevelError:   fontRed,
}

func printSomething(logLevel string, contents ...interface{}) {
	fmt.Print(logFontColor[logLevel] + time.Now().Format(time.UnixDate) + "] " + logLevel + ": ")
	fmt.Print(contents...)
	fmt.Println(fontDefault)
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
