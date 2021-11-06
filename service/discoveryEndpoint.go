package service

import (
	"encoding/json"
	"net/http"
)

type OpenIDConfiguration struct {
	Issuer                                     string   `json:"issuer"`
	AuthorizationEndpoint                      string   `json:"authorization_endpoint"`
	TokenEndpoint                              string   `json:"token_endpoint"`
	UserInfoEndpoint                           string   `json:"userinfo_endpoint,omitempty"`
	JwksUri                                    string   `json:"jwks_uri"`
	RegistrationEndpoint                       string   `json:"registration_endpoint,omitempty"`
	ScopesSupported                            []string `json:"scopes_supported,omitempty"`
	ResponseTypesSupported                     []string `json:"response_types_supported"`
	ResponseModesSupported                     []string `json:"response_modes_supported,omitempty"`
	GrantTypesSupported                        []string `json:"grant_types_supported,omitempty"`
	AcrValuesSupported                         []string `json:"acr_values_supported,omitempty"`
	SubjectTypesSupported                      []string `json:"subject_types_supported,omitempty"`
	IdTokenSigningAlgValuesSupported           []string `json:"id_token_signing_alg_values_supported"`
	IdTokenEncryptionAlgValuesSupported        []string `json:"id_token_encryption_alg_values_supported,omitempty"`
	IdTokenEncryptionEncValuesSupported        []string `json:"id_token_encryption_enc_values_supported,omitempty"`
	UserinfoSigningAlgValuesSupported          []string `json:"userinfo_signing_alg_values_supported,omitempty"`
	UserinfoEncryptionAlgValuesSupported       []string `json:"userinfo_encryption_alg_values_supported,omitempty"`
	UserinfoEncryptionEncValuesSupported       []string `json:"userinfo_encryption_enc_values_supported,omitempty"`
	RequestObjectSigningAlgValuesSupported     []string `json:"request_object_signing_alg_values_supported,omitempty"`
	RequestObjectEncryptionAlgValuesSupported  []string `json:"request_object_encryption_alg_values_supported,omitempty"`
	RequestObjectEncryptionEncValuesSupported  []string `json:"request_object_encryption_enc_values_supported,omitempty"`
	TokenEndpointAuthMethodsSupported          []string `json:"token_endpoint_auth_methods_supported,omitempty"`
	TokenEndpointAuthSigningAlgValuesSupported []string `json:"token_endpoint_auth_signing_alg_values_supported,omitempty"`
	DisplayValuesSupported                     []string `json:"display_values_supported,omitempty"`
	ClaimTypesSupported                        []string `json:"claim_types_supported,omitempty"`
	ClaimsSupported                            []string `json:"claims_supported,omitempty"`
	ServiceDocumentation                       string   `json:"service_documentation,omitempty"`
	ClaimsLocalesSupported                     []string `json:"claims_locales_supported,omitempty"`
	UiLocalesSupported                         []string `json:"ui_locales_supported,omitempty"`
	ClaimsParameterSupported                   bool     `json:"claims_parameter_supported,omitempty"`
	RequestParameterSupported                  bool     `json:"request_parameter_supported,omitempty"`
	RequestUriParameterSupported               bool     `json:"request_uri_parameter_supported,omitempty"`
	RequireRequestUriRegistration              bool     `json:"require_request_uri_registration,omitempty"`
	OpPolicyUri                                string   `json:"op_policy_uri,omitempty"`
	OpTosUri                                   string   `json:"op_tos_uri,omitempty"`
}

func DiscoveryEndpointHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	baseUrl := "http://localhost"
	jsonResp, err := json.Marshal(OpenIDConfiguration{
		Issuer:                                     baseUrl,
		AuthorizationEndpoint:                      baseUrl + AuthorizeEndpoint,
		TokenEndpoint:                              baseUrl + TokenEndpoint,
		UserInfoEndpoint:                           "",
		JwksUri:                                    baseUrl + JwksUriEndpoint,
		RegistrationEndpoint:                       "",
		ScopesSupported:                            nil,
		ResponseTypesSupported:                     []string{"code"},
		ResponseModesSupported:                     nil,
		GrantTypesSupported:                        []string{"authorization_code"},
		AcrValuesSupported:                         nil,
		SubjectTypesSupported:                      nil,
		IdTokenSigningAlgValuesSupported:           []string{"RS256"},
		IdTokenEncryptionAlgValuesSupported:        nil,
		IdTokenEncryptionEncValuesSupported:        nil,
		UserinfoSigningAlgValuesSupported:          nil,
		UserinfoEncryptionAlgValuesSupported:       nil,
		UserinfoEncryptionEncValuesSupported:       nil,
		RequestObjectSigningAlgValuesSupported:     nil,
		RequestObjectEncryptionAlgValuesSupported:  nil,
		RequestObjectEncryptionEncValuesSupported:  nil,
		TokenEndpointAuthMethodsSupported:          nil,
		TokenEndpointAuthSigningAlgValuesSupported: nil,
		DisplayValuesSupported:                     nil,
		ClaimTypesSupported:                        nil,
		ClaimsSupported:                            nil,
		ServiceDocumentation:                       "",
		ClaimsLocalesSupported:                     nil,
		UiLocalesSupported:                         nil,
		ClaimsParameterSupported:                   false,
		RequestParameterSupported:                  false,
		RequestUriParameterSupported:               false,
		RequireRequestUriRegistration:              false,
		OpPolicyUri:                                "",
		OpTosUri:                                   "",
	})
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Write(jsonResp)
}
