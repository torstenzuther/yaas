package logic

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/url"
)

const (
	Plain CodeMethod = "plain"
	S256  CodeMethod = "S256"
)

type CodeMethod string

func IsCodeVerifierOrChallenge(codeVerifier string) bool {
	if len(codeVerifier) < 43 || len(codeVerifier) > 128 {
		return false
	}
	for _, c := range codeVerifier {
		if !isALPHA(c) && !isDIGIT(c) && c != '-' && c != '.' && c != '_' && c != '~' {
			return false
		}
	}
	return true
}

func createCodeChallenge(codeVerifier []byte, method CodeMethod) ([]byte, error) {
	if !IsCodeVerifierOrChallenge(string(codeVerifier)) {
		return nil, errors.New("invalid code verifier")
	}
	switch method {
	case Plain:
		return codeVerifier, nil
	case S256:
		hash := sha256.Sum256(codeVerifier)
		encodedLen := base64.RawURLEncoding.EncodedLen(len(hash))
		codeChallenge := make([]byte, encodedLen)
		base64.RawURLEncoding.Encode(codeChallenge, hash[:])
		return codeChallenge, nil
	default:
		return nil, errors.New("not implemented")
	}
}

func isVSCHAR(char rune) bool {
	return char >= 0x20 && char <= 0x7e
}

func isNQCHAR(char rune) bool {
	return char == 0x21 ||
		(char >= 0x23 && char <= 0x5b) ||
		(char >= 0x5d && char <= 0x7e)
}

func isNQSCHAR(char rune) bool {
	return char == 0x20 ||
		char == 0x21 ||
		(char >= 0x23 && char <= 0x5b) ||
		(char >= 0x5d && char <= 0x7e)
}

func isUNICODECHARNOCRLF(char rune) bool {
	return char == 0x09 ||
		(char >= 0x20 && char <= 0x7e) ||
		(char >= 0x80 && char <= 0xd7ff) ||
		(char >= 0xe000 && char <= 0xfffd) ||
		(char >= 0x10000 && char <= 0x10ffff)
}

func isOnlyVSCHAR(value string) bool {
	for _, char := range value {
		if !isVSCHAR(char) {
			return false
		}
	}
	return true
}

func isOnlyDIGIT(value string) bool {
	for _, char := range value {
		if !isDIGIT(char) {
			return false
		}
	}
	return true
}

func isOnlyNQSCHAR(value string) bool {
	for _, char := range value {
		if !isNQSCHAR(char) {
			return false
		}
	}
	return true
}

func IsClientID(clientId string) bool {
	return isOnlyVSCHAR(clientId)
}

func isClientSecret(clientSecret string) bool {
	return isOnlyVSCHAR(clientSecret)
}

func isDIGIT(char rune) bool {
	return char >= 0x30 && char <= 0x39
}

func isALPHA(char rune) bool {
	return char >= 0x41 && char <= 0x5a ||
		char >= 0x61 && char <= 0x7a
}

func isResponseType(responseType string) bool {
	var spaceAllowed bool
	for _, char := range responseType {
		if char == '_' || isDIGIT(char) || isALPHA(char) {
			spaceAllowed = true
			continue
		}
		if spaceAllowed && char == 0x20 {
			spaceAllowed = false
			continue
		}
		return false
	}
	return spaceAllowed
}

func ParseScope(scope string) ([]string, bool) {
	var spaceAllowed bool
	var scopes []string
	var last int
	for i, char := range scope {
		if isNQCHAR(char) {
			spaceAllowed = true
			continue
		}
		if spaceAllowed && char == 0x20 {
			spaceAllowed = false
			scopes = append(scopes, scope[last:i])
			last = i + 1
			continue
		}
		return nil, false
	}
	if !spaceAllowed {
		return nil, false
	}
	return append(scopes, scope[last:]), spaceAllowed
}

func IsState(state string) bool {
	return len(state) > 0 && isOnlyVSCHAR(state)
}

func isError(error string) bool {
	return len(error) > 0 && isOnlyNQSCHAR(error)
}

func isErrorDescription(errorDescription string) bool {
	return len(errorDescription) > 0 && isOnlyNQSCHAR(errorDescription)
}

func isGrantType(grantType string) bool {
	isName := len(grantType) > 0 && isOnlyNameChar(grantType)
	if !isName {
		_, err := url.ParseRequestURI(grantType)
		return err != nil
	}
	return true
}

func isCode(code string) bool {
	return len(code) > 0 && isOnlyVSCHAR(code)
}

func isAccessToken(accessToken string) bool {
	return len(accessToken) > 0 && isOnlyVSCHAR(accessToken)
}

func isTokenType(tokenType string) bool {
	isName := len(tokenType) > 0 && isOnlyNameChar(tokenType)
	if !isName {
		_, err := url.ParseRequestURI(tokenType)
		return err != nil
	}
	return true
}

func isExpiresIn(expiresIn string) bool {
	return len(expiresIn) > 0 && isOnlyDIGIT(expiresIn)
}

func isUserName(userName string) bool {
	return isOnlyUNICODECHARNOLF(userName)
}

func isPassword(password string) bool {
	return isOnlyUNICODECHARNOLF(password)
}

func isRefreshToken(refreshToken string) bool {
	return len(refreshToken) > 0 && isOnlyVSCHAR(refreshToken)
}

func isEndpointParameter(endpointParameter string) bool {
	return len(endpointParameter) > 0 && isOnlyNameChar(endpointParameter)
}

func isNameChar(char rune) bool {
	return char == '-' || char == '.' || char == '_' || isDIGIT(char) || isALPHA(char)
}

func isOnlyNameChar(value string) bool {
	for _, char := range value {
		if !isNameChar(char) {
			return false
		}
	}
	return true
}

func isOnlyUNICODECHARNOLF(value string) bool {
	for _, char := range value {
		if !isUNICODECHARNOCRLF(char) {
			return false
		}
	}
	return true
}
