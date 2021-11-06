package service

import (
	"errors"
	"net/http"

	"yaas/dal"
)

const (
	authRequestCookie        = "yaas-auth-req"
	authRequestMaxAgeSeconds = 3600
)

func getAuthRequest(r *http.Request, sessionStore dal.SessionStore) (*authRequest, error) {
	session, err := sessionStore.Get(r, authRequestCookie)
	if err != nil {
		return nil, err
	}

	if len(session.Values()) == 0 {
		return nil, nil
	}
	var result = &authRequest{parameters: map[string]string{}}
	for sessionKey, sessionValue := range session.Values() {
		queryParam, ok := sessionKey.(string)
		if !ok {
			return nil, errors.New("invalid yaas-auth-req cookie")
		}
		queryParamValue, ok := sessionValue.(string)
		if !ok {
			return nil, errors.New("invalid yaas-auth-req cookie")
		}
		result.parameters[queryParam] = queryParamValue
	}

	return result, nil
}

func storeAuthRequest(r *http.Request, w http.ResponseWriter, sessionStore dal.SessionStore, request *authRequest) error {
	session, err := sessionStore.Get(r, authRequestCookie)
	if err != nil {
		return err
	}

	session.Clear()

	for queryParam, queryParamValue := range request.parameters {
		session.WithValue(queryParam, queryParamValue)
	}

	return sessionStore.Save(r, w, session.
		WithHttpOnly(true).
		WithSameSiteMode(http.SameSiteStrictMode).
		WithSecure(true).
		WithMaxAge(authRequestMaxAgeSeconds))
}

func removeAuthRequest(r *http.Request, w http.ResponseWriter, sessionStore dal.SessionStore) error {
	session, err := sessionStore.Get(r, authRequestCookie)
	if err != nil {
		return err
	}

	session.Clear()

	return sessionStore.Save(r, w, session.
		WithMaxAge(-1))
}
