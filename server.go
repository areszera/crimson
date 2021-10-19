// Copyright 2021 AreSZerA. All rights reserved.
// This file provides functions to start server.

package crimson

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type crimsonServer struct {
	port int
}

// NewServer creates new server instance.
func NewServer() *crimsonServer {
	return &crimsonServer{port: GetServerPort()}
}

// Start starts server in another goroutine and wait until received interrupt signal from keyboard (^C).
func (s *crimsonServer) Start() {
	PrintInfo("Start server on port " + strconv.Itoa(s.port))
	server := http.Server{Addr: ":" + strconv.Itoa(s.port)}
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			PrintError("Failed to start server: " + err.Error())
			os.Exit(1)
		}
	}()
	if IsBrowserAutoOpen() {
		OpenInBrowser()
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	wait := time.Duration(GetServerTimeout()) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		PrintInfo("There are unfinished works, will shut down in " + strconv.Itoa(int(GetServerTimeout())) + " seconds")
	}
	PrintInfo("Server shut down successfully")
}
