// Copyright 2021 AreSZerA. All rights reserved.
// This file provides a session manager to manage global sessions.
// Learnt from https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/06.0.md

package crimson

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

// Session is an interface with methods of set, get, delete, and get session ID.
type Session interface {
	Set(key, value interface{})
	Get(key interface{}) interface{}
	Delete(key interface{})
	SessionID() string
}

// sessionProvider is an interface with methods of initialize, read, destroy, update and garbage collect for sessions.
type sessionProvider interface {
	Init(sid string) Session
	Read(sid string) Session
	Destroy(sid string)
	Update(sid string)
	GC(maxLifeTime int64)
}

var seProviders = make(map[string]sessionProvider)

// register enables session provider by name.
// Provider must be initialized and session name must be unique.
func register(name string, provider sessionProvider) {
	if provider == nil {
		PrintError("Session provider is nil")
		os.Exit(3)
	}
	if _, dup := seProviders[name]; dup {
		PrintError("Session provider has been registered")
		os.Exit(3)
	}
	seProviders[name] = provider
}

// SessionManager for managing sessions.
type SessionManager struct {
	provider    sessionProvider
	cookieName  string
	maxTimeLife int64
	lock        sync.Mutex
}

// NewSessionManager creates new instance for SessionManager
func NewSessionManager() *SessionManager {
	manager := &SessionManager{
		provider:    seProviders[GetSessionProviderName()],
		cookieName:  GetSessionCookieName(),
		maxTimeLife: GetSessionTimeout(),
	}
	// As the manager is initialized, start garbage collecting.
	go manager.GC()
	return manager
}

// randSessionID generates a random ID for session.
func (m *SessionManager) randSessionID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// StartSession starts session according to session ID stored in cookie.
// If cookie not set or session id is empty, start a new session.
// Otherwise, use read session by ID.
func (m *SessionManager) StartSession(w http.ResponseWriter, r *http.Request) (session Session) {
	var id string
	m.lock.Lock()
	defer m.lock.Unlock()
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		id = m.randSessionID()
		session = m.provider.Init(id)
		http.SetCookie(w, &http.Cookie{
			Name:     m.cookieName,
			Value:    url.QueryEscape(id),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(m.maxTimeLife),
		})
	} else {
		id, _ = url.QueryUnescape(cookie.Value)
		session = m.provider.Read(id)
	}
	return
}

// DestroySession destroys session by making cookie expired.
func (m *SessionManager) DestroySession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(m.cookieName)
	if err == nil && cookie.Value != "" {
		m.lock.Lock()
		defer m.lock.Unlock()
		m.provider.Destroy(cookie.Value)
		http.SetCookie(w, &http.Cookie{
			Name:     m.cookieName,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   -1,
			Expires:  time.Now(),
		})
	}
}

// GC collects expired sessions.
func (m *SessionManager) GC() {
	m.lock.Lock()
	defer m.lock.Unlock()
	time.AfterFunc(time.Duration(m.maxTimeLife)*time.Second, func() { m.provider.GC(m.maxTimeLife) })
}
