package logic

import "strings"

func IsRedirectURIValid(redirectURI string, redirectURIs []string, allowHttp bool) bool {
	http := strings.HasPrefix(redirectURI, "http")
	https := http && len(redirectURI) >= 5 && redirectURI[4] == 's'
	protocolAllowed := https || http && allowHttp
	if !protocolAllowed {
		return false
	}
	for _, allowedRedirectURI := range redirectURIs {
		if redirectURI == allowedRedirectURI {
			return true
		}
	}
	return false
}
