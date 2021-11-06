package service

import (
	"net/http"
	"net/url"
	"strings"

	"yaas/dal"
	"yaas/service/templates"
)

func ConsentEndpointHandler(w http.ResponseWriter, r *http.Request, stores *dal.Stores) {
	userName, isAuthenticated, err := isAuthenticated(r, stores.SessionStore)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		if isAuthenticated {
			authReq, err := getAuthRequest(r, stores.SessionStore)
			if err != nil || authReq == nil {
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
			err = removeAuthRequest(r, w, stores.SessionStore)
			if err != nil {
				http.Error(w, "", http.StatusInternalServerError) // redirect error
				return
			}
			if r.Form.Get("decline") != "" {
				return // TODO redirect error
			}

			if r.Form.Get("accept") == "" {
				http.Error(w, "", http.StatusBadRequest) // redirect error?
				return
			}
			computedRedirectURI, ok := authReq.parameters["computedRedirectURI"]
			if !ok {
				http.Error(w, "", http.StatusInternalServerError) // redirect error?
				return
			}

			authCode, err := stores.GrantStore.StoreNewAuthorizationCode(dal.AuthCodeStoreRequest{
				CodeChallenge:       authReq.parameters[codeChallenge],
				CodeChallengeMethod: authReq.parameters[codeChallengeMethod],
				RedirectURI:         authReq.parameters[redirectURI],
				Scope:               authReq.parameters[scope],
				ClientID:            authReq.parameters[clientID],
				UserName:            userName,
			})
			if err != nil {
				// add error
				http.Redirect(w, r, computedRedirectURI, http.StatusSeeOther)
			} else {
				// add auth code
				u, _ := url.Parse(computedRedirectURI)
				addQueryParam(u, code, authCode)
				redirect(w, r, u)
			}
			return
		}
	}
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if isAuthenticated {
		authReq, err := getAuthRequest(r, stores.SessionStore)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		if authReq == nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		templates.ConsentHtml(w, r, strings.Split(authReq.parameters[scope], " "))
		return
	}
	http.Error(w, "", http.StatusUnauthorized)
}
