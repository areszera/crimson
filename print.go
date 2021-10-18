// Copyright 2021 AreSZerA. All rights reserved.
// This file provides functions for displaying information, warnings, and errors.

package crimson

import (
	"fmt"
	"time"
)

// PrintInfo displays green information in console.
func PrintInfo(content string) {
	fmt.Println("\033[32m[" + time.Now().Format(time.UnixDate) + "] Info: " + content + "\033[0m")
}

// PrintWarning displays yellow warnings in console.
func PrintWarning(content string) {
	fmt.Println("\033[33m[" + time.Now().Format(time.UnixDate) + "] Warning: " + content + "\033[0m")
}

// PrintError displays red error in console.
func PrintError(content string) {
	fmt.Println("\033[31m[" + time.Now().Format(time.UnixDate) + "] Error: " + content + "\033[0m")
}
