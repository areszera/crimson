// Copyright 2021 AreSZerA. All rights reserved.
// This file provides functions for displaying information, warnings, and errors.

package crimson

import (
	"fmt"
	"time"
)

// PrintInfo displays green information in console.
func PrintInfo(contents ...interface{}) {
	fmt.Print("\033[32m[" + time.Now().Format(time.UnixDate) + "] Info:")
	for _, content := range contents {
		fmt.Print("", content)
	}
	fmt.Println("\033[0m")
}

// PrintWarning displays yellow warnings in console.
func PrintWarning(contents ...interface{}) {
	fmt.Println("\033[33m[" + time.Now().Format(time.UnixDate) + "] Warning:")
	for _, content := range contents {
		fmt.Print("", content)
	}
	fmt.Println("\033[0m")
}

// PrintError displays red error in console.
func PrintError(contents ...interface{}) {
	fmt.Println("\033[31m[" + time.Now().Format(time.UnixDate) + "] Error:")
	for _, content := range contents {
		fmt.Print("", content)
	}
	fmt.Println("\033[0m")
}
