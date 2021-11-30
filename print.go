// Copyright 2021 AreSZerA. All rights reserved.
// This file provides functions for displaying information, warnings, and errors.

package crimson

import (
	"log"
)

var logFontColor = map[string]string{
	logLevelInfo:    fontGreen,
	logLevelWarning: fontYellow,
	logLevelError:   fontRed,
}

func printSomething(logLevel string, contents ...interface{}) {
	log.SetPrefix(logFontColor[logLevel])
	log.SetFlags(log.LstdFlags)
	contents = append([]interface{}{logLevel + ":"}, contents...)
	contents = append(contents, fontDefault)
	log.Println(contents...)
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
