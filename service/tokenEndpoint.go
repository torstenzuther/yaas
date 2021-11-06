package service

import (
	"net/http"

	"yaas/dal"
)

func TokenEndpointHandler(w http.ResponseWriter, r *http.Request, stores *dal.Stores) {
	http.Error(w, "", http.StatusInternalServerError)
}
