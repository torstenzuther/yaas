package main

import (
	"log"
	"net/http"
	"time"

	"yaas/dal"
	"yaas/service"
)

const (
	cleanupSessionStorePeriodSeconds = 1800
	sessionMaxAge                    = 3600
	hashKey                          = "ajklsdnlkasnl123nlö1nlk2nölkassdasdasdda"
	encryptionKey                    = "lasdälä123äölaslödmaaaaaaaasdasmdlmle"
)

type yaasHandlerFunc func(http.ResponseWriter, *http.Request, *dal.Stores)

func main() {
	stores, quit, err := dal.InitDAL(cleanupSessionStorePeriodSeconds*time.Second, sessionMaxAge, []byte(hashKey)[:32], []byte(encryptionKey)[:32])
	if err != nil {
		log.Fatal(err)
	}
	defer close(quit)

	handlerWithSessionStoreAccess := func(handler yaasHandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			handler(w, r, stores)
		}
	}

	handlers := map[string]yaasHandlerFunc{
		service.LoginEndpoint:     service.LoginEndpointHandler,
		service.AuthorizeEndpoint: service.AuthorizeEndpointHandler,
		service.TokenEndpoint:     service.TokenEndpointHandler,
		service.ConsentEndpoint:   service.ConsentEndpointHandler,
	}

	/*
		API

		Users (only admin):
		GET /api/user -> (also get for own user possible)
		POST /api/user
			all required user data in JSON
		PUT /api/user
		PATCH /api/user (for own user possible)
		DELETE /api/user (soft delete, for own user possible)

		Clients (only admin):
		GET /api/client... (+ redirect URI)
		POST, PATCH, DELETE, ...

		Grants (user can only change his own grants):
		GET /api/grant
		DELETE /api/grant

		later: user groups, client to user association
	*/

	mux := http.NewServeMux()
	for route, handler := range handlers {
		mux.HandleFunc(route, handlerWithSessionStoreAccess(handler))
	}
	mux.HandleFunc(service.DiscoveryEndpoint, service.DiscoveryEndpointHandler)
	mux.HandleFunc(service.JwksUriEndpoint, service.JwksUriEndpointHandler)
	//https://github.com/gorilla/csrf
	//log.Fatal(http.ListenAndServe(":80", csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(false))(mux)))
	log.Fatal(http.ListenAndServe(":80", mux))
}
