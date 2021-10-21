// Copyright 2021 AreSZerA. All rights reserved.
// This file provides default error page handler.

package crimson

import (
	"html/template"
	"net/http"
)

// ErrPageHandler sets response status code and error page.
func ErrPageHandler(w http.ResponseWriter, status int, info interface{}) {
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
