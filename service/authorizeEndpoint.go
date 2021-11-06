package service

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"yaas/dal"
	"yaas/logic"
)

func addQueryParam(u *url.URL, key string, value string) {
	q := u.Query()
	q.Add(key, value)
	u.RawQuery = q.Encode()
}

func redirect(w http.ResponseWriter, r *http.Request, redirectURL *url.URL) {
	http.Redirect(w, r, redirectURL.String(), http.StatusSeeOther)
}

func AuthorizeEndpointHandler(w http.ResponseWriter, r *http.Request, stores *dal.Stores) {
	switch r.Method {
	case http.MethodPost, http.MethodGet:
		if err := r.ParseForm(); err != nil {
			http.Error(w, invalidRequest, http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	request, err := parseAuthRequest(r.Form)
	if err != nil {
		http.Error(w, invalidRequest, http.StatusBadRequest)
		return
	}

	clientID := request.parameters[clientID]
	if !logic.IsClientID(clientID) {
		http.Error(w, invalidClient, http.StatusBadRequest)
		return
	}

	client, err := stores.ClientStore.GetClient(clientID)
	if err != nil {
		http.Error(w, invalidClient, http.StatusBadRequest)
		return
	}

	var redirectURL *url.URL
	rawRedirectURI, redirectURIDefined := request.parameters[redirectURI]

	if !redirectURIDefined {
		if len(client.RedirectURIs) > 1 {
			http.Error(w, invalidRedirectURI, http.StatusBadRequest)
			return
		}
		rawRedirectURI = client.RedirectURIs[0]
		redirectURIDefined = true
	}

	redirectURL, err = url.ParseRequestURI(rawRedirectURI)
	if err != nil {
		http.Error(w, invalidRedirectURI, http.StatusBadRequest)
		return
	}

	if !logic.IsRedirectURIValid(rawRedirectURI, client.RedirectURIs, client.AllowHTTP && !client.IsPublic) {
		http.Error(w, invalidRedirectURI, http.StatusBadRequest)
		return
	}

	stateValue, stateDefined := request.parameters[state]
	if stateDefined {
		if !logic.IsState(stateValue) {
			addQueryParam(redirectURL, errParam, invalidRequest)
			redirect(w, r, redirectURL)
			return
		}

		addQueryParam(redirectURL, state, stateValue)
	}

	var scopes []string
	var scopesDefined bool
	if rawScope, ok := request.parameters[scope]; ok {
		scopes, ok = logic.ParseScope(rawScope)
		if !ok {
			addQueryParam(redirectURL, errParam, invalidScope)
			redirect(w, r, redirectURL)
			return
		}
		scopesDefined = true
	}
	fmt.Printf("%v %v\n", scopes, scopesDefined)

	switch request.parameters[responseType] {
	case code:
		challenge, method, err := handlePKCE(request, client.IsPublic)
		if err != nil {
			addQueryParam(redirectURL, errParam, invalidRequest)
			redirect(w, r, redirectURL)
			return
		}

		fmt.Printf("%v %v\n", challenge, method)
	default:
		addQueryParam(redirectURL, errParam, unsupportedResponseType)
		redirect(w, r, redirectURL)
		return
	}

	request.parameters["computedRedirectURI"] = redirectURL.String()
	if err := storeAuthRequest(r, w, stores.SessionStore, request); err != nil {
		addQueryParam(redirectURL, errParam, serverError)
		redirect(w, r, redirectURL)
		return
	}

	http.Redirect(w, r, LoginEndpoint, http.StatusSeeOther)
}

func handlePKCE(request *authRequest, clientIsPublic bool) (*string, logic.CodeMethod, error) {
	var (
		challenge       *string
		challengeMethod logic.CodeMethod
	)

	if codeChallengeRequestParam, ok := request.parameters[codeChallenge]; ok {
		challenge = &codeChallengeRequestParam
		if !logic.IsCodeVerifierOrChallenge(codeChallengeRequestParam) {
			return nil, logic.Plain, errors.New("invalid code challenge")
		}
		if codeChallengeMethodRequestParam, ok := request.parameters[codeChallengeMethod]; ok {
			challengeMethod = logic.CodeMethod(codeChallengeMethodRequestParam)
		}
	} else {
		if clientIsPublic {
			return nil, logic.Plain, errors.New("PKCE required for public client")
		}
	}

	return challenge, challengeMethod, nil
}
