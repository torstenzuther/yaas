package service

const (
	// AuthorizeEndpoint is the endpoint to send authentication requests to (e.g. for Authorization Code Flow)
	AuthorizeEndpoint = "/authorize"
	// TokenEndpoint is the endpoint to send token requests to
	TokenEndpoint = "/token"
	// LoginEndpoint is the endpoint for user authentication, e.g. login form
	LoginEndpoint     = "/login"
	ConsentEndpoint   = "/consent"
	AdminUI           = "/admin"
	DiscoveryEndpoint = "/.well-known/openid-configuration"
	JwksUriEndpoint   = "/.well-known/jwks.json"
)

// openID scope
const (
	openID = "openid"
)

// error related values
const (
	errParam         = "error"
	errorDescription = "error_description"
	errorURI         = "error_uri"
)

// errors (custom)
const (
	// invalidClient is not defined in OAuth2
	invalidClient = "invalid_client"
	// invalidRedirectURI is not defined in OAuth2
	invalidRedirectURI = "invalid_redirect_uri"
)

// errors (OpenID Connect)
const (
	loginRequired = "login_required"
)

// errors (OAuth2)
const (
	invalidRequest          = "invalid_request"
	unauthorizedClient      = "unauthorized_client"
	accessDenied            = "access_denied"
	unsupportedResponseType = "unsupported_response_type"
	invalidScope            = "invalid_scope"
	serverError             = "server_error"
	temporarilyUnavailable  = "temporarily_unavailable"
)

// response_type values
const (
	code = "code"
)

// Auth request parameters
const (
	responseType        = "response_type"
	scope               = "scope"
	clientID            = "client_id"
	codeChallenge       = "code_challenge"
	codeChallengeMethod = "code_challenge_method"
	redirectURI         = "redirect_uri"
	state               = "state"
	responseMode        = "response_mode"
	nonce               = "nonce"
	display             = "display"
	prompt              = "prompt"
	maxAge              = "max_age"
	uiLocales           = "ui_locales"
	idTokenHint         = "id_token_hint"
	loginHint           = "login_hint"
	acrValues           = "acr_values"
)
