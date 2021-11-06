package service

import (
	"encoding/json"
	"net/http"
)

type Jwk struct {
	KeyType              string   `json:"kty"`
	PublicKeyUse         string   `json:"use,omitempty"`
	KeyOperations        []string `json:"key_ops,omitempty"`
	Algorithm            string   `json:"alg,omitempty"`
	KeyID                string   `json:"kid,omitempty"`
	X509URL              string   `json:"x5u,omitempty"`
	X509CertificateChain []string `json:"x5c,omitempty"`
	X509Thumbprint       string   `json:"x5t,omitempty""`
	X509SHA256Thumbprint string   `json:"x5t#S256,omitempty""`
}

type Jwks struct {
	Keys []Jwk `json:"keys"`
}

func JwksUriEndpointHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	jsonResp, err := json.Marshal(Jwks{
		Keys: []Jwk{
			{
				KeyType:      "RSA",
				PublicKeyUse: "sig",
				//KeyOperations: []string{"sign", "verify"},
				Algorithm: "RS256",
				//KeyID: "",
			}},
	})
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Write(jsonResp)
}
