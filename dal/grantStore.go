package dal

import (
	"time"

	"gorm.io/gorm"
	"yaas/dal/model"
	"yaas/logic"
)

type (
	AuthCodeStoreRequest struct {
		CodeChallenge       string
		CodeChallengeMethod string
		RedirectURI         string
		Scope               string
		ClientID            string
		UserName            string
	}

	GrantStore interface {
		StoreNewAuthorizationCode(request AuthCodeStoreRequest) (string, error)
	}

	gormGrantStore struct {
		db *gorm.DB
	}
)

func (store *gormGrantStore) StoreNewAuthorizationCode(request AuthCodeStoreRequest) (string, error) {
	authCode, err := logic.NewAuthCode()
	if err != nil {
		return "", err
	}
	var user model.User
	var client model.Client

	err = store.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Where(&model.User{UserName: request.UserName}).First(&user).Error
		if err != nil {
			return err
		}
		err = tx.Where(&model.Client{ClientID: request.ClientID}).First(&client).Error
		if err != nil {
			return err
		}

		err = tx.Create(&model.Grant{
			GrantValue:          authCode,
			GrantType:           "authorization_code",
			ExpiresAt:           time.Now().Add(time.Hour),
			CodeChallenge:       request.CodeChallenge,
			CodeChallengeMethod: request.CodeChallengeMethod,
			RedirectUri:         request.RedirectURI,
			Scope:               request.Scope,
			UserID:              user.ID,
			ClientID:            client.ID,
			Invalidated:         false,
		}).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return authCode, nil
}

func newGrantStore(db *gorm.DB) *gormGrantStore {
	return &gormGrantStore{
		db: db,
	}
}
