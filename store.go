// Copyright 2021 AreSZerA. All rights reserved.
// This file provides a session manager to manage global sessions.
// Learnt from https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/06.0.md

package crimson

import (
	"container/list"
	"sync"
	"time"
)

// Initialize session and session store provider.
func init() {
	stProvider.sessions = make(map[string]*list.Element, 0)
	register(GetSessionProviderName(), stProvider)
}

// storeProvider implements sessionProvider interface.
type storeProvider struct {
	lock     sync.Mutex
	sessions map[string]*list.Element
	list     *list.List
}

var stProvider = &storeProvider{list: list.New()}

// Init initializes new session by appending into list.
func (p *storeProvider) Init(sid string) (session Session) {
	stProvider.lock.Lock()
	defer stProvider.lock.Unlock()
	value := make(map[interface{}]interface{}, 0)
	session = &store{sid: sid, timeAccessed: time.Now(), value: value}
	element := stProvider.list.PushBack(session)
	stProvider.sessions[sid] = element
	return
}

// Read reads session by ID from list.
func (*storeProvider) Read(sid string) (session Session) {
	if element, ok := stProvider.sessions[sid]; ok {
		session = element.Value.(*store)
	} else {
		session = stProvider.Init(sid)
	}
	return
}

// Update updates session in list.
func (*storeProvider) Update(sid string) {
	stProvider.lock.Lock()
	defer stProvider.lock.Unlock()
	if element, ok := stProvider.sessions[sid]; ok {
		element.Value.(*store).timeAccessed = time.Now()
		stProvider.list.MoveToFront(element)
	}
}

// Destroy destroys session by removing from list.
func (*storeProvider) Destroy(sid string) {
	if element, ok := stProvider.sessions[sid]; ok {
		delete(stProvider.sessions, sid)
		stProvider.list.Remove(element)
	}
}

// GC collects session by removing from list.
func (*storeProvider) GC(maxLifeTime int64) {
	stProvider.lock.Lock()
	defer stProvider.lock.Unlock()
	for {
		element := stProvider.list.Back()
		if element != nil {
			if element.Value.(*store).timeAccessed.Unix()+maxLifeTime < time.Now().Unix() {
				stProvider.list.Remove(element)
				delete(stProvider.sessions, element.Value.(*store).sid)
			} else {
				break
			}
		} else {
			break
		}
	}
}

// store implements Session interface.
type store struct {
	sid          string
	timeAccessed time.Time
	value        map[interface{}]interface{}
}

// Set sets key and value in session.
func (s *store) Set(key, value interface{}) {
	s.value[key] = value
	stProvider.Update(s.sid)
}

// Get gets value according to key from session, returns nil if key is unset.
func (s *store) Get(key interface{}) interface{} {
	stProvider.Update(s.sid)
	if value, ok := s.value[key]; ok {
		return value
	}
	return nil
}

// Delete removes key and value from session.
func (s *store) Delete(key interface{}) {
	delete(s.value, key)
	stProvider.Update(s.sid)
}

// SessionID returns session ID.
func (s *store) SessionID() string {
	return s.sid
}
