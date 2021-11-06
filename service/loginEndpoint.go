package service

import (
	"net/http"

	"yaas/dal"
	"yaas/service/templates"
)

func LoginEndpointHandler(w http.ResponseWriter, r *http.Request, stores *dal.Stores) {
	switch r.Method {
	case http.MethodPost:
		handleLoginFormPost(w, r, stores)
	case http.MethodGet:
		handleLoginGet(w, r, stores.SessionStore)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

func handleLoginGet(w http.ResponseWriter, r *http.Request, sessionStore dal.SessionStore) {
	userName, isAuthenticated, err := isAuthenticated(r, sessionStore)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	if isAuthenticated {
		authReq, err := getAuthRequest(r, sessionStore)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		if authReq == nil {
			handleWelcomePage(w, userName)
			return
		}
		http.Redirect(w, r, ConsentEndpoint, http.StatusSeeOther)
	} else {
		templates.LoginHtml(w, r)
	}
}

func handleWelcomePage(w http.ResponseWriter, userName string) {
	if err := templates.WelcomeHtml(w, userName); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func handleLoginFormPost(w http.ResponseWriter, r *http.Request, stores *dal.Stores) {
	if err := r.ParseForm(); err == nil {
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")

		if err := authenticate(r, w, stores, username, password); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
