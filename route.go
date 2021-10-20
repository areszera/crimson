// Copyright 2021 AreSZerA. All rights reserved.
// This file provides functions for configuring web route.

package crimson

import (
	"net/http"
	"os"
	"regexp"
)

var httpReqMethods = map[string]bool{
	"GET":     false,
	"POST":    false,
	"HEAD":    false,
	"OPTIONS": false,
	"PUT":     false,
	"PATCH":   false,
	"DELETE":  false,
	"TRACE":   false,
	"CONNECT": false,
}

// route consists of a regular expression matcher for URL, a handler for this URL, and available methods for it.
type route struct {
	matcher *regexp.Regexp
	handler func(http.ResponseWriter, *http.Request)
	methods map[string]bool
}

var routes []route

// NewRoute creates new route instance.
// Available methods will be GET if not entered.
func NewRoute(pattern string, handler func(http.ResponseWriter, *http.Request), methods ...string) *route {
	rMethods := httpReqMethods
	if len(methods) == 0 {
		rMethods["GET"] = true
	} else {
		for _, m := range methods {
			rMethods[m] = true
		}
	}
	matcher, err := regexp.Compile(pattern)
	if err != nil {
		PrintError("Failed to compile pattern " + pattern + ": " + err.Error())
		os.Exit(2)
	}
	return &route{
		matcher: matcher,
		methods: rMethods,
		handler: handler,
	}
}

// AddRoutes appends new route(s).
func AddRoutes(r ...route) {
	routes = append(routes, r...)
}

// AddRoute append one new route.
func AddRoute(pattern string, handler func(http.ResponseWriter, *http.Request), methods ...string) {
	r := NewRoute(pattern, handler, methods...)
	routes = append(routes, *r)
}

// AddStaticRoute configures static file server, e.g. css and js files.
func AddStaticRoute(prefix, dir string) {
	http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(dir))))
}

// All the requests will firstly be directed to "/", then according r.URL.Path, do their corresponding actions.
func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, rt := range routes {
			if rt.matcher.MatchString(r.URL.Path) {
				if ok, exist := rt.methods[r.Method]; ok && exist {
					rt.handler(w, r)
				}
				ErrPageHandler(w, 405, "Method "+r.Method+" is not allowed here")
				return
			}
		}
		ErrPageHandler(w, 404, "Cannot find resource for "+r.URL.Path)
	})
}
