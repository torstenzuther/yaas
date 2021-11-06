package service

import (
	"net/http"

	"github.com/gorilla/sessions"
	"yaas/dal"
)

const (
	userName   = "username"
	authCookie = "yaas-auth"
	authMaxAge = 3600
)

func authenticate(r *http.Request, w http.ResponseWriter, stores *dal.Stores, username string, password string) error {
	if stores.UserStore.Authenticate(username, password) {
		session, err := stores.SessionStore.New(r, authCookie)
		if err != nil {
			return err
		}

		if err := stores.SessionStore.Save(r, w, session.
			WithValue(userName, username).
			WithHttpOnly(true).
			WithSameSiteMode(http.SameSiteStrictMode).
			WithSecure(true).
			WithMaxAge(authMaxAge)); err != nil {
			return err
		}
	}
	return nil
}

func isAuthenticated(r *http.Request, sessionStore dal.SessionStore) (string, bool, error) {
	session, err := sessionStore.Get(r, authCookie)
	if err != nil {
		return "", false, err
	}
	if userName, userAuthenticated := session.ValueAsString(userName); userAuthenticated {
		return userName, true, nil
	} else {
		return "", false, nil
	}
}

func logout(r *http.Request, w http.ResponseWriter, sessionStore dal.SessionStore) error {
	session, err := sessionStore.Get(r, authCookie)
	if err != nil {
		return err
	}
	session.WithMaxAge(-1)
	if err := sessions.Save(r, w); err != nil {
		return err
	}
	return nil
}
