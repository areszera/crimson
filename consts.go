// Copyright 2021 AreSZerA. All rights reserved.
// This file defines constants.

package crimson

const (
	// exit codes
	exitOK = iota
	exitErrStartServer
	exitErrAddRoute
	exitErrRegSession

	// log levels
	logLevelInfo    = "Info"
	logLevelWarning = "Warning"
	logLevelError   = "Error"

	// font colors in console
	fontGreen   = "\u001B[32m["
	fontYellow  = "\u001B[33m["
	fontRed     = "\u001B[31m["
	fontDefault = "\033[0m"

	// Version is the current version of Crimson
	Version = "0.1"

	// errPageTpl is the HTML template for displaying error page with status code, status code information and error information.
	errPageTpl = `<!DOCTYPE html>

<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta content="width=device-width,initial-scale=1" name="viewport">
    <title>Crimson - {{ .status }} {{ .sInfo }}</title>
    <style>
    * {
        font-family: serif;
        margin: 0;
        padding: 0;
    }
    header {
        border-bottom: 1px solid #c0c0c0;
        padding: 1rem 0;
        width: 100%;
    }
    header h1 {
        color: crimson;
        text-align: center;
    }
    main {
        padding-top: 1rem;
        width: 100%;
    }
    main p {
        color: black;
        line-height: 1.5rem;
        text-align: center;
    }
    main a {
        color: black;
        text-decoration: none;
    }
    </style>
</head>


<body>

<header>
    <h1>{{ .status }} {{ .sInfo }}</h1>
</header>

<main>
    <p>{{ .info }}</p>
    <p><a href="https://github.com/AreSZerA/crimson" target="_blank">Crimson v` + Version + `</a></p>
</main>

</body>

</html>`
)
