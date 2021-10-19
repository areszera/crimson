// Copyright 2021 AreSZerA. All rights reserved.
// This file provides default error page handler.

package crimson

import (
	"html/template"
	"net/http"
)

// statusInfo maps part of HTTP status codes to their representations.
var statusInfo = map[int]string{
	400: "Bad Request",
	401: "Unauthorized",
	402: "Payment Required",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Time-out",
	409: "Conflict",
	410: "Gone",
	411: "Length Required",
	412: "Precondition Failed",
	413: "Request Entity Too Large",
	414: "Request-URI Too Large",
	415: "Unsupported Media Type",
	416: "Requested range not satisfiable",
	417: "Expectation Failed",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Time-out",
	505: "HTTP Version not supported",
}

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
    <p><a href="https://github.com/AreSZerA/crimson">Crimson v0.1</a></p>
</main>

</body>

</html>`

// ErrPageHandler sets response status code and error page.
func ErrPageHandler(w http.ResponseWriter, status int, info interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/html")
	sInfo, ok := statusInfo[status]
	if !ok {
		sInfo = "(Unknown)"
	}
	tpl, _ := template.New("ErrorPage").Parse(errPageTpl)
	err := tpl.Execute(w, map[string]interface{}{
		"status": status,
		"sInfo":  sInfo,
		"info":   info,
	})
	if err != nil {
		PrintWarning("Failed to execute template: " + err.Error())
	}
}
