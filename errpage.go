// Copyright 2021 AreSZerA. All rights reserved.
// This file provides default error page handler.

package crimson

import (
	"html/template"
	"net/http"
)

// errPageTpl is the HTML template for displaying error page with status code, status code information and error information.
const errPageTpl = `<!DOCTYPE html>

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
    <p><a href="https://github.com/AreSZerA/crimson" target="_blank">Crimson v` + version + `</a></p>
</main>

</body>

</html>`

// ErrPageHandler sets response status code and error page.
func ErrPageHandler(w http.ResponseWriter, status int, info interface{}) {
	// TODO : http: superfluous response.WriteHeader call from github.com/AreSZerA/crimson.ErrPageHandler
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/html")
	tpl, _ := template.New("ErrorPage").Parse(errPageTpl)
	err := tpl.Execute(w, map[string]interface{}{
		"status": status,
		"sInfo":  http.StatusText(status),
		"info":   info,
	})
	if err != nil {
		PrintWarning("Failed to execute template: " + err.Error())
	}
}
