package service

import "errors"

// authRequest contains all parameters of the authorization request
// sent to the authorization endpoint.
type authRequest struct {
	parameters map[string]string
}

func parseAuthRequest(args map[string][]string) (*authRequest, error) {
	var authRequest = &authRequest{
		parameters: map[string]string{},
	}
	for key, value := range args {
		switch key {
		case
			responseType,
			scope,
			state,
			responseMode,
			nonce,
			display,
			prompt,
			maxAge,
			uiLocales,
			idTokenHint,
			loginHint,
			acrValues,
			clientID,
			codeChallenge,
			codeChallengeMethod,
			redirectURI:
			if len(value) != 1 {
				return nil, errors.New(invalidRequest)
			}
			authRequest.parameters[key] = value[0]
		}
	}
	return authRequest, nil
}
