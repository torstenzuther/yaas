package dal

import (
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/wader/gormstore/v2"
	"gorm.io/gorm"
)

type (
	SessionStore interface {
		Get(r *http.Request, name string) (*Session, error)
		New(r *http.Request, name string) (*Session, error)
		Save(r *http.Request, w http.ResponseWriter, session *Session) error
	}

	Session struct {
		session *sessions.Session
	}
)

type (
	sessionStore struct {
		store *gormstore.Store
	}
)

func (s *Session) WithMaxAge(maxAge int) *Session {
	s.session.Options.MaxAge = maxAge
	return s
}

func (s *Session) WithValue(key string, value string) *Session {
	s.session.Values[key] = value
	return s
}

func (s *Session) WithHttpOnly(httpOnly bool) *Session {
	s.session.Options.HttpOnly = httpOnly
	return s
}

func (s *Session) WithSecure(secure bool) *Session {
	s.session.Options.Secure = secure
	return s
}

func (s *Session) WithDomain(domain string) *Session {
	s.session.Options.Domain = domain
	return s
}

func (s *Session) WithPath(path string) *Session {
	s.session.Options.Path = path
	return s
}

func (s *Session) WithSameSiteMode(sameSite http.SameSite) *Session {
	s.session.Options.SameSite = sameSite
	return s
}

func (s *Session) ValueAsString(key string) (string, bool) {
	if value, ok := s.session.Values[key]; ok {
		if stringValue, ok := value.(string); ok {
			return stringValue, true
		}
	}
	return "", false
}

func (s *Session) Values() map[interface{}]interface{} {
	return s.session.Values
}

func (s *Session) Clear() {
	s.session.Values = map[interface{}]interface{}{}
}

func (s *sessionStore) Get(r *http.Request, name string) (*Session, error) {
	internalSession, err := s.store.Get(r, name)
	if err != nil {
		return nil, err
	}
	return &Session{session: internalSession}, nil
}

func (s *sessionStore) New(r *http.Request, name string) (*Session, error) {
	internalSession, err := s.store.New(r, name)
	if err != nil {
		return nil, err
	}
	return &Session{session: internalSession}, nil
}

func (s *sessionStore) Save(r *http.Request, w http.ResponseWriter, session *Session) error {
	if err := s.store.Save(r, w, session.session); err != nil {
		return err
	}
	return nil
}

func initSessionStore(db *gorm.DB, cleanupPeriod time.Duration, maxAge int, keys ...[]byte) (*sessionStore, chan struct{}) {
	quit := make(chan struct{})
	store := gormstore.New(db, keys...)
	store.SessionOpts.Secure = true
	store.SessionOpts.HttpOnly = true
	store.SessionOpts.MaxAge = maxAge
	go store.PeriodicCleanup(cleanupPeriod, quit)
	return &sessionStore{store: store}, quit
}
