package main

import (
	"log"
	"sync"

	"max/pb"
)

type SessionStore struct {
	lock      sync.Mutex
	sessCache map[string]*Session
}

// Create creates a new session and adds it to store.
func (ss *SessionStore) Create(conn pb.Service_MessageServer) *Session {
	s := NewSession(conn)
	ss.sessCache[s.sid] = s

	return s
}

// Get fetches a session from store by session ID.
func (ss *SessionStore) Get(sid string) *Session {
	ss.lock.Lock()
	defer ss.lock.Unlock()

	if sess := ss.sessCache[sid]; sess != nil {
		return sess
	}

	return nil
}

// Delete removes session from store.
func (ss *SessionStore) Delete(s *Session) int {
	ss.lock.Lock()
	defer ss.lock.Unlock()

	delete(ss.sessCache, s.sid)
	return len(ss.sessCache)
}

func (ss *SessionStore) Shutdown() {
	ss.lock.Lock()
	defer ss.lock.Unlock()

	for _, s := range ss.sessCache {
		s.stop <- struct{}{}
	}

	log.Printf("SessionStore shut down, sessions terminated: %d", len(ss.sessCache))
}

// NewSessionStore initializes a session store.
func NewSessionStore() *SessionStore {
	ss := &SessionStore{
		sessCache: make(map[string]*Session),
		lock:      sync.Mutex{},
	}

	return ss
}
