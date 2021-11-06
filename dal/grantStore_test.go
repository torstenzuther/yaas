package dal

import "testing"

func TestGormGrantStore_StoreNewAuthorizationCode(t *testing.T) {
	db, err := initDatabase(":memory:")
	if err != nil {
		t.Error(err)
		return
	}
	grantStore := newGrantStore(db)
	grantStore.StoreNewAuthorizationCode(AuthCodeStoreRequest{
		CodeChallenge:       "",
		CodeChallengeMethod: "",
		RedirectURI:         "",
		Scope:               "",
		ClientID:            "",
		UserName:            "",
	})
}
