package session

import "time"

const refreshSkew = 10 * time.Second

// Store caches a bearer token for multi-stage client usage.
type Store struct {
	token     string
	expiresAt time.Time
}

// Token returns the cached token when still valid.
func (s *Store) Token() (string, bool) {
	if s == nil || s.token == "" || !time.Now().Before(s.expiresAt.Add(-refreshSkew)) {
		return "", false
	}
	return s.token, true
}

// Set stores a token and TTL in seconds.
func (s *Store) Set(token string, expiresInSec int64) {
	if s == nil {
		return
	}
	s.token = token
	s.expiresAt = time.Now().Add(time.Duration(expiresInSec) * time.Second)
}

// Clear removes the cached token.
func (s *Store) Clear() {
	if s == nil {
		return
	}
	s.token = ""
	s.expiresAt = time.Time{}
}
